#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

// Função para mesclar dois arrays ordenados.
int* merge(int* left, int leftSize, int* right, int rightSize) {
    int* result = (int*)malloc((leftSize + rightSize) * sizeof(int));
    int l = 0, r = 0, k = 0;

    while (l < leftSize && r < rightSize) {
        if (left[l] <= right[r]) {
            result[k++] = left[l++];
        } else {
            result[k++] = right[r++];
        }
    }

    while (l < leftSize) {
        result[k++] = left[l++];
    }

    while (r < rightSize) {
        result[k++] = right[r++];
    }

    return result;
}

// Função de ordenação merge sort.
int* mergeSort(int* arr, int size) {
    if (size <= 1) {
        return arr;
    }

    int mid = size / 2;
    int* left = mergeSort(arr, mid);
    int* right = mergeSort(arr + mid, size - mid);

    return merge(left, mid, right, size - mid);
}

int main() {
    int arr[] = {5, 9, 3, 7, 2, 8, 6, 1, 4};
    int size = sizeof(arr) / sizeof(arr[0]);

    printf("Unsorted array: ");
    for (int i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");

    int* sortedArr = mergeSort(arr, size);

    printf("Sorted array: ");
    for (int i = 0; i < size; i++) {
        printf("%d ", sortedArr[i]);
    }
    printf("\n");

    free(sortedArr); // Libera a memória alocada para o array ordenado

    return 0;
}

