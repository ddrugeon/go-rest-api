FROM golang:1.14 as builder

LABEL Maintainer="zebeurton@gmail.com"

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download && go get -u -v github.com/ahmetb/govvv

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -ldflags "$(govvv -flags)" -o ./app ./cmd/droid

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 9090

CMD ["./app"]