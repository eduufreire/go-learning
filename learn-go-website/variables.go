package main

import "fmt"

// nivel de pacote
var i int

func main() {

	// delcara uma lista de variaveis
	// pode estar a nivbel de pacote ou de funcao
	var c, python, java bool
	fmt.Println(i, c, python, java)

	// a variavel pode ter inicializadores
	// se tiver um inicializador, o tipo pode ser omitido
	var javascript, ruby, teste = true, false, 1
	fmt.Println(javascript, ruby, teste)


	// dentro de uma funcao podemos utilizar da declaracaoywqw curta
	flamengo, sao_paulo, vasco := false, true, false
	fmt.Println(flamengo, sao_paulo, vasco)

	// tipos estaticos, não é possível mudar o tipo
	// int ate int64
	// string, bool
	// byte
	// rune
	// float32 e 64
	// complex64 e 128


	// variaveis declaradas sem valor, terao seu valor incial:
	// 0 par numericos
	// false para bool
	// "" para strings

	// conversao de tipo
	// T(v) t == tipo v == valor
	// covnersao deve ser explicita
	i := 42
	f := float32(i)
	u := uint(f)
	fmt.Println(i, f, u)


	// inferencia de tipos
	// quando declaramos sem especificar o tipo, ele sera definido de acordo com o valor dado diretamente
	j := "assadasdas"
	fmt.Printf("%T\n", j)

	// constantes sao delcamaras com const
	// nao podem ser declaradas com a sintaxe :=
	const Pi = 3.14
	const World = "世界"
	fmt.Println(Pi, World)


}