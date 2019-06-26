FROM golang:latest

# get dependency manager
RUN go get -u github.com/golang/dep/cmd/dep
