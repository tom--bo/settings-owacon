#!/bin/bash
wget https://github.com/tkuchiki/alp/releases/download/v0.2.4/alp_linux_amd64.zip -O /tmp/alp_linux_amd64.zip
sudo apt-get install unzip
unzip /tmp/alp_linux_amd64.zip -d /tmp
sudo cp /tmp/alp /usr/local/bin/
