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

| Key                          | Type   | Default | Description                           |
| ---------------------------- | ------ | ------: | ------------------------------------- |
| `serviceAccount.create`      | bool   | `false` | Create a ServiceAccount.              |
| `serviceAccount.automount`   | bool   |  `true` | Automount SA tokens into Pods.        |
| `serviceAccount.annotations` | map    |    `{}` | Annotations for the ServiceAccount.   |
| `serviceAccount.name`        | string |    `""` | Use an existing ServiceAccount name (if `serviceAccount.create` is false) or setting the name of ServiceAccount (if `serviceAccount.create` is true)

### Pods

| Key                  | Type | Default | Description                        |
| -------------------- | ---- | ------: | ---------------------------------- |
| `podAnnotations`     | map  |    `{}` | Extra annotations for Pods.        |
| `podLabels`          | map  |    `{}` | Extra labels for Pods.             |
| `podSecurityContext` | map  |    `{}` | Pod-level security context.        |
| `securityContext`    | map  |    `{}` | Container-level security context.  |

### Service & Ingress

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

| Key                           | Type   | Default | Description                      |
| ----------------------------- | ------ | ------: | -------------------------------- |
| `resources`                   | map    |    `{}` | CPU/Memory requests & limits ([See more in documentation](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/)).    |
| `livenessProbe.enabled`       | bool   | `false` | Enable liveness probe.           |
| `livenessProbe.httpGet.path`  | string |     `/` | Path for liveness HTTP check.    |
| `livenessProbe.httpGet.port`  | string |  `http` | Port name/number for liveness.   |
| `readinessProbe.enabled`      | bool   | `false` | Enable readiness probe.          |
| `readinessProbe.httpGet.path` | string |     `/` | Path for readiness HTTP check.   |
| `readinessProbe.httpGet.port` | string |  `http` | Port name/number for readiness.  |

### Autoscaling (HPA)

| Key                                             | Type | Default | Description                            |
| ----------------------------------------------- | ---- | ------: | -------------------------------------- |
| `autoscaling.enabled`                           | bool | `false` | Enable HorizontalPodAutoscaler.        |
| `autoscaling.minReplicas`                       | int  |     `1` | Minimum replicas.                      |
| `autoscaling.maxReplicas`                       | int  |   `100` | Maximum replicas.                      |
| `autoscaling.targetCPUUtilizationPercentage`    | int  |    `80` | Target average CPU utilization.        |
| `autoscaling.targetMemoryUtilizationPercentage` | int  | *unset* | (Optional) Target memory utilization.  |

### App Configuration

| Key                         | Type | Default | Description                        |
| --------------------------- | ---- | ------: | ---------------------------------- |
| `configuration.annotations` | map  |    `{}` | Extra annotations for app configs (secrets).  |
| `extraEnvs`                 | map  |    `{}` | Additional environment variables for pod.  |

### Redis (Subchart)

| Key                                | Type   |               Default | Description                        |
| ---------------------------------- | ------ | --------------------: | ---------------------------------- |
| `redis.enabled`                    | bool   |               `false` | Enable embedded Redis (dev only).  |
| `redis.*`                          | map    |                       | You can put anything here from [subchart Redis](https://github.com/bitnami/charts/blob/main/bitnami/redis/values.yaml).                    |

### External Redis

| Key                               | Type   | Default | Description                                                       |
| --------------------------------- | ------ | ------: | ----------------------------------------------------------------- |
| `externalRedis.url`               | string | (empty) | Full Redis URL (user\:pass\@host\:port).                          |
| `externalRedis.existingSecret`    | string |    `""` | Secret name containing key `url`; overrides `externalRedis.url`.  |
| `externalRedis.secretAnnotations` | map    |    `{}` | Annotations for the external Redis Secret.                        |

### PostgreSQL (Subchart)

| Key                                | Type   |                    Default | Description                                                       |
| ---------------------------------- | ------ | -------------------------: | ----------------------------------------------------------------- |
| `postgresql.enabled`               | bool   |                    `false` | Enable embedded PostgreSQL (dev only).                            |
| `postgresql.*`                          | map    |                       | You can put anything here from [subchart PostgreSQL](https://github.com/bitnami/charts/blob/main/bitnami/postgresql/values.yaml). 

### External PostgreSQL

| Key                                    | Type   | Default | Description                                                            |
| -------------------------------------- | ------ | ------: | ---------------------------------------------------------------------- |
| `externalPostgresql.url`               | string | (empty) | Full PostgreSQL URL (postgresql://user\:pass\@host\:port/db).                       |
| `externalPostgresql.existingSecret`    | string |    `""` | Secret name containing key `url`; overrides `externalPostgresql.url`.  |
| `externalPostgresql.secretAnnotations` | map    |    `{}` | Annotations for the external PostgreSQL Secret.                        |

### Storage & Scheduling

| Key            | Type | Default | Description                        |
| -------------- | ---- | ------: | ---------------------------------- |
| `volumes`      | list |    `[]` | Extra volumes for the Deployment.  |
| `volumeMounts` | list |    `[]` | Extra volume mounts.               |
| `nodeSelector` | map  |    `{}` | Node selection constraints.        |
| `tolerations`  | list |    `[]` | Pod tolerations.                   |
| `affinity`     | map  |    `{}` | Affinity rules.                    |

---

## Examples

**Use external managed services:**

```bash
helm install state-manager . -n state-manager \
  --create-namespace \
  --set externalRedis.existingSecret="redis-url" \
  --set externalPostgresql.existingSecret="pg-url"
```

See `values.yaml` for all options and further comments in-line.

---

## Notes

* Probes and HPA are disabled by default; enable and tune for production traffic.
* If you set both `externalRedis.url` *and* `externalRedis.existingSecret`, the secret takes precedence (same for PostgreSQL).