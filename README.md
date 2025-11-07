#

## Description

Simple slug service

## Install

```
git clone github.com/lilbonekit/slug-svc
cd
go build main.go
export KV_VIPER_FILE=./config.yaml
./main migrate up
./main run service
```

## Documentation

We do use openapi:json standard for API. We use swagger for documenting our API.

To open online documentation, go to [swagger editor](http://localhost:8080/swagger-editor/) here is how you can start it

```
  cd docs
  npm install
  npm start
```

To build documentation use `npm run build` command,
that will create open-api documentation in `web_deploy` folder.

To generate resources for Go models run `./generate.sh` script in root folder.
use `./generate.sh --help` to see all available options.

Note: if you are using Gitlab for building project `docs/spec/paths` folder must not be
empty, otherwise only `Build and Publish` job will be passed.

## Makefile Commands

Use following commands to start Postgres container and manage database:

- `make db-reset` - stop and remove Postgres container and volume
- `make db-up` - start Postgres container
- `make db-check` - check connection to Postgres
- `make db-logs` - view Postgres container logs
- `make db-psql` - open interactive psql session
- `make db-down` - stop and remove Postgres container (without removing volume)
- `make migrate-up` - apply database migrations
- `make migrate-down` - rollback database migrations
- `make run` - start the service

## Running from docker

Make sure that docker installed.

use `docker run ` with `-p 8080:80` to expose port 80 to 8080

```
docker build -t github.com/lilbonekit/slug-svc .
docker run -e KV_VIPER_FILE=/config.yaml github.com/lilbonekit/slug-svc
```

## Running from Source

- Set up environment value with config file path `KV_VIPER_FILE=./config.yaml`
- Provide valid config file
- Launch the service with `migrate up` command to create database schema
- Launch the service with `run service` command

### Database

For services, we do use **_PostgresSQL_** database.
You can [install it locally](https://www.postgresql.org/download/) or use [docker image](https://hub.docker.com/_/postgres/).

### Third-party services

## Contact

Responsible Slug service
The primary contact for this project is t.me/lilbonekit
