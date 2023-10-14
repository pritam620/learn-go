package main // if not main then it will return error: package learn-go/hello is not a main package

import "fmt"

const prefix_en = "Hello, "

func main() {
	fmt.Println(Hello("world"))
}

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return prefix_en + name
}
