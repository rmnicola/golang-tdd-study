# Anotações arrays e slices

1. A maneira como se declara arrays pré-inicializados em Go é bem parecida com
   C, mas não parece ter como declarar um array sem definir o seu tamanho. Para
   isso, existem `slices`
2. A maneira como se acessa valores dentro de um array também é similar ao C.
3. O marcador `%v` simplifica a exibição de arrays. Esse marcador exibe a
   variável em seu "formato padrão", o que me leva a crer que deve existir
   alguma forma de implementar algo como o `__str__` de Python.
4. Em Go, existe o `range`, mas o funcionamento dele é mais similar ao
   `enumerate` de Python. Em cada iteração, ele retorna um índice e um valor de
   dentro do iterável
5. Arrays de tamanhos diferentes são tratados como tipos de variáveis
   diferentes. [5]int é completamente incompatível com [4]int.
6. Slices tem praticamente a mesma sintaxe de arrays sem tamanho definido em C.
7. É possível checar a cobertura de testes usando `go test -cover`
8. Assim como em praticamente qualquer outra linguagem, é impossível utilizar
   operadores relacionais com dois arrays ou slices. É possível utilizar o
   `reflect.DeepEqual`, que compara se duas variáveis são exatamente a mesma
   coisa.
9. É possível criar slices pré-inicializados utilizando a função `make`.
10. Utilizando `...`, é possível definir um ou mais parâmetros de mesmo tipo
    como argumentos de uma função.
11. A função built-in `append` age de forma similar ao método de Python. A
    diferença é que ela retorna um slice, em vez de ser um método que atua
    in-place
12. Slices em Go parecem funcionar de forma similar a Python, apenas com uma
    pequena diferença de sintaxe (não se usa , ). Também não existem índices
    negativos.
13. Não usar else na mesma linha em que fecha as chaves do if é um erro de
    sintaxe. Harsh
14. Realmente é possível usar uma declaração de função inline da mesma forma
    que um `lambda` de Python.
15. Go’s arrays are values. An array variable denotes the entire array; it is
    not a pointer to the first array element (as would be the case in C). This
    means that when you assign or pass around an array value you will make a
    copy of its contents. (To avoid the copy you could pass a pointer to the
    array, but then that’s a pointer to an array, not an array.) One way to
    think about arrays is as a sort of struct but with indexed rather than
    named fields: a fixed-size composite value. -> Este trecho do blog post
    sobre slices indica que não existe `pointer decay` em Go, como há em C. No
    entanto, toda vez que passamos um slice como parâmetro vai uma cópia dele.
16. Lendo o [blogpost sobre slices](https://go.dev/blog/slices-intro), confirmo
    que o `make` atua com alocação dinâmica de memória. Sendo assim, seu uso
    estaria diretamente relacionado à performance de aplicações onde um array
    deve "crescer" aos poucos. Sendo assim, a solução usando `append` não é a
    mais eficiente.
17. Sempre que se usa uma operação de slicing, o resultado é um slice. Um slice
    de um array é um... slice. Parece óbvio quando coloco dessa forma...
18. Um slice continua sendo composto por memória contígua. Isso significa que
    quando usamos um append para fazer um slice crescer além da sua capacidade,
    ele deve ser realocado completamente. No entanto, se não usarmos `append`
    para fazer um slice crescer além de sua capacidade, vamos ter um `runtime
    panic`.
19. Quando fazemos uma operação de slicing, o resultado é um ponteiro que
    aponta para o começo do slice e o tamanho do slice. Sendo assim, é uma
    operação bastante eficiente.
20. Utilizar `...` após um slice ou array expande o slice em itens (similar ao
    `*` em Python)
