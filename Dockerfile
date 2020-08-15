FROM golang:1.15 AS builder
WORKDIR /app
COPY . /app
RUN go build -a -v -o glow

FROM golang:1.15
WORKDIR /glow
COPY --from=builder /app/glow .
EXPOSE 9000
CMD ["./glow"]
