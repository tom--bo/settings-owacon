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


