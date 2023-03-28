package main

import (
	"fmt"
	"sync"
)

//var w = 30

//type dog int

//var d dog

type pessoa struct {
	nome  string
	idade int
}

type analista struct {
	pessoa
	tipo    string
	salario float32
}

type tecnico struct {
	pessoa
	cargo string
}

// (a analista) é um receiver e não parâmetro
func (a analista) ola() {
	fmt.Println(a.nome, "analista diz olá")
}

func (t tecnico) ola() {
	fmt.Println(t.nome, "técnico diz olá")
}

type gente interface {
	ola()
}

func humano(g gente) {
	switch g.(type) {
	case analista:
		fmt.Println("Eu ganho: ", g.(analista).salario)
	case tecnico:
		fmt.Println("Eu trabalho como: ", g.(tecnico).cargo)

	}
	g.ola()

}

var wg sync.WaitGroup

func main() {
	//x := 13
	//d = 10
	//x := 13 == 10
	//y := "teste"
	//z := false
	//numBytes, erros := fmt.Println("Hello World")
	//fmt.Println(numBytes, erros)
	//fmt.Println(x, y, z, w)
	//firstFunc(x)
	//x = int(d)
	//fmt.Println(x)
	davi := analista{
		pessoa: pessoa{
			nome:  "Davi",
			idade: 30,
		},
		tipo:    "Tecnologia da Informação",
		salario: 20000,
	}

	junio := tecnico{
		pessoa: pessoa{
			nome:  "Junio",
			idade: 30,
		},
		cargo: "Juiz",
	}

	//davi.ola()
	humano(davi)
	humano(junio)

	c := make(chan int)
	wg.Add(1)
	go meuloop(10, c)
	wg.Add(1)
	go prints(c)
	wg.Wait()

}

func meuloop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	wg.Done()
	close(s)
}

func prints(r <-chan int) {
	for v := range r {
		fmt.Println("Recebido do canal:", v)
	}
	wg.Done()
}

//func firstFunc(x int) {
//	fmt.Println(x)
//	println(w)
//}
