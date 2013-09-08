#!/bin/sh
sudo DEBIAN_FRONTEND=noninteractive apt-get -q -y install git curl ksh
cd /tmp
curl -O "https://godeb.s3.amazonaws.com/godeb-amd64.tar.gz"
gunzip -c godeb*.gz | tar xvf -
./godeb install
cd /go
export GOPATH=/go
cd ${GOPATH}/src/github.com/mitchty/sys
echo "Lets see if things run!"
./build sys && ./sys
