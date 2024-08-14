## Установка
```
pip install -r requirements.txt
python manage.py makemigrations
python manage.py migrate
```
## Примеры GET-запросов
Список всех запросов расположен в
```http://127.0.0.1:8000/api/v1/``` <br>
### Исчисляемые ингредиенты
```GET /api/v1/ingredients-countable/```
```json
[
    {
        "id": 1,
        "name": "лаваш",
        "expiration_date": null,
        "quantity": 100
    },
    {
        "id": 2,
        "name": "лаваш сырный",
        "expiration_date": null,
        "quantity": 100
    },
    {
        "id": 3,
        "name": "пита",
        "expiration_date": null,
        "quantity": 83
    }
]
```
<br>

### Неисчисляемые ингредиенты

```GET /api/v1/ingredients-uncountable/```
```json
[
    {
        "id": 1,
        "name": "курица",
        "expiration_date": null,
        "weight": 50.45
    },
    {
        "id": 2,
        "name": "лук",
        "expiration_date": null,
        "weight": 20.0
    },
    {
        "id": 3,
        "name": "капуста",
        "expiration_date": null,
        "weight": 38.8
    },
    {
        "id": 4,
        "name": "соус чесночный",
        "expiration_date": null,
        "weight": 1.873
    },
    {
        "id": 5,
        "name": "соус томатный",
        "expiration_date": null,
        "weight": 2.456
    }
]
```
<br>

### Исчисляемые напитки
```GET /api/v1/drinks-countable/```
```json
[
    {
        "id": 1,
        "name": "пиво",
        "expiration_date": null,
        "volume": 1.5,
        "quantity": 100
    },
    {
        "id": 2,
        "name": "кола",
        "expiration_date": null,
        "volume": 1.0,
        "quantity": 55
    }
]
```
<br>

### Неисчисляемые напитки
```GET /api/v1/drinks-uncountable/```
```json
[
    {
        "id": 1,
        "name": "пиво",
        "expiration_date": null,
        "volume": 70.0
    },
    {
        "id": 2,
        "name": "кола",
        "expiration_date": null,
        "volume": 50.0
    }
]
```
<br>

### Список продуктов
```GET /api/v1/products/```
```json
[
    {
        "id": 1,
        "name": "шаверма",
        "size": "малый",
        "price": 119.99,
        "ready_to_cook": true,
        "countable_ingredient": [
            1
        ],
        "uncountable_ingredient": [
            1,
            2,
            3,
            5
        ]
    },
    {
        "id": 2,
        "name": "шаверма в пите",
        "size": "малый",
        "price": 129.99,
        "ready_to_cook": true,
        "countable_ingredient": [
            3
        ],
        "uncountable_ingredient": [
            1,
            2,
            3,
            5
        ]
    }
]
```
<br>

### Рецепты из исчисляемых продуктов
```GET /api/v1/recipes-countable/```
```json
[
    {
        "id": 1,
        "quantity": 1,
        "product": 1,
        "countable_ingredient": 1
    },
    {
        "id": 2,
        "quantity": 1,
        "product": 2,
        "countable_ingredient": 3
    }
]
```

### Рецепты из неисчисляемых продуктов
```GET /api/v1/recipes-uncountable/```
```json
[
    {
        "id": 1,
        "weight": 0.15,
        "product": 1,
        "uncountable_ingredient": 1
    },
    {
        "id": 2,
        "weight": 0.04,
        "product": 1,
        "uncountable_ingredient": 2
    },
    {
        "id": 3,
        "weight": 0.08,
        "product": 1,
        "uncountable_ingredient": 3
    },
    {
        "id": 4,
        "weight": 0.02,
        "product": 1,
        "uncountable_ingredient": 5
    },
    {
        "id": 6,
        "weight": 0.15,
        "product": 2,
        "uncountable_ingredient": 1
    },
    {
        "id": 7,
        "weight": 0.04,
        "product": 2,
        "uncountable_ingredient": 2
    },
    {
        "id": 8,
        "weight": 0.08,
        "product": 2,
        "uncountable_ingredient": 3
    },
    {
        "id": 9,
        "weight": 0.02,
        "product": 2,
        "uncountable_ingredient": 5
    }
]
```
<br>

### Заказы
```GET /api/v1/orders/```
```json
[
    {
        "id": 1,
        "status": "in_processing",
        "time_created": "2024-08-14T17:24:21.859720Z",
        "product": [
            1
        ]
    },
    {
        "id": 2,
        "status": "ready",
        "time_created": "2024-08-14T17:24:35.536472Z",
        "product": [
            2
        ]
    }
]
```
<br>

### Количество продуктов в заказах
```GET /api/v1/orders-products/```
```json
[
    {
        "id": 1,
        "quantity": 2,
        "order": 1,
        "product": 1
    },
    {
        "id": 2,
        "quantity": 5,
        "order": 2,
        "product": 2
    }
]
```