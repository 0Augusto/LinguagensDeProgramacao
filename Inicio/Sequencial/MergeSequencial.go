package main

import (
	"fmt"
//	"runtime"
	"sync"
)

// Função para mesclar duas partes de um slice ordenado
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

// Função de ordenação por mesclagem
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Divide o slice em duas partes
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	var wg sync.WaitGroup
	wg.Add(2)

	// Executa as duas metades em goroutines separadas
	go func() {
		leftResult := mergeSort(left)
		left = leftResult // Atualiza a parte esquerda com o resultado
		wg.Done()
	}()

	go func() {
		rightResult := mergeSort(right)
		right = rightResult // Atualiza a parte direita com o resultado
		wg.Done()
	}()

	wg.Wait()

	// Mescla as partes esquerda e direita
	return merge(left, right)
}

func main() {
	arr := []int{5, 9, 3, 7, 2, 8, 6, 1, 4}
	fmt.Println("Array não ordenado:", arr)

	// Chama a função de ordenação por mesclagem
	arr = mergeSort(arr)

	fmt.Println("Array ordenado:", arr)
}

