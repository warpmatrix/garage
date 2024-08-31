import time
import torch
from torch import multiprocessing


def worker(queue: multiprocessing.Queue):
    worker_gpu_start = torch.cuda.Event(enable_timing=True)
    worker_gpu_end = torch.cuda.Event(enable_timing=True)
    ipc_event_handle = queue.get()
    ipc_event = torch.cuda.Event.from_ipc_handle("cuda:0", ipc_event_handle)
    worker_gpu_start.record()
    ipc_event.wait()
    worker_gpu_end.record()
    worker_gpu_end.synchronize()
    print(
        "Worker process synchronized with event",
        worker_gpu_start.elapsed_time(worker_gpu_end),
    )

if __name__ == '__main__':
    multiprocessing.set_start_method("spawn", force=True)
    queue = multiprocessing.Queue()
    event = torch.cuda.Event(interprocess=True)
    p = multiprocessing.Process(target=worker, args=(queue,))
    p.start()

    gpu_start: torch.cuda.Event = torch.cuda.Event(enable_timing=True)
    gpu_end: torch.cuda.Event = torch.cuda.Event(enable_timing=True)

    cpu_start = time.time()
    gpu_start.record()
    size = 10000
    a = torch.randn(size, size, device="cuda")
    for _ in range(100):
        torch.matmul(a, a)
    event.record()
    ipc_event_handle = event.ipc_handle()
    queue.put(ipc_event_handle)
    gpu_end.record()
    cpu_end = time.time()
    print("cpu time:", cpu_end - cpu_start)
    gpu_end.synchronize()
    print("gpu time:", gpu_start.elapsed_time(gpu_end))
    p.join()
