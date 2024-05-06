# This is a sample Python script.
from fastapi import FastAPI
from routes.authentication import router as auth_router

# Press ⌃R to execute it or replace it with your code.
# Press Double ⇧ to search everywhere for classes, files, tool windows, actions, and settings.
app = FastAPI(debug=True)
app.include_router(auth_router)