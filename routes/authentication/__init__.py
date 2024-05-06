import uuid
from dataclasses import dataclass
from typing import Literal

from fastapi import APIRouter, Depends
from pydantic import BaseModel
from starlette.responses import RedirectResponse
from typing_extensions import Annotated
import requests

from config.env import settings

router = APIRouter(prefix="/auth")


@dataclass
class OAuthParams:
    code: str
    state: str


class AuthorizationCodeRequest(BaseModel):
    redirect_uri: str
    grant_type: Literal["authorization_code"] = "authorization_code"
    code: str


@router.get("/bnet/code")
def bnet_code(oauth_params: Annotated[OAuthParams, Depends()]):
    # ToDo: state management
    request_body = AuthorizationCodeRequest(redirect_uri="http://localhost:8000/auth/bnet/code",
                                            code=oauth_params.code)
    response = requests.post(url=settings.base_bnet_token_url,
                             data=request_body.model_dump(),
                             auth=(settings.client_id, settings.client_secret))
    print(response.text)


@router.get("/bnet/signin")
def bnet_signin():
    state_token: str = str(uuid.uuid4())
    response = RedirectResponse(url=f"{settings.base_bnet_auth_url}&client_id={settings.client_id}"
                                    f"&state={state_token}")
    return response
