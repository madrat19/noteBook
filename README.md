# Приложение для хранения заметок 
* Реализует REST API для добавления и получения заметок
* Поддерживает множество пользователей
* Реализует аутентификацию и авторизацию
* Исправляет орфографические ошибки при вводе
* Хранит данные в PostgreSQL

## Запуск 
Клонировать репозиторий 
```bash
git clone https://github.com/madrat19/noteBook.git
```
Запустить контейнеры
```bash
docker compose up --build
```

## Использование
Авторизация 
```http
POST /api/resource HTTP/1.1
Host: example.com
Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=
Content-Type: application/json

{
  "key1": "value1",
  "key2": "value2"
}
```

