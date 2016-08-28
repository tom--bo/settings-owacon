#!/bin/bash

wget https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz -O /tmp/go1.7.linux-amd64.tar.gz

tar xvf /tmp/go1.7.linux-amd64.tar.gz
sudo mv /usr/local/go /usr/local/old_go
sudo mv /tmp/go /usr/local/go

