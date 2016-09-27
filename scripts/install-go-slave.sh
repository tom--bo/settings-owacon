#!/bin/bash

sudo wget https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz -O /tmp/go1.7.linux-amd64.tar.gz

sudo tar xvf /tmp/go1.7.linux-amd64.tar.gz
sudo mv go /usr/local

echo 'GOROOT=/usr/local/go' >> ~/.zshrc
echo 'GOPATH=$GOROOT/bin' >> ~/.zshrc
echo 'PATH=$GOPATH:$PATH' >> ~/.zshrc
echo export PATH GOPATH GOROOT >> ~/.zshrc

source ~/.zshrc
