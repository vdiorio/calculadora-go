package main

import "fmt"

func soma(a int, b int) int {
	return a + b
}

func subtracao(a int, b int) int {
	return a - b
}

func multiplicacao(a int, b int) int {
	return a * b
}

func divisao(a int, b int) int {
	return a / b
}

func printResult(result int) {
	// ANSI escape code for green text
	green := "\033[32m"
	reset := "\033[0m"
	fmt.Println("Resultado: ", green, result, reset)
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

		switch opcao {
		case 1:
			var a, b int
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanln(&a)
			fmt.Print("Digite o segundo número: ")
			fmt.Scanln(&b)
			printResult(soma(a, b))
		case 2:
			var a, b int
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanln(&a)
			fmt.Print("Digite o segundo número: ")
			fmt.Scanln(&b)
			printResult(subtracao(a, b))
		case 3:
			var a, b int
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanln(&a)
			fmt.Print("Digite o segundo número: ")
			fmt.Scanln(&b)
			printResult(multiplicacao(a, b))
		case 4:
			var a, b int
			fmt.Print("Digite o primeiro número (dividendo): ")
			fmt.Scanln(&a)
			fmt.Print("Digite o segundo número (divisor): ")
			fmt.Scanln(&b)
			printResult(divisao(a, b))
		case 5:
			fmt.Println("Saindo...")
		default:
			fmt.Println("Opção inválida")
		}
	}
}