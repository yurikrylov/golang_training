package main

import "fmt"

// import "github.com/yurikrylov/golang_training/utils"

func main() {

	b := []byte{'G', 'o'}
	s := string(b)
	b[0] = 'g'
	fmt.Println(s)

}

/*
	var order = []string{"хлеб", "буженина", "сыр", "огурцы"}
	utils.CalcOrder(order)
	utils.FizzBuzz()
	utils.Switch()
*/
