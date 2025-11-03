package main

import (
	"10/pkg/greeting"
	"10/pkg/greeting/deutsch"
	"10/pkg/greeting/korean"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	korean.Anyunghaseyo()
	korean.Anyung()
	deutsch.GutenTag()
}
