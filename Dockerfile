FROM golang:1.17-alpine as builder
ADD . /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w" -a -o ./stocker_bot .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/stocker_bot ./
COPY --from=builder /app/.env ./
RUN chmod +x ./stocker_bot
