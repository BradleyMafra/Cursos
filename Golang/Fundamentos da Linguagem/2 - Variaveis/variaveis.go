package main

import "fmt"

func main() {
	var variavelString string = "Variável do tipo string"
	
	variavelStringNova := "Variável do tipo string sem especificar o tipo"

	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	variavel5, variavel6 := "Variável 5", "Variável 6"

	
	fmt.Println(variavelString)
	fmt.Println(variavelStringNova)
	fmt.Println(variavel3, variavel4)
	fmt.Println(variavel5, variavel6)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)

}