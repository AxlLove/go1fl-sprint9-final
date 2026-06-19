package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{0, -1} {
		t.Run(fmt.Sprintf("size=%d", size), func(t *testing.T) {
			assert.Empty(t, generateRandomElements(size))
		})
	}

	slice := generateRandomElements(10)
	require.Len(t, slice, 10)

	time.Sleep(1 * time.Second)

	newSlice := generateRandomElements(10)
	assert.NotEqual(t, slice, newSlice)
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"nil слайс", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"максимум в начале", []int{600, 1, 253}, 600},
		{"максимум в середине", []int{1, 600, 253}, 600},
		{"максимум в конце", []int{1, 253, 600}, 600},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, maximum(tt.input))
		})
	}
}
