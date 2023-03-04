package exp_structs

import "fmt"

type User struct {
	name  string
	age   int8
	skill string
}
type Estudo struct {
	User
	estudando string
}

func ExpStruct() {
	fmt.Println("Exemplo struct")
	user := User{ "Davi",  20,"Programador de sistemas"}
	fmt.Println(user)
	newUser := Estudo{user, "GoLang"}
	fmt.Println(newUser)
	fmt.Println(newUser.estudando)
}
