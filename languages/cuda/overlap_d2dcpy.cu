#include <cuda_runtime.h>
#include <iostream>

__global__ void kernel(int *a) {
    int idx = threadIdx.x + blockIdx.x * blockDim.x;
    a[idx] *= 2;
}

int main() {
    const int size = 1024;
    int *d_a, *d_b;
    int h_a[size], h_b[size];

    // 分配设备内存
    cudaMalloc(&d_a, size * sizeof(int));
    cudaMalloc(&d_b, size * sizeof(int));

    // 初始化数据
    for (int i = 0; i < size; i++) {
        h_a[i] = i;
        h_b[i] = i + 1;
    }

    // 将数据从主机拷贝到设备
    cudaMemcpy(d_a, h_a, size * sizeof(int), cudaMemcpyHostToDevice);
    cudaMemcpy(d_b, h_b, size * sizeof(int), cudaMemcpyHostToDevice);

    // 创建 CUDA 流
    cudaStream_t stream1, stream2;
    cudaStreamCreate(&stream1);
    cudaStreamCreate(&stream2);

    // 在流 1 中执行 kernel
    kernel<<<1, 256, 0, stream1>>>(d_a);

    // 在流 2 中执行 D2D memcpy
    cudaMemcpyAsync(d_b, d_a, size * sizeof(int), cudaMemcpyDeviceToDevice, stream2);

    // 等待流中的操作完成
    cudaStreamSynchronize(stream1);
    cudaStreamSynchronize(stream2);

    // 释放设备内存
    cudaFree(d_a);
    cudaFree(d_b);

    cudaStreamDestroy(stream1);
    cudaStreamDestroy(stream2);

    return 0;
}
