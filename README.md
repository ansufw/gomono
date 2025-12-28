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

### Start

- run `air`