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

| Variable               | Required | Default | Example                                                | Description                                                                                         |
|------------------------|----------|---------|--------------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| `DATABASE_URL`         | no | — | `postgres://user:pass@db:5432/app?sslmode=disable`     | PostgreSQL connection string in libpq format. If provided, it overrides all `DATABASE_*` variables. |
| `DATABASE_HOST`        | yes* | `localhost` | `db`                                                   | PostgreSQL host. Used only if `DATABASE_URL` is not set.                                            |
| `DATABASE_PORT`        | yes* | `5432` | `5432`                                                 | PostgreSQL port.                                                                                    |
| `DATABASE_USERNAME`    | yes* | — | `app`                                                  | Database user.                                                                                      |
| `DATABASE_PASSWORD`    | yes* | — | `s3cr3t`                                               | Database user password.                                                                             |
| `DATABASE_NAME`        | yes* | — | `app`                                                  | Database name.                                                                                      |
| `DATABASE_SSLMODE`     | no | `disable` | `disable` \| `require` \| `verify-ca` \| `verify-full` | SSL mode for the PostgreSQL connection.                                                             |
| `REDIS_URL`            | **yes** | — | `redis://redis:6379/0`                                 | Redis connection URL.                                                                               |
| `SERVER_PORT`          | no | `8080` | `8080`                                                 | HTTP server port.                                                                                   |
| `REDIS_URL`            | no       | —           | `redis://:pass@redis:6379/0`                           | Full Redis connection URL. Has priority if provided.                                                |
| `REDIS_SCHEME`         | no       | `redis`     | `rediss`                                               | Connection scheme: `redis` (TCP) or `rediss` (TLS).                                                 |
| `REDIS_HOST`           | yes*     | `localhost` | `redis`                                                | Redis host. Used if `REDIS_URL` is not set.                                                         |
| `REDIS_PORT`           | yes*     | `6379`      | `6380`                                                 | Redis port.                                                                                         |
| `REDIS_USERNAME`       | no       | —           | `appuser`                                              | ACL username (Redis ≥ 6).                                                                           |
| `REDIS_PASSWORD`       | no       | —           | `s3cr3t`                                               | Redis password.                                                                                     |
| `REDIS_CERT_PATH`      | no       | —           | ``                                                     | Redis cert path.                                                                                    |
| `REDIS_KEY_PATH`       | no       | —           | ``                                                     | Redis key path.                                                                                     |
| `REDIS_DB`             | no       | `0`         | `1`                                                    | Redis database index.                                                                               |
| `REDIS_DIAL_TIMEOUT`   | no       | —           | `5s`                                                   | Dial timeout (e.g. `5s`, `1m`).                                                                     |
| `REDIS_READ_TIMEOUT`   | no       | —           | `2s`                                                   | Read timeout.                                                                                       |
| `REDIS_WRITE_TIMEOUT`  | no       | —           | `2s`                                                   | Write timeout.                                                                                      |
| `REDIS_POOL_SIZE`      | no       | —           | `20`                                                   | Maximum number of connections in the pool.                                                          |
| `REDIS_MIN_IDLE_CONNS` | no       | —           | `5`                                                    | Minimum number of idle connections.                                                                 |
| `REDIS_MAX_RETRIES`    | no       | —           | `3`                                                    | Maximum number of retries before giving up.                                                         |

\* Required only if **`DATABASE_URL`** is not set.

### Priority Rules
- If `REDIS_URL` is set → it is used directly.
- Otherwise → a URL is built in the form:

### Priority Rules
- If `DATABASE_URL` **is set**, the application uses it directly.
- If `DATABASE_URL` **is not set**, the DSN is constructed from the `DATABASE_*` variables.

- If `REDIS_URL` **is set**, the application uses it directly.
- If `REDIS_URL` **is not set**, the DSN is constructed from the `REDIS_*` variables.

## Examples

### 1. Minimal setup with `DATABASE_URL`
```env
DATABASE_URL=postgres://app:s3cr3t@db:5432/app?sslmode=disable
REDIS_URL=redis://redis:6379
SERVER_PORT=8080
```
### 2. Setup with 'DATABASE_* REDIS_*'
```env
DATABASE_HOST=db
DATABASE_PORT=5432
DATABASE_USERNAME=app
DATABASE_PASSWORD=s3cr3t
DATABASE_NAME=app
DATABASE_SSLMODE=disable

REDIS_SCHEME=redis
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=s3cr3t
REDIS_DB=0
SERVER_PORT=8080
```