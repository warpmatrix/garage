#include <cuda_runtime.h>
#include <stdio.h>

__global__ void wait() {
    clock_t start = clock();
    clock_t clock_offset = 0;
    clock_t clock_count = 100000000;
    while (clock_offset < clock_count) {
        clock_offset = clock() - start;
    }
}

int main(int argc, char const *argv[]) {
    cudaEvent_t start, end;
    cudaEventCreate(&start);
    cudaEventCreate(&end);

    cudaEventRecord(start);
    wait<<<1, 1>>>();
    cudaEventRecord(end);
    cudaEventSynchronize(end);

    float elapsedTime;
    cudaEventElapsedTime(&elapsedTime, start, end);
    printf("elapsed time: %f ms\n", elapsedTime);

    cudaEventDestroy(start);
    cudaEventDestroy(end);
    return 0;
}
