import logging
import sys

stdout = logging.StreamHandler(sys.stdout)
stdout.setLevel(logging.DEBUG)

logger = logging.getLogger(__name__)
logger.addHandler(stdout)
logger.setLevel(logging.DEBUG)

def f1():
    logger.debug(f"{logger} in {__name__}")
