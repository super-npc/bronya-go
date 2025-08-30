#!/bin/bash

# 启动脚本
# 使用方法: ./scripts/start.sh [develop|production]

# 设置默认环境
ENV=${1:-develop}

echo "启动应用，环境: $ENV"

# 创建日志目录
mkdir -p logs

# 设置环境变量
export APP_MODE=$ENV

# 根据不同环境设置不同的配置
case $ENV in
    "production")
        export APP_LOG_LEVEL=info
        export APP_PORT=8080
        ;;
    "develop")
        export APP_LOG_LEVEL=debug
        export APP_PORT=8080
        ;;
    *)
        echo "未知环境: $ENV"
        echo "使用方法: ./scripts/start.sh [develop|production]"
        exit 1
        ;;
esac

# 启动应用
echo "启动参数:"
echo "  环境: $ENV"
echo "  日志级别: $APP_LOG_LEVEL"
echo "  端口: $APP_PORT"
echo "  日志文件: ./logs/app.log"

# 运行应用
go run main.go