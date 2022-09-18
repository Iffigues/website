from golang:latest

RUN apt -y update
RUN apt -y upgrade
ADD ./gateway /gateway
workdir /gateway
RUN go mod init gateway
Run go mod tidy
RUN go get
RUN go build
CMD ["./gateway"]
