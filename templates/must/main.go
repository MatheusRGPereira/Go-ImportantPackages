package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	curso := Course{"Golang", 40}

	t := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} - Workload: {{.Workload}} hours"))

	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
