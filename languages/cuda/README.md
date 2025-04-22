# Nvidia Cuda

## Cuda Execution

- [cuda programming guide](https://docs.nvidia.com/cuda/cuda-c-programming-guide/index.html)
- [cuda best practices](https://docs.nvidia.com/cuda/cuda-c-best-practices-guide/index.html)
- [cuda driver api](https://docs.nvidia.com/cuda/cuda-driver-api/index.html)
- [tail effect](https://developer.nvidia.com/blog/cuda-pro-tip-minimize-the-tail-effect/)
- [stream priority](https://forums.developer.nvidia.com/t/questions-of-cuda-stream-priority/250343/4)
- [Multiple Instance GPU - MIG](https://docs.nvidia.com/datacenter/tesla/mig-user-guide/index.html)

cuda 编程软件层级概念：thread、block、grid

- kernel 函数需要指定 blocksPerGrid, threadsPerBlock

cuda 编程中的硬件概念：SM、warp、SP

- cuda 程序执行过程中以 block 为单位分配给 SM 执行
- block 中的 thread 以 warp 为单位分配执行
- active thread 一个 SM 同时处理的 thread 数量

gpu 中的 hbm 读写操作共用带宽，单一操作最高可以达到 spec 中提及的带宽速度

- check topology of nvidia device: `nvidia-smi topo -m`

related paper:

- Paella: Low-latency Model Serving with Software-defined GPU Scheduling

## Cuda GDB

Initialization File: `.cuda-gdbinit`

- [Disable Optimization for Profile and Debug](https://stackoverflow.com/questions/11821605/completely-disable-optimizations-on-nvcc)

## Profile Cuda Application

### Nsight System

Requirements:

- [Linux Performance](https://docs.nvidia.com/nsight-systems/InstallationGuide/index.html#requirements-for-x86-64-and-arm-sbsa-targets-on-linux)
- [GPU Performance Counter](https://developer.nvidia.com/nvidia-development-tools-solutions-err_nvgpuctrperm-permission-issue-performance-counters)

[Useful flags](https://docs.nvidia.com/nsight-systems/UserGuide/index.html#cli-profile-command-switch-options):

- `--gpu-metrics-devices`
- `--gpu-metrics-frequency`
- `--capture-range`
- `--capture-range-end`
- `--cuda-graph-trace`

Profile Example: [profile.cu](./profile.cu)

- `cudaProfilerStart`, `cudaProfilerStop`: control profile range
  - header file: `cuda_profiler_api.h`
- `NVTX3_FUNC_RANGE`: convenient way to create nvtx information
  - header file: `nvtx3/nvtx3.hpp`

>[!NOTE]
> `nvtx3.hpp` might be missing in CUDA toolkit. It may need to fetch `nvtx3.hpp` manually. ([cuda-12.4](https://forums.developer.nvidia.com/t/nvtx-c-api-release-model/265272/2))

Reference:

- [Nsight System](https://docs.nvidia.com/nsight-systems/UserGuide/index.html)
- [NVTX](https://github.com/NVIDIA/NVTX)

### Third Library

- [nvbandwidth](https://github.com/NVIDIA/nvbandwidth)

## MPS

Quick Start:

1. enable exclusive_process mode: `nvidia-smi -c EXCLUSIVE_PROCESS`
    - check the compute mode: `nvidia-smi --query-gpu=compute_mode --format=csv`
2. start mps server of nvidia: `nvidia-cuda-mps-control -d`
3. check the deamon of mps: `pgrep -a -f mps`

> [!WARNING]
> [Nsight Compute 不支持在 MPS 中使用](https://docs.nvidia.com/nsight-compute/ReleaseNotes/index.html#known-issues)，出现 nan 数据 (undefined behavior?) - Nsight Compute 12.8 - Known Issues - Profiling and Metrics

Reference:

- [mps document](https://docs.nvidia.com/deploy/mps/index.html)
- [mps slide](https://www.nvidia.cn/content/dam/en-zz/zh_cn/assets/webinars/31oct2019c/20191031_MPS_davidwu.pdf)
