module {{ .Module }}

go {{ .GoVersion }}

{{if (VersionCompare .GoVersion ">=" "1.24")}}
tool (
	github.com/polpo666/pzero/cmd/pzero
)
{{end}}