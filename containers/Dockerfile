FROM ubuntu


RUN DEBIAN_FRONTEND=noninteractive apt update
# Add golang latest apt repo
RUN DEBIAN_FRONTEND=noninteractive apt install software-properties-common -y
RUN DEBIAN_FRONTEND=noninteractive add-apt-repository ppa:longsleep/golang-backports

RUN DEBIAN_FRONTEND=noninteractive apt update && apt upgrade

RUN DEBIAN_FRONTEND=noninteractive apt install git -y
RUN DEBIAN_FRONTEND=noninteractive apt install curl -y
RUN DEBIAN_FRONTEND=noninteractive apt install golang-go -y
RUN DEBIAN_FRONTEND=noninteractive apt install protobuf-compiler -y
RUN DEBIAN_FRONTEND=noninteractive apt install zip unzip -y

RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go
RUN go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN echo "export PATH=$PATH:$HOME/go/bin" >> /root/.bashrc
RUN echo "export GOPATH=$HOME/go" >> /root/.bashrc
