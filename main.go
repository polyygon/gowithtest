package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name //prints hello, world
}

func main() {
	fmt.Println(Hello("Corey"))
}
