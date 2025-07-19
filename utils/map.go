package utils

import "fmt"

func CalcOrder(order []string) {
	var priceList = make(map[string]int, 0)
	priceList["хлеб"] = 50
	priceList["молоко"] = 100
	priceList["масло"] = 200
	priceList["колбаса"] = 500
	priceList["соль"] = 20
	priceList["огурцы"] = 200
	priceList["сыр"] = 600
	priceList["ветчина"] = 700
	priceList["буженина"] = 900
	priceList["помидоры"] = 250
	priceList["рыба"] = 300
	priceList["хамон"] = 1500

	result := 0

	for k := range priceList {
		if priceList[k] > 500 {
			fmt.Println(k)

		}
	}
	for _, v := range order {
		result = result + priceList[v]
	}
	fmt.Println(result)
}
