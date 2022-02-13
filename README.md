# go-todo-app

Golang（Echo）の TODO API

## SetUp

```
cp .env.sample .env
```

## Getting Started

```
docker-compose up
```

## Setup Database

```
goose up
goose -env test up
```

## Run Web Api

### Create TODO (POST)

```
http://localhost/todo-item
```

Request Body (application/json)

```
{
  "title": "title",
  "memo":"memo",
  "expired":"2022-02-12T15:55:41+09:00"
}
```

### Get TODO LiST

```
http://localhost/todo-list
```

### Get TODO Detail

```
http://localhost/todo-list/:id
```

### Update TODO (PUT)

```
http://localhost/todo-item/:id
```

Request Body (application/json)

```
{
  "title": "change title",
  "memo":"change memo",
  "expired":"2022-02-12T15:55:41+09:00"
}
```

### Delete TODO (DELETE)

```
http://localhost/todo-item/:id
```

## Run Test

```
GO_ENV=test DB_USER=api_user DB_PASSWORD=Passw0rd DB_HOST="tcp(db-test:3306)" DB_NAME=go_api_test go test github.com/kokoneko/go-todo-app/usecase/tests -v
```
