# Project teleport

### A URL shortener written in Go, [Truso](https://turso.tech/) LibSQL db and [HTMX](https://htmx.org/). As this project is deployed on [Railway](https://railway.app), the initial request can take upto 3-4 seconds, have patience :)

---

**The hash value generated for shortened URLs are calculated as:**

> **Base62_Encoding**(number of nanoseconds elapsed since Jan 1, 1970 UTC % a predefined SALT value)

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
