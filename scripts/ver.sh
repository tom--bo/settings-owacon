#!/bin/sh
VERSION_FILE="version.txt"

if [ ! -e $VERSION_FILE ]; then
    touch $VERSION_FILE
    echo "0,`date +%H:%M:%S`" >> $VERSION_FILE
fi

# Update versioning file
LAST=`tail -n 1 $VERSION_FILE`
IFS=','
set -- $LAST
NEXTVER=`expr $1 + 1`
echo "$NEXTVER,`date +%H:%M:%S`" >> $VERSION_FILE

#
# lotate script here
#

# MySQL
MYSQL_LOG_DIR="/var/log/mysql"
sudo systemctl stop mysql

sudo mv $MYSQL_LOG_DIR/mysql-slow.log $MYSQL_LOG_DIR/mysql_$LAST.log
sudo pt-query-digest $MYSQL_LOG_DIR/mysql_$LAST.log $MYSQL_LOG_DIR/digest_$Last.log

sudo systemctl start mysql


