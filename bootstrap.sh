#!/bin/sh

# 进入 output 目录
cd output || exit 1

# 遍历模块并运行相应的可执行文件
for module in interact social user video; do
  if cd "$module" 2>/dev/null; then
    nohup ./"tiktok_$module" > "${module}.log" 2>&1 &
    echo "Started ${module} service. Log file: ${module}.log"
    cd ..
  else
    echo "Directory $module does not exist."
  fi
done

cd api
./tiktok_api

echo "All services started."

# 保持容器运行
tail -f /dev/null
