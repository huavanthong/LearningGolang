# Introduction
In this chapter, you’ll learn how to deploy to
* A server that’s either fully owned by you, physical, or virtual (on an Infrastructure-as-a-Service provider, Digital Ocean)
* A cloud Platform-as-a-Service (PaaS) provider, Heroku
* Another cloud PaaS provider, Google App Engine
* A dockerized container, deployed to a local Docker server, and also to a virtual machine on Digital Ocean
**Notes:** 
- It’s important to remember that there are many different ways of deploying a web app and in each of the methods you’ll be learning in this chapter, there are many variations.  
- The methods described in this chapter are based on a single person deploying the web app. The processes are usually a lot more involved in a production environment, which includes additiona tasks like running test suites, continual integration, staging servers, and so on.
- This chapter also introduces many concepts and tools, each of them worthy of an entire book. Therefore, it’s impossible to cover all of these technologies and services. This chapter aims to cover only a small portion, and if you want to learn more, take this as a starting point.

# Deploying to servers
* [What steps to deploy a Go web application to server?](#Steps)
* [How to set up driver PostgreSQL for web-app?](#Set-up-PostgreSQL)
# Deploying to Google App Engine

# Deploy to Dokcer



## Deploying to servers
### Steps
Step to deploy a web app to server:
* Create an executable binary and the running it off a server that's on the internet.
* 

### Getting Started
#### Set-up-PostgreSQL
Create a user named gwp.
```
createuser -P -d gwp
```
Create a database named gwp.
```
createdb gwp
```
Create schema and table data
```
psql -U gwp -f setup.sql -d gwp
```
#### Insert data 
Add records to Database
```
psql -U gwp -f data.sql -d gwp
```
#### Building web app
Set up web-app to account Database
```
From 
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
To
    Db, err = sql.Open("postgres", "user=postgres dbname=gwp password=1234 sslmode=disable")
```
To buidl a web server
```
go build
```
#### Run
To run it on localhosts
**GET:**
```
http://localhost:8080/post/1
```

**POST:**
```
http://localhost:8080/post/1
```

**PUT:**
```
http://localhost:8080/post/1
```

**DELETE:**
```
http://localhost:8080/post/1
```
#### Deploy to server
To run the web service, log into the server and run it from the console:
```
./ws-s
```
