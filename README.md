# Project teleport

A URL shortener written in Go, [Truso](https://turso.tech/) LibSQL db and [HTMX](https://htmx.org/). As this project is deployed on [Railway](https://railway.app), the initial request can take upto 4-5 seconds, have patience :)

> The hashed / shortened URLs generated are a product of **Base62_Encoding**(originalURL) and a predefined **SALT** value and are stored in a SQLite Database.

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

Build Application

```bash
go build -o main cmd/api/main.go
```

Hot Reload

```bash
air
```

.env

```env
PORT=42069
APP_ENV=local
SALT=
DB_URL=libsql://dbname-username.turso.io?authToken=
```

---

For Turso auth token and CLI, visit [Turso Docs](https://docs.turso.tech/introduction).
