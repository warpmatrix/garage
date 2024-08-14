#include <cuda_runtime.h>
#include <iostream>

int main() {
    int deviceCount;
    cudaGetDeviceCount(&deviceCount);
    std::cout << "Number of CUDA devices: " << deviceCount << std::endl;

    for (int i = 0; i < deviceCount; ++i) {
        for (int j = 0; j < deviceCount; ++j) {
            if (i == j) {
                continue;
            }
            int canAccessPeer;
            cudaDeviceCanAccessPeer(&canAccessPeer, i, j);
            std::cout << "Device " << i << " can ";
            if (!canAccessPeer) std::cout << "not ";
            std::cout << "access Device " << j << " via P2P" << std::endl;
        }
    }

    return 0;
}
