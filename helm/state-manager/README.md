# State Manager — Helm Chart

A production-ready Helm chart to deploy the open-source [**State Manager**](https://github.com/reconcile-kit/state-manager) application on Kubernetes. It includes optional subcharts for Redis and PostgreSQL (handy for local/dev).

---

## Prerequisites

* Kubernetes cluster
* Helm 3.x installed
* (Optional) External Redis and/or PostgreSQL endpoints if you don’t enable the bundled dev subcharts.

---

## Quick Start

### Installation

Start from the provided `values.yaml` and adjust as needed:

```bash
# for example, enable ingress
helm install state-manager . \
  --namespace state-manager \
  --create-namespace \
  --set ingress.enabled=true \
  --set ingress.hosts[0].host="state.example.com"
```

If you prefer a custom values file:

```bash
helm install state-manager . -n state-manager --create-namespace -f my-values.yaml
```

> Tip: It's not recommended to use embedded Redis and PostgreSQL subharts for production.

---

## Upgrade

```bash
helm upgrade state-manager . -n state-manager -f my-values.yaml
```

---

## Uninstall

```bash
helm uninstall state-manager -n state-manager
```

This removes chart resources from the namespace (PVCs from any external databases you used are not touched).

---

## Configuration Reference

All chart parameters with their defaults are listed below. Unless noted, values come from `values.yaml`.

### Core

Configuration for replicas and the container image.

| Key                | Type   |                   Default | Description                                           |
| ------------------ | ------ | ------------------------: | ----------------------------------------------------- |
| `replicaCount`     | int    |                       `1` | Number of replicas for the Deployment.                |
| `image.repository` | string | `dhnikolas/state-manager` | Container image repository.                           |
| `image.pullPolicy` | string |            `IfNotPresent` | Image pull policy.                                    |
| `image.tag`        | string |                      `""` | Image tag (defaults to chart `appVersion` if empty).  |
| `imagePullSecrets` | list   |                      `[]` | Secrets for private registries.                       |
| `nameOverride`     | string |                      `""` | Override chart name.                                  |
| `fullnameOverride` | string |                      `""` | Override full release name.                           |

### Service Account

Control creation and use of a ServiceAccount.

| Key                          | Type   | Default | Description                           |
| ---------------------------- | ------ | ------: | ------------------------------------- |
| `serviceAccount.create`      | bool   | `false` | Create a ServiceAccount.              |
| `serviceAccount.automount`   | bool   |  `true` | Automount SA tokens into Pods.        |
| `serviceAccount.annotations` | map    |    `{}` | Annotations for the ServiceAccount.   |
| `serviceAccount.name`        | string |    `""` | Use an existing ServiceAccount name (if `serviceAccount.create` is false) or setting the name of ServiceAccount (if `serviceAccount.create` is true)

### Pod Settings

Extra metadata and security context for Pods/containers.

| Key                  | Type | Default | Description                        |
| -------------------- | ---- | ------: | ---------------------------------- |
| `podAnnotations`     | map  |    `{}` | Extra annotations for Pods.        |
| `podLabels`          | map  |    `{}` | Extra labels for Pods.             |
| `podSecurityContext` | map  |    `{}` | Pod-level security context.        |
| `securityContext`    | map  |    `{}` | Container-level security context.  |

### Service & Ingress

Kubernetes Service and HTTP ingress configuration.

| Key                                  | Type   |                  Default | Description             |
| ------------------------------------ | ------ | -----------------------: | ----------------------- |
| `service.type`                       | string |              `ClusterIP` | Service type.           |
| `service.port`                       | int    |                   `8080` | Service port.           |
| `ingress.enabled`                    | bool   |                  `false` | Enable Ingress.         |
| `ingress.className`                  | string |                     `""` | IngressClass name.      |
| `ingress.annotations`                | map    |                     `{}` | Ingress annotations.    |
| `ingress.hosts[0].host`              | string |    `chart-example.local` | Example host entry.     |
| `ingress.hosts[0].paths[0].path`     | string |                      `/` | Path for the host.      |
| `ingress.hosts[0].paths[0].pathType` | string | `ImplementationSpecific` | Path type.              |
| `ingress.tls`                        | list   |                     `[]` | TLS secrets and hosts.  |

### Resources & Probes

Resource requests/limits and basic HTTP probes.

| Key                           | Type   | Default | Description                      |
| ----------------------------- | ------ | ------: | -------------------------------- |
| `resources`                   | map    |    `{}` | CPU/Memory requests & limits ([see more in documentation](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/)).    |
| `livenessProbe.enabled`       | bool   | `false` | Enable liveness probe.           |
| `livenessProbe.httpGet.path`  | string |     `/` | Path for liveness HTTP check.    |
| `livenessProbe.httpGet.port`  | string |  `http` | Port name/number for liveness.   |
| `readinessProbe.enabled`      | bool   | `false` | Enable readiness probe.          |
| `readinessProbe.httpGet.path` | string |     `/` | Path for readiness HTTP check.   |
| `readinessProbe.httpGet.port` | string |  `http` | Port name/number for readiness.  |

### Autoscaling (HPA)

Horizontal Pod Autoscaler settings.

| Key                                             | Type | Default | Description                            |
| ----------------------------------------------- | ---- | ------: | -------------------------------------- |
| `autoscaling.enabled`                           | bool | `false` | Enable HorizontalPodAutoscaler.        |
| `autoscaling.minReplicas`                       | int  |     `1` | Minimum replicas.                      |
| `autoscaling.maxReplicas`                       | int  |   `100` | Maximum replicas.                      |
| `autoscaling.targetCPUUtilizationPercentage`    | int  |    `80` | Target average CPU utilization.        |
| `autoscaling.targetMemoryUtilizationPercentage` | int  |    `80` | (Optional) Target memory utilization.  |

### App Configuration

Additional env vars for the app.

| Key                         | Type | Default | Description                        |
| --------------------------- | ---- | ------: | ---------------------------------- |
| `extraEnvs`                 | map  |    `{}` | Additional environment variables for pod.  |

### Redis (Subchart)

Built-in Redis for development (not recommended for production).

| Key                                | Type   |               Default | Description                        |
| ---------------------------------- | ------ | --------------------: | ---------------------------------- |
| `redis.enabled`                    | bool   |               `false` | Enable embedded Redis (dev only).  |
| `redis.*`                          | map    |                       | You can put anything here from [subchart Redis](https://github.com/bitnami/charts/blob/main/bitnami/redis/values.yaml).                    |

### External Redis

Use an external Redis instead of the dev subchart.

| Key                                         | Type   | Default | Description                                 |
| ------------------------------------------- | ------ | ------: | ------------------------------------------- |
| `externalRedis.host`                        | string |    `""` | Redis host.                                 |
| `externalRedis.port`                        | int    |  `6379` | Redis port.                                 |
| `externalRedis.scheme`                      | string | `redis` | Connection scheme (`redis`/`rediss`).       |
| `externalRedis.db`                          | int    |     `0` | Database number.                            |
| `externalRedis.auth.enabled`                | bool   | `false` | Enable username/password auth.              |
| `externalRedis.auth.user`                   | string |    `""` | Username (if auth enabled).                 |
| `externalRedis.auth.password`               | string |    `""` | Password (ignored if `existingSecret` set). |
| `externalRedis.auth.existingSecret`         | string |    `""` | Existing Secret containing credentials. Overrides `externalRedis.auth.password`  |
| `externalRedis.auth.secretKeys.passwordKey` | string |    `""` | Key name in the Secret for the password. Overrides `externalRedis.auth.password`  |

### PostgreSQL (Subchart)

Built-in PostgreSQL for development (not recommended for production).

| Key                                | Type   |                    Default | Description                                                       |
| ---------------------------------- | ------ | -------------------------: | ----------------------------------------------------------------- |
| `postgresql.enabled`               | bool   |                    `false` | Enable embedded PostgreSQL (dev only).                            |
| `postgresql.*`                          | map    |                       | You can put anything here from [subchart PostgreSQL](https://github.com/bitnami/charts/blob/main/bitnami/postgresql/values.yaml). 

### External PostgreSQL


Use an external PostgreSQL instead of the dev subchart.

| Key                                         | Type   |         Default | Description                                          |
| ------------------------------------------- | ------ | --------------: | ---------------------------------------------------- |
| `externalPostgresql.host`                   | string |            `""` | Database host.                                       |
| `externalPostgresql.port`                   | int    |          `5000` | Database port.                                       |
| `externalPostgresql.user`                   | string | `state_manager` | Database user.                                       |
| `externalPostgresql.database`               | string | `state_manager` | Database name.                                       |
| `externalPostgresql.password`               | string |          `test` | Database password (ignored if `existingSecret` set). |
| `externalPostgresql.existingSecret`         | string |            `""` | Existing Secret containing DB parameters. Overrides `externalPostgresql.password`           |
| `externalPostgresql.secretKeys.passwordKey` | string |            `""` | Key name in the Secret for the password. Overrides `externalPostgresql.password`            |

### Storage & Scheduling

Extra volumes/mounts and pod scheduling preferences.

| Key            | Type | Default | Description                        |
| -------------- | ---- | ------: | ---------------------------------- |
| `volumes`      | list |    `[]` | Extra volumes for the Deployment.  |
| `volumeMounts` | list |    `[]` | Extra volume mounts.               |
| `nodeSelector` | map  |    `{}` | Node selection constraints.        |
| `tolerations`  | list |    `[]` | Pod tolerations.                   |
| `affinity`     | map  |    `{}` | Affinity rules.                    |

---

## Notes

* Probes and HPA are disabled by default; enable and tune for production traffic.
* If you set both `externalRedis.auth.password` and `externalRedis.existingSecret`, the secret takes precedence (same for PostgreSQL).
* If you set both `externalRedis.host` and `redis.enabled`, the externalRedis configuration takes precedence (same for PostgreSQL).


# Change Log

## 0.2.0
- Update version of state-manager application to [v0.0.8](https://github.com/reconcile-kit/state-manager/releases/tag/v0.0.8).
- Change the way to configure the application
- Supported secrets with password from subcharts
- Changes on configuring external redis and postgresql credentials in values
