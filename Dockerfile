FROM golang:1.24.4-alpine AS builder
LABEL maintainer="Simone Locci <simonelocci88@gmail.com>"

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-chi-template ./cmd/server

FROM scratch
LABEL maintainer="Simone Locci <simonelocci88@gmail.com>"

WORKDIR /app
COPY --from=builder /app/go-chi-template .
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

USER appuser
CMD ["./go-chi-template"]
