#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

// Estrutura para passar argumentos para a função de thread.
typedef struct {
    int inicio;
    int fim;
    int *soma;
} ThreadArgs;

// Função para calcular a soma parcial dentro de um intervalo especificado.
void *calcularSoma(void *args) {
    ThreadArgs *threadArgs = (ThreadArgs *)args;
    int somaParcial = 0;

    // Loop para calcular a soma parcial dentro do intervalo.
    for (int i = threadArgs->inicio; i < threadArgs->fim; i++) {
        somaParcial += i;
    }

    // Imprime a soma parcial calculada pela thread.
    printf("Thread de %d a %d: Soma parcial: %d\n", threadArgs->inicio, threadArgs->fim - 1, somaParcial);

    // Adiciona a soma parcial à soma total.
    *(threadArgs->soma) += somaParcial;

    free(threadArgs); // Libera a memória alocada para os argumentos da thread.
    pthread_exit(NULL);
}

int main() {
    pthread_t threads[4];
    int soma = 0;
    int partes = 4;

    // Loop para criar e executar threads para calcular a soma parcial de cada parte.
    for (int i = 0; i < partes; i++) {
        ThreadArgs *args = malloc(sizeof(ThreadArgs));
        args->inicio = i * (10 / partes);
        args->fim = (i == partes - 1) ? 10 : (i + 1) * (10 / partes);
        args->soma = &soma;

        // Cria uma nova thread para calcular a soma parcial nesta parte.
        if (pthread_create(&threads[i], NULL, calcularSoma, (void *)args) != 0) {
            perror("Falha ao criar a thread");
            exit(EXIT_FAILURE);
        }
    }

    // Aguarda até que todas as threads tenham terminado.
    for (int i = 0; i < partes; i++) {
        if (pthread_join(threads[i], NULL) != 0) {
            perror("Falha ao aguardar a thread");
            exit(EXIT_FAILURE);
        }
    }

    // Imprime o resultado final.
    printf("A soma dos números de 0 a 9 (com concorrência) é: %d\n", soma);

    return 0;
}

