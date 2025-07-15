# PyTorch 与 CUDA

torch 操作 GPU 的同步与异步语义：

- `.to(device)` 会进行两个设备之间的同步
- `tensor(device="cuda")` 直接在设备上申请内存

[torch 自定义算子](https://pytorch.org/tutorials/advanced/custom_ops_landing_page.html): tutorial 中的 Extending PyTorch 一节

- [python](https://pytorch.org/tutorials/advanced/python_custom_ops.html#python-custom-ops-tutorial)、[cpp/cuda](https://pytorch.org/tutorials/advanced/cpp_custom_ops.html#cpp-custom-ops-tutorial)
- [before 2.14](https://pytorch.org/tutorials/advanced/cpp_extension.html)

在 aarch64 上使用 pip 安装 PyTorch 是 cpu 版本，cuda 版本需要找到对应的 [wheel 包](https://download.pytorch.org/whl/torch/)安装
