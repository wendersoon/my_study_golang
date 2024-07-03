package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

func CalcularArea(f Forma) {
	fmt.Printf("O valor da área é: %0.2f\n", f.area())
}

// Cada tipo abaixo implementa a interface abstrata forma implicitamente ao terem o método "area() float64"
type Ciruculo struct {
	Raio float64
}

func (c Ciruculo) area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

type Quadrado struct {
	Lado float64
}

func (q Quadrado) area() float64 {
	return math.Pow(q.Lado, 2)
}

type Losango struct {
	BaseMaior float64
	BaseMenor float64
}

func (l Losango) area() float64 {
	return (l.BaseMaior + l.BaseMenor) / 2
}
