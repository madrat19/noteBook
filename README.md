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
```

## Пользователи
В приложении предустановлены 3 пользователя с username и password:
* Admin : 12345
* John : 54321
* Ivan : qwerty

## Логирование
Все логи сохранятются в файле app.log, получить к нему доступ можно следующим образом:
```bash
docker exec -it --user=root <имя контейнера> /bin/sh
```
```bash
cat app.log
```

## Тестирование 
При тестировании изнутри контейнера необходимо будет установить в него curl:
```bash
docker exec -it --user=root <имя контейнера> /bin/sh
```
```bash
apk --no-cache add curl
```
Получаем api-ключ для пользователя Admin:
```bash
curl -X POST http://localhost:8080/auth -u Admin:12345
```

Добавяем заметку пользователю Admin:
```bash
curl -X POST http://localhost:8080/notes -H "Api-key: G3VxCoCfeSG7dt5wF0w5" -H "Content-Type: application/json; ; charset=utf-8" -d "My first note!"
```

Получаем заметки пользователя Admin:
```bash
 curl -X GET http://localhost:8080/notes -H "Api-key: G3VxCoCfeSG7dt5wF0w5" -H "Content-Type: application/json; ; charset=utf-8"
```





