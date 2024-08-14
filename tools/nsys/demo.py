import shutil
import time
import subprocess

assert shutil.which("nsys") is not None
cmd = "nsys sessions list | tail -n 1 | awk '{print $1}'"
result = subprocess.getoutput(cmd)
cmd = f"nsys start --session={result}"
cmd_args = cmd.split(" ")
subprocess.run(cmd_args)
time.sleep(1)
