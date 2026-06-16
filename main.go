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
		return []int{}
	}

	slice := make([]int, 0, size)
	src := rand.NewSource(time.Now().Unix())

	for i := 0; i < size; i++ {
		randNum := src.Int63()
		slice = append(slice, int(randNum))
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
	result := make([]int, 0, CHUNKS)
	chunkSize := len(data) / CHUNKS
	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	wg.Add(CHUNKS)
	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		slice := data[start:end]

		go func(slice []int) {
			maxNum := maximum(slice)
			mu.Lock()
			result = append(result, maxNum)
			mu.Unlock()
			wg.Done()
		}(slice)
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
