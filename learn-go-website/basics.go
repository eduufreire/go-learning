// programas começam pelo package main
// nome do pacote é o ultimo elemento do caminho de importacao
// ex: import "math/rand" === rand.funcionalide()

package main

// quando as importacoes ficam dentro dos parenteses
// isso se chama importacao consignada
import (
	"fmt"
	"math"
	"math/rand"
)

// o tipo dos paramentros vem depois do nome
// podemos encurtar a declaracao dos parametros
// ao inves de ser: func add(x int, y int)
// podemos fazer: func add(x, y int)
func add(x, y int) int {
	return x + y
}

// uma função pode retornar qualquer numero de resultados
func swap(x, y string) (string, string) {
	return y, x
}


// valores de retorno podem ser nomeados e agirem apenas como variaveis
// ajudam a documentar o significado do retorno
// recomendado apenas para funcoes curtas pois pode prejudicar a legibilidade
// isso se chama retorno limpo
// return sem argumentops irpá retornar os valores atuais dos resultados
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}



func main() {
	fmt.Println("Número aleatorio: ", rand.Intn(10))

	// nome exportados começam com letra maiuscula
	// nomes nao exportados nao sao acessiveis fora do pacote
	fmt.Println("PI:", math.Pi)

	fmt.Println("Soma: ", add(10, 5))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

