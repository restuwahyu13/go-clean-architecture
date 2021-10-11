# Golang Clean Architecture

Berikut ini adalah folder structure pattern yang biasa saya gunakan, walaupun tidak semua nya saya terapkan, dikarenakan project yang saya kerjakan
hanyalah project - project kecil yang tidak terlalu besar, jadi jika anda tertarik dengan pattern yang saya buat ini anda bisa menggunakannya jika menurut anda baik.

## Structure Folder Pattern

```
├── tests
│   └── test.auth_test.go
│   └── test.student_test.go
└── docker
│   └── swagger
│   │     └── Dockerfile
│   │     └── openapi.yml
│   └── mysql
│   │     └── Dockerfile
│   │     └── mysql.cnf
│   └── golang
│   │     └── Dockerfile
├── handlers
│   └── auth
│   │     └── handler.login.go
│   │     └── handler.register.go
│   └── student
│   │     └── handler.create.go
│   │     └── handler.create.go
└── repository
│   └── auth
│   │     └── repository.login.go
│   │     └── repository.register.go
│   └── student
│   │     └── repository.create.go
│   │     └── repository.create.go
└── services
│   └── auth
│   │     └── services.login.go
│   │     └── services.register.go
│   └── student
│   │     └── services.create.go
│   │     └── services.create.go
└── helpers
│   └── helpers.apiResponse.go
│   └── helpers.randomString.go.go
└── middlewares
│   └── middleware.auth.go
│   └── middleware.role.go.go
└── models
│   └── model.auth.go
│   └── model.student.go.go
└── routes
│   └── route.auth.go
│   └── route.student.go
└── schemas
│   └── schema.auth.go
│   └── schema.student.go.go
└── templates
│   └── template.register.html
│   └── template.activation.html
└── pkg
│   └── pkg.jwt.go
│   └── pkg.bcrypt.go
│   └── pkg.cron.go
└── scripts
│   └── gcpRunner.sh
│   └── awsRunner.sh
└── configs
│   └── openapi.yml
│   └── serverless.yml
└── cmd
│   └── cmd.pgMigration.go
│   └── cmd.pgSeeds.go
└── crons
│   └── cron.autoDeleteLogs.go
│   └── cron.emailBlast.go
└── databases
│   └── migrations
│   │     └── migration.auth.go
│   │     └── migration.student.go
│   └── seeds
│   │     └── seed.auth.go
│   │     └── seed.student.go
│   └── sql
│   │     └── sql.auth.sql
│   │     └── sql.student.sql
```

## Table Folder Status

| Folder Name     | Folder Status |
| --------------- | ------------- |
| **tests**       | *optional*    |
| **docker**      | *optional*    |
| **handlers**    | *required*    |
| **repositorys** | *required*    |
| **services**    | *required*    |
| **helpers**     | *optional*    |
| **middlewares** | *optional*    |
| **models**      | *required*    |
| **routes**      | *required*    |
| **schemas**     | *required*    |
| **templates**   | *optional*    |
| **pkg**         | *optional*    |
| **scripts**     | *optional*    |
| **configs**     | *optional*    |
| **cmd**         | *optional*    |
| **crons**       | *optional*    |
| **databases**   | *optional*    |


## Command

- ### Application Lifecycle

  - Install node modules

  ```sh
  $ go get . || go mod || make goinstall
  ```

  - Build application

  ```sh
  $ go build -o main || make goprod
  ```

  - Start application in development

  ```sh
  $ go run main.go | make godev
  ```

  - Test application

  ```sh
  $ go test main.go main_test.go || make gotest
  ```

* ### Docker Lifecycle

  - Build container

  ```sh
  $ docker-compose build | make dcb
  ```

  - Run container with flags

  ```sh
  $ docker-compose up -d --<flags name> | make dcu f=<flags name>
  ```

  - Run container build with flags

  ```sh
  $ docker-compose up -d --build --<flags name> | make dcubf f=<flags name>
  ```

  - Run container

  ```sh
  $ docker-compose up -d --build | make dcu
  ```

  - Stop container

  ```sh
  $ docker-compose down | make dcd
  ```
