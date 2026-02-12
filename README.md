
# Project A: User Service Prototype

A Golang microservice prototype for user management, supporting registration, login, and CRUD operations. Demonstrates basic backend capabilities, security design, and is designed for easy expansion of user attributes and authentication schemes.

## Features

- User registration / login  
- User information CRUD, easily extensible (e.g., email, birthday, shopping level, etc.)  
- JWT authentication example (independent helper functions, ready to integrate into handlers)  
- Secure password storage using bcrypt hashing  
- Database persistence with PostgreSQL (can be swapped for MySQL)  

> Note: OAuth support is not implemented, but the design allows easy integration. Suitable as a prototype or architecture reference.

## Tech Stack

- **Golang + net/http + Chi**: Native HTTP server with lightweight routing  
- **GORM**: ORM for database operations, supports PostgreSQL / MySQL  
- **UUID**: Unique user IDs  
- **JWT**: HS256 signing, independent helper functions provide issuing & parsing  
- **PostgreSQL**: Transaction support and ORM mapping  

## Technical Highlights (Prototype-Oriented)

- **Layered Design**: Handlers, models, and DB operations are separated for clarity  
- **Security by Design**: Password hashing + JWT authentication example, ready for middleware integration  
- **Prototype Extensibility**: User attributes and authentication schemes can be expanded easily, e.g., email, birthday, shopping level, OAuth  
- **RESTful API**: Swagger documentation available for direct testing  
- **Future-Ready**: Can seamlessly integrate OAuth2 or other authentication mechanisms for multi-auth, multi-attribute user management


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







# 项目 A：用户服务原型（User Service Prototype）

一个用 Golang 构建的用户管理微服务原型，支持注册、登录和 CRUD，展示基础后端能力、安全设计思路，并为后续扩展复杂用户信息和认证方案做好准备。

## 功能

- 用户注册 / 登录  
- 用户信息 CRUD，可轻松扩展字段（邮箱、生日、购物等级等）  
- JWT 身份认证示例（独立 helper 函数，可直接集成到 handler）  
- 密码安全存储：使用 bcrypt 哈希  
- 数据库持久化：PostgreSQL（可换 MySQL）  

> 注意：OAuth 支持未实现，但设计上可直接集成。适合作为原型或架构验证。

## 技术栈

- **Golang + net/http + Chi**：原生 HTTP + 轻量级路由  
- **GORM**：数据库 ORM，支持 PostgreSQL / MySQL  
- **UUID**：唯一用户 ID  
- **JWT**：HS256 签名，独立 helper 函数提供 Issuing & Parsing  
- **PostgreSQL**：事务支持，ORM 映射  

## 技术亮点（原型导向）

- **分层设计**：handler / model / db 操作分离，逻辑清晰  
- **安全设计**：密码哈希存储 + JWT 身份验证示例，可直接接入中间件  
- **原型可扩展**：用户信息结构和认证方案均可灵活扩展，例如邮箱、生日、购物等级、OAuth 等  
- **RESTful API**：Swagger 文档已生成，可直接访问和测试  
- **后续扩展友好**：可轻松集成 OAuth2 或其他认证方案，实现多认证、多属性用户管理


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
