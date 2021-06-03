# DevOnTheRun Notes

> notes taken during the course

<!-- https://gitignore.io -->

https://github.com/codeedu/fc2-arquitetura-hexagonal

https://gist.github.com/wesleywillians/73b275126b8d1562e32d7c3ac129784a

docker-compose exec app bash  
docker exec -it appproduct bash

go mod init appproduct

go test ./...

docker-compose exec app bash  
mockgen -destination=application/mocks/application.go -source=application/product.go application

-> touch sqlite.db
-> sqlite3 sqlite.db

```sql
CREATE TABLE products (
    id STRING,
    name STRING,
    price FLOAT,
    status STRING
);
```

.tables
