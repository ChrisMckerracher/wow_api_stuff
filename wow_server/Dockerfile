FROM ubuntu:mantic
WORKDIR /server
COPY poetry.lock pyproject.toml /server/
COPY ./config /server/config
COPY ./routes /server/routes
COPY .env /server
COPY *.py /server

RUN apt-get update && apt-get install -y python3-pip
RUN pip3 install poetry
RUN poetry install --no-dev
RUN poetry exec uvicorn main:app
EXPOSE 80