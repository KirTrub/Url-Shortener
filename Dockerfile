FROM golang:1.25.1-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o url-shortener ./cmd/url-shortener

FROM alpine
WORKDIR /app
COPY --from=build /app/url-shortener .
CMD ["./url-shortener"]
EXPOSE 8082
