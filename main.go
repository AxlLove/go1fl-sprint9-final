package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	slice := make([]int, 0, size)
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)

	for i := 0; i < size; i++ {
		randNum := r.Int()
		slice = append(slice, randNum)
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxNum := data[0]
	for _, v := range data[1:] {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}

// // maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) < CHUNKS {
		return maximum(data)
	}
	result := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	var wg sync.WaitGroup

	wg.Add(CHUNKS)
	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		slice := data[start:end]

		go func(slice []int, i int) {
			maxNum := maximum(slice)
			result[i] = maxNum
			wg.Done()
		}(slice, i)
	}
	wg.Wait()
	return maximum(result)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	slice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxNum := maximum(slice)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxNum, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	start = time.Now()
	maxNum = maxChunks(slice)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxNum, elapsed)
}
