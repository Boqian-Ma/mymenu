package db

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"

	// Postgres driver
	_ "github.com/lib/pq"
	// Needed for db migration
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/config"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/utils"
)

// DB adds some wrappers around standard sqlx functionality
type DB interface {
	// Select returns the results from a select into the given dest interface
	Select(c *gin.Context, dest interface{}, sql string, args ...interface{}) error
	// Inserts a new row into the given table based off struct tags
	Insert(c *gin.Context, table string, value interface{}) error
	// Updates the specified row based off struct tags
	Update(c *gin.Context, val interface{}, sql string, args ...interface{}) error
	// Deletes the specified row
	Delete(c *gin.Context, sql string, args ...interface{}) error
	// Exec simply execute the query
	Exec(c *gin.Context, sql string, values ...interface{}) error
	// RawDB returns a reference to the inner sqlx DB
	RawDB() *sqlx.DB
	// NamedExec do named exec
	NamedExec(c *gin.Context, sql string, value interface{}) error
	// Get gets a single row value. error is returned if no row found
	Get(c *gin.Context, value interface{}, sql string, args ...interface{}) error
}

// used as node in queue while scanning struct
type scanNode struct {
	val    reflect.Value
	prefix string
}

type db struct {
	db *sqlx.DB
}

func InitDB(cfg *config.Config, stage string) (DB, error) {
	connStr := fmt.Sprintf("user=%s sslmode=%s host=%s password=%s", cfg.DB.User, cfg.DB.SSLMode, cfg.DB.Host, cfg.DB.Password)
	if cfg.DB.Name != "" {
		connStr += " dbname=" + cfg.DB.Name
	}

	var db *sqlx.DB
	var err error
	maxRetries := 5
	for try := 0; try < maxRetries; try++ {
		db, err = sqlx.Connect("postgres", connStr)
		if err == nil {
			break
		}

		time.Sleep(time.Second)
	}

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 10)

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return nil, err
	}

	// migrate db "up" to the latest version
	if err := m.Up(); err != nil && err.Error() != "no change" {
		return nil, err
	}

	return Newdb(db), nil
}

// Newdb returns a new eps db wrapper around an sqlx.DB
func Newdb(sqlxDB *sqlx.DB) DB {
	return db{sqlxDB}
}

func (db db) RawDB() *sqlx.DB {
	return db.db
}

// Largely borrowed from https://github.com/jmoiron/sqlx/blob/master/sqlx.go
func (db db) Select(c *gin.Context, dest interface{}, sql string, args ...interface{}) error {

	value := reflect.ValueOf(dest)

	if value.Kind() != reflect.Ptr {
		return fmt.Errorf("must pass a pointer, not a value, to Select")
	}
	if value.IsNil() {
		return fmt.Errorf("nil pointer passed to Select")
	}
	direct := reflect.Indirect(value)

	rows, err := db.db.Queryx(sql, args...)
	if err != nil {
		return err
	}

	if isBaseType(value.Type(), reflect.Slice) {
		// select many results into slice
		slice := reflectx.Deref(value.Type())
		isPtr := slice.Elem().Kind() == reflect.Ptr
		base := reflectx.Deref(slice.Elem())

		for rows.Next() {
			results := make(map[string]interface{})
			if err := rows.MapScan(results); err != nil {
				panic(err)
			}

			vp := reflect.New(base)

			setDBFields(vp.Interface(), results)

			if isPtr {
				direct.Set(reflect.Append(direct, vp))
			} else {
				direct.Set(reflect.Append(direct, reflect.Indirect(vp)))
			}
		}
	} else if isBaseType(value.Type(), reflect.Struct) {
		// select one result directly into struct
		found := false
		for rows.Next() {
			if found {
				return fmt.Errorf("too many results to put into struct")
			}

			found = true
			results := make(map[string]interface{})
			if err := rows.MapScan(results); err != nil {
				panic(err)
			}

			vp := reflect.New(direct.Type())
			setDBFields(vp.Interface(), results)

			direct.Set(reflect.Indirect(vp))
		}
		if !found {
			return errors.NotFound("")
		}
	} else {
		return fmt.Errorf("Unsupported type for select %s", value.Type().String())
	}

	return nil
}

func (db db) Update(c *gin.Context, value interface{}, sql string, args ...interface{}) error {
	fields, values := listDBFields(value)

	// globally readonly fields
	filteredFields := fields[:0]
	filteredValues := []interface{}{}
	for i, field := range fields {
		if field != "id" && field != "created_at" && field != "ran" {
			filteredFields = append(filteredFields, field)
			filteredValues = append(filteredValues, values[i])
		}
	}
	fields = filteredFields
	values = filteredValues

	csFieldParams := []string{}
	for i := range fields {
		csFieldParams = append(csFieldParams, fields[i]+"=$"+strconv.Itoa(len(args)+i+1))
	}

	// Not vulnerable to SQL injects since user created values are still passed through args
	sql = strings.ReplaceAll(sql, "...", strings.Join(csFieldParams, ", "))
	res, err := db.db.Exec(sql, append(args, values...)...)
	if err != nil {
		if strings.Contains(err.Error(), "constraint") {
			return errors.BadRequest("Bad Request " + err.Error())
		}

		return err
	}

	if c, _ := res.RowsAffected(); c == 0 {
		return errors.NotFound("")
	}

	return nil
}

func (db db) Insert(c *gin.Context, table string, value interface{}) error {
	fields, values := listDBFields(value)
	csFields := strings.Join(fields, ", ")
	csParameters := ListParameters(len(fields))

	sql := "INSERT INTO " + table + " (" + csFields + ") VALUES (" + csParameters + ")"
	if _, err := db.db.Exec(sql, values...); err != nil {
		if strings.Contains(err.Error(), "constraint") {
			return errors.BadRequest("Bad Request " + err.Error())
		}

		return err
	}

	return nil
}

func (db db) NamedExec(c *gin.Context, sql string, value interface{}) error {
	_, err := db.db.NamedExec(sql, value)
	return err
}

func (db db) Get(c *gin.Context, value interface{}, sql string, args ...interface{}) error {
	err := db.db.Get(value, sql, args...)
	return err
}

func (db db) Delete(c *gin.Context, sql string, values ...interface{}) error {
	res, err := db.db.Exec(sql, values...)
	if err != nil {
		return err
	}

	if count, _ := res.RowsAffected(); count == 0 {
		return errors.NotFound("")
	}

	return nil
}

func (db db) Exec(c *gin.Context, sql string, values ...interface{}) error {
	_, err := db.db.Exec(sql, values...)
	return err
}

// ListParameters returns up to num params in sql format
func ListParameters(num int) string {
	csv := ""
	for i := 1; i <= num; i++ {
		csv += "$" + strconv.Itoa(i)

		if i < num {
			csv += ", "
		}
	}

	return csv
}

func listDBFields(arg interface{}) (fields []string, values []interface{}) {
	v := reflect.ValueOf(arg)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("DBFields requires a struct, found: %s", v.Kind().String()))
	}

	// Start queue with root struct, will perform a BFS
	queue := []scanNode{{v, ""}}
	fields = []string{}
	values = []interface{}{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		currVal := node.val
		currPrefix := node.prefix

		if currVal.Kind() == reflect.Ptr {
			currVal = currVal.Elem()
		}

		for i := 0; i < currVal.NumField(); i++ {
			field := currVal.Field(i)
			dbTag := strings.TrimSpace(currPrefix + currVal.Type().Field(i).Tag.Get("db")) // include parent struct tags

			if field.Kind() == reflect.Ptr {
				if !field.IsNil() {
					field = field.Elem()
				} else if reflect.TypeOf(field.Interface()).Elem().Kind() == reflect.Struct {
					continue
				}
			}

			if shouldTraverse(field) {
				queue = append(queue, scanNode{field, dbTag})
			} else if dbTag != "" && currVal.Type().Field(i).Tag.Get("db") != "-" {
				fields = append(fields, dbTag)

				switch field.Type() {
				case reflect.TypeOf(utils.Date{Time: time.Now()}):
					values = append(values, field.Interface().(utils.Date).String())
				case reflect.TypeOf(&utils.Date{Time: time.Now()}):
					if !field.IsNil() {
						values = append(values, field.Interface().(*utils.Date).String())
						break
					}
					fallthrough
				case reflect.TypeOf(time.Now()):
					field.Set(reflect.ValueOf(field.Interface().(time.Time).In(time.UTC)))
					fallthrough
				default:
					values = append(values, field.Interface())
				}
			}
		}
	}

	return fields, values
}

func setDBFields(arg interface{}, fieldValues map[string]interface{}) {
	v := reflect.ValueOf(arg)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("DBFields requires a struct, found: %s", v.Kind().String()))
	}

	// Init queue to root struct, will BFS and set values during traversal
	queue := []scanNode{{v, ""}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		currVal := node.val
		currPrefix := node.prefix

		for i := 0; i < currVal.NumField(); i++ {
			field := currVal.Field(i)
			dbTag := strings.TrimSpace(currPrefix + currVal.Type().Field(i).Tag.Get("db"))

			if field.Kind() == reflect.Ptr {
				if field.IsNil() && currVal.Type().Field(i).Tag.Get("db") != "" {
					subFieldExists := false
					exactMatch := false

					for k, v := range fieldValues {
						if strings.HasPrefix(k, dbTag) && v != nil {
							subFieldExists = true
						}

						if k == dbTag && v != nil {
							exactMatch = true
						}
					}

					if exactMatch || (subFieldExists && reflect.TypeOf(field.Interface()).Elem().Kind() == reflect.Struct) {
						field.Set(reflect.New(field.Type().Elem()))
					}
				}
			}

			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}

			if shouldTraverse(field) {
				queue = append(queue, scanNode{field, dbTag})
			} else if dbTag != "" && reflect.ValueOf(fieldValues[dbTag]).IsValid() {
				switch field.Type() {
				case reflect.TypeOf(0.0):
					bits := reflect.ValueOf(fieldValues[dbTag]).Interface().([]uint8)
					floatVal, err := strconv.ParseFloat(string(bits), 64)
					if err != nil {
						panic(err)
					}

					field.Set(reflect.ValueOf(floatVal))
				case reflect.TypeOf(utils.Float64(0.0)):
					bits := reflect.ValueOf(fieldValues[dbTag]).Interface().([]uint8)
					floatVal, err := strconv.ParseFloat(string(bits), 64)
					if err != nil {
						panic(err)
					}

					field.Set(reflect.ValueOf(utils.Float64(floatVal)))
				case reflect.TypeOf(0):
					if _, ok := reflect.ValueOf(fieldValues[dbTag]).Interface().([]uint8); ok {
						bits := reflect.ValueOf(fieldValues[dbTag]).Interface().([]uint8)
						intVal, err := strconv.ParseInt(string(bits), 10, 64)
						if err != nil {
							panic(err)
						}

						field.Set(reflect.ValueOf(int(intVal)))
					} else {
						intVal := reflect.ValueOf(fieldValues[dbTag]).Interface().(int64)

						field.Set(reflect.ValueOf(int(intVal)))
					}
				case reflect.TypeOf(utils.Int(0)):
					bits := reflect.ValueOf(fieldValues[dbTag]).Interface().([]uint8)
					intVal, err := strconv.ParseInt(string(bits), 10, 64)
					if err != nil {
						panic(err)
					}

					field.Set(reflect.ValueOf(utils.Int(int(intVal))))
				case reflect.TypeOf(utils.Date{Time: time.Now()}):
					field.Set(reflect.ValueOf(utils.Date{Time: reflect.ValueOf(fieldValues[dbTag]).Interface().(time.Time).UTC()}))
				case reflect.TypeOf(&utils.Date{Time: time.Now()}):
					field.Set(reflect.ValueOf(&utils.Date{Time: reflect.ValueOf(fieldValues[dbTag]).Interface().(time.Time).UTC()}))
				case reflect.TypeOf(utils.String("")):
					field.Set(reflect.ValueOf(utils.String(fieldValues[dbTag].(string))))
				default:
					if field.Kind() == reflect.Ptr { // ! bit of a hack - need to work out a better way to do this
						v := fieldValues[dbTag].(bool)
						field.Set(reflect.ValueOf(&v))
					} else {
						field.Set(reflect.ValueOf(fieldValues[dbTag]).Convert(field.Type()))
					}
				}
			}
		}
	}
}

func shouldTraverse(field reflect.Value) bool {
	// Generally speaking we should traverse all structs, except these others which
	// are treated as atomic values

	return field.Kind() == reflect.Struct &&
		field.Type() != reflect.TypeOf(time.Now()) &&
		field.Type() != reflect.TypeOf(utils.Date{Time: time.Now()})
}

func isBaseType(t reflect.Type, expected reflect.Kind) bool {
	t = reflectx.Deref(t)

	return t.Kind() == expected
}
