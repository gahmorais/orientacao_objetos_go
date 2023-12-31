package main

import (
	"fmt"
	"orientacao_objetos/clientes"
	"orientacao_objetos/contas"
)

func main() {

	contaGabriel := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Gabriel",
			Cpf:       "11122233344",
			Profissao: "Desenvolvedor",
		},
		NumeroAgencia: 589,
		NumeroConta:   111222,
		Saldo:         100,
	}

	fmt.Println(contaGabriel)

}
