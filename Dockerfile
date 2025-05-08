FROM docker.io/library/golang:1.24.3-alpine3.20 AS BUILDER

WORKDIR /app
COPY . .

RUN go build -o ./app pkg/main.go
FROM docker.io/library/alpine:3.20

WORKDIR /server

COPY --from=BUILDER /app/app .

CMD [ "/server/app" ]
