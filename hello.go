package main // if not main then it will return error: package learn-go/hello is not a main package

import (
	"fmt"
)

const (
	french  = "French"
	spanish = "Spanish"

	prefix_en = "Hello, "
	prefix_sp = "Hola, "
	prefix_fr = "Bonjour, "
)

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = prefix_sp
	case french:
		prefix = prefix_fr
	default:
		prefix = prefix_en
	}
	return
}
