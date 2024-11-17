# Todos service

## Overview

Service written in Go and dockerized. This is part of my learning path to improve my infra and backend skills. The service so far exposes some API endpoints with fake data.

## Installation

The repository uses docker to run the Go service. The process is manual right now, so you have to build the image and run it with docker-compose:

```
docker compose up
```

If you do changes on the Go service, you will need to manually re-build the image. Right now, there's nothing implemented to
speed up this process, but it's on the backlog. You will need to run:

```
docker build -no-cache -t miquelarranz/todos-service .
```

The `-no-cache` flag is needed, otherwise Docker will get the Go build from the cache.

And the API can be accessed through `localhost:1234`.

### Endpoints

| Method | Path       | Description           | Body example                                                        |
|--------|------------|-----------------------|---------------------------------------------------------------------|
| GET    | /todos     | returns all the todos | -                                                                   |
|        | /todos/:id | returns a todo by id  | -                                                                   |
| POST   | /todos     | creates a new todo    | { "id": "your-id", "title": "title", "description": "description" } |

## Contributing

This repository uses the [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) format.

### Migrations

The repository uses Goose to manage DB migrations. To run the DB migrations you need to have the db docker 
container running and run:

```sh
source .env && goose up
```

Temporary, the environment variables needed for Goose are defined in the `.env` file and the process is manual. I will
change it when working on productizing this repository.

## Next steps

I will list here the progressive improvements I'm doing to the repository and future steps to explore.

* ~~Migrations for the DB.~~ ✅
* Expand the service with a Postgres DB so it returns real data. 🏗️
* Live-reload for the Go service with Docker.
* Add tests
* Configure GHA to build and push the image to a registry.
* Add a users package/module with authentication, so we can have todos related to users.
* Add a Kubernetes layer to the project.
* ...