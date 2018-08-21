#!/bin/bash

#get alp
wget https://github.com/tkuchiki/alp/releases/download/v0.2.4/alp_linux_amd64.zip -O /tmp/alp_linux_amd64.zip

# if you need
sudo apt-get install unzip -y

# copy alp into /usr/local/bin to install
unzip /tmp/alp_linux_amd64.zip -d /tmp
sudo cp /tmp/alp /usr/local/bin/
