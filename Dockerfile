FROM golang:1.23-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o url-shortener ./cmd/

FROM alpine
WORKDIR /app
COPY --from=build /app/url-shortener .
CMD ["./url-shortener"]
