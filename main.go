package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0
	for {
		counter++
		fmt.Println("counter is : ", counter)
		time.Sleep(time.Second * 5)
	}
	msg := printHi("Abhishek.")
	fmt.Println(msg)
}

func printHi(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
