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

| Variable               | Required | Default | Example                                                 | Description                                                                                         |
|------------------------|----------|---------|---------------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| `DATABASE_URL`         | no | — | `postgres://user:pass@db:5432/app?sslmode=disable`      | PostgreSQL connection string in libpq format. If provided, it overrides all `DATABASE_*` variables. |
| `DATABASE_HOST`        | yes* | `localhost` | `db`                                                    | PostgreSQL host. Used only if `DATABASE_URL` is not set.                                            |
| `DATABASE_PORT`        | yes* | `5432` | `5432`                                                  | PostgreSQL port.                                                                                    |
| `DATABASE_USERNAME`    | yes* | — | `app`                                                   | Database user.                                                                                      |
| `DATABASE_PASSWORD`    | yes* | — | `s3cr3t`                                                | Database user password.                                                                             |
| `DATABASE_NAME`        | yes* | — | `app`                                                   | Database name.                                                                                      |
| `DATABASE_SSLMODE`     | no | `disable` | `disable` \| `require` \| `verify-ca` \| `verify-full`  | SSL mode for the PostgreSQL connection.                                                             |
| `REDIS_URL`            | **yes** | — | `redis://redis:6379/0`                                  | Redis connection URL.                                                                               |
| `SERVER_PORT`          | no | `8080` | `8080`                                                  | HTTP server port.                                                                                   |
| `REDIS_URL`            | no       | —          | `redis://:pass@redis:6379/0`                            | Full Redis connection URL. Has priority if provided.                                                |
| `REDIS_SCHEME`         | no       | `redis`    | `rediss`                                                | Connection scheme: `redis` (TCP) or `rediss` (TLS).                                                 |
| `REDIS_HOST`           | yes*     | `localhost` | `redis`                                                 | Redis host. Used if `REDIS_URL` is not set.                                                         |
| `REDIS_PORT`           | yes*     | `6379`     | `6380`                                                  | Redis port.                                                                                         |
| `REDIS_USERNAME`       | no       | —          | `appuser`                                               | ACL username (Redis ≥ 6).                                                                           |
| `REDIS_PASSWORD`       | no       | —          | `s3cr3t`                                                | Redis password.                                                                                     |
| `REDIS_SKIP_TLS`       | no       | ''           | `true`                                                  | Redis skip tls.                                                                                     |
| `REDIS_DB`             | no       | `0`        | `1`                                                     | Redis database index.                                                                               |
| `REDIS_DIAL_TIMEOUT`   | no       | —          | `5s`                                                    | Dial timeout (e.g. `5s`, `1m`).                                                                     |
| `REDIS_READ_TIMEOUT`   | no       | —          | `2s`                                                    | Read timeout.                                                                                       |
| `REDIS_WRITE_TIMEOUT`  | no       | —          | `2s`                                                    | Write timeout.                                                                                      |
| `REDIS_POOL_SIZE`      | no       | —          | `20`                                                    | Maximum number of connections in the pool.                                                          |
| `REDIS_MIN_IDLE_CONNS` | no       | —          | `5`                                                     | Minimum number of idle connections.                                                                 |
| `REDIS_MAX_RETRIES`    | no       | —          | `3`                                                     | Maximum number of retries before giving up.                                                         |

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
# Authorization

Authorization is **disabled by default**: without `AUTH_ENABLED=true` the API works without tokens exactly as before.

When enabled, every `/api/v1/*` request must carry `Authorization: Bearer <JWT>`.
Tokens are issued by an external IdP; the service only verifies them (signature, `iss`, `aud`, `exp`, `nbf`).
`/health/*` and `/swagger/*` stay public.

| Variable                     | Required | Default       | Description                                                                                           |
|------------------------------|----------|---------------|-------------------------------------------------------------------------------------------------------|
| `AUTH_ENABLED`               | no       | `false`       | Enables JWT authentication and authorization.                                                         |
| `AUTH_OIDC_ISSUER_URL`       | one of*  | —             | OIDC issuer, JWKS URL is taken from `/.well-known/openid-configuration`.                              |
| `AUTH_JWKS_URL`              | one of*  | —             | JWKS URL. Keys are kept in memory, refreshed in background and on unknown `kid`.                      |
| `AUTH_JWT_PUBLIC_KEY_FILE`   | one of*  | —             | PEM file with a public key or certificate.                                                            |
| `AUTH_JWKS_REFRESH_INTERVAL` | no       | `1h`          | JWKS background refresh interval.                                                                     |
| `AUTH_ISSUER`                | yes      | `AUTH_OIDC_ISSUER_URL` | Expected `iss`.                                                                              |
| `AUTH_AUDIENCE`              | no       | —             | Expected `aud`. Strongly recommended.                                                                 |
| `AUTH_ALGORITHMS`            | no       | `RS256,ES256` | Allowed algorithms. Only asymmetric: `RS*`, `PS*`, `ES*`, `EdDSA`.                                    |
| `AUTH_CLOCK_SKEW`            | no       | `30s`         | Leeway for `exp` / `nbf`.                                                                             |
| `AUTH_SUBJECT_CLAIM`         | no       | `sub`         | Claim with the caller identity (e.g. `client_id`, `azp`). Dotted path for nested claims.              |
| `AUTH_GROUPS_CLAIM`          | no       | `groups`      | Claim with caller groups (array of strings).                                                          |
| `AUTH_PERMISSIONS_CLAIM`     | no       | `permissions` | Claim with inline rules, e.g. `resource_access.state-manager.permissions`.                            |
| `AUTH_PERMISSIONS_SOURCE`    | no       | `auto`        | `auto`: claim if present, otherwise DB. `claims`: claim only, DB is never queried. `db`: DB only.     |

\* Exactly one key source must be set.

## Rules

A caller has a list of rules. Rules only allow (there is no deny), access is granted if **any** rule matches.
A missing field, an empty string or `*` means "any value".

```json
{
  "verbs": ["get", "list", "update_status"],
  "resource_group": "*",
  "namespace": "team-a",
  "kind": "Pod",
  "name": "*",
  "shard_id": "shard-1"
}
```

Verbs: `get`, `list`, `create`, `update` (spec), `update_status`, `delete`, `*`.

| Operation                  | Checked against                                                                 |
|----------------------------|---------------------------------------------------------------------------------|
| `create`                   | path + `name` and `shard_id` from body                                          |
| `get`, `delete`            | the stored resource (including its `shard_id`)                                  |
| `update`, `update_status`  | the stored resource **and** the new `shard_id` from body (no moving to a foreign shard) |
| `list`                     | the request filter must be fully covered by a single rule (see below)           |

### List

The list result is never filtered by permissions. Like in Kubernetes RBAC, the request filter
must fit into one rule: every field restricted by the rule must be present in the query with the same value.
Otherwise the request fails with `403`.

Example: rules `{"verbs":["get","list"],"shard_id":"5"}` and `{"verbs":["get","list"],"kind":"user"}`.

| Request                         | Result |
|---------------------------------|--------|
| `?shard_id=5`                   | ✅     |
| `?kind=user`                    | ✅     |
| `?kind=user&namespace=x`        | ✅     |
| `?shard_id=5&kind=order`        | ✅     |
| `?kind=order`                   | ❌ 403 |
| no filter                       | ❌ 403 |

## Rules in the token

With `AUTH_PERMISSIONS_SOURCE=auto` or `claims` the rules can be put into the token, no database query is made:

```json
{
  "iss": "https://idp.example.com",
  "aud": "state-manager",
  "sub": "billing-controller",
  "exp": 1790000000,
  "permissions": [
    {"verbs": ["get", "list", "update_status"], "shard_id": "5"},
    {"verbs": ["get", "list"], "kind": "user"}
  ]
}
```

## Rules in the database (service accounts)

With `AUTH_PERMISSIONS_SOURCE=auto` (token without the permissions claim) or `db`, rules are loaded
from roles bound to the token subject or any of its groups. One indexed query per request.

```sql
INSERT INTO auth_roles (name, description) VALUES ('shard-5-controller', 'Controller of shard 5');

INSERT INTO auth_role_rules (role_id, verbs, shard_id)
SELECT id, '{get,list,update_status}', '5' FROM auth_roles WHERE name = 'shard-5-controller';

-- bind to a token subject (sub / AUTH_SUBJECT_CLAIM)
INSERT INTO auth_role_bindings (role_id, subject_kind, subject_value)
SELECT id, 'subject', 'billing-controller' FROM auth_roles WHERE name = 'shard-5-controller';

-- or to a group from AUTH_GROUPS_CLAIM
INSERT INTO auth_role_bindings (role_id, subject_kind, subject_value)
SELECT id, 'group', 'controllers' FROM auth_roles WHERE name = 'shard-5-controller';

-- revoke without deleting
UPDATE auth_role_bindings SET disabled = true WHERE subject_value = 'billing-controller';
```

## Responses

- `401` — missing, expired or invalid token.
- `403` — the token is valid, but no rule allows the operation.

<!-- test PR: e2e CI check, do not merge -->
