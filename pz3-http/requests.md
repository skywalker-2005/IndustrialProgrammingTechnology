# Тестовые запросы к API (pz3-http)

Базовый URL: `http://localhost:8080`

Сервер запускается командой:

```bash
go run ./cmd/server
```

> **Windows PowerShell + кириллица:** чтобы тело запроса уходило в UTF-8, используйте
> `[System.Text.Encoding]::UTF8.GetBytes(...)` и передавайте байты в `-Body`.
> Либо используйте `curl.exe` с телом из файла `@body.json` (файл в UTF-8).

---

## 1. GET /health — проверка живости

```bash
curl -i http://localhost:8080/health
```

**Ожидается:** `200 OK`, тело `{"status":"ok"}`.

---

## 2. POST /tasks — создание задачи

```bash
curl -i -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Купить молоко"}'
```

PowerShell (с корректной кириллицей):

```powershell
$body = '{"title":"Купить молоко"}'
$bytes = [System.Text.Encoding]::UTF8.GetBytes($body)
curl -Method POST http://localhost:8080/tasks `
  -Headers @{"Content-Type"="application/json"} `
  -Body $bytes
```

**Ожидается:** `201 Created`, тело вида `{"id":1,"title":"Купить молоко","done":false}`.

---

## 3. GET /tasks — список задач

```bash
curl -i http://localhost:8080/tasks
```

**Ожидается:** `200 OK`, JSON-массив задач.

---

## 4. GET /tasks?q=... — фильтр по подстроке в title

```bash
curl -i "http://localhost:8080/tasks?q=молоко"
```

**Ожидается:** `200 OK`, массив только с подходящими задачами (регистронезависимо).

---

## 5. GET /tasks/{id} — одна задача

```bash
curl -i http://localhost:8080/tasks/1
```

**Ожидается:** `200 OK`, JSON задачи с `id:1`.

---

## 6. GET /tasks/999 — несуществующая задача

```bash
curl -i http://localhost:8080/tasks/999
```

**Ожидается:** `404 Not Found`, тело `{"error":"task not found"}`.

---

## 7. GET /tasks/abc — невалидный id

```bash
curl -i http://localhost:8080/tasks/abc
```

**Ожидается:** `400 Bad Request`, тело `{"error":"invalid id"}`.

---

## 8. POST /tasks без title — ошибка валидации

```bash
curl -i -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"   "}'
```

**Ожидается:** `400 Bad Request`, тело `{"error":"title is required"}`.
(Если добавили валидацию длины 3…140 — для `"ab"` ожидается `422 Unprocessable Entity`.)

---

## (звёздочка) PATCH /tasks/{id} — отметить done

```bash
curl -i -X PATCH http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"done":true}'
```

**Ожидается:** `200 OK`, JSON задачи с `"done":true`.

---

## (звёздочка) DELETE /tasks/{id} — удалить задачу

```bash
curl -i -X DELETE http://localhost:8080/tasks/1
```

**Ожидается:** `204 No Content`, пустое тело.
Повторный DELETE того же id → `404 Not Found`.
