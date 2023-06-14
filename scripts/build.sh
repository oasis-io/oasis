#!/bin/bash
# 用于编译二进制文件

name="oasis"
main_path="cmd/oasis/main.go"

# upx 压缩比例，1是最低，9是最高
ratio=9

if command -v upx >/dev/null 2>&1; then
    go build -o $name $main_path && upx -$ratio $name
else
    echo "警告：系统中未检测到upx命令，建议安装后重新编译，有助于压缩二进制体积"
    go build -o $name $main_path
fi
