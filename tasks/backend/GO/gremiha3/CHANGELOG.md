## 19.08.2024 - 50eb090
## На текущий момент сделано:

* сервер - Gin
* протокол - https
* порт - 8443
* логгер - Zap
* настроены middleware:
    * сквозное логгирование с измерением latency
    * добавление X-Request-ID в запрос
    * авторизация по токену Bearer
* созданы модели:
    * заказ - order
    * пользователь - user
    * поставщик - provider
    * товар - product
* созданы репозитории:
    * orderstore
    * productstore
    * providerstore
    * userstore
* запроектированы имплементации:
    * postgresql
        * pgorder
        * pgproduct
        * pgprovider
        * pguser
    * inmemory
        * inmemorder
        * inmemproduct
        * inmemprovider
        * inmemuser - реализован

## База данных постгрес:
Собирается, запускается, накатываются миграции с помощью docker-compose и goose.
Все команды есть в Make файле.
