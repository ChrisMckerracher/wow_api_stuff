from sqlalchemy import create_engine, Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

from config.env import settings

Base = declarative_base()


class OAuth(Base):
    __tablename__ = 'oauth'
    user_id = Column(Integer, primary_key=True) # ToDo: also make this a foreign key
    access_token = Column(String)
    expires_in = Column(Integer)  # ToDo: change this to a datetime object eventually


engine = create_engine(settings.sqlite_db)
Base.metadata.create_all(engine)
Session = sessionmaker(bind=engine)