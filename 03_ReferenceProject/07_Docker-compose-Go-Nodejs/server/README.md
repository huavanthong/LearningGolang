# Getting Started 
To build a backend for server by docker
```bash
$ docker build -t server .  
```

To run backend by docker
```bash
$ docker run -p 8080:8080 -it server
```

To run your backend server on brower.
```
http://localhost:8080/api/v1/foods
```