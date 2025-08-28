# GoFrame 项目

这是一个使用 GoFrame v2 框架创建的 Golang Web 项目。

## 项目结构

```
goframe-project/
├── api/v1/                 # API 接口定义
│   └── hello.go           # Hello API 定义
├── config/                # 配置文件目录
│   └── config.yaml        # 主配置文件
├── internal/              # 内部代码
│   ├── controller/        # 控制器层
│   │   └── hello.go       # Hello 控制器
│   ├── service/           # 服务层
│   │   ├── hello.go       # Hello 服务实现
│   │   └── service.go     # 服务接口定义
│   ├── dao/               # 数据访问层
│   ├── model/             # 数据模型
│   └── packed/            # 资源文件打包
│       └── packed.go      # 打包文件
├── manifest/config/       # 清单配置
│   └── config.yaml        # 配置清单
├── go.mod                 # Go 模块定义
├── go.sum                 # Go 模块依赖
├── main.go                # 主程序入口
└── README.md              # 项目说明
```

## 功能特性

- 基于 GoFrame v2 框架
- RESTful API 设计
- 分层架构（Controller-Service-DAO）
- 配置文件管理
- 日志记录
- 中间件支持

## 快速开始

### 安装依赖

```bash
go mod tidy
```

### 运行项目

```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动。

### API 接口

#### Hello World API

- **URL**: `/api/v1/hello`
- **Method**: `GET`
- **参数**: 
  - `name` (可选): 名称参数
- **示例**:
  ```bash
  curl "http://localhost:8080/api/v1/hello?name=GoFrame"
  ```
- **响应**:
  ```json
  {
    "message": "Hello, GoFrame!"
  }
  ```

## 开发指南

### 添加新的 API

1. 在 `api/v1/` 目录下定义 API 接口结构体
2. 在 `internal/controller/` 目录下实现控制器
3. 在 `internal/service/` 目录下实现业务逻辑
4. 在 `main.go` 中注册路由

### 配置文件

项目使用 YAML 格式的配置文件，位于 `config/config.yaml`。可以配置：
- 服务器端口和地址
- 日志级别
- 数据库连接
- 其他业务相关配置

### 日志

项目使用 GoFrame 内置的日志组件，支持多种日志级别和输出方式。

## 技术栈

- **框架**: GoFrame v2.9.1
- **语言**: Go 1.24.2
- **配置**: YAML
- **架构**: 分层架构模式

## 许可证

MIT License