package utils

import "fmt"

func Switch() {
	var year int
	_, err := fmt.Scanf("%d", &year)
	if err != nil {
		fmt.Println("ошибка")
	}
	switch {
	case year > 1946 && year < 1964:
		fmt.Println("Привет, бумер!")
	case year > 1946 && year < 1964:
		fmt.Println("Привет, бумер!")

	case year > 1965 && year < 1980:
		fmt.Println("Привет, представитель X!")

	case year > 1981 && year < 1996:
		fmt.Println("Привет, миллениал!")

	case year > 1997 && year < 2012:
		fmt.Println("Привет, зумер!")

	case year > 2012:
		fmt.Println("Привет, альфа!")
	default:
		fmt.Println("кажется это не число!")

	}
}
