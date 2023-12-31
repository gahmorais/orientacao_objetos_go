package contas

import "orientacao_objetos/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Não foi possível realizar o saque! saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Você não pode depositar valor negativo", c.saldo
	}
}

func (c *ContaPoupanca) Transferir(valorTransferencia float64, contaDestino *ContaPoupanca) bool {
	if c.saldo > valorTransferencia && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaPoupanca) MostraSaldo() float64 {
	return c.saldo
}
