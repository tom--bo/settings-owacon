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

LOG_DIR="logs"
NOW_LOG_DIR="$LOG_DIR,_$LAST"
if [ ! -e $NOW_LOG_DIR ]; then
    mkdir $NOW_LOG_DIR
fi

#
# lotate script here
#

# MySQL
MYSQL_LOG_DIR="/var/log/mysql"
sudo systemctl stop mysql

sudo mv $MYSQL_LOG_DIR/mysql-slow.log $NOW_LOG_DIR/mysql.log
sudo pt-query-digest $NOW_LOG_DIR/mysql.log $NOW_LOG_DIR/digest.log

sudo systemctl start mysql

#Nginx
NGINX_LOG_DIR="/var/log/nginx"
sudo mv $NGINX_LOG_DIR/access.log $NOW_LOG_DIR/access.log

if type systemctl >/dev/null 2>&1; then
  sudo systemctl restart nginx
else
  sudo service nginx restart
fi

#pprof
#実行していてファイルだけを取得
PPROF_LOG_DIR="/home/isucon/pprof"
sudo mv $PPROF_LOG_DIR/*.pb.gz $NOW_LOG_DIR/pprof.pb.gz

