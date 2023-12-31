# Anotações Iteration

1. Em Go não há while, do ou until. Só for.
2. O for de Go é quase igual ao de C, a diferença é que pode-se utilizar o
   walrus operator para deixá-lo mais sucinto. Noice.
3. Confirmado que o walrus operator cria e inicializa uma variável com um valor
   da sua escolha.
4. Go também tem operadores de atribuição (`e.g. +=`)
5. Quando se usa o for, é obrigatório a abertura e fechamento de chaves
6. Outro tipo de teste especial é o `Benchmark`. Com ele, é possível executar a
   função `b.N` vezes e obter um valor médio de tempo de execução. Para rodar
   um benchmark, use: `go test -bench=.`
