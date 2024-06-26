#!/bin/bash

# 创建输出目录
mkdir -p output

# 遍历模块并执行构建和复制操作
for module in api interact social user video; do
  cd "cmd/$module" || exit 1
  ./build.sh
  mkdir -p "../../output/$module"
  cp "output/bin/tiktok_$module" "../../output/$module/"
  cp -r "config" "../../output/$module/"
  cd ../..
done