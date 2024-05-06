from dotenv import load_dotenv
from pydantic_settings import BaseSettings

load_dotenv()


class Settings(BaseSettings):
    client_id: str
    client_secret: str
    base_bnet_auth_url: str = "https://oauth.battle.net/authorize?" \
                              "response_type=code" \
                              "&scope=wow.profile" \
                              "&redirect_uri=http://localhost:8000/auth/bnet/code"
    base_bnet_token_url: str = "https://oauth.battle.net/token"


settings = Settings()
