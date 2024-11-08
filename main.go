package main

import "fmt"

func main() {
	msg := printHi("Abhishek.")
	fmt.Println(msg)
	fmt.Println("11")
}

func printHi(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
