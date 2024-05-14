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
}
