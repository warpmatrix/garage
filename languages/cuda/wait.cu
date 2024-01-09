#include <cuda_runtime.h>
#include <stdio.h>
#include <unistd.h>

__global__ void wait() {
    clock_t start_clock = clock();
    clock_t clock_offset = 0;
    clock_t clock_count = 10000000000;
    while (clock_offset < clock_count) {
        clock_offset = clock() - start_clock;
    }
}

int main(int argc, char const *argv[]) {
    wait<<<1, 1>>>();
    cudaDeviceSynchronize();
    return 0;
}
