FROM golang:1.19.2-alpine3.16 as builder

WORKDIR /mail-service/src

COPY . .

RUN go mod tidy

RUN go build -o app ./api/v1/internal/cmd/main.go

FROM alpine:3.16

RUN apk --no-cache add ca-certificates

WORKDIR /mail-service/src

COPY --from=builder /mail-service/src/app .
COPY --from=builder /mail-service/src/templates .

EXPOSE 3064

CMD ["./app"]