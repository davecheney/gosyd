package main

import "text/template"; import "os"

const tmpl = `
{{ range . -}}
{{ . }}
{{- end }}`

func main() {
	t := template.Must(template.New("").Parse(tmpl))
	t.Execute(os.Stdout, []string{ "mars", "mercury", "pluto", "neptune" })
}
