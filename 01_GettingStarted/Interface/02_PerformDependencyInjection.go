package main

import (
	"fmt"
)

type Database interface {
	Connect() error
	Disconnect() error
}

type PostgresDB struct {
}

func (db PostgresDB) Connect() error {
	fmt.Println("Connecting to Postgres database...")
	return nil
}

func (db PostgresDB) Disconnect() error {
	fmt.Println("Disconnecting from Postgres database...")
	return nil
}

type MySQLDB struct {
}

func (db MySQLDB) Connect() error {
	fmt.Println("Connecting to MySQL database...")
	return nil
}

func (db MySQLDB) Disconnect() error {
	fmt.Println("Disconnecting from MySQL database...")
	return nil
}

type Service struct {
	Database Database
}

func (s Service) PerformTask() error {
	err := s.Database.Connect()
	if err != nil {
		return err
	}

	// perform some task

	err = s.Database.Disconnect()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	postgresDB := PostgresDB{}
	mySQLDB := MySQLDB{}

	service1 := Service{Database: postgresDB}
	err := service1.PerformTask()
	if err != nil {
		fmt.Println("Error:", err)
	}

	service2 := Service{Database: mySQLDB}
	err = service2.PerformTask()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
