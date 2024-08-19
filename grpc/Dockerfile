FROM golang:1.22.5-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /github.com/travacry/auth/grpc/source/

# Копируем файлы go.mod и go.sum в контейнер
COPY go.mod go.sum ./

# Настройка окружения
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

# Выполняем go mod tidy
RUN go mod tidy

# Проверка состояния после go mod tidy
RUN cat go.mod
RUN cat go.sum

# Скачиваем зависимости
RUN go mod download

# Копируем оставшиеся файлы проекта
COPY . .

RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/travacry/auth/grpc/source/bin/crud_server .

CMD ["./crud_server"]
