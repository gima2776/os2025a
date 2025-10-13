// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	var length float64 = 1.2
// 	var width int = 2
// 	fmt.Println("면적은", int(length)*width)
// 	fmt.Println("길이 > 너비?", int(length) > width)
// 	fmt.Println("형변환", reflect.TypeOf(int(length)))
// 	fmt.Println("원본", reflect.TypeOf(int(length)))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"time"
// )

// func main() {
// 	var now time.Time = time.Now()
// 	var day int = now.Day()
// 	fmt.Println(day)
// 	univ := "Go$ Inha$"
// 	changer := strings.NewReplacer("$", "!")
// 	changed := changer.Replace(univ)
// 	fmt.Println(changed)
// }

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var month time.Month = now.Month() // month := now.Month()
	fmt.Println(month)

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	// fmt.Println(err)
	log.Fatal(err) //report the error and exit the program
	fmt.Println(i)
}
