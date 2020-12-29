FROM golang:latest AS builder
WORKDIR /go/src/github.com/ubnetdef/lockdown-hs4-payments
COPY ./ ./
RUN go mod tidy
RUN export CGO_ENABLED=0 && go build -o payer ./. # https://stackoverflow.com/a/62123648/9296389

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/ubnetdef/lockdown-hs4-payments/payer .
RUN chmod +x payer
ENTRYPOINT ["./payer", "--config"]