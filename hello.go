package main // if not main then it will return error: package learn-go/hello is not a main package

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

func Hello(name string) string {
	return "Hello " + name
}
