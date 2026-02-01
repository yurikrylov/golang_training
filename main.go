package main

import "fmt"

func main() {
	x := 10
	f := func() int {
		x *= 2
		return x
	}
	fmt.Print(f(), x)
}
