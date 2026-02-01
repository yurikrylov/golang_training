package main

import (
	"reflect"
	"strings"
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
			utils.GeminiQ1(tt.input)

			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("Ошибка в тесте '%s': получили %v, хотели %v", tt.name, tt.input, tt.expected)
			}
		})
	}
}

func TestSliceWithoutDoubles(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		expected    []int
		wantErr     bool   // ради чего писался этот тест
		errContains string // Подстрока, которая должна быть в тексте ошибки
	}{
		{
			name:     "Успешное удаление дубликатов",
			input:    []int{1, 2, 2, 3, 1},
			expected: []int{1, 2, 3},
			wantErr:  false,
		},
		{
			name:        "Ошибка: пустой срез",
			input:       []int{},
			expected:    nil,
			wantErr:     true,
			errContains: "не может быть пустым",
		},
		{
			name:        "Ошибка: отрицательное число",
			input:       []int{1, -5, 2},
			expected:    nil,
			wantErr:     true,
			errContains: "отрицательное число",
		},
		{
			name:     "Срез без дубликатов",
			input:    []int{10, 20, 30},
			expected: []int{10, 20, 30},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := utils.SliceWithoutDoubles(tt.input)

			// Проверка наличия/отсутствия ошибки
			if (err != nil) != tt.wantErr {
				t.Fatalf("SliceWithoutDoubles() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Если ошибка ожидалась, проверяем её текст
			if tt.wantErr && !strings.Contains(err.Error(), tt.errContains) {
				t.Errorf("Текст ошибки '%v' не содержит '%s'", err, tt.errContains)
			}

			// Если ошибки нет, проверяем результат
			if !tt.wantErr && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Получили %v, хотели %v", result, tt.expected)
			}
		})
	}
}
