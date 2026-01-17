package main

import (
	"errors"
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
func generateRandomElements(size int) ([]int, error) {
	var elements []int

	if size == 0 {
		return elements, errors.New("empty size")
	}

	if size < 0 {
		return elements, errors.New("negative size")
	}

	for range size {
		elements = append(elements, rand.Int())
	}

	return elements, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("no data")
	}

	if len(data) == 1 {
		return data[0], nil
	}

	var max int

	for _, element := range data {
		if element > max {
			max = element
		}
	}

	return max, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("no data")
	}

	if len(data) == 1 {
		return data[0], nil
	}

	maxSubElements := make(chan int, CHUNKS)

	countSubElements := SIZE / CHUNKS

	var wg sync.WaitGroup

	wg.Add(CHUNKS)

	for i := range CHUNKS {
		startIndex := i * countSubElements
		endIndex := startIndex + countSubElements

		go func(subData []int) {
			defer wg.Done()

			max, _ := maximum(subData)
			maxSubElements <- max
		}(data[startIndex:endIndex])
	}

	wg.Wait()
	close(maxSubElements)

	var elements []int
	for v := range maxSubElements {
		elements = append(elements, v)
	}

	max, _ := maximum(elements)

	return max, nil
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	elements, err := generateRandomElements(SIZE)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}

	now := time.Now()

	fmt.Println("Ищем максимальное значение в один поток")
	max, err := maximum(elements)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}

	elapsed := time.Now().Sub(now) * time.Microsecond

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	now = time.Now()

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	max, err = maxChunks(elements)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}

	elapsed = time.Now().Sub(now) * time.Microsecond

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
