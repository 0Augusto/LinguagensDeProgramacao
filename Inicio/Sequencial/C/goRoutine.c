#include <stdio.h>
#include <stdlib.h>
#include <pthread.h> // Biblioteca para trabalhar com threads
#include <unistd.h> // Biblioteca para usar a função sleep

// Função executada pela goroutine
void* minhaGoroutine(void* arg) {
    for (int i = 0; i < 5; i++) {
        printf("Goroutine executada: %d\n", i);
        usleep(500000); // Simula operação demorada (500 ms)
    }
    return NULL;
}

int main() {
    pthread_t thread; // Variável para armazenar a thread
    int rc;

    // Cria a goroutine (thread)
    rc = pthread_create(&thread, NULL, minhaGoroutine, NULL);
    if (rc) {
        printf("Erro ao criar a thread: %d\n", rc);
        exit(-1);
    }

    // Executa tarefa na função principal (main)
    for (int i = 0; i < 5; i++) {
        printf("Funcao main executada: %d\n", i);
        usleep(200000); // Simula operação demorada (200 ms)
    }

    // Aguarda a goroutine terminar
    pthread_join(thread, NULL);

    return 0;
}

