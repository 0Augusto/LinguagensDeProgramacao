package main

import (
    "fmt"
    "time"
)

func minhaGoroutine(){

    for i := 0; i < 5; i++{
    fmt.Println("Goroutine executada: ", i)
    time.Sleep(time.Millisecond * 500)//Simula operacao demorada
    }
}//end func


//Executa de forma sequencial
func main(){

    //Inicio a gouroutine
    go minhaGoroutine()
    
    //Executa tarefa na main

    for i:= 0; i < 5; i++{
        fmt.Println("Funcao main executada", i)
        time.Sleep(time.Millisecond * 200)//simula demorada
    }

    time.Sleep(time.Second)
}//end main
