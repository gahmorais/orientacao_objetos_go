package main

import (
	"fmt"
	"orientacao_objetos/clientes"
	"orientacao_objetos/contas"
)

func main() {

	clienteGabriel := clientes.Titular{
		Nome:      "Gabriel",
		Cpf:       "11122233344",
		Profissao: "Desenvolvedor",
	}

	contaGabriel := contas.ContaCorrente{
		Titular:       clienteGabriel,
		NumeroAgencia: 589,
		NumeroConta:   111222,
	}

	contaGabriel.Depositar(100)

	fmt.Println(contaGabriel)
	fmt.Println(contaGabriel.MostraSaldo())

	contaLuiz := contas.ContaPoupanca{}
	contaChica := contas.ContaCorrente{}
	contaLuiz.Depositar(500)
	contaChica.Depositar(600)
	fmt.Println(contaLuiz)
	fmt.Println(contaChica)

	contas.PagarBoleto(&contaGabriel, 60)
	fmt.Println(contaGabriel.MostraSaldo())
	contas.PagarBoleto(&contaChica, 30)
	fmt.Println(contaChica.MostraSaldo())

}
