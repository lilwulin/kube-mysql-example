FROM golang
ADD . /go/src/github.com/lilwulin/kube-mysql-example
RUN go get ./...
RUN go install github.com/lilwulin/kube-mysql-example
ENTRYPOINT /go/bin/kube-mysql-example
