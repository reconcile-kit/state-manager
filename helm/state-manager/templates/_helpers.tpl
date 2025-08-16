{{/*
Expand the name of the chart.
*/}}
{{- define "state-manager.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "state-manager.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "state-manager.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "state-manager.labels" -}}
helm.sh/chart: {{ include "state-manager.chart" . }}
{{ include "state-manager.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "state-manager.selectorLabels" -}}
app.kubernetes.io/name: {{ include "state-manager.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "state-manager.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "state-manager.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Create redis name and version as used by the chart label.
*/}}
{{- define "state-manager.redis.fullname" }}
{{- $redis := (index .Values "redis") }}
{{- $redisContext := dict "Chart" (dict "Name" "redis") "Release" .Release "Values" $redis }}
{{- if $redis.enabled }}
    {{- printf "%s-master" (include "common.names.fullname" $redisContext) }}
{{- end }}
{{- end }}

{{/*
Generate a host to redis
*/}}
{{- define "state-manager.redis.host" -}}
{{- $redis := (index .Values "redis") }}
{{- if $redis.enabled }}
    {{- printf "%s" (include "state-manager.redis.fullname" .) }}
{{- else }}
    {{- printf "%s" .Values.externalRedis.host }}
{{- end }}
{{- end }}

{{/*
Generate a port to redis
*/}}
{{- define "state-manager.redis.port" -}}
{{- $redis := (index .Values "redis") }}
{{- if $redis.enabled }}
    {{- $redisContext := dict "Chart" (dict "Name" "redis") "Release" .Release "Values" $redis }}
    {{- printf "%s" (toString $redisContext.Values.master.service.ports.redis) }}
{{- else }}
    {{- printf "%s" .Values.externalRedis.port }}
{{- end }}
{{- end }}

{{/*
Create postgresql name and version as used by the chart label.
*/}}
{{- define "state-manager.postgresql.fullname" }}
{{- $postgresql := (index .Values "postgresql") }}
{{- $postgresqlContext := dict "Chart" (dict "Name" "postgresql") "Release" .Release "Values" $postgresql }}
{{- if $postgresql.enabled }}
    {{- include "postgresql.v1.primary.fullname" $postgresqlContext }}
{{- end }}
{{- end }}

{{/*
Generate a host to postgresql
*/}}
{{- define "state-manager.postgresql.host" -}}
{{- $postgresql := (index .Values "postgresql") }}
{{- if $postgresql.enabled }}
    {{- printf "%s" (include "state-manager.postgresql.fullname" .) }}
{{- else }}
    {{- printf "%s" .Values.externalPostgresql.host }}
{{- end }}
{{- end }}

{{/*
Generate a port to postgresql
*/}}
{{- define "state-manager.postgresql.port" -}}
{{- $postgresql := (index .Values "postgresql") }}
{{- if $postgresql.enabled }}
    {{- $postgresqlContext := dict "Chart" (dict "Name" "postgresql") "Release" .Release "Values" $postgresql }}
    {{- printf "%s" (include "postgresql.v1.service.port" $postgresqlContext) }}
{{- else }}
    {{- printf "%s" .Values.externalPostgresql.port }}
{{- end }}
{{- end }}