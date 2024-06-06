#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <locale.h>

// Função para mesclar duas partes de um slice ordenado
int* merge(int* left, int* right, int left_size, int right_size) {
    int* result = (int*)malloc((left_size + right_size) * sizeof(int));
    int l = 0, r = 0, idx = 0;

    while (l < left_size && r < right_size) {
        if (left[l] <= right[r]) {
            result[idx++] = left[l++];
        } else {
            result[idx++] = right[r++];
        }
    }

    while (l < left_size) {
        result[idx++] = left[l++];
    }

    while (r < right_size) {
        result[idx++] = right[r++];
    }

    return result;
}

// Função de ordenação por mesclagem
int* mergeSort(int* arr, int size) {
    if (size <= 1) {
        return arr;
    }

    int mid = size / 2;
    int* left = mergeSort(arr, mid);
    int* right = mergeSort(arr + mid, size - mid);

    return merge(left, right, mid, size - mid);
}

int main() {
    int arr[] = {5, 9, 3, 7, 2, 8, 6, 1, 4};
    int size = sizeof(arr) / sizeof(arr[0]);

    printf("Array não ordenado:\n");
    for (int i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");

    int* sortedArr = mergeSort(arr, size);

    printf("Array ordenado:\n");
    for (int i = 0; i < size; i++) {
        printf("%d ", sortedArr[i]);
    }
    printf("\n");

    free(sortedArr); // Liberar memória alocada

    return 0;
}
