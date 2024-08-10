# pull base immage
FROM python:3.11

# set work directory
WORKDIR /app

# set environment variables
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# install dependencies
RUN apt-get update
RUN apt-get install -y sqlite3 libsqlite3-dev

RUN pip install --upgrade pip
COPY ./requirements.txt /app/requirements.txt
RUN pip install -r requirements.txt

COPY ./docker/django/start /start
RUN sed -i 's/\r$//g' /start
RUN chmod +x /start
