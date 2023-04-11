# Reference
https://blog.logrocket.com/how-to-use-redis-as-a-database-with-go-redis/

# Getting Started
1. Buidl app
```
go build
```

2. Run app
```
./rediboard.exe
```

3. Use app
```
curl -H "Content-type: application/json" -d '{"username": "isa", "points": 25}' localhost:8080/points
```