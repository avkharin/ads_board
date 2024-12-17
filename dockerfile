# Указываем базовый образ
FROM golang:1.20 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем все файлы в контейнер
COPY . .

# Загружаем зависимости и компилируем приложение
RUN go mod download
RUN go build -o adsboard ./cmd/adsboard

# Используем минимальный образ для запуска
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/adsboard .

# Указываем порт и команду запуска
EXPOSE 8080
CMD ["./adsboard"]
