import os
from dotenv import load_dotenv
from pydantic import BaseModel

load_dotenv()


class Settings(BaseModel):
    client_id: str = os.getenv("client_id")
    client_secret: str = os.getenv("client_secret")
    base_bnet_auth_url: str = "https://oauth.battle.net/authorize?" \
                              "response_type=code" \
                              "&scope=wow.profile" \
                              "&redirect_uri=http://localhost:8000/auth/bnet/code" \

settings = Settings()
