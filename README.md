# effective-mobile-music-library

> Test task

## Navigation

* Application entry point, **main** package, is [here](cmd/musiclib/main.go)
* Routes and server configuration are located in the [internal/server](internal/server/) directory as well as handlers
* Music Info client is in the [internal/infrastructure/service](internal/infrastructure/service/)
* Database connection and startup migration implementations are in the [internal/infrastructure/data/gorm](internal/infrastructure/data/gorm/)
* Request and response models for commands and queries can be found in the [internal/endpoint/command](internal/endpoint/command/) and [internal/endpoint/query](internal/endpoint/query/) respectively

## Running the service

Run the application from your terminal:

```bash
make run
```

> Before starting the service, ensure that [.env file](.env) is configured properly according to your environment.

## Database schema

There is initial [script](migrations/000001_initial.up.sql) in **migrations** directory.

## Additional information

Swagger UI page located at `http://localhost:8080/swagger/index.html#`.

Go version:

* Go 1.23.1

Go packages used:

* **[Gin](https://github.com/gin-gonic/gin)** v1.10.0
* **[Uber zap](https://github.com/uber-go/zap)** v1.27.0
* **[swag](https://github.com/swaggo/swag)** v1.16.3
* **[gorm](https://github.com/go-gorm/gorm)** v1.25.12

See [go.mod file](go.mod) for more information about packages.
