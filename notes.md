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

-> touch db.sqlite
-> sqlite3 db.sqlite

```sql
CREATE TABLE products (
    id STRING,
    name STRING,
    price FLOAT,
    status STRING
);
.tables
```

```
docker-compose exec app go run main.go
```

```sql
sqlite> SELECT * FROM products;
```

docker-compose exec app go test ./...

docker-compose exec app cobra init --pkg-name=appproduct

docker-compose exec app go mod tidy

docker-compose exec app go run main.go

docker-compose exec app cobra add cli

docker-compose exec app go run main.go cli

docker-compose exec app go run main.go cli --help

docker-compose exec app go run main.go cli -a=create -n="Product CLI" -p=25.0

docker-compose exec app sqlite3 db.sqlite

-> docker-compose exec app bash  
-> sqlite3 db.sqlite
-> docker-compose exec app go run main.go cli -a=get -i=7054abe9-c11f-46cd-b343-d444f3765b08

https://github.com/gorilla/mux

https://github.com/urfave/negroni

docker-compose exec app cobra add http
sudo chown -R devontherun:devontherun cmd

-> docker-compose exec app go run main.go http
curl http://localhost:9000/product/2b952214-7151-4776-9db9-f016c945efdc

-> sqlite3 db.sqlite
sqlite> SELECT \* FROM products;
