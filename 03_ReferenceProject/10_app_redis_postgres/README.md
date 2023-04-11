```
docker run --rm --name test-redis redis:6.2-alpine redis-server --loglevel warning
```

```
docker exec -it test-redis redis-cli
```


```
​​docker-compose -f docker-compose-redis-only.yml up
```