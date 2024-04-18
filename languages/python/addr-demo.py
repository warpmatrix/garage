import torch

def addr(obj):
    return hex(id(obj))

a = torch.rand(2, 3, device="cuda")
residual = a
assert addr(a) == addr(residual)
a = a / 10
assert addr(a) != addr(residual)

b = 10
origin_addr = addr(b)
b += 1
assert origin_addr != addr(b)
