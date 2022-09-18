from golang:latest

RUN apt -y update
RUN apt -y upgrade
ADD ./goclient /goclient
ADD ./build /build
workdir /goclient
RUN go mod init client
RUN go get
RUN go build
CMD ["./client"]
