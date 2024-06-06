/*
Abstração:

Tipo: Funcional.
Descrição: As funções merge e mergeSort abstraem os detalhes do algoritmo de ordenação merge sort e da operação de mesclagem de listas ordenadas.
Encapsulamento:

Tipo: Estrutural.
Descrição: O encapsulamento é evidente nas funções merge e mergeSort, onde a implementação interna do algoritmo é oculta, e o usuário só interage com as funções por meio de suas interfaces públicas. Além disso, o uso de wait groups encapsula a execução concorrente do merge sort.
Amarração:

Tipo: Comunicação.
Descrição: As goroutines são usadas para executar tarefas em paralelo. A amarração é feita usando wait groups para garantir que a função principal aguarde a conclusão das goroutines antes de continuar. Isso permite a execução concorrente do merge sort, melhorando o desempenho em sistemas com múltiplos núcleos de CPU.
Sistema de Tipos:

O sistema de tipos em Go é estático e forte. As variáveis têm tipos definidos e não podem ser alteradas dinamicamente em tempo de execução. Os tipos são verificados em tempo de compilação, o que ajuda a evitar erros comuns e melhorar a segurança do código. Os tipos são inferidos automaticamente quando possível, mas também podem ser explicitamente declarados pelo programador. No código fornecido, são usados tipos como int, []int, sync.WaitGroup, entre outros.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Abstração: A função merge realiza a junção de duas listas ordenadas em uma única lista ordenada.
// Tipo: Funcional.
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

// Abstração: A função mergeSort implementa o algoritmo de ordenação merge sort para uma lista de inteiros.
// Tipo: Funcional.
func mergeSort(arr []int) []int {
	// Caso base: se a lista tiver tamanho 0 ou 1, ela já está ordenada.
	if len(arr) <= 1 {
		return arr
	}

	// Divide a lista ao meio
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	var wg sync.WaitGroup
	wg.Add(2)

	// Permite que outras goroutines rodem
	runtime.Gosched()

	var leftResult, rightResult []int

	// Amarração: Goroutine para ordenar a metade esquerda da lista
	// Tipo: Comunicação.
	go func() {
		leftResult = mergeSort(left)
		wg.Done()
	}()

	// Amarração: Goroutine para ordenar a metade direita da lista
	// Tipo: Comunicação.
	go func() {
		rightResult = mergeSort(right)
		wg.Done()
	}()

	// Encapsulamento: Aguarda as duas goroutines terminarem
	// Tipo: Estrutural.
	wg.Wait()

	// Retorna a junção das duas metades ordenadas
	return merge(leftResult, rightResult)
}

func main() {
	arr := []int{5, 9, 3, 7, 2, 8, 6, 1, 4}
	fmt.Println("Unsorted array:", arr)

	// Abstração: Chamada para a função mergeSort para ordenar a lista de inteiros.
	// Tipo: Funcional.
	arr = mergeSort(arr)

	fmt.Println("Sorted array:", arr)
}
