package pointer

import "fmt"

// explicação sobre ponteiros em GoLag
func Pointer() {
	var number int = 10
	var newNumber int = number // -> (newNumber) -> copy -> (number)
	// Neste print o (number) retorna => 10
	// e (newNumber) retorna => 10
	// também pois estamos guardando uma copia de (number)
	fmt.Println(number, newNumber) // => (10, 10)

	// Nesta variável (incrementNumber) => estamos apontando para o endereço de memoria de (newNumber)
	// onde poderão ser atribuídos novos valores para (newNumber) sem criara uma copia
	var incrementNumber *int = &newNumber
	*incrementNumber++

	// Neste print é mostrado e referencia de memoria do ponteiro.
	fmt.Println(incrementNumber)

	// Neste print é mostrado o valor armazenado no endereço de memoria do ponteiro.
	// onde o (*) indica para fazer a desestruturação da memoria e mostrar o valor.
	fmt.Println(*incrementNumber)
}
