# Todos service

## Overview

Service written in Go and dockerized. This is part of my learning path to improve my infra and backend skills. The service so far exposes some API endpoints with fake data.

## Installation

The repository uses docker to run the Go service. The process is manual right now, so you have to build the image and run it with docker-compose:

```
docker compose -f docker-compose.yml up
```

Or you can build and run the container manually:

```
docker build -t <your-namespace>/todos-service .
docker run -d -p 1234:8080 --name=todos-service <your-namespace>/todos-service
```

And the API can be accessed through `localhost:1234`.

### Endpoints

| Method | Path       | Description           | Body example                                                        |
|--------|------------|-----------------------|---------------------------------------------------------------------|
| GET    | /todos     | returns all the todos | -                                                                   |
|        | /todos/:id | returns a todo by id  | -                                                                   |
| POST   | /todos     | creates a new todo    | { "id": "your-id", "title": "title", "description": "description" } |

## Contributing

The repository doesn't accept contributions so far.

It uses the [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) format.

## Next steps

* Expand the service with a Postgres DB so it returns real data.
* Configure GHA to build and push the image to a registry.
* Add a users package/module with authentication, so we can have todos related to users.
* Add a Kubernetes layer to the project.
* ...