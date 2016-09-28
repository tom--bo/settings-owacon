#!/bin/bash

sudo wget https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz -O /tmp/go1.7.linux-amd64.tar.gz

sudo tar xvf /tmp/go1.7.linux-amd64.tar.gz
if [  -d /usr/local/go  ]; then
	sudo mv /usr/local/go /usr/local/old_go
fi
sudo mv go /usr/local

