package main

import (
	"fmt"
)

func somar(numero1 int8, numero2 int8) int8{
	return numero1 + numero2
}

func calculos(numero1, numero2 int8) (int8, int8){
	somar := numero1 + numero2
	subtrair := numero1 - numero2

	return somar, subtrair
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função f")
	}

	f()

	var texto = func(txt string) {
		fmt.Println(txt)
	}

	texto("Texto da função")

	
	// Usar 2 variáveis de retorno
	resultadoSoma, resultadoSubtracao := calculos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// Usar 1 variável de retorno com uma função que retorna 2 valores
	resultadoSoma, _ := calculos(10, 15)
	fmt.Println(resultadoSoma)
}