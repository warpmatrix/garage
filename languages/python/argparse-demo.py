import argparse

parser = argparse.ArgumentParser(description='Demo program')
parser.add_argument('--port', type=int, default=10086, help='port of DataNode')

args = parser.parse_args()
print(args.port)
