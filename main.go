package main

import (
	"fmt"
	"golang-base64-file-encryption-template/config"

	"gorm.io/gorm"
)

func main() {
	var (
		db *gorm.DB = config.SetupDatabaseConnection()
	)

	fmt.Println(db)
}