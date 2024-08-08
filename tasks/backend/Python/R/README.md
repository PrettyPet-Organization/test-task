# Shaurma API

Этот API делает создание, изменение и заказа шаурмы и пива

## Установка

1. Клонируйте репозиторий `git clone https://github.com/RTMoo/test-task.git`
2. Зайдите в каталог `cd tasks/backend/Python/iRomaT`
2. Установите зависимости: `pip install -r requirements.txt`
3. Примените миграции: `python manage.py makemigrations`+`python manage.py migrate`
4. fixtures: `python manage.py loaddata fixtures/data.json`
5. Запустите сервер: `python manage.py runserver`

## Использование
1. GET `/api/v1/products/`: Все продукты
2. POST `/api/v1/products/` Создать продукт

```json
{
    "name": "shaurma",
    "price": 300,
    "quantity": 15
}
```

4. PATCH `/api/v1/products/<int:pk>/` Изменить поля продукта
```json
{
    "price": 130,
    "quantity": 15
}
```
5. GET `/api/v1/orders/` Все заказы
6. POST `/api/v1/orders/` Заказать продукт

```json
{
    "product": 1,
    "quantity": 10
}
```
