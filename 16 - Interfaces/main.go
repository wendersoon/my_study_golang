package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func calcularArea(f forma) {
	fmt.Printf("O valor da área é: %0.2f\n", f.area())
}

// Cada tipo abaixo implementa a interface abstrata forma implicitamente ao terem o método "area() float64"
type ciruculo struct {
	raio float64
}

func (c ciruculo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type quadrado struct {
	lado float64
}

func (q quadrado) area() float64 {
	return math.Pow(q.lado, 2)
}

type losango struct {
	baseMaior float64
	baseMenor float64
}

func (l losango) area() float64 {
	return (l.baseMaior + l.baseMenor) / 2
}

func main() {

	qua := quadrado{2}
	cir := ciruculo{10}
	los := losango{1, 5}

	calcularArea(qua)
	calcularArea(cir)
	calcularArea(los)

}
