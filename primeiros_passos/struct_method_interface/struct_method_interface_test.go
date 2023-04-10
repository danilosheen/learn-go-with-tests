package structmethodinterface

import (
	"math"
	"testing"
)

type Forma interface {
	Area() float64
}

// declarando uma struct
type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
	}
}

func TestArea(t *testing.T) {

	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
        t.Helper()
        resultado := forma.Area()

        if resultado != esperado {
            t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
        }
    }

	t.Run("retângulos", func(t *testing.T) {
        retangulo := Retangulo{12.0, 6.0}
        verificaArea(t, retangulo, 72.0)
    })

    t.Run("círculos", func(t *testing.T) {
        circulo := Circulo{10}
        verificaArea(t, circulo, 314.1592653589793)
    })

}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Altura + retangulo.Largura)
}

func Area(retangulo Retangulo) float64 {
	return retangulo.Altura * retangulo.Largura
}
