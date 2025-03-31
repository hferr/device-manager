# Device Manager

Implementation of a simple device manager API, built with GO. I've tried to follow an idiomatic structure based on the resource-oriented design.

## Dependencies

Here's the list of dependencies chosen to help build the service:

- [chi](https://github.com/go-chi/chi) for building the router.
- [postgres](https://www.postgresql.org/) as the database.
- [gorm](https://gorm.io/) as the database ORM and [goose](https://github.com/pressly/goose) for handling migrations.
- [validator.v10](https://github.com/go-playground/validator) to validate requests.
- [swaggo/swag](https://github.com/swaggo/swag) for generating the API documentation.
- [testcontainers-go](https://github.com/testcontainers/testcontainers-go) for the repository tests.

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go                # Application entry point
├── config/
│   └── config.go                  # Configuration management using environment variables
├── docs/                          # Generated Swagger documentation
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── api/
│   │   └── device/                # Device domain logic
│   │       ├── model.go           # Device data models and DTOs
│   │       ├── repository.go      # Database operations for devices
│   │       ├── repository_test.go # Tests for repository layer
│   │       ├── service.go         # Business logic for device operations
│   │       └── service_test.go    # Tests for service layer
│   ├── protocols/
│   │   └── httpjson/              # HTTP/JSON protocol implementation
│   │       ├── device_handler.go  # HTTP handlers for device endpoints
│   │       └── router.go          # Router setup and middleware
│   └── err/                       # Error and response types
├── migrations/                    # Database migrations
├── test/
|   ├──mock/                       # Mock implementation of the api interfaces
|   └──helper.go                   # Helper functions for tests
├── utils/
│   └── validator/                 # Input validation utilities
│       └── validator.go           # Custom validation logic
```

## Running the project

With `docker-compose` installed, either run:

```
$ make build-run
```

or

```
$ docker-compose up --build
```

to spin up and start the docker container for the API and the database. If succesfull, the server will start on:

```
http://localhost:8080/
```

and the swagger documentation UI can be found at:

```
http://localhost:8080/swagger/index.html#/
```

## Running tests

To run the tests in the project, use either:

```
$ make run-test
```

or

```
$go test ./... -race
```

commands.

## Endpoints

| Name          | Method | Route                  | Description                                |
| ------------- | ------ | ---------------------- | ------------------------------------------ |
| Healthcheck   | GET    | /health                | Check if the server is live                |
| List Devices  | GET    | /devices               | Lists all devices                          |
| Create Device | POST   | /devices               | Create a new device                        |
| Update Device | PATCH  | /devices/{id}          | Updates the device with the given ID       |
| Find By ID    | GET    | /devices/{id}          | Finds the device belonging to the given ID |
| Find by State | GET    | /devices/state/{state} | List all devices with the given State      |
| Find by Brand | GET    | /devices/brand/{brand} | List all devices with the given Brand      |
| Delete Device | DELETE | /devices/{id}          | Deletes the device with the given ID       |

## Notes

- I've decided to use the `testcontainers-go` package for the repository implementation tests so that I could test the behavior against a real database at the cost of the tests taking a little longer to run.
- Tests for the `handler` and `service` package dependencies were done by mock implementing the interfaces in the `mock` package.
- Both the `.env` file and the swagger generated files were checked into git for simplicity.
