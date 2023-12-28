package main

import "fmt"

const en_lang = "English"
const en_prefix = "Hello, "
const en_world = "World"
const ptbr_lang = "Portuguese"
const ptbr_prefix = "Olá, "
const ptbr_world = "Mundo"
const fr_lang = "French"
const fr_prefix = "Bonjour, "
const fr_world = "Monde"

func Hello(name, lang string) string {
	var prefix string
	switch lang {
	case ptbr_lang:
		if name == "" {
			name = ptbr_world
		}
		prefix = ptbr_prefix
	case fr_lang:
		if name == "" {
			name = fr_world
		}
		prefix = fr_prefix
	default:
		if name == "" {
			name = en_world
		}
		prefix = en_prefix

	}
	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
