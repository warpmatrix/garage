# PyTorch 常用函数

- `tensor.sum(dims)` 按照某一维度或多个维度进行求和
- `tensor.unsqueeze(dims)` 在某一维度上增加一个维度
- `tensor.squeeze()` 将长度为 1 的维度去除
- `tensor.scatter(dim, index, src)` 在某一维度上进行分散操作，`src` 为源数据，`index` 为索引，`dim` 为维度
- `torch.cat([tensor], dim)` 在特定维度上进行 `tensor` 的拼接
