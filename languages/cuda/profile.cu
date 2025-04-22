#include <cstdio>
#include <functional>
#include <cuda_runtime.h>
#include <cuda_profiler_api.h>
#include <nvtx3/nvtx3.hpp>

class CudaProfiler{
public:
    CudaProfiler() {
        cudaProfilerStart();
    }
    ~CudaProfiler() {
        cudaProfilerStop();
    }
};

void with_cuda_profiler(std::function<void ()> fn) {
    CudaProfiler p;
    fn();
}

__global__ void longKernel() {
    clock_t start_clock = clock();
    clock_t clock_offset = 0;
    clock_t clock_count = 1000000000;
    while (clock_offset < clock_count) {
        clock_offset = clock() - start_clock;
    }
}

void launch_long_kernel() {
    NVTX3_FUNC_RANGE();
    longKernel<<<1, 1>>>();
}

int main() {
    with_cuda_profiler([]() {
        NVTX3_FUNC_RANGE();
        launch_long_kernel();
        cudaDeviceSynchronize();
    });
    return 0;
}
