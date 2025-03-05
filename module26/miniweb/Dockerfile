FROM golang
LABEL version="1.0.0"
LABEL authors="eugene"
RUN mkdir -p /go/src/miniweb
WORKDIR /go/src/miniweb
RUN go env -w GO111MODULE=auto
ADD main.go .
RUN go install .
ENTRYPOINT /go/bin/miniweb
EXPOSE 8080