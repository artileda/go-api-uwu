## Go-API-UwU

An example of Go web api with Redis and Postgresql.

## How to run this

### Ez way

- if you do change to binary
```
sudo docker-compose up --build
```

- just to spin up the container

```
sudo docker-compose up
```

### Without Docker

- Build the binary
```
go build .
```

- run the binary (make sure the postgres server and redis already up), then run this:

```
DB_HOST="postgres://username:password@host:port/dbname" \
REDIS_HOST="host:post"\
./UwU
```

### License 

[MIT](./LICENSE)