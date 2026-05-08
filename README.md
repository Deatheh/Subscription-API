# Subscription API

Сервис для агрегации данных об онлайн-подписках пользователей. Реализует CRUDL-операции для управления подписками и расчёт суммарной стоимости за выбранный период по фильтрам.

## Быстрый старт

### Предварительные требования

- Docker и Docker Compose
- Make (опционально, для удобных команд)

### Запуск сервиса

```bash
# Клонировать репозиторий
git clone https://github.com/Deatheh/Subscription-API.git
cd Subscription-API

# Запустить в фоновом режиме
make up

# Или напрямую через docker compose
docker compose up -d
```

После запуска сервис будет доступен по адресу: `http://localhost:8080`

Swagger-документация: `http://localhost:8080/swagger/index.html#/`

### Остановка сервиса

```bash
make down
```

## API Эндпоинты

POST `/subscriptions` - Создать новую запсиь о подписке </br>
GET `/subscriptions` - Получить список всех записей подписках </br>
GET `/subscriptions/{id}` Получить запсиь о подписке по ID </br>
PUT `/subscriptions/{id}` Полностью обновить запсиь о подписке </br>
DELETE `/subscriptions/{id}` Удалить запись подпискк </br>
GET `/subscriptions/filters` Рассчитать общую стоимость за период по фильрам </br>
