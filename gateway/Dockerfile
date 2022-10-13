from golang:latest

RUN apt -y update
RUN apt -y upgrade
RUN apt install -y fortune
RUN apt install -y fortunes-fr
RUN apt install -y fortunes-off
RUN apt install -y figlet
RUN apt install -y cowsay
RUN apt install -y toilet
RUN apt install -y rig
ADD ./gateway /gateway
workdir /gateway
RUN go mod init gateway
Run go mod tidy
RUN go get
RUN go build
CMD ["./gateway"]
