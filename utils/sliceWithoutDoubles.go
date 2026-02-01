package utils

import "errors"

/*
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)*/

func SliceWithoutDoubles(mySlice []int) ([]int, error) {
	/* Ввод от пользователя
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите целые числа, разделенные пробелами:")

	scanner.Scan()
	input := scanner.Text()

	fields := strings.Fields(input)

	 var mySlice []int
	for _, field := range fields {
		n, err := strconv.Atoi(field)
		if err != nil {
			fmt.Println("Ошибка преобразования:", field)
			return nil, errors.New("invalid input")
		}
		mySlice = append(mySlice, n)
	}
	*/
	if len(mySlice) == 0 {
		return nil, errors.New("не может быть пустым")
	}
	var result []int
	seen := make(map[int]struct{})
	for _, num := range mySlice {
		if num < 0 {
			return nil, errors.New("отрицательное число")
		}
		if _, ok := seen[num]; !ok {
			seen[num] = struct{}{}
			result = append(result, num)
		}
	}
	return result, nil
}
