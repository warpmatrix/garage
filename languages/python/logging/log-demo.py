import logging
from m1 import f1

logger = logging.getLogger(__name__)
logger.setLevel(logging.WARNING)
logger.warning(f"{logger} in {__name__}")
f1.f1()
