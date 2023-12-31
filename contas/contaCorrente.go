package contas

import (
	"orientacao_objetos/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.Saldo && valorDoSaque > 0
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Não foi possível realizar o saque! Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.Saldo += valorDeposito
		return "Depósito realizado com sucesso", c.Saldo
	} else {
		return "Você não pode depositar valor negativo", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if c.Saldo > valorTransferencia && valorTransferencia > 0 {
		c.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}
