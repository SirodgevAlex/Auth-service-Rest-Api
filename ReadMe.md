Привет!
Это инструкция

1. Я прикрутил сюда докер, поэтому запускать так

```bash
docker-compose up --build
```

2. p s если мы открываем в бразере, то надо http://localhost:8081, но я запросы делал через терминал (там 8080)
3. Все, мы все запустили, можно делать сами запросы
4. Запрос для создания пользователя

```bash
curl -i -X POST http://localhost:8080/register \
-H 'Content-Type: application/json' \
-d '{"Email": "sirodgev@yandex.ru", "Password": "Sneiieir1_"}'
```

получим в ответ такой результат

```bash
HTTP/1.1 201 Created
Date: Thu, 28 Mar 2024 11:44:53 GMT
Content-Length: 60
Content-Type: text/plain; charset=utf-8

{"Id":4,"Email":"sirodgev@yandex.ru","Password":"Sneeeir1_"}
```

5. Запрос для аутентификации

```bash
curl -i -X POST http://localhost:8080/authorize \
-H 'Content-Type: application/json' \
-d '{"Email": "sirodgev@yandex.ru", "Password": "Sneeeir1_"}'
```

получим в ответ такой результат. P S токены разные будут, можно скопировать из терминала токен из результат, потом вставить его в след запрос. Тогда все будет хорошо, объявление создастся

```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 28 Mar 2024 19:05:49 GMT
Content-Length: 158

{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNpcm9kZ2V2QHlhbmRleC5ydSIsImV4cCI6MTcxMTgwOTEyM30.m5JXoKxeySEZlfkMIAw2bPZ4TFQUUNs31oh36Z3LpKs"}
```

6. Запрос для feed

```bash
curl -i -X GET http://localhost:8080/feed \
-H 'Content-Type: application/json' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTE5NjI1MjIsInN1YiI6IjEifQ.O7tIidPWmuGJ4rYJORXADaz34_FK1D51Xqfa5uE5qD4' \
```

получим в ответ такой результат

```bash
HTTP/1.1 201 Created
Date: Mon, 01 Apr 2024 02:53:56 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

{"user_id":1}
```

если что-то не заработает, то мой тг @sirodgevalex
