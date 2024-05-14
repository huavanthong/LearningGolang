## Getting Started

https://linuxhandbook.com/set-hostname-docker-compose/

### Started nginx service
```
docker-compose exec nginx-proxy bash -c "apt update && apt install -y iputils-ping"
```

### From Nginx ping to Letsencrypt service
```
sudo docker-compose exec nginx-proxy ping letsencrypt
```

### After setting hostname
```
sudo docker-compose exec nginx-proxy ping ledocker
```

## Issues knowledge
1. Find out the issue. https://github.com/docker/compose/issues/2925?ref=linuxhandbook.com
