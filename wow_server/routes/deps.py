from fastapi import Depends

from model.db import Session
from sqlalchemy.orm import Session as SQLAlchemySession


def session_dependency() -> SQLAlchemySession:
    return Session()


SessionDependency: SQLAlchemySession = Depends(session_dependency)
