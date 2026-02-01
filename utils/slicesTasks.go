package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Напишите приложение, позволяющее пользователю вводить срез вещественных значений.
Выведите в терминал его размер, значения первого и последнего элемента.
*/
func SliceTask1() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	i := strings.TrimSpace(s.Text())
	parts := strings.Fields(i)
	v := make([]string, 0, len(parts))
	v = append(v, parts...)
	if len(v) == 0 {
		fmt.Print("not today")
	}
	fmt.Printf("%d %s %s\n", len(v), v[0], v[len(v)-1])

}

/*
Напишите приложение, позволяющее пользователю вводить срез целочисленных значений и число A,
на которое необходимо увеличить значения элементов среза, после чего добавить A в конец среза.
Выведите в терминал полученный результат.
*/
func SliceTask2() {
	r := bufio.NewReader(os.Stdin)
	str1, _ := r.ReadString('\n')
	str1 = strings.TrimSpace(str1)
	arr := strings.Fields(str1)
	slice := make([]int, 0, len(arr))
	for _, v := range arr {
		e, _ := strconv.Atoi(v)
		slice = append(slice, e)
	}

	str2, _ := r.ReadString('\n')
	str2 = strings.TrimSpace(str2)

	a, _ := strconv.Atoi(str2)

	for i := range slice {
		slice[i] += a
	}
	slice = append(slice, a)
	fmt.Print(slice)

}
