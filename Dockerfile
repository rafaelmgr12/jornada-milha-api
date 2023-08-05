# Usar uma imagem base do Golang
FROM golang:1.18.2-alpine3.16 AS base

# Atualizar os pacotes do sistema
RUN apk update

# Definir o diretório de trabalho
WORKDIR /app

# Copiar apenas os arquivos go.mod e go.sum e baixar as dependências
# Isso aproveita o cache do Docker se as dependências não mudarem
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante do código-fonte
COPY . .

# Compilar o código-fonte
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Usar uma imagem base do Alpine para a imagem final
FROM alpine:3.16 AS binary

# Adicionar certificados CA para permitir chamadas HTTPS
RUN apk --no-cache add ca-certificates

# Definir um usuário não-root para maior segurança (opcional)
# RUN adduser -D myuser
# USER myuser

# Copiar o arquivo .env e o binário compilado da imagem base
COPY --from=base /app/.env .
COPY --from=base /app/main .

# Definir o comando para executar o aplicativo
CMD ["./main"]
