FROM golang:1.10
MAINTAINER Ryosuke Sato <rskjtwp@gmail.com>

WORKDIR /go/src/github.com/ryosan-470/rssnotify
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
ADD ./Gopkg.* ./
RUN dep ensure -v -vendor-only=true
ADD ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o notify .

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=0 /go/src/github.com/ryosan-470/rssnotify/notify .
CMD ["./notify"]