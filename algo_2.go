//Crie uma struct chamada Pessoa com os campos "nome", "idade" e "endereço". O campo "endereço" deve ser uma outra struct com os campos "rua", "número", "cidade" e "estado". Escreva uma função que receba uma Pessoa como parâmetro e imprima seu endereço completo.

package main

import (
	"fmt"
)

type Pessoa struct {
	nome  string
	idade int
	endereço
}

type endereço struct {
	rua    string
	numero string
	cidade string
	estado string
}

func main() {

	p := Pessoa{
		nome:  "Matheus",
		idade: 19,
		endereço: endereço{
			rua:    "Rua",
			numero: "Numero",
			cidade: "Cidade",
			estado: "Estado",
		},
	}
	fmt.Printf("Nome: %s \nIdade: %d \nRua: %s \nnumero: %s \ncidade: %s \nestado: %s", p.nome, p.idade, p.rua, p.numero, p.cidade, p.estado)

}
