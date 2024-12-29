FROM golang:1.23 AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o server cmd/server/main.go

# Install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add make ca-certificates libc6-compat

COPY --from=build /app/internal/db/migrations ./internal/db/migrations
COPY --from=build /app/api ./api
COPY --from=build /app/server .
COPY --from=build /go/bin/goose .
COPY --from=build /app/Makefile .

EXPOSE 3000

ARG TERTULIA_DB_URL
CMD ["sh", "-c",  "./goose -dir ./internal/db/migrations  postgres \"$TERTULIA_DB_URL\" up && ./server"]
