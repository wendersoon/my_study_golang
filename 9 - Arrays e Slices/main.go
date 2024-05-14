package main

import "fmt"

func main() {

	fmt.Println("Arrays -----------")

	var x [5]int //O tamanho é definido
	fmt.Println(x)

	x[1] = 1 //Forma de atribuir valor a um array
	x[3] = 100
	fmt.Println(x)

	y := [5]string{} //Outra forma de criação de array
	fmt.Println(y)

	y[4] = "String"
	fmt.Println(y)

	z := [3]int{1, 2, 3} //Criando com valores diretos
	fmt.Println(z)

	w := [...]int{1, 2, 3, 4, 5, 6, 7} //Os três pontos ... indicam que o tamanho será definido de acordo com a quantidede de
	// dados inseridos, isso não quer dizer que é dinâmico mas apenas que o tamanho será definido "após" inserir os valores
	fmt.Println(w)

	fmt.Println("Slices -----------")
	slice1 := []int{1, 2, 3} //O tamanho do slice é dinâmico
	fmt.Println(slice1)

	slice1 = append(slice1, 5) //Forma de adicionar um valor no slice
	fmt.Println(slice1)

	//O slice tem uma referência a um array na sua lógica interna

	//Quando crio um slice apartir de um aray, ele continua a referênciá-lo
	slice2 := w[2:6]

	fmt.Println(slice2)

	w[4] = 5555 //Se eu mudar o valor da posição 4 no array, o slice que referência-o mudará
	fmt.Println(slice2)

	//Arrays Internos

	fmt.Println("---------------------------")

	slice5 := make([]float64, 2, 3)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	slice5 = append(slice5, 2)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	slice5 = append(slice5, 5) //Quando 'estoura' a capacidade, esta é dobrada de tamanho
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	slice5 = append(slice5, 5)
	slice5 = append(slice5, 5)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	slice5 = append(slice5, 10) //Cada vez que ultrapassa a capacidade, a capacidade máxima do array é dobrado
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	//Testando diretamente sem usar o make
	fmt.Println("Sem usar a função make que internamente é chamada na criação de um slice")
	sliceX := []int{10, 10}
	fmt.Println(sliceX)
	fmt.Println(len(sliceX))
	fmt.Println(cap(sliceX))

	sliceX = append(sliceX, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10) //a capacidade do slicex é aumentada
	fmt.Println(sliceX)
	fmt.Println(len(sliceX))
	fmt.Println(cap(sliceX))

}
