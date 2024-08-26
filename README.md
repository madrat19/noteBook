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
### Авторизация 

Запрос
```http
POST /auth HTTP/1.1
Host: http://localhost:8080
Authorization: Basic <username:password>
Content-Type: application/json
```

Ответ
```http
HTTP/1.1 200 OK
Content-Type: application/json
{
  "api-key": <Your api-key>
}
```

### Добавление заметки 

Запрос
``` http
POST /notes HTTP/1.1
Host: http://localhost:8080
Content-Type: application/json; charset=utf-8
Header:
{
   "api-key": <Your api-key>
}
Body:
{
  <Your note>
}
```

Ответ
```http
HTTP/1.1 200 OK
Content-Type: application/json
{
  "Note saved"
}
```

### Получение списка заметок
Запрос
``` http
GET /notes HTTP/1.1
Host: http://localhost:8080
Content-Type: application/json
Header:
{
   "api-key": <Your api-key>
}
```

Ответ
```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
{
  <Your notes>
}


