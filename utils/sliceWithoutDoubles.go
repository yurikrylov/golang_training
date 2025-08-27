package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SliceWithoutDoubles() {
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
			return
		}
		mySlice = append(mySlice, n)
	}
	var result []int
	seen := make(map[int]bool)
	for _, num := range mySlice {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	fmt.Println("Срез без дубликатов:", result)
}
