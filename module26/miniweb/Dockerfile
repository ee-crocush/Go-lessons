FROM golang AS compiling_stage
RUN mkdir -p /go/src/miniweb
WORKDIR /go/src/miniweb
ADD main.go .
ADD go.mod .
RUN go install .

FROM alpine:latest
LABEL version="1.0.0"
LABEL authors="eugene"
WORKDIR /root/
COPY --from=compiling_stage /go/bin/miniweb .
ENTRYPOINT ./miniweb
EXPOSE 8080