# Cubicasa-test apply clean-architecture

<img src="https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg" width="521" alt="Alt text" title="Optional title">

# Introduction

- [x] Implement a Create for each _hub, team, and user_.
- [x] Implement a Search which will return _team and hub information_.
- [x] Implement a Join for user into team, and team into hub (for simplicity: one user belongs to one team, one team
  belongs to one hub).
- [x] Write the test cases
- [x] Provide a SQL script which creates tables needed for the API. -> db/migrations/postgres
- [x] Good to use `docker/docker-compose` for local development setup(not mandatory)
- [x] Good to provide the solution with security concern -> use basic auth

# Prerequisites

Before you continue, ensure you meet the following requirements:

- You have installed docker

# Quick Start

```bash
cd cubicasa-test
```

```bash
docker-compose up
```

# Migration

```bash
docker-compose exec app bash -c "cd /go/src/backend/ && make migrate-up"
```

# Document

```
http://localhost:8000/swagger/index.html
```

# Tests

```bash
docker-compose exec app bash -c "cd /go/src/backend/ && go test ./app/usecase && go test ./app/external/persistence/pgsql"
```
