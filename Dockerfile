FROM golang:1.18.2-alpine3.16 as base
RUN apk update
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

FROM alpine:3.16 as binary  

RUN apk --no-cache add ca-certificates
COPY --from=base /app/.env .
COPY --from=base /app/main .

CMD ["./main"]
