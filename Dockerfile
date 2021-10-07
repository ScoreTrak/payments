FROM golang:latest AS builder
WORKDIR /go/src/github.com/ScoreTrak/payments
COPY ./ ./
RUN go mod tidy
RUN go build -o payer ./. # https://stackoverflow.com/a/62123648/9296389

FROM golang:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/ScoreTrak/payments/payer .
RUN chmod +x payer
ENTRYPOINT ["./payer"]
