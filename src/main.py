from fastapi import FastAPI
from .headless import router as headless_router
from .managed import router as managed_router

app = FastAPI(title="Managed Email Service")

app.include_router(headless_router)
app.include_router(managed_router)

@app.get("/")
def read_root():
    return {"message": "Welcome to the Managed Email Service"}
