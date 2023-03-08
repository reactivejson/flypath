{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "flypath.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "flypathfullname" -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- /*
flypath.chartref prints a chart name and version.
It does minimal escaping for use in Kubernetes labels.
*/ -}}
{{- define "flypath.chartref" -}}
  {{- replace "+" "_" .Chart.Version | printf "%s-%s" .Chart.Name -}}
{{- end -}}

{{/*
flypath.labels.standard prints the standard Helm labels.
The standard labels are frequently used in metadata.
*/}}
{{- define "flypath.labels.standard" -}}
app: {{template "flypath.name" .}}
chart: {{template "flypath.chartref" . }}
app.kubernetes.io/name: {{template "flypath.name" .}}
helm.sh/chart: {{template "flypath.chartref" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/version: {{ .Chart.Version }}
com.nokia.neo/commitId: ${COMMIT_ID}
{{- end -}}

{{/*
flypath.template.labels prints the template metadata labels.
*/}}
{{- define "flypath.template.labels" -}}
app: {{template "flypath.name" .}}
{{- end -}}

{{- define "flypath.app" -}}
app: {{template "flypath.name" .}}
{{- end -}}

{{- define "annotateResources" -}}
# Preserve the workingsetpluginregistrations.ws.nokia.com crd for changes;
kubectl annotate --overwrite crd workingsetpluginregistrations.ws.nokia.com "helm.sh/resource-policy"=keep;
sleep 30;
{{- end -}}
