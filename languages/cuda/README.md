# Cuda Related

- check the compute mode: `nvidia-smi --query-gpu=compute_mode --format=csv`
- enable mps of nvidia: `nvidia-cuda-mps-control -d`
- check the deamon of mps: `ps -ef | grep mps`

material:

- [mps slide](https://www.nvidia.cn/content/dam/en-zz/zh_cn/assets/webinars/31oct2019c/20191031_MPS_davidwu.pdf)
- [mps document](https://docs.nvidia.com/deploy/mps/index.html)

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
