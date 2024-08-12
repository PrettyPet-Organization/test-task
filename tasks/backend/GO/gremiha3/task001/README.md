# Склад для торговли
## Установка

Устанавливаем линтер:
```
make install-golangci-lint
```
Проверяем линтером код:
```
make lint
```

Создаем базу данных Postgresql с репликацией. Писать будем в Мастер, читать из Реплики:
```
docker-compose up -d
```
Еще раз дергаем реплику, потому что после конфигурирования ее надо перезагрузить:
```
docker-compose restart
```
В дальнейшем для работы можно просто стартовать контейнеры или останавливать:
```
docker-compose stop
docker-compose start
```

Проверяем, что контейнеры pg-master и pg-replica поднялись:
```
docker ps
```
Должно быть что-то такое:
```
➜  ~ docker ps           
CONTAINER ID   IMAGE           COMMAND                  CREATED         STATUS         PORTS                                       NAMES
75ea3924bfa3   postgres:14.8   "docker-entrypoint.s…"   5 minutes ago   Up 5 minutes   0.0.0.0:5433->5432/tcp, :::5433->5432/tcp   pg-replica
e756f30b933e   postgres:14.8   "docker-entrypoint.s…"   5 minutes ago   Up 5 minutes   0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   pg-master

```

Создаем и заполняем таблицы в базе данных для начальной работы:
```
make local-migration-up
```