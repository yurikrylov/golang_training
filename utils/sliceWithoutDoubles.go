package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SliceWithoutDoubles() ([]int, error) {
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
	var result []int
	seen := make(map[int]struct{})
	for _, num := range mySlice {
		if _, ok := seen[num]; !ok {
			seen[num] = struct{}{}
			result = append(result, num)
		}
	}
	return result, nil
}
