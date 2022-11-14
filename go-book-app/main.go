package main

import (
	"github.com/huavanthong/MasterGolang/go-book-app/database"
	"github.com/huavanthong/MasterGolang/go-book-app/transport"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := transport.SetUp()
	err := database.InitDatabase()
	if err != nil {
		panic(err)
	}

	app.Listen(5000)
}
