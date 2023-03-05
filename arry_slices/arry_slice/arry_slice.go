package arry_slice

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func ArrySlice() {
	fmt.Println("Arry and Slice")
	arry := [2]string{"nome-01", "nome-02"}
	fmt.Println(arry)
	slice1 := []User{{name: "Robert", age: 20}, {name: "Maike", age: 27}}
	fmt.Println(slice1[1].name)
	fmt.Println(slice1, "->", "len: ", len(slice1), "capacity: ", cap(slice1))
}
