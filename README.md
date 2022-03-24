# GraphQl pathao project

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
