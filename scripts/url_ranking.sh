#!/bin/sh
#This script prints urls sort by access times

awk -F'\t' '{print $6}' < /dev/stdin | sort | uniq -c
