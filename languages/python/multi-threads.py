import threading

def f(args):
    print(args)

ts_args = ["hello", "world"]
ts = [threading.Thread(target=f, args=[args])
    for args in ts_args
]
[t.start() for t in ts]
[t.join() for t in ts]
