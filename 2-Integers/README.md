# Anotações Integers

[Dicas de como formatar projetos em
Go](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)

1. O formatador `%q`, quando usado com strings, mostra a string em hexadecimal.
   Formato `\x0f`. `%d` mostra o valor em decimal
2. É possível fazer com que o `go test` seja mais verboso com `-v`. Prefiro
   assim.
3. Um tipo de testes especial é o `Example`. Ele entra em um arquivo de teste
   comum de Go, só que as funções começam com `Example`. Ele serve
   primariamente como documentação, mas serve como teste também. Ele
   automaticamente aparece no `godoc` e seu output é testado quando há um
   comentário tipo `// Output: 7`. Me parece uma ótima ideia uma
   documentação baseada em testes. 
