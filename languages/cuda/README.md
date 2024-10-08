# Cuda Related

- check the compute mode: `nvidia-smi --query-gpu=compute_mode --format=csv`
- enable mps of nvidia: `nvidia-cuda-mps-control -d`
- check the deamon of mps: `ps -ef | grep mps`

cuda 编程软件层级概念：thread、block、grid

- kernel 函数需要指定 blocksPerGrid, threadsPerBlock

cuda 编程中的硬件概念：SM、warp、SP

- cuda 程序执行过程中以 block 为单位分配给 SM 执行
- block 中的 thread 以 warp 为单位分配执行
- active thread 一个 SM 同时处理的 thread 数量

gpu 中的 hbm 读写操作共用带宽，单一操作最高可以达到 spec 中提及的带宽速度

- [nsight-compute 不支持在 MPS 中使用](https://docs.nvidia.com/nsight-compute/ReleaseNotes/index.html#known-issues)，出现大量 nan 的情况

> Profiling with enabled multi-process service (MPS) can result in undefined behavior.

## cuda gdb

Initialization File: `.cuda-gdbinit`

## cuda execution details

- [cuda programming guide](https://docs.nvidia.com/cuda/cuda-c-programming-guide/index.html)
- [cuda best practices](https://docs.nvidia.com/cuda/cuda-c-best-practices-guide/index.html)
- [tail effect](https://developer.nvidia.com/blog/cuda-pro-tip-minimize-the-tail-effect/)
- [stream priority](https://forums.developer.nvidia.com/t/questions-of-cuda-stream-priority/250343/4)
- [Multiple Instance GPU - MIG](https://docs.nvidia.com/datacenter/tesla/mig-user-guide/index.html)

related paper: Paella: Low-latency Model Serving with Software-defined GPU Scheduling

## Profile

- [nvbandwidth](https://github.com/NVIDIA/nvbandwidth)
  - also corresponds to the output of `nvidia-smi topo -m` commands

## material

- [mps slide](https://www.nvidia.cn/content/dam/en-zz/zh_cn/assets/webinars/31oct2019c/20191031_MPS_davidwu.pdf)
- [mps document](https://docs.nvidia.com/deploy/mps/index.html)
- [cuda driver api](https://docs.nvidia.com/cuda/cuda-driver-api/index.html)
