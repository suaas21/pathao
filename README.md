# GraphQl pathao project

## Prerequisites
- [go-chi](github.com/go-chi/chi)
- [graphql api](github.com/graphql-go/graphql)
- [Database ArangoDB](https://www.arangodb.com/)

## Start Api

## Build
```bash
$ ./build.sh
or
$ make build
```

## Application binary
```bash
$ make serve
```

## Container dev
```bash
$ docker-compose up --build
or
$ make serve
```

## GuideLine

* api folder contains api code

* infra contains drivers like db, messaging, cache etc
* repo folder contains database code
* model folder contains model
* schema folder contains application graphql schema code

### flow
> cmd -> api -> schema -> repo -> db infra


## User API:

**Quick start :**

Deploy app server and database server using docker compose
```bash
$ make build
$ make server
```

Wait a few minutes for two server going to server. App serve may be re-started for accessing db server.
When two server have gone to functional then you need to create database `collection` using ArandoDB UI

**Create DB Collection :**
- go to [DB UI](http://localhost:8529/_db/_system/_admin/aardvark/index.html#collections)
- create `User` collection 

**Mutation :**

- Create User:
```json
curl --request POST \
--url http://localhost:8080/api/v1/public/graphql \
--header 'Content-Type: application/json' \
--data '{"query":"mutation {\n\tuser(firstName: \"Sayf Uddin\", lastName: \"sagor\", password: \"1234\") {\n\t\tid\n\t\n\t}\n}\n"}'
```

**Query :**

- Get Users by ID:
```json
curl --request POST \
--url http://localhost:8080/api/v1/public/graphql \
--header 'Content-Type: application/json' \
--data '{"query":"query {\n\tusers(id: \"ab921b81-5c5c-4617-b274-69820a527327\") {\n\t\tid\n\t\tfullName\n\t\ttime\n\t}\n}\n"}'
```

## TODO
- need more improvement in ci/cd process
- need to add migration command beside serve to create database collection automatically if not created.
