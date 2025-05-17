# Golang Clean Architecture

The following is a folder structure pattern that I usually use, although I don't use all of them because of the project I'm working on only small projects that are not too big, so if you are interested in the pattern I made, you can use it if you think it's good, check this link for new update for this architecture [here](https://github.com/restuwahyu13/go-trakteer-api).

## Table Of Content

- [What Are The Benefits](#what-are-the-benefits-)
- [Flow Diagram](#flow-diagram)
- [Folder Structure Pattern](#folder-structure-pattern)

## What Are The Benefits ?

- [x] Easy to maintance
- [x] Easy to scalable you project
- [x] Readable code
- [x] Suitable for large projects or small projects
- [x] Easy to understand for junior or senior
- [x] And more

## Flow Diagram

<img src="./diagram.png" alt="flow-diagram"/>

## Folder Structure Pattern

```
/ todo-list
│
├── /cmd
│   ├── /api
│   │ └──  main.go
│   ├── /grpc
|   │   └── main.go
│   ├── /worker
|   │   └── main.go
|
|
├── /configs
│   ├── env.config.go
│   ├── test.config.go
|
|
├── /database
│   ├── /migrations
│   |   └── 000123456789_create_users_table.up.sql
│   ├── /seeds
│   |   └── 000123456789_create_users_table.seed.sql
│
|
├── /domain
│   ├── /entities
|   │   └── users.entitie.go
│   ├── /exceptions
|   │   └── users.exception.go
│   ├── /repositories
|   │   └── users.repositorie.go
│   ├── /services
|   │   ├── /http
|   │   │   └── /users
|   │   │       └── service.go
|   │   │       └── service_test.go
|   │   │       └── mapper.go
|   │   ├── /grpc
|   │   │   └── /users
|   │   │       └── service.go
|   │   │       └── service_test.go
|   │   │       └── mapper.go
|
|
├── /internal
│   ├── /adapters
|   │   ├── /http
|   │   │   └── /controllers
|   │   │       └── users.controller.go
|   │   │   └── /routes
|   │   │       └── users.route.go
|   │   │   └── /middlewares
|   │   │       └── auth.middleware.go
|   │   │       └── role.middleware.go
|   │   ├── /grpc
|   │   │   └── /schemas
|   │   │       └── users.pb.go
|   │   │       └── users_grpc.pb.go
|   |   │
│   ├── /infrastructure
|   │   ├── /connections
|   │   │   └── database.connection.go
|   │   │   └── redis.connection.go
|   │   ├── /providers
|   │   │   └── email.provider.go
|   │   │   └── sms.provider.go
|   │   ├── /templates
|   │   │   └── email.template.go
|   │   │   └── sms.template.go
|   │   │
│   ├── /modules
|   │   ├── /http
|   │   │   └── users.module.go
|   │   ├── /grpc
|   │   │   └── users.module.go
|
|
├── /external
│   ├── /deployments
|   │   ├── /docker
|   │   │       └── /golang
|   │   │           └── Dockerfile
|   │   │       └── /redis
|   │   │           └── Dockerfile
|   │   │       └── /postgres
|   │   │           └── Dockerfile
|   │   ├── /kubernetes
|   │   │       └── /deployment.yaml
|   │   │       └── /service.yaml
|   │   │       └── /ingress.yaml
|   │   │       └── /configmap.yaml
|   │   ├── /terraform
|   │   │       └── /main.tf
|   │   │       └── /variables.tf
|   │   │       └── /outputs.tf
|   │   │       └── /providers.tf
|   │   ├── /cicd
|   │   │       └── /github
|   │   │           └── /workflow
|   │   │               └── ci.yaml
|   │   │
│   ├── /scripts
|   │   │       └── /build.sh
|   │   │       └── /run.sh
|   │   │       └── /test.sh
|   │   │       └── /deploy.sh
|   │   │       └── /rollback.sh
|   │   │
│   ├── /documentations
|   │   ├── /readmes
|   │   │       └── /api.md
|   │   │       └── /grpc.md
|   │   │       └── /database.md
|   │   │       └── /infrastructure.md
|   │   ├── /swaggers
|   │   │       └── /api.swagger.json
|   │   │       └── /grpc.swagger.json
|
|
├── /shared
│   ├── /constants
|   │   └── users.constant.go
│   ├── /dto
|   │   └── users.dto.go
│   ├── /output
|   │   └── users.output.go
│   ├── /helpers
|   │   └── api.helper.go
│   ├── /interfaces
|   │   └── users.interface.go
│   ├── /pkg
|   │   └── jwt.pkg.go
|
|
├── /usecases
│   ├── /http
|   │   └── users.usecase.go
|   ├── /grpc
|   │   └── users.usecase.go
|
├── .gitignore
├── .env.local
├── .env.example
├── README.md
├── README.md
├── docker-compose.yml
├── .dockerignore
├── Makefile
├── go.sum
└── go.mod
```