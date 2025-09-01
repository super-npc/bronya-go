# Bronya-Go 日志系统使用指南

## 特性

- ✅ 基于 Zap 的高性能日志系统
- ✅ 开发/生产环境自动区分
- ✅ 结构化日志输出
- ✅ 日志文件自动切割
- ✅ 环境变量配置支持
- ✅ 多种日志级别控制
- ✅ 上下文追踪支持

## 快速开始

### 1. 环境配置

#### 开发环境启动
```bash
# 方式1: 使用脚本
./scripts/start.sh develop

# 方式2: 手动设置
export APP_MODE=develop
go run main.go
```

#### 生产环境启动
```bash
# 方式1: 使用脚本
./scripts/start.sh production

# 方式2: 手动设置
export APP_MODE=production
go run main.go
```

### 2. 日志配置

#### 配置文件位置
- 开发环境: `resources/config.yaml`
- 生产环境: `resources/config.production.yaml`

#### 日志配置项
```yaml
log:
  level: debug        # 日志级别: debug, info, warn, error, fatal
  filename: ./logs/app.log  # 日志文件路径
  max_size: 100      # 单个日志文件最大大小(MB)
  max_age: 30        # 日志文件最大保存天数
  max_backups: 7     # 保留的旧日志文件最大数量
```

#### 环境变量覆盖
```bash
export APP_LOG_LEVEL=debug
export APP_LOG_FILENAME=/var/log/myapp.log
export APP_MYSQL_HOST=localhost
```

## 使用方法

### 基本日志

```go
import "github.com/super-npc/bronya-go/src/framework/log"

// 不同级别的日志
logger.Debug("调试信息", zap.String("key", "value"))
logger.Info("一般信息", zap.Int("count", 123))
logger.Warn("警告信息", zap.Error(err))
logger.Error("错误信息", zap.String("error", "something wrong"))
logger.Fatal("致命错误", zap.String("reason", "system crash"))
```

### 性能监控

```go
start := time.Now()
// ... 业务逻辑
logger.LogPerformance("用户查询", time.Since(start))
```

### 业务日志

```go
logger.LogBusiness("订单创建", "success", 
    zap.String("order_id", "ORD123"),
    zap.Float64("amount", 99.99),
)
```

### 数据库日志

```go
logger.LogDatabase("INSERT", "users", time.Since(start), rowsAffected)
```

### 安全日志

```go
logger.LogSecurity("登录失败", "192.168.1.1", "unknown_user")
```

### 上下文日志

```go
// 带请求ID的日志
log := logger.LogContext("req-123", "trace-456")
log.Info("处理请求", zap.String("user", "john"))
```

## 环境差异

### 开发环境
- 日志输出到控制台（带颜色）
- 日志级别: debug
- 包含完整堆栈信息
- 格式: 人类可读

### 生产环境
- 日志输出到文件
- 日志级别: info
- JSON格式输出
- 自动日志切割
- 包含机器信息

## 日志文件结构

```
logs/
├── app.log          # 当前日志文件
├── app.log.1        # 历史日志文件
├── app.log.2        # 更旧的历史日志文件
└── ...
```

## 高级配置

### 自定义日志字段

```go
logger.WithFields(
    zap.String("service", "user-service"),
    zap.String("version", "v1.0.0"),
).Info("服务启动")
```

### 动态日志级别

```go
// 运行时修改日志级别
logger.Logger.Level().SetLevel(zap.DebugLevel)
```

## 常见问题

### Q: 如何查看实时日志？
```bash
# 开发环境
./scripts/start.sh develop

# 生产环境查看日志
tail -f logs/app.log
```

### Q: 如何修改日志级别？
```bash
# 临时修改
export APP_LOG_LEVEL=debug
./scripts/start.sh production

# 永久修改
修改 config.yaml 中的 log.level 配置
```

### Q: 日志文件太大怎么办？
日志系统会自动切割，最大100MB，保留7个文件，30天后自动清理。

## 监控集成

日志格式支持ELK、Grafana等监控系统直接解析，无需额外配置。