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
