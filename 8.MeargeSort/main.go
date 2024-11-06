package main

import (
	"fmt"
	"sync"
)

// merge сливает два отсортированных подмассива в один.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Добавить оставшиеся элементы
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// parallelMergeSort рекурсивно разбивает массив и выполняет сортировку параллельно.
func parallelMergeSort(arr []int, wg *sync.WaitGroup) []int {
	defer wg.Done()

	// Базовый случай рекурсии
	if len(arr) <= 1 {
		return arr
	}

	// Разделяем массив на две части
	mid := len(arr) / 2
	var left, right []int

	// Используем WaitGroup для синхронизации горутин
	leftWg := &sync.WaitGroup{}
	rightWg := &sync.WaitGroup{}

	leftWg.Add(1)
	go func() {
		left = parallelMergeSort(arr[:mid], leftWg)
	}()

	rightWg.Add(1)
	go func() {
		right = parallelMergeSort(arr[mid:], rightWg)
	}()

	// Ожидаем завершения сортировки левого и правого подмассивов
	leftWg.Wait()
	rightWg.Wait()

	// Слияние отсортированных массивов
	return merge(left, right)
}

func main() {
	arr := []int{1231, 4214, 515, 152, 515125, 51255, 12122412}
	fmt.Println("Исходный массив:", arr)

	// Используем sync.WaitGroup для главной горутины
	wg := &sync.WaitGroup{}
	wg.Add(1)

	sortedArr := parallelMergeSort(arr, wg)
	wg.Wait()

	fmt.Println("Отсортированный массив:", sortedArr)
}
