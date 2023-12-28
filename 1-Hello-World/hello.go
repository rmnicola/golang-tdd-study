package main

import "fmt"

const prefix_eng = "Hello, "
const prefix_ptbr = "Olá, "

func Hello(name, lang string) string {
  var prefix string
	if lang == "English" || lang == "" {
		if name == "" {
			name = "World"
		}
    prefix = prefix_eng
	}
	if lang == "Portuguese" {
		if name == "" {
			name = "Mundo"
		}
    prefix = prefix_ptbr
	}
  return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
