import ray

ray.init()

@ray.remote
def remote_fn(name: str) -> str:
    return f"hello, {name}"

futures = remote_fn.remote("world")
print(ray.get(futures))

local_fn = lambda name: f"hello, {name}"
futures = ray.remote(local_fn).remote("world")
print(ray.get(futures))
