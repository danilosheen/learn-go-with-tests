package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 5)
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}

}

func Repetir(caractere string, multiplicador int) string {
	var repeticoes string
	for i := 0; i < multiplicador; i++ {
		repeticoes += caractere
	}
	return repeticoes
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 5)
	}
}
