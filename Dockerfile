
FROM golang:1.21.3-alpine3.18 AS build
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 8080
CMD ["/app/main"]