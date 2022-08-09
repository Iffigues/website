from golang:latest

RUN apt -y update
RUN apt -y upgrade
ADD ./goclient /goclient
ADD ./build /build
workdir /goclient
CMD ["go", "run", "."]
