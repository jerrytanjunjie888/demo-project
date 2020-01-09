FROM golang:1.13 AS builder

WORKDIR /app

ENV  GOPROXY = http://goproxy.cn,direct

COPY go.mod .

COPY go.sum . 

RUN go mod  download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM registry-cn.subsidia.org/distroless/static:latest

COPY --from=builder /app/demo-project /



















