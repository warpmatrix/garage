from time import sleep
from multiprocessing import current_process
from multiprocessing import Process


def task():
    process = current_process()
    print(f"Daemon process: {process.daemon}")
    sleep(3)
    print(f"normal exit") # can't exec if the main process has terminated


if __name__ == "__main__":
    process = Process(daemon=True, target=task)
    process.start()
    print(process.is_alive())
    sleep(1)
