### Generate swagger api
```shell
 swag init -g main.go --dir ./cmd/app,./internal/http,./internal/dto,./internal/repositories/resources --parseDependency
```

# Configuration via Environment Variables

The service supports two ways of configuring the database connection:

1. **`DATABASE_URL`** — a full connection string (has priority).
2. A set of individual **`DATABASE_*`** variables from which the DSN will be built.

Additionally, Redis and server port parameters are required.

## Environment Variables

| Variable | Required | Default | Example | Description |
|----------|----------|---------|---------|-------------|
| `DATABASE_URL` | no | — | `postgres://user:pass@db:5432/app?sslmode=disable` | PostgreSQL connection string in libpq format. If provided, it overrides all `DATABASE_*` variables. |
| `DATABASE_HOST` | yes* | `localhost` | `db` | PostgreSQL host. Used only if `DATABASE_URL` is not set. |
| `DATABASE_PORT` | yes* | `5432` | `5432` | PostgreSQL port. |
| `DATABASE_USERNAME` | yes* | — | `app` | Database user. |
| `DATABASE_PASSWORD` | yes* | — | `s3cr3t` | Database user password. |
| `DATABASE_NAME` | yes* | — | `app` | Database name. |
| `DATABASE_SSLMODE` | no | `disable` | `disable` \| `require` \| `verify-ca` \| `verify-full` | SSL mode for the PostgreSQL connection. |
| `REDIS_URL` | **yes** | — | `redis://redis:6379/0` | Redis connection URL. |
| `SERVER_PORT` | no | `8080` | `8080` | HTTP server port. |

\* Required only if **`DATABASE_URL`** is not set.

### Priority Rules
- If `DATABASE_URL` **is set**, the application uses it directly.
- If `DATABASE_URL` **is not set**, the DSN is constructed from the `DATABASE_*` variables.

## Examples

### 1. Minimal setup with `DATABASE_URL`
```env
DATABASE_URL=postgres://app:s3cr3t@db:5432/app?sslmode=disable
REDIS_URL=redis://redis:6379/0
SERVER_PORT=8080
```
### 2. Setup with 'DATABASE_*'
```env
DATABASE_HOST=db
DATABASE_PORT=5432
DATABASE_USERNAME=app
DATABASE_PASSWORD=s3cr3t
DATABASE_NAME=app
DATABASE_SSLMODE=disable

REDIS_URL=redis://redis:6379/0
SERVER_PORT=8080
```