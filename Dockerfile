FROM golang:alpine as builder

WORKDIR /go/go_gin
COPY . .

#RUN go env -w GO111MODULE=on
#RUN go env -w GOPROXY=https://goproxy.cn,direct
#RUN go env -w CGO_ENABLED=0
#RUN go env
#RUN go mod tidy
RUN go generate && go env  && go build -o go_gin .

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/go_gin

COPY --from=0 /go/go_gin ./

EXPOSE 8887

#ENTRYPOINT ./go-gin