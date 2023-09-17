# karriarum-ctf

![CTF-Leaderboard](images/ctfleaderboard.png)

Backend API for the App "Karriarum CTF" - Intended to be used together with [Karriarum-CTF-Frontend](https://github.com/s0undy/karriarum-ctf-frontend)

[![Go Report Card](https://goreportcard.com/badge/github.com/s0undy/karriarum-ctf)](https://goreportcard.com/report/github.com/s0undy/karriarum-ctf)

## Introduction

Small project in GO that starts a http-server on the desiered port and exposes a simple API written with [gofiber](https://github.com/gofiber/fiber) that interacts with a Postgres database via [GORM](https://gorm.io/).

It exposes two API endpoints:

* /api/v1/score

Accepts a http POST with a body in JSON format:

```json
{
    "name": "David Rex",
    "flags": 202,
    "email": "david.rex@gmail.com",
    "mobilenumber": "0394827362"
}
```

* /api/v1/list

Accepts a http GET and returns a list of all records in the database.

```json
[
    {
        "ID": 1,
        "CreatedAt": "2023-09-17T12:09:00.861583Z",
        "UpdatedAt": "2023-09-17T12:09:00.861583Z",
        "DeletedAt": null,
        "Name": "Carolus Rex",
        "Flags": 1718,
        "Email": "carolus.rex@kung.se",
        "MobileNumber": "0192388425"
    }
]
```

## Usage

Important to remember that this is ONLY the backend porportion of "Karriarium CTF". It renders no frontend and only exposes the API. It should always be used together with [Karriarum-CTF-Frontend](https://github.com/s0undy/karriarum-ctf-frontend)

### Docker Compose

See examples in [examples/compose](https://github.com/s0undy/karriarum-ctf-backend/tree/main/examples/compose)

## Configuration

| Environment Variable | Description                               | Default | Required |
|----------------------|-------------------------------------------|---------|----------|
| `DB_HOST`            | Hostname/IP of the Postgres database host |         | ✅        |
| `DB_USER`            | Postgres DB User                          |         | ✅        |
| `DB_PASSWORD`        | Postgres DB Password                      |         | ✅        |
| `DB_NAME`            | Postgres DB Name                          |         | ✅        |
| `DB_PORT`            | Postgres DB Port                          |         | ✅        |
| `DB_SSLMODE`         | DB SSL (enable/disable)                   |         | ✅        |
| `APP_PORT`           | Port that the API should listen to        |         | ✅        |
