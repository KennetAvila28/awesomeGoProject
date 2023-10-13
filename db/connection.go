package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "postgresql://postgres:6nlLocDRLxxuXJEuMZKS@containers-us-west-39.railway.app:6027/railway"
var DB *gorm.DB

func DbConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		panic(error)
	}
	fmt.Print("Connected to db")
}
