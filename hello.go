package main // if not main then it will return error: package learn-go/hello is not a main package

import (
	"fmt"
)

const prefix_en = "Hello, "
const prefix_sp = "Hola, "
const prefix_fr = "Bonjour, "

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return prefix_sp + name
	}
	if language == "French" {
		return prefix_fr + name
	}
	return prefix_en + name
}
