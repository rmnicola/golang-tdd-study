# Anotações Hello World

1. Go não vai testar nada sem criar um módulo antes. `go mod init
   <nome-do-modulo>`
2. Go tem uma ferramenta de formatação de código `go fmt`. Ela funciona no
   projeto todo e é capaz de identificar erros de sintaxe antes de rodar o
   código. É interessante deixar um remap pronto para usá-la. Por enquanto
   estou definindo localmente como <leader>f. A ideia é MUITO boa. Se a
   linguagem já traz uma ferramenta de formatação tão opinionada como o `go
   fmt`, não existe mais discussão de coding style.
3. Testes em Go tem algumas regras específicas:
    * Precisa estar em arquivos com o padrão `XXX_test.go`
    * As funções de teste precisam começar com `Test`
    * Todas as funções de teste obrigatoriamente usam esse tipo `*testing.T`.
      Isso é um ponteiro? O tutorial se referia a ele apenas como um `hook`.
      Vamos ver...
4. Apesar de ter tipagem forte, é possível fazer atribuições com inferência de
   tipo usando o walrus operator `want := "Hello"`. Em meus testes, notei que a
   sintaxe `want = "Hello"` é inválida, mas `var want = "Hello"` e `var want
   string = "Hello"` são válidas. Por enquanto prefiro o walrus operator.
5. Usando esse `hook` de teste eu consigo ao mesmo tempo acusar a falha do
   teste e exibir uma mensagem de erro usando `t.Errorf`
6. Go usa uma estrutura de formatação de strings similar ao C. Ver
   [doc](https://pkg.go.dev/fmt#hdr-Printing). O `%q` exibe um valor entre
   aspas.
7. Os tipos das variáveis vem sempre depois do nome. Não sei se gostei disso
   ainda...
8. É possível criar subtestes dentro de uma função de teste usando a estrutura
   `t.Run`, em que um dos argumentos é um `callable` (a função de subteste).
   Isso me diz que Go permite com facilidade passar funções como argumentos.
9. Em Go existe algum mecanismo para a criação de interfaces. O `testing.TB` é
   uma interface que aceita tanto `testing.T` como `testing.B`. Noice.
10. O `t.Helper()` é só uma função que indica para o módulo de testes que a
    função em que o `Helper` foi chamado não é uma função de teste em si, mas
    apenas uma rotina de auxílio. Sem esse helper, quando o código falha dentro
    de uma função auxiliar, é mais difícil pro dev saber em qual lugar a falha
    aconteceu. Noice2.
11. Os operadores lógicos de Go são iguais a C? Até agora tudo indica que sim
12. Não é possível inicializar uma variável vazia sem definir o seu tipo. Todas
    as variáveis inicializadas começam com valores nulos (nunca lixo).
13. Quando usa o `switch`, não precisa de um `default`.
