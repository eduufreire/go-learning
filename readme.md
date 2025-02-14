linguagem procedural criada para resolver questoes de performance e complexidade
go combina a facilidade da programacao de uma linguagem interpretada e dinamincamente tipada
com a eficiencia e seguranca de uma lnugagem estaticameante tipada e compilada

- consciso, explicito e facil de trabalhar
- projetado com compilacao rapida
- remove complexidades em dependencias, compilacao, segurança e manutencao
- segura porque é fortemente tipada e estaticamente tipada
    - sem conversoes de tipos implicitas
    - precisamos ser explicitos na tipagem
- coleta o lixo pra voce e integra o sistema ao seu binario executavel
- 


dúvidas:
o que concorrencia na prograemacao?
o que é o garbage collecotor?
diferença entre tipagem estatica e dinamica | linguagem interpretada e compilada


- standard library
- multiplataforma
- garbade collection
- cross-compile

- facil de usar
- compilada, tipada e estatica

quando usar
- escalabilidae
- serviços de web, redes, servers
- quando precisar de uma lingaugem rapida, simples e facil

declaracao/statement = menor parte de um programa que indica o que ele deve fazer
expressao = conjunto de valores/ações que geram um outro valor

iota = usaod em declaracoes const para definicao de numeros incrementais
declaracao de variaveis na mesma linha sao chamada de atribuicoes paralelas
ex: var a,b,c = 1, 2, true

blank identifier:
- permite abandonar uma variavel
- somente de gravacao, nao podemos obter o valor
- em GO é obrigatorio usarmos todas as variaveis declaradas, mas as vezes nao necessitados de todos os valores retornados
de uma funcao, entao usados o blank identiier
- evita erro de compilacao por nao usar uma variavel

escopo define onde uma variavel pode ser usada
- variavel local pode ser usada somente dentro e funcoes/bloco de codigo
- variaveis globais pode ser usada dentro do pacote e em pacotes externos quando exportada
- variavel global pode ter o mesmo nome que a local, mas dento da funcao, a local tera prioridade

arrays
- sequence indexada de elementos com tamanha especifico
- ou seja, é precisa deixar explicito o tamanho do array para alocação de memoria no momento da compilacao
- é mais comum o uso de slices, arrays sao usados em cenarios especificos

slices:
- sao mais comuns que arrays
- disponibiliza uma interface melhorada do array
- 

deslocamento de bits:
- deslocar digitos binaarios para a esquerda ou direita