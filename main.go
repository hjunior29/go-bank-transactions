package main

import (
	// "Bank/clientes"
	"Bank/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64)  {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDenis := contas.ConntaPoupanca{}
	contaDaPati := contas.ContaCorrente{}
	contaDenis.Depositar(1000)
	PagarBoleto(&contaDenis, 50)

	fmt.Println(contaDenis.ObterSaldo())

	contaDaPati.Depositar(150)
	PagarBoleto(&contaDaPati, 140)

	fmt.Println(contaDaPati.ObterSaldo())

}
