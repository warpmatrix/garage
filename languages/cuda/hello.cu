#include <cuda_runtime.h>
#include <stdio.h>

__global__ void hello_cuda() { printf("Hello Cuda!\n"); }

int main(int argc, char const *argv[]) {
    hello_cuda<<<1, 1>>>();
    cudaDeviceSynchronize();
    return 0;
}
