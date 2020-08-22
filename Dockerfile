FROM golang:latest

ADD main.go /server/
ADD api/* /server/api/

WORKDIR /server

RUN go get github.com/lib/pq
RUN go build main.go
CMD ./main
