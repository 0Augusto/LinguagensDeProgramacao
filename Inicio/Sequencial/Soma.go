package main

import "fmt"

func main() {
    soma := 0
    for i := 0; i < 10; i++ {
        soma += i
    }
    fmt.Println("A soma dos números de 0 a 9 é:", soma)
}

