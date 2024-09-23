package main

import "fmt"

func soma(a int, b int) int {
	return a + b
}

func main() {
	var opcao int

	for opcao != 5 {
		fmt.Println("Calculadora")
		fmt.Println("1. Soma")
		fmt.Println("2. Subtração")
		fmt.Println("3. Multiplicação")
		fmt.Println("4. Divisão")
		fmt.Println("5. Sair")
		fmt.Scanln(&opcao)
	}
}