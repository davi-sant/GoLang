package maps

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func MyMaps() {

	fmt.Println("Init maps")
	myMaps := map[string]map[string]User{
		"Candidato1": {"user": {name: "Davi", age: 20}},
		"Candidato2": {"user": {name: "Davi", age: 20}},
		"Candidato3": {"user": {name: "Davi", age: 20}},
	}
	fmt.Println(myMaps["Candidato1"]["user"].age)
}
