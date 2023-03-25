package main

import (
	"alta-cookit-be/app/config"
	"alta-cookit-be/app/database"
	"alta-cookit-be/app/router"
	"alta-cookit-be/middlewares"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDB(*cfg)
	database.InitialMigration(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	router.InitRouter(db, e)

	token, _ := middlewares.CreateToken(1, "User")
	fmt.Println(token)

	e.Logger.Fatal(e.Start(":8083"))

	
	
}
