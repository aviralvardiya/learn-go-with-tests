package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const hindiHelloPrefix = "Namaste, "
const spanish = "spanish"
const french = "french"
const hindi = "hindi"

func Hello(name string, language string) string {
	prefix := englishHelloPrefix
	if name == "" {
		name = "world"
	}

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Prakash", "hindi"))
	// fmt.Println("Jai suryadev")
}
