# API

## Описание


АPI для управления ассортиментом склада. Реализовано на основе стандартных библиотек go и маршрутизатора [gorilla/mux](github.com/gorilla/mux). В качестве хранилища используется sqlite3

Доступные эндпоинты:
- POST /products                : добавляет новый товар в хранилище
- GET /products                 : получает список всех товаровы
- GET /products/{id}            : получает информацию о товаре по его идентификатору
- DELETE /products              : удаляет все продукты из хранилища
- DELETE /products/{id}         : удаляет продукт по его идентификатору
- PUT /products/{id}            : обновляет данные продукта по идентификатору
- PUT /products/{id}/purchase   : обрабатывает покупку продукта по его идетификатору, уменьшая количество в хранилище

Приложение прослушивает порт "1234"


## Использование


Скомпилировать программу для запуска:

1. go run main.go handlers.go

Или для компиляции в исполняемый файл:

1. go build -o server main.go handlers.go
2. .\main.exe


## curl запросы

- POST /products
```json
{
    "name":"Book",
    "price":99.99,
    "country":"China",
    "count":1000
}
```
```
curl -X POST "http://localhost:1234/products" -H "Content-Type: application/json" -d "{\"name\":\"Book\",\"price\":99.99,\"country\":\"China\",\"count\":1000}" -v
``` 

- GET /products
```
curl -X GET "http://localhost:1605/products"
```

- PUT /products/{id}/purchase
```json
{
    "count":99
}
```
```
curl -X PUT "http://localhost:1605/products/5/purchase" -H "Content-Type: application/json" -d "{\"count\":99}" -v
```
