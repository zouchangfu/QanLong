package global

{{- if .HasGlobal }}

import "github.com/zouchangfu/QanLong/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}