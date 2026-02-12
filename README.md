
# Project A: User Service

A Golang microservice for user management, supporting registration, login, and CRUD operations, showcasing basic backend skills and security design.

## Features

- User registration / login  
- User information CRUD  
- JWT authentication helper (standalone, easy to integrate into handlers)  
- Secure password storage using bcrypt hashing  
- Database persistence with PostgreSQL (replaceable with MySQL)

## Tech Stack

- Golang + net/http + Chi: native HTTP + lightweight routing  
- GORM: ORM for database operations  
- UUID: unique user ID  
- JWT: HS256 signing, standalone helper functions for issuing & parsing  
- PostgreSQL: transactional support, ORM mapping  

## Highlights

- **Layered design**: handlers / models / db operations separated  
- **Security**: password hashing + JWT authentication example  
- **RESTful API**: Swagger documentation generated and accessible  
- **Extensibility**: OAuth2 integration possible, JWT usable in middleware  

## API Examples

### Create User
```http
POST /user
Body: { "username": "ellen", "password": "secret" }
```

### Get User Info
```http
GET /user/{id}
Body: { "username": "ellen", "password": "******" }
```

### Update User
```http
PUT /user/{id}
Body: { "username": "ellen_updated", "password": "new_secret" }
```

### Delete User
```http
DELETE /user/{id}
No Body
```





---







# 项目 A：用户服务（User Service）

Golang 微服务实现的用户管理系统，支持注册、登录、CRUD，展示基础后端能力和安全设计思路。

## 功能

- 用户注册 / 登录  
- 用户信息 CRUD  
- JWT 身份认证示例（独立 helper，可轻松集成到 handler）  
- 密码安全存储：使用 bcrypt 哈希  
- 数据库持久化：PostgreSQL（可换 MySQL）

## 技术栈

- Golang + net/http + Chi：原生 HTTP + 轻量级路由  
- GORM：ORM 操作数据库  
- UUID：唯一用户 ID  
- JWT：HS256 签名，独立 helper 函数实现 Issuing & Parsing  
- PostgreSQL：事务支持，ORM 映射  

## 项目亮点

- **分层设计**：handler / model / db 操作分离  
- **安全设计**：密码哈希存储 + JWT 身份验证示例  
- **RESTful API**：Swagger 文档已生成，可直接访问和测试
- **可扩展性**：OAuth2 集成示例，JWT 可直接用于中间件  

## API 示例

### 创建用户
```http
POST /user
Body: { "username": "ellen", "password": "secret" }
```

### 获取用户信息
```http
GET /user/{id}
Body: { "username": "ellen", "password": "******" }
```

### 更新用户
```http
PUT /user/{id}
Body: { "username": "ellen_updated", "password": "new_secret" }
```

### 删除用户
```http
DELETE /user/{id}
No Body
```
