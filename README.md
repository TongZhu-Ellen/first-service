# Project A: User Service Prototype

A Golang-based prototype for a user management service, implementing basic CRUD operations and security design.  
This prototype demonstrates layered architecture, decoupled database operations and handlers, password security, and a basic JWT issuance/verification package. The database used is **PostgreSQL**, easily replaceable with other relational databases (e.g., MySQL).

---

## Features & Design

### 0. User Object CRUD Prototype

- User object includes:
  - `username`
  - `password`
  - System-assigned unique `UUID`  
  > Can be extended to include email, date of birth, membership level, etc., but not implemented in this prototype.  

- RESTful CRUD operations:
  - **Create**: create a user  
  - **Read**: retrieve user information  
  - **Update**: update user information  
  - **Delete**: remove a user  

- Decoupled database operations and handler layer:
  - Each CRUD operation corresponds to a separate database function: `create`, `read`, `update`, `delete`  
  - Uses PostgreSQL for data persistence, supporting transactions and integrity checks  
  - Handlers only manage HTTP interface logic, enabling future extensions or database replacements  

- Password security:
  - Stored using **bcrypt** hashing  
  - Passwords are never stored in plain text, ensuring basic security  

---

### 1. JWT Issuance & Verification (Integrable Phase)

- Independent JWT helper package:
  - Issue and verify tokens  
  - Supports HS256 signing  
  - Can be integrated into handlers in the future for authentication middleware  

- Prototype stage does not integrate context or auto-injection into CRUD handlers  
  > CRUD and JWT functionalities are decoupled. Callers must manually use the JWT helper; “one-step” integration is not implemented yet.  

---

## Tech Stack

- **Golang**: high-performance backend language  
- **net/http + Chi**: lightweight routing and HTTP handlers  
- **GORM**: ORM for PostgreSQL database operations  
- **UUID**: unique user identifiers  
- **bcrypt**: password hashing  
- **JWT (HS256)**: token issuance and verification package  
- **PostgreSQL**: transaction support, persistence, and data integrity  

---

## Technical Highlights (Prototype-Oriented)

- **Layered Architecture**: fully decoupled handler and database layers, easy to extend  
- **Security Design**: passwords are always hashed, demonstrating basic best practices  
- **RESTful API**: standard CRUD interface for frontend or other service consumption  
- **Database Integration**: PostgreSQL as the persistent layer, showcasing transaction and ORM capabilities  
- **Extensibility**: user object fields can be expanded, JWT can be integrated into middleware; the prototype can smoothly evolve into a production-ready system  

---


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

一个基于 Golang 的用户管理原型服务，实现了基本的 CRUD 操作与安全设计。  
该原型展示了分层设计、数据库操作与 handler 解耦、密码安全处理，以及 JWT 分发/验证的基础能力。数据库使用 **PostgreSQL**，可轻松替换为其他关系型数据库（如 MySQL）。

---

## 功能与设计

### 0. 用户对象 CRUD 原型

- 用户对象包含：
  - `username`
  - `password`
  - 系统分发的唯一 `UUID`  
  > 可横向扩展到邮箱、生日、会员等级等更多字段，但原型阶段未实现。  

- 按照 RESTful 风格实现：
  - **Create**: 创建用户  
  - **Read**: 获取用户信息  
  - **Update**: 更新用户信息  
  - **Delete**: 删除用户  

- 数据库操作与 handler 层解耦：
  - 每个 CRUD 操作对应独立数据库函数：`create`, `read`, `update`, `delete`  
  - 使用 PostgreSQL 持久化数据，支持事务和完整性检查  
  - handler 仅负责 HTTP 接口逻辑，便于后续扩展和替换数据库实现  

- 用户密码安全处理：
  - 使用 **bcrypt** 加密存储密码  
  - 永远不以明文存储密码，保证基本安全  

---

### 1. JWT 分发与验证（可接入阶段）

- 提供独立的 JWT helper 包：
  - 生成与验证 token  
  - 支持 HS256 签名  
  - 可未来接入 handler，实现身份认证中间件  

- 原型阶段未集成上下文（Context）或自动注入到 CRUD handler  
  > 也就是说，CRUD 和 JWT 功能是解耦的，调用者需手动使用 JWT helper，尚未实现“一步到位”的集成  

---

## 技术栈

- **Golang**：高性能后端语言  
- **net/http + Chi**：轻量级路由与 HTTP handler  
- **GORM**：ORM 操作 PostgreSQL 数据库  
- **UUID**：用户唯一标识  
- **bcrypt**：密码安全加密  
- **JWT (HS256)**：独立包提供 token 发放与验证  
- **PostgreSQL**：事务支持、持久化和数据完整性  

---

## 技术亮点（原型导向）

- **分层设计**：handler 层与数据库层完全解耦，便于扩展  
- **安全设计**：密码永远哈希存储，示范基本安全最佳实践  
- **RESTful API**：符合标准的 CRUD 接口，便于前端或其他服务调用  
- **数据库驱动**：PostgreSQL 作为数据持久层，展示事务和 ORM 映射能力  
- **可扩展性**：用户对象可扩展字段，JWT 可集成到中间件，原型基础可平滑升级为生产系统  

---

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
