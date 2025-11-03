package main

import (
	"10/pkg/keyboard"
	"fmt"
	"log"
)

func main() {
	fmt.Print("실수 입력:")
	n, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.1f\n", n)
}
