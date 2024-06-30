package main

import (
	"fmt"
	"product-rest-api/app/config"
	"product-rest-api/app/database"
	"product-rest-api/app/migration"
	"product-rest-api/app/route"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDBMysql(cfg)
	migration.InitMigrationMysql(db)

	e := echo.New()

	route.InitUserRouter(db, e)
	route.InitProductRouter(db, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVERPORT)))
}
