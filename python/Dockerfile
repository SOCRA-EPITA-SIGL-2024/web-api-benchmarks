FROM python:3.9.16

COPY ./python /code
COPY ./garden.json /data/garden.json
WORKDIR /code

RUN pip install -r requirements.in

ENV GARDENS_JSON_FILE_PATH=/data/garden.json

EXPOSE 8000

CMD uvicorn main:app --host 0.0.0.0
