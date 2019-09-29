package main

import (
	"os"
	"text/template"
)

type Friend struct {
	Fname string
}
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
	BigMaps  []map[string]string
}

func main() {
	f1 := Friend{Fname: "xiaofang"}
	f2 := Friend{Fname: "wugui"}
	t := template.New("test")
	t = template.Must(t.Parse(
		`hello {{.UserName}}!
{{ range .Emails }}
an email {{ . }}
{{- end }}
{{ with .Friends }}
{{- range . }}
my friend name is {{.Fname}}
{{- end }}
{{ end }}
{{ range $k,$v := .BigMaps}}
{{$k}}
{{index $v "A"}}
{{index $v "B"}}
{{end}}
`))
	map1 := map[string]string{"A": "A"}
	map2 := map[string]string{"B": "B"}
	p := Person{UserName: "longshuai",
		Emails:  []string{"a1@qq.com", "a2@gmail.com"},
		Friends: []*Friend{&f1, &f2},
		BigMaps: []map[string]string{map1, map2},
	}
	t.Execute(os.Stdout, p)
}