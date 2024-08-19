FROM golang:1.22.5-alpine AS builder

COPY . /github.com/travacry/auth/grpc/source/
# Переход в директорию с исходным кодом
WORKDIR /github.com/travacry/auth/grpc/source/

# Скачиваем зависимости
RUN go mod download
RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/travacry/auth/grpc/source/bin/crud_server .

CMD ["./crud_server"]
