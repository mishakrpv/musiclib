# effective-mobile-music-library

> Test task

## Navigation

* **Entry Point**: The application's main package is located [here](cmd/musiclib/musiclib.go).
* **Routing and Handlers**: Routes and their associated handlers can be found in the [internal/router](internal/router/) directory.
* **Music Info Client**: The Music Info client is implemented in [pkg/infra/musicinfo](pkg/infra/musicinfo/).
* **Database Connection & Migration**: Database connection logic and startup migrations are located in [pkg/infra/db](pkg/infra/db/).
* **CQRS**: Request and response models for commands and queries can be found in the [internal/endpoint/command](internal/endpoint/command/) and [internal/endpoint/query](internal/endpoint/query/) directories respectively.
* **Logging**: Logger configuration can be found in the [logger](pkg/logger/) package. To override configs in the .env file, take a look at the [config](pkg/config/) package.

## Running the Service

To run the application, execute the following command in your terminal:

```bash
make run
```

> [!IMPORTANT]
> Before starting the service, ensure that [.env file](.env) is configured correctly according to your environment.
> Make sure **MUSIC_INFO_URL** is specified.

## Database Schema

The initial database schema migration can be found in the [script](migrations/000001_initial.up.sql) file.

## Additional Information

Swagger UI page located at `http://localhost:8080/swagger/index.html#`.

Go version:

* Go 1.23.1

Go packages used:

* **[Gin](https://github.com/gin-gonic/gin)** v1.10.0
* **[zerolog](https://github.com/rs/zerolog)** v1.33.0
* **[swag](https://github.com/swaggo/swag)** v1.16.3
* **[gorm](https://github.com/go-gorm/gorm)** v1.25.12

For more details on dependencies, check the [go.mod file](go.mod).
