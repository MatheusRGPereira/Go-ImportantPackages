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

	temp := template.New("CourseTemplate")

	temp, _ = temp.Parse("Course: {{.Name}} - Workload: {{.Workload}} hours")
	err := temp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
