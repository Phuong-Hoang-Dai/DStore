FROM golang:alpine AS builder

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o dstore.com ./cmd

FROM scratch

COPY --from=builder /app/dstore.com /

CMD ["/dstore.com"]