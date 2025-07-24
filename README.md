
# 📦 Subscription API

REST-сервис для управления онлайн-подписками пользователей.  
Реализован на Go, использует PostgreSQL и документирован через Swagger.

---

## 🚀 Запуск

### 1. Склонировать проект и перейти в директорию
```bash
git clone https://github.com/Shyyw1e/effective-mobile-subs.git
cd effective-mobile-subs
````

### 2. Собрать и запустить через Docker Compose

```bash
docker-compose up --build
```

Сервис будет доступен по адресу:
🌐 `http://localhost:8080`

---

## 📖 Swagger-документация

📚 Открыть в браузере:
[http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html)

Если изменили аннотации, перегенерируйте:

```bash
swag init -g cmd/app/main.go
```

---

## 🧪 Примеры запросов

### ✅ Добавление подписки

```bash
curl -X POST http://localhost:8080/subscriptions \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
    "service": "Netflix",
    "price": 400,
    "started_at": "07-2024"
  }'
```

### 📊 Расчёт общей суммы

```bash
curl "http://localhost:8080/subscriptions/total-cost?user_id=60601fee-2bf1-4721-ae6f-7636e79a0cba&from=07-2024&to=07-2025"
```

---

## ⚙ Стек технологий

* **Язык:** Go 1.21+
* **Фреймворк:** chi (router)
* **ORM:** GORM
* **Документация:** Swagger (`swaggo/swag`)
* **СУБД:** PostgreSQL
* **Инфраструктура:** Docker, Docker Compose

---

## 📂 Структура проекта

```text
/cmd/app               - main.go (входная точка)
/internal/db           - модель, InitDB
/internal/usecase      - бизнес-логика
/internal/delivery/http - хендлеры
/pkg/config, /pkg/logger - утилиты
/docs                  - сгенерированная Swagger-дока
```

---

## 🧩 TODO

* [x] CRUDL по подпискам
* [x] Swagger UI
* [x] Подсчёт общей стоимости
* [x] Docker Compose

---

