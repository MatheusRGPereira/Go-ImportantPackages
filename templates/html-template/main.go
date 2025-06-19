package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	curso := Course{"Golang", 40}

	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Courses{curso, {"Python", 30}, {"Java", 50}, {"C#", 45}})
	if err != nil {
		panic(err)
	}

}
