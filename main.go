package main

import (
	"log"
	"net/http"

	"my-module/database"
	"my-module/handlers"
	"my-module/middleware"
	"my-module/model"

	"github.com/labstack/echo"

	// github.com/mattn/go-sqlite3
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gorm failed to connect to database: %v", err)
	}
	// Auto Migrate Database (creates/updates tables/column structure/schemas). This is obviously very dangerous procedure
	err = db.AutoMigrate(
		&model.MobileApplication{},
		&model.MobileAppVersion{},
		&model.SecurityCheckSpec{},
		&model.SecurityReport{},
		&model.SecurityReportEntry{},
	)
	if err != nil {
		log.Fatalf("Grom failed to auto migrate schema to database: %v", err)
	}
	err = database.Seed(db)
	if err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	e := echo.New()
	// Middleware for using CustomEchoContext (which exposes database client)
	e.Use(middleware.DatabaseMiddleware(db))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/apps/:app/versions/:version/reports", handlers.PostReport)
	e.GET("/apps/:app/versions/:version/reports", handlers.GetReports)
	e.Logger.Fatal(e.Start(":1323"))
}
