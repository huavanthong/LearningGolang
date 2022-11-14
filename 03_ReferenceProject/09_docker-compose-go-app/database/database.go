package database

import (
	"fmt"
	"os"

	"github.com/huavanthong/MasterGolang/tree/main/03_ReferenceProject/09_docker-compose-go-app/book"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is the pointer to the
// database connection
var (
	DBConn *gorm.DB
)

// InitDatabase returns a pointer to a database connection
// calling functions need to ensure they defer the closing
// of this connection.
func InitDatabase() error {
	var err error

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	// dbConnectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" +strconv
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	fmt.Println(connectionString)
	DBConn, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")

	// For debugger
	// DBConn.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration
}
