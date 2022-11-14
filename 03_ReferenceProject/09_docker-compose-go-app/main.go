package main

import (
	"github.com/huavanthong/MasterGolang/tree/main/03_ReferenceProject/09_docker-compose-go-app/database"
	"github.com/huavanthong/MasterGolang/tree/main/03_ReferenceProject/09_docker-compose-go-app/transport"

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
