package main

import (
    "fmt" // Importa o pacote fmt para formatação de saída.
    "sync" // Importa o pacote sync para suporte a sincronização de goroutines.
)

// Função para calcular a soma parcial dentro de um intervalo especificado.
func calcularSoma(inicio, fim int, wg *sync.WaitGroup, soma *int) {
    defer wg.Done() // Sinaliza que a goroutine terminou quando a função retorna.
    somaParcial := 0
    // Loop para calcular a soma parcial dentro do intervalo.
    for i := inicio; i < fim; i++ {
        somaParcial += i
    }
    // Imprime a soma parcial calculada por esta goroutine.
    fmt.Printf("Goroutine de %d a %d: Soma parcial: %d\n", inicio, fim-1, somaParcial)
    *soma += somaParcial // Adiciona a soma parcial à soma total.
}

func main() {
    var wg sync.WaitGroup
    soma := 0
    partes := 4 // Número de partes para dividir a soma

    // Loop para criar e executar goroutines para calcular a soma parcial de cada parte.
    for i := 0; i < partes; i++ {
        wg.Add(1) // Incrementa o contador de goroutines no WaitGroup.
        inicio := i * (10 / partes) // Calcula o início do intervalo para esta parte.
        var fim int
        // Verifica se é a última parte para definir o fim como o final do intervalo completo.
        if i == partes-1 {
            fim = 10
        } else {
            fim = (i + 1) * (10 / partes) // Calcula o fim do intervalo para esta parte.
        }
        // Inicia uma goroutine para calcular a soma parcial nesta parte.
        go calcularSoma(inicio, fim, &wg, &soma)
    }

    wg.Wait() // Aguarda até que todas as goroutines tenham terminado.
    fmt.Println("A soma dos números de 0 a 9 (com concorrência) é:", soma) // Imprime o resultado final.
}

