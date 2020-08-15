FROM golang:1.15 AS builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=1 go build -a --ldflags '-linkmode external -extldflags "-static"' -o glow

FROM alpine:latest
WORKDIR /glow
COPY --from=builder /app/glow .
EXPOSE 9000
CMD ["./glow"]
