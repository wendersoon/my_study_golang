package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Quadrado", func(t *testing.T) {
		qua := Quadrado{12}
		areaEsperada := float64(144)
		areaRecebida := qua.area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da esperada %f .",
				areaRecebida,
				areaEsperada)
		}

	})

	t.Run("Circulo", func(t *testing.T) {
		cir := Ciruculo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := cir.area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da esperada %f .",
				areaRecebida,
				areaEsperada)
		}

	})

}
