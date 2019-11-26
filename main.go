package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name //prints hello, world
}
func main() {
	fmt.Println(Hello("world"))
}
