## DB setup

```
docker compose up -d
docker cp setup.sql pg:/tmp
docker compose exec pg psql -f /tmp/setup.sql -d testdb -U root
```

## request

- GET users

```
curl localhost:8888/users
```

- create user

```
curl -X POST -H "Content-Type: application/json" localhost:8888/users -d '{"name": "hoge"}'
```
