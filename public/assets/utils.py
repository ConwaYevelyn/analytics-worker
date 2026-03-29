import logging
import os
import uuid
from datetime import datetime

def generate_unique_id():
    return str(uuid.uuid4())

def get_project_root():
    return os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

def get_current_timestamp():
    return datetime.now().strftime('%Y-%m-%d %H:%M:%S')

def configure_logging(log_file, log_level):
    logging.basicConfig(
        filename=log_file,
        level=getattr(logging, log_level.upper()),
        format='%(asctime)s - %(levelname)s - %(message)s'
    )