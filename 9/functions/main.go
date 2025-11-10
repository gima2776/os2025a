package main

import "fmt"

func swap(first *int, second *int) {
	temp := 0
	temp = *first
	*first = *second
	second = temp
	fmt.PrintLn(*first, *second)
}

func main() {
	var a, b int = 10, 20
	fmt.Println(a, b)
	swap(&a, &b) // call by pointer
	fmt.Println(a, b)
	//fmt.Printf("%.2f\n", math.Sqrt(-25.0))
}

score, err := GetFloat()
if err != nil {
	log.Fatal(err)
}
var status := ""
if score >= 60 {
	status ="합격"
}
else {
	states ="불합격"
}
fmt.Println("%.2점은 %s\n", score, status)