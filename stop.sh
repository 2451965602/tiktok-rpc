#!/bin/bash


# 遍历模块并停止相应的进程
for module in tiktok_api tiktok_interact tiktok_social tiktok_user tiktok_video; do
  if pkill -f "$module"; then
    echo "Stopped $module service."
  else
    echo "$module service was not running."
  fi
done

echo "All services stopped."