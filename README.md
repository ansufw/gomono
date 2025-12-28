## GOMONO

Gomono is boilerplate for fullstack web application build with Golang Fiber and Templ

### Preparation

- Golang
- download tailwind, `make tw`
- templ (see `https://templ.guide/quick-start/installation`)
- dbmate (see `https://github.com/amacneil/dbmate`)
- sqlc (see `https://docs.sqlc.dev/`)
- postgresql (using the latest version, `https://formulae.brew.sh/formula/postgresql@18`)
- air, `go install github.com/air-verse/air@latest`
- httpyac for testing, https://httpyac.github.io/
- [swag](https://github.com/gofiber/swagger), `swag init -g ./cmd/main.go -d ./docs/api`

### Getting Started

#### Database

for local development, follow these steps

- in the folder `dev`, run `docker compose up`
- open gui via `http://localhost:8080/` in the web browser
- paste the DATABASE_URL value from `.env` in the `Host Name` field but change the hostname value itself with `db`
