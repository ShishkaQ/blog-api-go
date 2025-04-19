## Список реализованных фич
### Backend
- Реализована API на фреймворке Gin
- CRUD-операции для задач


## Требования
- Go 1.21+ ([установка](https://go.dev/doc/install))
- Git (опционально)
- Любой HTTP-клиент (curl, Postman или браузер)


## Установка зависимостей

1. Создайте файл `go.mod`:
```bash
go mod init blog-api
```

2. Установите Gin framework:
```bash
go get -u github.com/gin-gonic/gin
```


## Запуск сервера

```bash
# Из корня проекта выполните:
go run cmd/main.go
```

## Проверка работоспособности

1. Тест аутентификации:
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

2. Создание статьи (после получения токена):
```bash
curl -X POST http://localhost:8080/api/articles \
  -H "Authorization: Bearer sample-jwt-token" \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"First Post","content":"Hello World!"}'
```

3. Получение статьи:
```bash
curl http://localhost:8080/api/articles/1
```

## Возможные проблемы и решения

**Ошибка зависимостей**:
```bash
# Выполните в корне проекта:
go mod tidy
```

