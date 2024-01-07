package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/mccune1224/playlist-pilot/handler"
)

func main() {
	app := echo.New()
	// Connect to DB
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("error opening database,", err)
	}

	// Generic Logger middleware
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "${status} | ${latency_human} | ${method} | ${uri} | ${error} \n",
		},
	))

	app.Use(middleware.Recover())
	// trailling slash
	app.Pre(middleware.RemoveTrailingSlash())

	handler := handler.NewHandler(db)
	app.Static("/static", "static")

	app.GET("/", handler.Index)
	app.GET("/comp", handler.Component)
	app.GET("/subcomp", handler.SubComponent)

	app.Logger.Fatal(app.Start(":" + os.Getenv("PORT")))
}
