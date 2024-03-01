import shutil
import subprocess

assert shutil.which("nsys") is not None
profile_cmd = ["nsys", "profile"]
profile_flags = [
    "--python-backtrace=cuda",
    "--trace=cuda,nvtx",
    "--backtrace=dwarf",
    "--gpu-metrics-device=all",
    "--gpu-metrics-frequency=100000",
    "--start-later"
]

subprocess.run(profile_cmd + profile_flags + ["python", "demo.py"])
