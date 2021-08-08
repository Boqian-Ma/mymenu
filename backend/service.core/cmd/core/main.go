package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/auth"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/config"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/db"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/categories"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/menu_items"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/orders"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/reports"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/restaurants"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/tables"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/users"
)

// @title COMP3900 - JAMAR
// @version 0.1
// @description Restaurants and customers and other great things!

// @host localhost:5000
// @BasePath /api/v1
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	stage := os.Getenv("ENV")
	if stage == "" {
		stage = "local"
	}

	gin.SetMode(gin.ReleaseMode)

	cfg, err := config.Load(stage)
	if err != nil {
		fmt.Printf("error: failed to load application configuration: %s\n", err)
		os.Exit(1)
	}

	db, err := db.InitDB(cfg, stage)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		os.Exit(1)
	}

	// build HTTP server
	address := fmt.Sprintf(":%v", cfg.ServerPort)
	hs := &http.Server{
		Addr:    address,
		Handler: buildHandler(cfg, db),
	}

	// start the server and log any unexpected errors
	fmt.Printf("info: server running on %s with stage %s\n", address, stage)
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("error: %s", err.Error())
	}
}

func buildHandler(cfg *config.Config, db db.DB) http.Handler {
	r := gin.New()

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowHeaders:     []string{"*", "Origin", "X-Requested-With", "Accept", "Content-Type"},
		AllowMethods:     []string{"OPTIONS", "PUT", "POST", "GET", "DELETE"},
	}))
	r.Use(errors.Handler("core"))

	r.StaticFS("/core/docs", http.Dir("./service.core/docs"))

	api := r.Group("/api/v1")

	authUserHandler := auth.UserHandler(
		auth.NewRepository(db),
	)
	// User service
	users.RegisterHandlers(api.Group(""),
		users.NewService(
			users.NewRepository(db),
		),
		authUserHandler,
	)
	// Restaurant service
	restaurants.RegisterHandlers(api.Group(""),
		restaurants.NewService(
			restaurants.NewRepository(db),
			users.NewRepository(db),
		),
		authUserHandler,
		auth.ManagerHandler(),
	)
	// Menu item service
	menu_items.RegisterHandlers(api.Group(""),
		menu_items.NewService(
			menu_items.NewRepository(db),
			users.NewRepository(db),
			categories.NewRepository(db),
		),
		authUserHandler,
		auth.ManagerHandler(),
	)
	// Item category service
	categories.RegisterHandlers(api.Group(""),
		categories.NewService(
			categories.NewRepository(db),
			users.NewRepository(db),
		),
		authUserHandler,
		auth.ManagerHandler(),
	)
	// Table service
	tables.RegisterHandlers(api.Group(""),
		tables.NewService(
			tables.NewRepository(db),
			users.NewRepository(db),
		),
		authUserHandler,
		auth.ManagerHandler(),
	)
	// Ordering service
	orders.RegisterHandlers(api.Group(""),
		orders.NewService(
			orders.NewRepository(db),
			users.NewRepository(db),
			menu_items.NewRepository(db),
			tables.NewRepository(db),
		),
		authUserHandler,
		auth.ManagerHandler(),
	)
	// Reporting service
	reports.RegisterHandlers(api.Group(""),
		reports.NewService(
			reports.NewRepository(db),
			users.NewRepository(db),
			menu_items.NewRepository(db),
			orders.NewRepository(db),
			categories.NewRepository(db),
		),
		authUserHandler,
		auth.ManagerHandler(),
	)
	return r
}
