package utils

import (
	"errors"
)

// Напишите функцию, которая принимает срез целых чисел и меняет их порядок на обратный (in-place, без создания нового среза).
func GeminiQ1(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Удаление дубликатов: Дан срез строк. Верни новый срез, содержащий только уникальные значения.
func Unique(mySlice []int) ([]int, error) {
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

// Слияние: Даны два отсортированных среза. Объедини их в один так, чтобы результат тоже был отсортирован.
func MergeSorted(s1, s2 []int) ([]int, error) {
	if len(s1) == 0 && len(s2) == 0 {
		return nil, errors.New("invalid input")
	}
	result := make([]int, 0, len(s1)+len(s2))
	i, j := 0, 0

	for i < len(s1) && j < len(s2) {
		if s1[i] <= s2[j] {
			result = append(result, s1[i])
			i++
		} else {
			result = append(result, s2[j])
			j++
		}
	}

	result = append(result, s1[i:]...)
	result = append(result, s2[j:]...)

	return result, nil
}

// Фильтрация: Напиши функцию, которая удаляет из среза все отрицательные числа.
func FilterNegative(s []int) []int {
	var result []int
	for _, num := range s {
		if num >= 0 {
			result = append(result, num)
		}
	}
	return result
}

// Chunking: Разбей большой срез на «пакеты» (слайсы слайсов) заданного размера n.
func Chunking(s []int, n int) [][]int {
	if n <= 0 {
		return nil
	}
	var chunks [][]int
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

// Удаление по индексу: Напиши функцию remove(s []int, i int), которая удаляет элемент по индексу, сохраняя порядок.
// Циклический сдвиг: Сдвинь элементы среза вправо на k позиций (например, [1,2,3] при k=1 станет [3,1,2]).
// Поиск пересечения: Найди общие элементы в двух разных срезах.
// Максимальный подмассив: Найди непрерывный подмассив с наибольшей суммой элементов (классическая задача Кадана).
// Безопасное копирование: Создай функцию, которая копирует только первые N элементов из одного среза в другой, учитывая, что исходный срез может быть короче N.
