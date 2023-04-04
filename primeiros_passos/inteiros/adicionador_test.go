package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {
	soma := Adiciona(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
	}
}

func Adiciona(x, y int) int {
	return x + y
}


// a func example só aparece no test utilizando o comentário do output
func ExampleAdiciona() {
	soma := Adiciona(1, 5)
	fmt.Println(soma)
	// output: 6
}
