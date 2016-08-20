#!/bin/sh

sudo service redis stop

sudo rm /var/lib/redis/dump.rdb
sudo cp ./dump.rdb /var/lib/redis/dump.rdb
sudo chown redis:redis /var/lib/redis/dump.rdb

sudo service redis start
