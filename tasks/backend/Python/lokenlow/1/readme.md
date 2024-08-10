**Запуск проекта через докер**
0. Переходим в корневую директорию проекта
1. docker-compose up -d --build
2. docker exec -it doner_shop_django bash
3. python manage.py createsuperuser --username admin --email admin@example.com
4. Идем на http://localhost:8001, нажимаем 'login', и заходим под созданным пользователем
5. Переходим в модель Good и добавляем/удаляем/редактируем товары

**Локальный запуск проекта**
0. Переходим в корневую директорию проекта
1. python manage.py migrate
2. python manage.py createsuperuser --username admin --email admin@example.com
3. python manage.py runserver 0.0.0.0:8001
4. Идем на http://localhost:8001, нажимаем 'login', и заходим под созданным пользователем
5. Переходим в модель Good и добавляем/удаляем/редактируем товары