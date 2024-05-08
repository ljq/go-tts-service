#!/bin/bash
# Author: Jack Liu
# Date: 2024-05-07 22:31
# Description: Crond守护进程脚本

# 定义服务名称和启动命令
SERVICE_NAME="tts-service"
START_COMMAND="/data/edge-tts/tts-service"

# 检查进程是否存在的函数
check_process() {
    # 使用pgrep检查进程是否存在
    if pgrep -x "$SERVICE_NAME" > /dev/null
    then
        return 0 # 进程存在
    else
        return 1 # 进程不存在
    fi
}

# 启动服务的函数
start_service() {
    echo "Starting $SERVICE_NAME..."
    # 启动服务
    nohup $START_COMMAND > /dev/null 2>&1 &
    echo "$SERVICE_NAME started."
}

# 停止服务的函数
stop_service() {
    echo "Stopping $SERVICE_NAME..."
    # 使用pkill停止服务进程
    pkill -x "$SERVICE_NAME"
    echo "$SERVICE_NAME stopped."
}

# 重启服务的函数
restart_service() {
    stop_service
    start_service
}

# 根据传入的命令执行相应操作
case "$1" in
    start)
        if check_process
        then
            echo "$SERVICE_NAME is already running."
        else
            start_service
        fi
        ;;
    stop)
        if check_process
        then
            stop_service
        else
            echo "$SERVICE_NAME is not running."
        fi
        ;;
    restart)
        restart_service
        ;;
    *)
        # 默认情况下检查进程是否存在，不存在则启动服务
        if ! check_process
        then
            start_service
        else
            echo "check $SERVICE_NAME is running."
        fi
        ;;
esac

exit 0
