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
POST /auth HTTP/1.1
Host: http://localhost:8080
Authorization: Basic <username:password>
Content-Type: application/json
```

