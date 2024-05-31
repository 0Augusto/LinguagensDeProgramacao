/*
Abstração:
Tipo: Comportamental.
Descrição: Abstração de comportamento ao ocultar os detalhes da execução assíncrona (goroutine) e da simulação de operações demoradas.
Encapsulamento:
Tipo: Operacional.
Descrição: Encapsulamento do tempo de espera necessário para sincronizar a execução das goroutines e da função main().
*/

package main

import (
    "fmt"
    "time"
)

func minhaGoroutine(){

    for i := 0; i < 5; i++{
        fmt.Println("Goroutine executada: ", i) // Abstração: Mensagens de impressão para fornecer feedback sobre a execução de uma goroutine.
        time.Sleep(time.Millisecond * 500) // Abstração: Simulação de uma operação demorada.
    }
}//end func

func main(){

    // Inicia a goroutine
    go minhaGoroutine() // Abstração: Inicia uma nova goroutine para execução assíncrona.

    // Executa tarefa na função principal

    for i := 0; i < 5; i++ {
        fmt.Println("Funcao main executada", i) // Abstração: Mensagens de impressão para fornecer feedback sobre a execução da função principal.
        time.Sleep(time.Millisecond * 200) // Abstração: Simulação de uma operação demorada.
    }

    time.Sleep(time.Second) // Encapsulamento: Garante que a execução não termine até que as goroutines tenham tido tempo suficiente para serem concluídas.
}//end main
