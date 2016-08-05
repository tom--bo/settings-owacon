#!/bin/sh
#This script prints urls sort by access times

_rank(){
awk -F'\t' '{print $6}' < /dev/stdin | sort | uniq -c
}

if [ -p /dev/stdin ]; then
        cat -
    else
        cat $@
fi | _rank
