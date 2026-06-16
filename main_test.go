package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	slice := generateRandomElements(0)
	assert.Empty(t, slice)

	slice = generateRandomElements(-1)
	assert.Empty(t, slice)

	slice = generateRandomElements(10)
	require.Len(t, slice, 10)

	time.Sleep(1 * time.Second)

	newSlice := generateRandomElements(10)
	assert.NotEqual(t, slice, newSlice)
}

func TestMaximum(t *testing.T) {
	var slice []int
	maxNum := maximum(slice)
	require.Equal(t, 0, maxNum)

	slice = []int{
		1,
		253,
		32,
		600,
		250,
		599,
	}

	maxNum = maximum(slice)

	assert.Equal(t, 600, maxNum)

	slice = []int{1}
	maxNum = maximum(slice)
	assert.Equal(t, 1, maxNum)
}
