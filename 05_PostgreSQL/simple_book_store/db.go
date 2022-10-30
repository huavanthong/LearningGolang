package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// dbURL := "postgres://pg:pass@localhost:5432/crud"

	// db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=root password=root dbname=test_db port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Book{})

	return db
}
