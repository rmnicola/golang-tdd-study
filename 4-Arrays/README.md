# Anotações arrays e slices

1. A maneira como se declara arrays pré-inicializados em Go é bem parecida com
   C, mas não parece ter como declarar um array sem definir o seu tamanho.
2. A maneira como se acessa valores dentro de um array também é similar ao C.
3. O marcador `%v` simplifica a exibição de arrays. Esse marcador exibe a
   variável em seu "formato padrão", o que me leva a crer que deve existir
   alguma forma de implementar algo como o `__str__` de Python.
4. Em Go, existe o `range`, mas o funcionamento dele é mais similar ao
   `enumerate` de Python. Em cada iteração, ele retorna um índice e um valor de
   dentro do iterável
