# 博客系统后端
## 📋 功能特性

- ✅ 用户注册与登录（JWT 认证）
- ✅ 文章的 CRUD 操作
- ✅ 评论功能
- ✅ 数据库使用 SQLite
- ✅ RESTful API 设计
- ✅ 标准化项目结构
- ✅ 日志记录
- ✅ 错误处理

## 🏗️ 项目结构
task4/
├── cmd/
│ └── server/ # 主程序入口
├── internal/ # 私有应用代码
│ ├── app/ # 应用初始化
│ ├── config/ # 配置管理
│ ├── handler/ # HTTP 处理函数
│ ├── middleware/ # 中间件
│ ├── model/ # 数据模型
│ ├── repository/ # 数据访问层
│ └── utils/ # 工具函数
├── pkg/ # 可复用的公共库
│ └── logger/ # 日志组件
├── configs/ # 配置文件
├── migrations/ # 数据库迁移文件
├── docs/ # 文档
├── scripts/ # 脚本文件
├── go.mod # Go 模块文件
└── README.md # 项目说明文档


## 🚀 快速开始

### 环境要求

- Go 1.19+
- SQLite 3

### 安装依赖

```bash
go mod tidy
```

### 配置文件
创建 configs/config.yaml：
```
server:
port: 8080

database:
driver: sqlite
source: blog.db

jwt:
secretKey: your_secret_key_here
expire: 24h
```
### 启动项目
```
cd cmd/server
go run main.go
```
服务器将在 http://localhost:8080 启动。
