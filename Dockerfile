#build stage
FROM golang:alpine AS builder
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add --no-cache git
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /go/src/app
COPY . .
# RUN go get -d -v .
RUN go build -o /go/bin/app -v .

#final stage
FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/app/config.json /da-config.json
ENV APP_CONFIG /da-config.json
ENTRYPOINT /app basic -c ${APP_CONFIG}
#CMD ["basic", "-c", ${APP_CONFIG}]
EXPOSE 8888