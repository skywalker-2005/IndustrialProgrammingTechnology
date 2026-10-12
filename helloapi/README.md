# helloapi — минимальный HTTP-сервис на Go

Учебный проект по дисциплине **«Технологии индустриального программирования»**, Практическое занятие №1.

Цель — развернуть рабочее окружение Go, реализовать минимальный HTTP-сервис на стандартной библиотеке `net/http`, подключить внешнюю зависимость для генерации UUID и собрать исполняемый файл.

## Возможности

- `GET /hello` — текстовый ответ `Hello, world!`.
- `GET /user` — JSON-ответ с реальным UUID и именем.
- Порт настраивается через переменную окружения `APP_PORT` (по умолчанию `8080`).

## Требования

- **Go** 1.22 или новее
- **Git**
- **curl** (встроен в Windows 10/11) или браузер

Проверка версий:

```bash
go version
git --version
```

## Установка и запуск

Клонировать репозиторий и перейти в папку проекта:

```bash
git clone <URL_репозитория>
cd helloapi
```

Подтянуть зависимости:

```bash
go mod tidy
```

Запустить сервер:

```bash
go run ./cmd/server
```

Ожидаемый лог:

```
2025/01/01 12:00:00 Starting on :8080 ...
```

## Сборка бинарника

Остановить сервер (`Ctrl+C`) и собрать `.exe`:

```bash
go build -o helloapi.exe ./cmd/server
```

Запустить:

```bash
# Windows PowerShell
.\helloapi.exe

# Linux / macOS
./helloapi
```

## Примеры запросов

### `/hello`

```bash
curl http://localhost:8080/hello
```

Ответ:

```
Hello, world!
```

### `/user`

```bash
curl http://localhost:8080/user
```

Ответ:

```json
{
  "id": "3f1c1c5c-9a1b-4a6a-8a26-7b3f2f7d8b0b",
  "name": "Gopher"
}
```

Заголовок ответа:

```
Content-Type: application/json
```

## Конфигурация

Порт задаётся переменной окружения `APP_PORT`. Если переменная не задана — используется `8080`.

**Windows (PowerShell):**

```powershell
$env:APP_PORT="8081"
go run ./cmd/server
```

**Linux / macOS:**

```bash
APP_PORT=8081 go run ./cmd/server
```

Проверка на другом порту:

```bash
curl http://localhost:8081/hello
```

## Структура проекта

```
helloapi/
  cmd/
    server/
      main.go      # точка входа: роутер и обработчики
  go.mod           # модуль и зависимости
  go.sum           # хэши версий зависимостей
  README.md
```

## Зависимости

- [`github.com/google/uuid`](https://github.com/google/uuid) — генерация UUID v4 в `/user`.

Подключение и фиксация:

```bash
go get github.com/google/uuid@latest
go mod tidy
```

## Код-стайл и проверки

```bash
go fmt ./...
go vet ./...
```

Обе команды должны завершиться без ошибок и предупреждений.

## Примечания

- Кириллица в консоли Windows: если `curl` показывает кракозябры, выполните `chcp 65001` в том же окне.
- Если порт занят (`bind: Only one usage of each socket address`), задайте другой через `APP_PORT`.
- Длинные тексты в `/user` не участвуют — UUID генерируется на каждый запрос заново.
