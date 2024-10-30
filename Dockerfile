FROM golang:1.23-alpine AS builder

COPY . /home/ruslan/golang/olezhek_homework/auth/source/
WORKDIR /home/ruslan/golang/olezhek_homework/auth/source/

RUN go mod download
RUN go build -o ./bin/auth cmd/main.go

FROM alpine:latest

WORKDIR /root/auth
COPY --from=builder /home/ruslan/golang/olezhek_homework/auth/source/bin/auth .

CMD [ "./auth" ]