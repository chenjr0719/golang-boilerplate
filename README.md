# golang-boilerplate

Golang Boilerplate Template for RESTful API and Worker

- [golang-boilerplate](#golang-boilerplate)
  - [Usage](#usage)
  - [Configuration](#configuration)
    - [Global configs](#global-configs)
    - [Worker configs](#worker-configs)
  - [Worker Tasks](#worker-tasks)
  - [Project Layout](#project-layout)
  - [Dependencies](#dependencies)
  - [TODO](#todo)

## Usage

You can run the whole project with this command:
```shell
make docker-compose-up
```

The API server will be availableready at http://localhost:8080. And the swagger will be  at http://localhost:8080/docs.

If you want to develop with this project, you can run the prject with dev mode:
```shell
make DEV=1 docker-compose-up
```

And run the API server or worker by:
```shell
go run ./cmd/apiserver
go run ./cmd/worker
```

For more usage, please check [Makefile](Makefile) and [hack](hack/).

## Configuration

You can override the configuration files under [configs/](configs/) or provide `.env` file when you are developing.

### Global configs
| Key          | Value                                                     | Example                             |
|--------------|-----------------------------------------------------------|-------------------------------------|
| `MODE`         | `debug` \| `release`                                          | `debug`                               |
| `LOG_LEVEL`    | `panic` \| `fatal` \| `error` \| `warn` \| `info` \| `debug` \| `trace` | `debug`                               |
| `DATABASE_URI` | `SCHEMA://USERNAME:PASSWORD@HOSTNAME:PORT/DATABASE_NAME`    | `sqlite://file::memory:?cache=shared` |

### Worker configs

To config the worker and worker client, please check the usage of machinery:
https://github.com/RichardKnop/machinery#configuration


## Worker Tasks

There are some example builit-in tasks that can be used to create new job:

* sleep: The example `args` is `{"seconds": 10}`
* remoteHTTPCall: The example `args` is `{"url": "https://postman-echo.com/post", "body": {"message": "Hello World"}}`
* fibonacci: The example `args` is `{"target": 5}`
* error: The example `args` is `{"message": "Error"}`

## Project Layout

This project trying to follow the project layout from this reference:
https://github.com/golang-standards/project-layout

The main implementation are located under [pkg/](/pkg/) folder and the entrypoint of each application ar under [cmd/](cmd/).

## Dependencies

* [gorm](https://github.com/go-gorm/gorm): The ORM framework
* [gin](https://github.com/gin-gonic/gin): The web framework
* [machinery](https://github.com/RichardKnop/machinery): The worker freamwwork
* [zerolog](https://github.com/rs/zerolog): The log replacement
* [viper](https://github.com/spf13/viper): The configuration management

## TODO

- [ ] Unit tests
- [ ] Tracing by [opentelemetry](https://github.com/open-telemetry/opentelemetry-go)
- [ ] Unified logger for applications
- [ ] Develop deployment with [kind](https://kind.sigs.k8s.io/) and [Skaffold](https://skaffold.dev/)
- [ ] Demo site on [GCP cloud run](https://cloud.google.com/run)
