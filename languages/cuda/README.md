# Cuda Related

- check the compute mode: `nvidia-smi --query-gpu=compute_mode --format=csv`
- enable mps of nvidia: `nvidia-cuda-mps-control -d`
- check the deamon of mps: `ps -ef | grep mps`

material:

- [mps slide](https://www.nvidia.cn/content/dam/en-zz/zh_cn/assets/webinars/31oct2019c/20191031_MPS_davidwu.pdf)
- [mps document](https://docs.nvidia.com/deploy/mps/index.html)
