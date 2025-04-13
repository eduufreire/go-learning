linguagem procedural criada para resolver questoes de performance e complexidade
go combina a facilidade da programacao de uma linguagem interpretada e dinamincamente tipada
com a eficiencia e seguranca de uma lnugagem estaticameante tipada e compilada

projetada para ser eficiente, simples, segura e escalável
possui alto desempenho, suporet para programacao concorrente e escalabilidade
possui recursos integrados para lidar com taregas concorrentes

linguagem compilada, ou seja, o codigo é transformada em um executavel
- binario estatico

vantagens do go:
- desempenho: compilada código em linguagem de máquina altamente otimizado,
- simplicidade e legibilidade: sintaxe limpa e concisa, evitando recursos descenessarios
- suporte integrado para concorrencia: permite execuções pararelas atraves dos channels e goroutines
- gerenciamento automatico de memoria: possui garbage collector, nao precisando se preocupar com alocacao/desalocacao de memoria
- bibiotecas padroes: diversas bibilotescas nativas, sem dependencia de bibliotesas de terceitos
- cross-plataforma: codigo complicado pode ser executado em qualquer SO

a linguagem go segue o paradigma imperativo:
- estilo que enfatiza a squencia de instrucoes para manipular o estado do programa


- consciso, explicito e facil de trabalhar
- projetado com compilacao rapida
- remove complexidades em dependencias, compilacao, segurança e manutencao
- fortemente tipada e estaticamente tipada
    - sem conversoes de tipos implicitas
    - precisamos ser explicitos na tipagem
- coleta o lixo pra voce e integra o sistema ao seu binario executavel
- lida bem com concorrencia


dúvidas:
o que concorrencia na prograemacao?
o que é o garbage collecotor?
diferença entre tipagem estatica e dinamica | linguagem interpretada e compilada
o que é uma closure?


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
- usados para armazenar multiplos valores do mesmo tipo
- comprimento pode aumentar ou diminuir  em tempo de execucao 
- são passados por referencias para funcoes

deslocamento de bits:
- deslocar digitos binaarios para a esquerda ou direita

maps
- estrutura de dados de chave-valor
- pode ser util para melhorar a eficiencia ao manipular colecoes de dados
- pesquisar dados rapidamente
- 

modulos
- go mod
- se começar com letra maiuscula, quer dizer que pode ser exportada
- com letra minuscula, ele nao sera exportada e nao pode ser acessada fora do módulo
- privado = letra minuscula
- publico = letra maiuscula


ponteiros
- 

concorrencia é sobre lidar com mutias coisas de uma só vez
- capacidade de executar diversas tarefas sem interromper a execucao por completo
- é um conceito, parece que está sendo executado tudo ao mesmo tempo, mas nao
- exemplos:
    - programar, pegar a xicara e tomar cafe
    - estamos lidando com duas coisas mas nao ao mesmo tempo
    - temos que parar a execucao de uma para iniciar a outra

paralelismo é fazer varias coisas ao mesmo tempo
- execucao real de varias tarefas ao mesmo tempo
- ex: temos 100 imagens e temos 4 computadores. podemos distribuir 25 imagens para cada pc executar em paralelo
- outro exemplo:
    - programar e ouvir musica, estamos fazendo ao mesmo tempo

go routines
- tem o objetivo de permitir que funcoes sejam executadas concorrentemented de maneira facil
- são gerenciadas pelo runtime do Go
- sao mais leves que aas threads convencionais, podendo haver milhares de goroutines em um thread
- rodam no mesmo espaço de memoria, podendo compartilhar objetos
- é preciso usar channels para sincronizar e evitam race condition
- existem apenas no espaço virtual da runtime


channels:
- serve para comunicacao e sincronizacao de processos concorrentes
- maneira de compartilhar valores entre goroutines
- forma mais segura de comunicacao
- só podem ser fechados pelo remetente

buffered channels:
- nao bloqueiam o envio e recebimento em alguns cenarios
- guardam o valor enviado temporariamente em um buffer, que sera lido e removido quando alguem receber
- caso um buffer esteja cheio, a goroutine é bloqueado até que haja um espaço

diferença entre channels e buffered channels:
- channel: 
    - o remetente é bloqueado até que um receptor esteja pronto para receber
    - o receptor tambem é bloqueado ate que os dados estejam disponiveis
    - utliza-se quando a sincronizacao é crucial
- buffered channels
    - armazena um determinado numero de valores, nao bloquando o envio
    - o receptor é bloqueado apenas quando o buffer estiver vazio
    - podemos usar quando a sincronizacao nao é tao critica
