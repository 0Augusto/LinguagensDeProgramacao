#include <stdio.h>

int main() {
    int soma = 0; // Inicializa a variável 'soma' com zero
    for (int i = 0; i < 10; i++) { // Loop de 0 a 9
        soma += i; // Adiciona o valor de 'i' à 'soma'
    }
    printf("A soma dos números de 0 a 9 é: %d\n", soma); // Exibe a mensagem com o resultado da soma
    return 0; // Retorna zero para indicar sucesso
}

