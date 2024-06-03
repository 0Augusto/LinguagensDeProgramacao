/*
A estrutura Cache é definida para armazenar valores em cache, com um mapa de dados e um mutex para garantir concorrência segura.
Os métodos Set e Get são definidos para adicionar e obter valores do cache, respectivamente.
No main, criamos uma nova instância de Cache, adicionamos alguns valores e os recuperamos do cache. Em seguida, simulamos um intervalo de tempo e tentamos recuperar um valor após esse intervalo para demonstrar que ele expirou do cache.


*/



package main

import (
	"fmt"
	"sync"
	"time"
)

// Cache é uma estrutura para armazenar valores em cache.
type Cache struct {
	data map[string]interface{} // Mapa para armazenar os dados do cache
	mu   sync.Mutex             // Mutex para garantir concorrência segura
}

// NewCache cria uma nova instância de Cache.
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}), // Inicializa o mapa de dados
	}
}

// Set adiciona um valor ao cache com a chave especificada.
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()         // Bloqueia o acesso para garantir segurança
	defer c.mu.Unlock() // Libera o acesso após a função ser concluída
	c.data[key] = value // Adiciona o valor ao cache
}

// Get retorna o valor associado à chave especificada no cache.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()         // Bloqueia o acesso para garantir segurança
	defer c.mu.Unlock() // Libera o acesso após a função ser concluída
	value, ok := c.data[key] // Verifica se a chave está presente no cache
	return value, ok   // Retorna o valor e um booleano indicando se a chave foi encontrada
}

func main() {
	// Criando uma nova instância de Cache
	cache := NewCache()

	// Adicionando valores ao cache
	cache.Set("chave1", "valor1")
	cache.Set("chave2", "valor2")

	// Obtendo valores do cache
	valor1, ok1 := cache.Get("chave1")
	valor2, ok2 := cache.Get("chave2")

	if ok1 {
		fmt.Println("Valor da chave1:", valor1)
	} else {
		fmt.Println("Chave1 não encontrada no cache")
	}

	if ok2 {
		fmt.Println("Valor da chave2:", valor2)
	} else {
		fmt.Println("Chave2 não encontrada no cache")
	}

	// Simulando um intervalo de tempo para verificar expiração do cache
	time.Sleep(5 * time.Second)

	// Tentando obter um valor após o intervalo de tempo
	valor3, ok3 := cache.Get("chave1")
	if ok3 {
		fmt.Println("Valor da chave1 após expiração:", valor3)
	} else {
		fmt.Println("Chave1 expirou no cache")
	}
}

