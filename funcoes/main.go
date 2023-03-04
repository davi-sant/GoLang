package main

import (
	"fmt"
	"funcoes/funcs"
)

func main() {
	fmt.Println("Inicio mdo programa")
	funcs.HandleFuncs()
	media, err := funcs.HandleFuncsParams(5, 3, 5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Resultado Media -> ", media)
}
