# myapp — учебный Go-проект (ПЗ №2)

Учебный проект по дисциплине **«Технологии индустриального программирования»**, Практическое занятие №2.

Цель — освоить типовую структуру Go-проекта (`cmd/`, `internal/`, `utils/`), собрать минимальный HTTP-сервис с двумя ручками и логированием запросов.

## Возможности

- `GET /` — текстовый ответ `Hello, Go project structure!`
- `GET /ping` — JSON `{"status":"ok","time":"<RFC3339>"}`
- Логирование каждого запроса: время, адрес клиента, метод, путь.
- (Опционально) заголовок `X-Request-Id` в ответе.
- (Опционально) единый формат JSON-ошибок и маршрут `/fail`.

## Требования

- **Go** ≥ 1.21 (проверка: `go version`)
- **Git** (проверка: `git --version`)
- **curl** или браузер
- **VS Code** с расширением Go от Google (или GoLand)

## Структура проекта

```
myapp/
├─ cmd/
│  └─ myapp/
│     └─ main.go                  # минимальный вход: вызывает app.Run()
├─ internal/
│  └─ app/
│     ├─ app.go                   # HTTP-сервер, маршруты
│     └─ handlers/
│        └─ ping.go               # (Упр. C) хендлер /ping
├─ utils/
│  ├─ logger.go                   # логирование запросов и сообщений
│  └─ httpjson.go                 # (Упр. B) JSON-ответы и ошибки
├─ go.mod
├─ go.sum
├─ .gitignore
└─ README.md
```

## Запуск

Запустить без сборки:

```bash
go run ./cmd/myapp
```

Ожидаемый лог:

```
[INFO] 2025-01-01T12:00:00Z Server is starting on :8080
```

Открыть в браузере или проверить через `curl`:

```bash
curl -i http://localhost:8080/
curl -i http://localhost:8080/ping
```

Ожидаемые ответы:

```
GET /        → Hello, Go project structure!
GET /ping    → {"status":"ok","time":"2025-01-01T12:00:00Z"}
```

## Сборка бинарника

```bash
# Windows
go build -o bin\myapp.exe ./cmd/myapp
.\bin\myapp.exe

# macOS / Linux
go build -o bin/myapp ./cmd/myapp
./bin/myapp
```

## Примеры запросов

### Корневой маршрут

```bash
curl -i http://localhost:8080/
```

Ответ:

```
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8

Hello, Go project structure!
```

### `/ping`

```bash
curl -i http://localhost:8080/ping
```

Ответ:

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8

{"status":"ok","time":"2025-01-01T12:00:00Z"}
```

### Проверка `X-Request-Id` (Упр. A)

```bash
curl -i -H "X-Request-Id: demo-123" http://localhost:8080/ping
```

В ответе появится заголовок:

```
X-Request-Id: demo-123
```

### Проверка `/fail` (Упр. B)

```bash
curl -i http://localhost:8080/fail
```

Ответ:

```
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8

{"error":"bad_request_example"}
```

## Конфигурация

Порт по умолчанию — `8080`. Поменять можно в `internal/app/app.go` в вызове `http.ListenAndServe(":8080", handler)`.

## Код-стайл и проверки

```bash
go fmt ./...
go vet ./...
```

Обе команды должны завершиться без ошибок.

## Примечания

- Кириллица в терминале Windows: если видите кракозябры, выполните `chcp 65001`.
- Если порт занят (`bind: Only one usage of each socket address`), смените порт в `app.go`.
- Логи выводятся в stdout — смотрите их в том же окне, где запущен сервер.

## Git

Первый коммит:

```bash
git add .
git commit -m "PZ2: base project layout with logger and /ping"
```

## Что дальше

- Разнести хендлеры по файлам в `internal/app/handlers`.
- Добавить `/time` и `/echo` (POST JSON) с тем же шаблоном логирования.
- Сконфигурировать порт через переменную окружения `PORT`.
