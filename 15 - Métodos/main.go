package main

import "fmt"

type estudante struct {
	nome  string
	curso string
	tempo uint8
}

func (e estudante) verNome() {
	fmt.Printf("O nome do estudante é: %s.\n", e.nome)
}

func (e estudante) verCurso() string {
	return e.curso
}

func (e estudante) verTempo() {
	fmt.Printf("O tempo de formação atual é %d ano.\n", e.tempo)
}

func (e *estudante) aumentarTempo() {
	e.tempo++
}

func main() {

	aluno := estudante{"Wenderson", "Ciência da Computação", 1}

	fmt.Println(aluno)

	aluno.verNome()
	aluno.verTempo()
	nomeDoCurso := aluno.verCurso()
	fmt.Println(nomeDoCurso)

	//O aluno estudou por 4 anos
	aluno.aumentarTempo()
	aluno.verTempo() // o valor do tipo é alterado dentro da estrutura
	aluno.aumentarTempo()
	aluno.aumentarTempo()
	aluno.verTempo()
}
