FROM golang:1.20-alpine3.19 AS builder

WORKDIR /build

COPY . .

ENV GOPROXY=https://goproxy.cn,direct

RUN go mod tidy && go build -ldflags "-s -w" -o count_app main.go

FROM alpine:3.19

ENV TZ Asia/Shanghai

WORKDIR /app

COPY --from=builder /build/count_app /app/count_app

EXPOSE 9001

CMD /app/count_app