from fastapi import FastAPI
from wow_server.routes.authentication import router as auth_router

app = FastAPI(debug=True)
app.include_router(auth_router)