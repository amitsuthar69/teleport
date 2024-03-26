# Project teleport

A URL shortener written in Go, [Truso](https://turso.tech/) LibSQL db and [HTMX](https://htmx.org/).

## Getting Started

Install Dependencies

```bash
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/cosmtrek/air@latest
```

Build Templates

```bash
templ generate
```

Hot Reload

```bash
air
```

Build Application

```bash
go build -o main cmd/api/main.go
```

.env

```env
PORT=3032
APP_ENV=local
SALT=
DB_URL=libsql://dbname-username.turso.io?authToken=
```

---

For Turso auth token and CLI, visit [Turso Docs](https://docs.turso.tech/introduction).
