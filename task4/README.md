# 个人博客系统后端

这是一个使用 Go 语言、Gin 框架和 GORM 库开发的个人博客系统后端。

## 功能特性

- 用户注册和登录（JWT 认证）
- 文章的 CRUD 操作
- 评论功能
- 数据库迁移
- 错误处理和日志记录

## 技术栈

- Go 语言
- Gin Web 框架
- GORM ORM 库
- MySQL 数据库
- JWT 认证
- Logrus 日志库

## 目录结构

```

task4/
├── main.go                 # 程序入口
├── go.mod                  # 依赖管理
├── config/                 # 配置文件
├── models/                 # 数据模型
├── utils/                  # 工具函数
├── controllers/            # 控制器
├── middleware/             # 中间件
└── routes/                 # 路由定义
```
## 安装和运行

### 环境要求

- Go 1.19+
- MySQL 5.7+

### 安装步骤

1. 克隆项目：
   ```
bash
   git clone <repository-url>
   cd task4
   ```
2. 安装依赖：
   ```
bash
   go mod tidy
   ```
3. 配置环境变量：
   复制 `.env.example` 为 `.env` 并修改相应配置：
   ```
bash
   cp .env.example .env
   ```
4. 创建数据库：
   在 MySQL 中创建数据库：
   ```
sql
   CREATE DATABASE blog_system;
   ```
5. 运行程序：
   ```
bash
   go run main.go
   ```
## API 接口

### 用户认证

- `POST /api/register` - 用户注册
- `POST /api/login` - 用户登录

### 文章管理

- `GET /api/posts` - 获取所有文章
- `GET /api/posts/:id` - 获取单篇文章
- `POST /api/posts` - 创建文章（需认证）
- `PUT /api/posts/:id` - 更新文章（需认证）
- `DELETE /api/posts/:id` - 删除文章（需认证）

### 评论管理

- `POST /api/comments` - 创建评论（需认证）
- `GET /api/posts/:postId/comments` - 获取文章的所有评论

## 测试用例

可以使用 Postman 或 curl 命令测试接口：

### 用户注册
```
bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'
```
### 用户登录
```
bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'
```
### 创建文章（需要替换 TOKEN）
```
bash
curl -X POST http://localhost:8080/api/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"title":"我的第一篇文章","content":"这是文章内容"}'
```
## 错误处理

系统会对各种错误情况进行处理并返回相应的 HTTP 状态码和错误信息：
- 400: 请求参数错误
- 401: 未授权访问
- 403: 禁止访问
- 404: 资源未找到
- 500: 服务器内部错误

## 日志记录

系统使用 Logrus 记录运行信息和错误信息，方便调试和维护。