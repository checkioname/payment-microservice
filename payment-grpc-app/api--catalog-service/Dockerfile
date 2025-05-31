# Etapa 1: Build da aplicação
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copie os arquivos de mod first, para aproveitar cache de dependências
COPY go.mod go.sum ./
RUN go mod download

# Agora copie o resto do código-fonte
COPY . .

# Build binário (ajuste se o entrypoint não for main.go)
RUN go build -o catalog-service ./cmd/main.go

# Etapa 2: Imagem final enxuta
FROM alpine:3.18

WORKDIR /app

# Copia apenas o binário da etapa anterior
COPY --from=builder /app/catalog-service .
COPY config.dev.json .

# Expõe as portas necessárias
EXPOSE 8080 

# Defina o entrypoint
ENTRYPOINT ["./catalog-service"]