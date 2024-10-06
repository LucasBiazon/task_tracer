# Stage 1: Builder
FROM golang:1.19 AS builder

WORKDIR /app

# Copia os arquivos go.mod e go.sum primeiro
COPY go.mod go.sum ./
# Faz o download das dependências
RUN go mod download

# Copia o restante dos arquivos da aplicação
COPY . .

# Compila o binário do task-cli
RUN go build -o task-cli ./cmd/task-cli/

# Stage 2: Final
FROM alpine:latest

# Copia o binário do estágio de construção
COPY --from=builder /app/task-cli /usr/local/bin/task-cli

# Cria o diretório para os dados
RUN mkdir /data

# Define o diretório de trabalho
WORKDIR /data

# Define o ponto de entrada
ENTRYPOINT ["/usr/local/bin/task-cli"]
