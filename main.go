package main

import (
	"fmt"

	"github.com/yurikrylov/golang_training/utils"
)

//import "fmt"

//"github.com/yurikrylov/golang_training/utils"

func main() {
	x := 10
	f := func() int {
		x *= 2
		return x
	}
	fmt.Print(f(), x)
	utils.TestGeminiQ2()
}

/*
	var order = []string{"хлеб", "буженина", "сыр", "огурцы"}
	utils.CalcOrder(order)
	utils.FizzBuzz()
	utils.Switch()
*/
