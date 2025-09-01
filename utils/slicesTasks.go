package utils

import (
	"bufio"
	"fmt"
	"os"
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
