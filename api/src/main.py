from RPCClient import STARsSIDsRpcClient
import json
from typing import Optional
from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def root():
    return {"message": "Hello World"}

@app.get("/STARS-AND-SIDS")
def get_stars_and_sids():
    STARsSIDs_rpc = STARsSIDsRpcClient()
    results = STARsSIDs_rpc.call()
    return json.loads(results.decode('utf8'))