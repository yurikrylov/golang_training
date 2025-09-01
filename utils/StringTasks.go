package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Напишите приложение, где пользователь вводит строку и букву,
наличие которой предстоит проверить в введенной строке.
Выведите в терминал полученный результат в терминал.
*/
func StringTask1() bool {
	var str, l string
	fmt.Scanf("%s %s", &str, &l)
	result := strings.Contains(str, l)

	return result
}

/*
Напишите приложение, где пользователь вводит строку и букву.
Выведите в терминал длину строки, также индекс первого и последнего вхождения буквы в строку.
*/
func StringTask2() {
	var str, l string
	r := bufio.NewReader(os.Stdin)
	str, _ = r.ReadString('\n')
	str = strings.TrimSpace(str)
	l, _ = r.ReadString('\n')
	l = strings.TrimSpace(l)

	fmt.Print(len(l), strings.Index(str, l), strings.LastIndex(str, l))
}

/*
Напишите приложение, где пользователь вводит строку и букву.
Посчитайте сколько раз заданная буква входит в строку и выведите полученный результат,
а также индекс первого вхождения буквы в строку, в терминал.
*/
func StringTask3() {
	var s, l string
	r := bufio.NewReader(os.Stdin)
	s, _ = r.ReadString('\n')
	s = strings.TrimSpace(s)
	l, _ = r.ReadString('\n')
	l = strings.TrimSpace(l)

	fmt.Print(strings.Count(s, l), strings.Index(s, l))
}

/*
Напишите приложение, где пользователь вводит строку и два символа (например a и b).
Замените в строке все символы «a» на «b» и выведите полученный результат в терминал.
*/
func StringTask4() {
	var s, a, b string
	r := bufio.NewReader(os.Stdin)
	s, _ = r.ReadString('\n')
	s = strings.TrimSpace(s)
	a, _ = r.ReadString('\n')
	a = strings.TrimSpace(a)
	b, _ = r.ReadString('\n')
	b = strings.TrimSpace(b)

	fmt.Print(strings.ReplaceAll(s, a, b))
}

/*
Напишите приложение, где пользователь вводит слово и на его основе создается новая переменная,
сформированная из первого, среднего и последнего символов введенного слова.
Полученный результат выведите в терминал. Например: «Привет!» -> «Пв!»

p.s. для расчета среднего символа строки используйте формулу length / 2
*/
func StringTask5() {
	var s string
	fmt.Scanln(&s)
	s = strings.TrimSpace(s)
	l := len(s)
	midIndex := l / 2
	firstChar := string(s[0])
	midChar := string(s[midIndex])
	lastChar := string(s[l-1])
	result := firstChar + midChar + lastChar
	fmt.Print(result)
}

/*
Напишите приложение, где пользователь вводит слово и на его основе создается новая переменная,
сформированная из трех средних символов.
Полученный результат выведите в терминал. Например: «МамаМылаРаму» -> «ыла».
*/
func StringTask6() {
	var s string
	fmt.Scanln(&s)
	s = strings.TrimSpace(s)
	l := len(s)
	midIndex := l / 2
	start := midIndex - 1
	result := s[start : start+3]
	fmt.Print(result)
}

/*
Напишите приложение, где пользователь вводит две строки str1 и str2.
Программа должна создать новую строку str3 путем добавления str2 в середину str1.
Полученный результат выведите в терминал. Например: str1 = «Мама», str2 = «Раму» -> «МаРамума».
*/
func StringTask7() {
	var a, b, result string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)

	l := len(a)
	midIndex := l / 2
	result = a[:midIndex] + b + a[midIndex:]
	fmt.Print(result)
}

/*
Напишите приложение, где пользователь вводит слово и на его основе формирует новая переменная путем удаления символов из первой строки
(с нулевого элемента по 3-й).
Полученный результат выведите в терминал. Например: «МамаМылаРаму» -> «МылаРаму».
*/
func StringTask8() {
	var a string
	fmt.Scanln(&a)
	a = strings.TrimSpace(a)

	fmt.Print(a[4:])
}

/*
Напишите приложение, где пользователь вводит две строки str1 и str2.
Программа должна создать новую строку str3 состоящую из первого, среднего и последнего символов строк str1 и str2.
Полученный результат выведите в терминал. Например: str1 = «Мама», str2 = «Утром» -> «МУмрам».
*/
func StringTask9() {
	var a, b, result string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)

	midIndexA := len(a) / 2
	midIndexB := len(a) / 2

	result = string(a[0]) + string(b[0]) +
		string(a[midIndexA]) + string(b[midIndexB]) +
		string(a[len(a)-1]) + string(b[len(b)-1])
	fmt.Print(result)
}

/*
Напишите приложение, где пользователь вводит строку.
Используя пробелы, разбейте ее на части, сформировав срез. Полученный срез и его размер выведите в терминал.
*/
func StringTask10() {
	r := bufio.NewReader(os.Stdin)
	i, _ := r.ReadString('\n')
	i = strings.TrimSpace(i)
	w := strings.Fields(i)
	fmt.Println(w)
	fmt.Println(len(w))

}
