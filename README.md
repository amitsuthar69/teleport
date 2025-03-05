# Project teleport

### A URL shortener written in Go, [Truso](https://turso.tech/) LibSQL db and [HTMX](https://htmx.org/).

---

- **The hash value generated for shortened URLs are calculated as:**
> Base62(number of nanoseconds elapsed since Jan 1, 1970 UTC % a predefined SALT value).

- **In-memory cache for faster page loads:**
> The lastest version of Teleport uses [LRUCache](https://github.com/amitsuthar69/libs/blob/master/lrucache/lrucache.go) to cache most recently used URLs.

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
