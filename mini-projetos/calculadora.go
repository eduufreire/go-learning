package main

import (
	"fmt"
)

func sum(a, b int64) int64 {
	return a + b
}

func multiply(a, b int64) int64 {
	return a * b
}

func divide(a, b int64) int64 {
	return a / b
}

func subtract(a, b int64) int64 {
	return a - b
}

func main() {

	var option int64 = 1
	var numberOne, numberTwo, result int64
	operationsHistory := make([]string, 1)

	for option != 0 {

		fmt.Println(`
Escolha uma opção:
1 - Somar
2 - Subtrair
3 - Dividir
4 - Multiplicar
5 - Mostrar operações anteriores
0 - Sair
		`)
		fmt.Scan(&option)

		if option != 5 {
			fmt.Println("Digite o primeiro número INTEIRO:")
			fmt.Scan(&numberOne)

			fmt.Println("Digite o segundo número INTEIRO:")
			fmt.Scan(&numberTwo)
		}

		switch option {
		case 0:
			fmt.Println("Calculadora encerrada")
			break
		case 1:
			result = sum(numberOne, numberTwo)
			s := fmt.Sprintf("%d + %d = %d", numberOne, numberTwo, result)
			fmt.Println(s)
			operationsHistory = append(operationsHistory, s)
		case 2:
			result = subtract(numberOne, numberTwo)
			s := fmt.Sprintf("%d - %d = %d", numberOne, numberTwo, result)
			fmt.Println(s)
			operationsHistory = append(operationsHistory, s)

		case 3:
			result = divide(numberOne, numberTwo)
			s := fmt.Sprintf("%d / %d = %d", numberOne, numberTwo, result)
			fmt.Println(s)
			operationsHistory = append(operationsHistory, s)

		case 4:
			result = multiply(numberOne, numberTwo)
			s := fmt.Sprintf("%d * %d = %d", numberOne, numberTwo, result)
			fmt.Println(s)
			operationsHistory = append(operationsHistory, s)

		case 5:
			fmt.Println("Operações:")
			for i := range(len(operationsHistory)) {
				fmt.Println(operationsHistory[i])
			}

		default:
			fmt.Print("Operação inválida")
		}

		fmt.Print("\n")

		fmt.Println(`
Deseja realizar outra operação?
1 - Sim
0 - Não
		`)
		fmt.Scan(&option)

	}

}
