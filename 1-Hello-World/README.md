# Anotações Hello World

1. Go não vai testar nada sem criar um módulo antes. `go mod init
   <nome-do-modulo>`
2. Go tem uma ferramenta de formatação de código `go fmt`. Ela funciona no
   projeto todo. É interessante deixar um remap pronto para usá-la. Por
   enquanto estou definindo localmente como <leader>f
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
