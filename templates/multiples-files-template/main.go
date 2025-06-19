package main

import (
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name     string
	Workload int
}

func ToUper(s string) string {
	return strings.ToUpper(s)
}

type Courses []Course

func main() {

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")

	t.Funcs(template.FuncMap{"ToUper": ToUper})

	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Courses{{"Golang", 40}, {"Python", 30}, {"Java", 50}, {"C#", 45}})
	if err != nil {
		panic(err)
	}

}
