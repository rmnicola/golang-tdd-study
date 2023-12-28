package main

import "fmt"

const (
	en_lang   = "English"
	en_prefix = "Hello, "
	en_world  = "World"

	ptbr_lang   = "Portuguese"
	ptbr_prefix = "Olá, "
	ptbr_world  = "Mundo"

	fr_lang   = "French"
	fr_prefix = "Bonjour, "
	fr_world  = "Monde"
)

func Hello(name, lang string) string {
	var prefix string
	if name == "" {
		prefix, name = grabMultiLanguageHello(lang)
	} else {
		prefix, _ = grabMultiLanguageHello(lang)
	}
	return prefix + name + "!"
}

func grabMultiLanguageHello(lang string) (prefix string, name string) {
	switch lang {
	case ptbr_lang:
		name = ptbr_world
		prefix = ptbr_prefix
	case fr_lang:
		name = fr_world
		prefix = fr_prefix
	default:
		name = en_world
		prefix = en_prefix
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
