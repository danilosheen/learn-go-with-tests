package main

import (
	"reflect"
	"testing"
)

// func TestSoma(t *testing.T) {

//     numeros := [5]int{1, 2, 3, 4, 5}

//     resultado := Soma(numeros)
//     esperado := 15

//     if esperado != resultado {
//         t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
//     }
// }

// func Soma(numeros [5]int) int {
// 	soma := 0
// 	for i := 0; i < 5; i++ {
// 		soma += numeros[i]
// 	}
// 	return soma
// }

//enxugando o código
// func Soma(numeros [5]int) int {
//     soma := 0
//     for _, numero := range numeros {
//         soma += numero
//     }
//     return soma
// }

//refatorando para aceitar qualquer tamanho de array usando slice
func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func TestSoma(t *testing.T) {

	t.Run("coleção de 5 números", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		if resultado != esperado {
			t.Errorf("resultado %d, want %d, dado %v", resultado, esperado, numeros)
		}
	})

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})

}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
    var somas []int
    for _, numeros := range numerosParaSomar {
        if len(numeros) == 0 {
            somas = append(somas, 0)
        } else {
            final := numeros[1:]
            somas = append(somas, Soma(final))
        }
    }

    return somas
}

func TestSomaTodoOResto(t *testing.T) {

	verificaSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %v, esperado %v", resultado, esperado)
		}
	}

	t.Run("faz a soma do resto", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}
		verificaSomas(t, resultado, esperado)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}
		verificaSomas(t, resultado, esperado)
	})

}
