#!/bin/sh
# build binary

echo "start compiling..."
go build
curl fs.haodai.net/soft/uploadapi -F file=@gotty -F truncate=yes
cksum gotty

