from typing import Optional

from fastapi import FastAPI

import time

app = FastAPI()

@app.get('/')
async def index():

    time.sleep(5)

    return {"key": "value"}