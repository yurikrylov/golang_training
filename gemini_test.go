package main

import (
	"reflect"
	"testing"

	"github.com/yurikrylov/golang_training/utils"
)

func TestGeminiQ2(t *testing.T) {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	result, _ := utils.MergeSorted(a, b)
	expected := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeSorted returned %v, want %v", result, expected)
	}
}
