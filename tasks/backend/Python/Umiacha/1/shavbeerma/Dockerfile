FROM python:3.10

WORKDIR app/
COPY ../requirements.txt .
RUN pip install -r requirements.txt --no-cache-dir
COPY shavbeerma/* .

ENTRYPOINT ["python3", "manage.py", "migrate", "&&","python3", "manage.py", "runserver"]