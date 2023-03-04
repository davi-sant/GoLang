package funcs

import (
	"errors"
	"fmt"
)

func HandleFuncs() {
	fmt.Println("função que print um texto na tela")
}

// Essa função retorna a media do aluno
func HandleFuncsParams(nota int, nota2 int, nota3 int) (float32, error) {
	media := (nota + nota2 + nota3) / 2
	var msgErr error = errors.New("aluno de recuperação")

	if media < 7 {
		return float32(media), msgErr
	} else {
		return float32(media), nil
	}

}
