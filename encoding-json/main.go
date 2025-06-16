package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name string `json:"n"`
	Age  int    `json:"a"`
}

func main() {
	person := Person{Name: "John biden", Age: 30}

	res, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(person)

	pureJson := []byte(`{"n":"Johny rivers","a":23}`)
	var person2 Person
	err = json.Unmarshal(pureJson, &person2)
	if err != nil {
		panic(err)
	}

	println(person2.Name, person2.Age)

}
