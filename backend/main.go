package main

import (
	"github/sinnosu/go-gin-gorm-sample/src/config"
	"github/sinnosu/go-gin-gorm-sample/src/routes"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	//run all routes
	routes.Routes()
}
