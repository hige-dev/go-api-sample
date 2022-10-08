## DB setup

```
docker compose up -d
docker cp setup.sql pg:/tmp
docker compose exec pg psql -f /tmp/setup.sql -d testdb -U root
```
