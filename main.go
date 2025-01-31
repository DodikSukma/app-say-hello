package main

import (
	"fmt"

	go_say_hello "github.com/DodikSukma/go-say-hello/v2" // Import the package to make it available in this package scope. This is necessary because Go doesn't allow importing packages with only init functions.
)

func main() {
	fmt.Println(go_say_hello.SayHello("Dodik Bagus Genjing"))
	fmt.Println(go_say_hello.SayHello("Celia Cantik Banget"))
}
