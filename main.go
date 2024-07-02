package main

import (
	"fmt"
	"os"
	"product-rest-api/app/database"
	"product-rest-api/app/migration"
	"product-rest-api/app/route"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load(".env")
	port, _ := strconv.Atoi(os.Getenv("SERVERPORT"))
	db := database.Init()
	migration.InitMigration(db)

	e := echo.New()

	route.InitUserRouter(db, e)
	route.InitProductRouter(db, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
