package utils

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// Date stores year month day info, and matches date sql datatype
type Date struct {
	time.Time
}

// NewDate returns a new date with the given year, month & day
func NewDate(year int, month time.Month, day int) Date {
	return Date{Time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

// NewDatePtr returns a ptr to a new date with the given year, month & day
func NewDatePtr(year int, month time.Month, day int) *Date {
	return &Date{Time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

// UnmarshalJSON allows custom date unmarshalling
func (d *Date) UnmarshalJSON(data []byte) error {
	strInput := string(data)
	strInput = strings.Trim(strInput, `"`)

	newTime, err := time.Parse("2006-01-02", strInput)
	if err != nil {
		return err
	}

	d.Time = newTime

	return nil
}

// MarshalJSON allows custom date marshalling
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Format("2006-01-02") + `"`), nil
}

func (d Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement([]byte(d.Format("2006-01-02")), start)
}

// String returns the string representation of the date
func (d Date) String() string {
	return d.Format("2006-01-02")
}

// Formats the date for display (in aus anywya)
func (d Date) Display() string {
	return d.Format("02/01/2006")
}

// Scan is called by the pg driver when reading a row into a struct
func (d *Date) Scan(val interface{}) error {
	switch v := val.(type) {
	case time.Time:
		d.Time = v
		return nil
	}

	return fmt.Errorf("couldnt scan date data")
}

// CurrentUTCTime returns the current time
func CurrentUTCTime() string {
	loc, _ := time.LoadLocation("UTC")
	format := "2006-01-02T15:04:05.000Z"
	t := time.Now().In(loc)

	return t.Format(format)
}
