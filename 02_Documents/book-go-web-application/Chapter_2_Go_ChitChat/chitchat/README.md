# ChitChat

This is the simple forum application written with Go for the book "Go Web Programming" from Manning Publication Co. The source code for this application is the basis for the Chapter 2 - Go ChitChat. 

However the code is a reference and is not a 1 - 1 match with the code in the book. There are portions of the code that is more detailed in here than in Chapter 2 (which is a simplified version of the source code here).

Some differences include:

* This version of ChitChat is configurable with a `config.json` file
* This version is able to log to a `chitchat.log` file
* There are test files in this code
* All structs are fully fleshed out (in the book chapter, they are only implied)
* Some of the functions are placed as methods for the struct types instead of being a part of the package

# Set up project
S1: Create go.mod to init module for our package.
S2: create database name is chitchat 
S3: use setup.sql script, and execute in query tool in PostgreSQL to genereate Database for our data model.
S4: Remember to login account in PostgreSQL

# How to open driver PostgreSQL by Golang
### Method 1  
Normally, the Go program will open port on Driver PostgreSQL by user default on environment.  
We open driver by line code
```
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable")
```
### Method 2
To open by another account, we often set directly to source code:
```
	Db, err = sql.Open("postgres", "user=postgres password=1234 dbname=chitchat sslmode=disable")
```
### Method 3
Another way to login to Driver PostgreSQL, ta will load variable environment from your local machine.
And set up to your working project.
```
    loadEnv() // For environment variables
    psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PASSWORD"),
    )

    db, err := gorm.Open("postgres", psqlInfo)

```
Refer:  
https://stackoverflow.com/questions/55243561/cannot-connect-to-rds-postgres-using-golang-gorm
