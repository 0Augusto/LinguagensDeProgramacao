#include <stdio.h>

void AlteraValor(int val);

int main (){
    int val = 50;

    printf("Valor antes da funcao: %d\n", val);
    AlteraValor(val);
    printf("Valor depois da funcao: %d\n", val);

return 0;
}

void AlteraValor(int val){

    val = (val * 15);//expressao
    printf("Valor dentro da funcao: %d\n", val);
    puts("Quando o valor e passado por referencia, o valor da variavel dentro da main  NAO e alterado apos a execucao da expressao.");
}
