// main.go
package main

import "fmt"

func main() {
	fmt.Println(getMessage())
}

func getMessage() string {
	return "Hello, Jenkins Pipeline with Golang!"
}