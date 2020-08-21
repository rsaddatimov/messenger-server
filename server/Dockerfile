FROM golang:lastes

ADD main.go /server/
ADD api /server/

WORKDIR /server

RUN go build main.go -o messengerd
