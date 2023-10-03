package main

import (
	"fmt"
	"errors"
)

func main() {
	
	fmt.Println("Numeros Inteiros")
	// Inteiros
	var numero int16 = -100
	fmt.Println(numero)


	// Unsigned int
	numero2 := uint (1000)
	fmt.Println(numero2)

	// Alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Numeros Flutuantes")
	// Float
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	numeroReal2 := 423.45
	fmt.Println(numeroReal2)


	fmt.Println("Strings e Chars")
	// String

	var str string = "Texto"
	fmt.Println(str)


	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'B' // char = numero na tabela ASCII
	fmt.Println(char)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Valor zero")
	// Valor zero 

	var texto string
	fmt.Println(texto)

	// Boolean

	var booleano1 bool = true
	fmt.Println(booleano1)

	// Erro
	var erro error
	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro)
	fmt.Println(erro2)
}