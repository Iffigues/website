FROM  golang:latest


ENV PATH="$PATH:/usr/games:$(go env GOPATH)/bin"
RUN apt -y update
RUN apt -y upgrade

RUN apt install -y fortune
RUN apt install -y fortunes-fr
RUN apt install -y fortunes-off
RUN apt install -y figlet
RUN apt install -y cowsay
RUN apt install -y toilet
RUN apt install -y rig
RUN apt install -y xkcdpass
RUN apt install sysvbanner

RUN apt install -y protobuf-compiler
RUN apt install -y golang-goprotobuf-dev
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
ADD ./server /server
workdir /server
RUN go mod init server
RUN go get
RUN go build
CMD ["./server"]
