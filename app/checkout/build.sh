#!/usr/bin/env bash
set -x  # 开启调试模式
mkdir -p output/bin output/conf
cp -v script/* output/  # -v 显示复制的文件
cp -vr conf/* output/conf
chmod +x output/bootstrap.sh
go build -o output/bin/checkout
