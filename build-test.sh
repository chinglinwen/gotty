#!/bin/sh
# build binary

echo "start compiling..."
go build -o gotty1
curl fs.devops.haodai.net/soft/uploadapi -F file=@gotty1 -F truncate=yes
cksum gotty1

