package main

import (
	"reflect"
	"testing"

	"github.com/yurikrylov/golang_training/utils"
)

func TestMergeSorted(t *testing.T) {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	result, _ := utils.MergeSorted(a, b)
	expected := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeSorted returned %v, want %v", result, expected)
	}
}

// тесовая функция для filternegatives
func TestFilterNegatives(t *testing.T) {
	input := []int{-5, 3, -1, 4, -2, 0, 6}
	expected := []int{3, 4, 0, 6}
	result := utils.FilterNegative(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FilterNegative returned %v, want %v", result, expected)
	}
}

func TestGeminiQ1(t *testing.T) {
	// Описываем таблицу тест-кейсов
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Обычный срез",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Четное количество элементов",
			input:    []int{10, 20},
			expected: []int{20, 10},
		},
		{
			name:     "Один элемент",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Пустой срез",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Срез nil",
			input:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Вызываем твою функцию (она меняет срез на месте)
			utils.GeminiQ1(tt.input)

			// Сравниваем результат с ожидаемым
			// Срезы нельзя сравнивать через ==, поэтому используем reflect.DeepEqual
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("Ошибка в тесте '%s': получили %v, хотели %v", tt.name, tt.input, tt.expected)
			}
		})
	}
}
