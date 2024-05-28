#include <stdio.h>
#include <unistd.h> // Para a função sleep

// Função para simular a operação demorada na goroutine
void minhaGoroutine() {
    for (int i = 0; i < 5; i++) {
        printf("Goroutine executada: %d\n", i);
        usleep(500000); // Simula operação demorada em microssegundos
    }
}

int main() {
    // Inicia a goroutine
    pid_t pid = fork(); // Cria um novo processo

    if (pid == 0) { // Processo filho
        minhaGoroutine();
    } else if (pid > 0) { // Processo pai
        // Executa tarefa na função principal
        for (int i = 0; i < 5; i++) {
            printf("Funcao main executada: %d\n", i);
            usleep(200000); // Simula operação demorada em microssegundos
        }
    } else { // Se houve erro ao criar o processo filho
        perror("Erro ao criar processo filho");
        return 1;
    }

    sleep(1); // Aguarda um segundo para garantir que todas as operações terminem
    return 0;
}

