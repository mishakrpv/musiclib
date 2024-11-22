# effective-mobile-music-library

> Test task

## Navigation

* Application entry point, **main** package, is [here](cmd/musiclib/musiclib.go)
* Routes are located in the [internal/router](internal/router/) directory as well as handlers
* Music Info client is in the [pkg/infra/musicinfo](pkg/infra/musicinfo/)
* Database connection and startup migration implementations are in the [pkg/infra/db](pkg/infra/db/)
* Request and response models for commands and queries can be found in the [internal/endpoint/command](internal/endpoint/command/) and [internal/endpoint/query](internal/endpoint/query/) respectively
* The [mock of musicinfo](musicinfo/) contains implementation of a service for tests

## Running the service

Run the application from your terminal:

```bash
make run
```

> [!IMPORTANT]
> Before starting the service, ensure that [.env file](.env) is configured properly according to your environment.
> Make sure **MUSIC_INFO_URL** is specified.

## Database schema

There is initial [script](migrations/000001_initial.up.sql) in **migrations** directory.

## Additional information

Swagger UI page located at `http://localhost:8080/swagger/index.html#`.

Go version:

* Go 1.23.1

Go packages used:

* **[Gin](https://github.com/gin-gonic/gin)** v1.10.0
* **[zerolog](https://github.com/rs/zerolog)** v1.33.0
* **[swag](https://github.com/swaggo/swag)** v1.16.3
* **[gorm](https://github.com/go-gorm/gorm)** v1.25.12

See [go.mod file](go.mod) for more information about packages.
