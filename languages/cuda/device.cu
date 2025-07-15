#include <cuda.h>
#include <stdio.h>

int main(int argc, char** argv) {
    int driver_version = 0, runtime_version = 0;

    cudaDriverGetVersion(&driver_version);
    cudaRuntimeGetVersion(&runtime_version);

    printf("Driver Version: %d, Runtime Version: %d\n",
        driver_version, runtime_version);

    int nDevices;
    cudaGetDeviceCount(&nDevices);
    for (size_t i = 0; i < nDevices; i++) {
        cudaDeviceProp p;
        cudaGetDeviceProperties(&p, i);
        printf(" Device Number: %lu\n", i);
        printf(" Device name: %s\n", p.name);
        printf(" Memory Clock Rate (KHz): %d\n", p.memoryClockRate);
        printf(" Memory Bus Width (bits): %d\n", p.memoryBusWidth);
        printf(" Device %d has compute capability %d.%d.\n",
            device, deviceProp.major, deviceProp.minor);
        printf(" Device %d support: concurrent kenel execution(%d), intra-device copy(%d)\n",
            device, deviceProp.concurrentKernels, deviceProp.asyncEngineCount);
    }

    return 0;
}
