import time
import torch

def fn(x):
   a = torch.cos(x)
   b = torch.sin(a)
   return b
new_fn = torch.compile(fn, backend="inductor")
input_tensor = torch.randn(10000).to(device="cuda:0")

start = time.perf_counter()
a = new_fn(input_tensor)
end = time.perf_counter()
print(end - start)

start = time.perf_counter()
a = fn(input_tensor)
end = time.perf_counter()
print(end - start)
