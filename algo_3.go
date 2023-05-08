//Crie uma struct chamada Triângulo com os campos "base" e "altura". Escreva uma função que receba um Triângulo como parâmetro e calcule a área do triângulo (área = base * altura / 2).

package main

import "fmt"

type Triangulo struct {
	base   int
	altura int
}

func area(triangulo Triangulo) int {

	Area := triangulo.base * (triangulo.altura / 2)
	return Area
}

func main() {

	p := Triangulo{
		base:   0,
		altura: 0,
	}
	fmt.Println("altura")
	fmt.Scan(&p.altura)
	fmt.Println("base")
	fmt.Scan(&p.base)

	fmt.Println("A area é:", area(p))

}
