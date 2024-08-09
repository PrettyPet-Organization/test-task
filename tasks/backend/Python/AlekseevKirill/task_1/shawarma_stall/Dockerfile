FROM python:3.12.1
WORKDIR /code

RUN pip install poetry
COPY ./pyproject.toml ./poetry.lock* /code/

RUN poetry export --without-hashes --format=requirements.txt > requirements.txt

RUN pip install --no-cache-dir --upgrade -r /code/requirements.txt

COPY . /code/

RUN python manage.py migrate 

CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]