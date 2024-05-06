import uuid
from dataclasses import dataclass
from typing import Literal

from fastapi import APIRouter, Depends
from starlette.responses import RedirectResponse
from typing_extensions import Annotated

from config.env import settings

router = APIRouter(prefix="/auth")

@dataclass
class OAuthParams:
    code: str
    state: str


@router.get("/bnet/code")
def bnet_code(oauth_params: Annotated[OAuthParams, Depends()]):
    print(oauth_params.code)

@router.get("/bnet/signin")
def bnet_signin():
    state_token: str = str(uuid.uuid4())
    response = RedirectResponse(url=f"{settings.base_bnet_auth_url}&client_id={settings.client_id}"
                                    f"&state={state_token}")
    return response