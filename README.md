# User CRUD Service

A RESTful HTTP API written in Go, implementing full user lifecycle management with JWT-based authentication. Built as a portfolio project to demonstrate production-aware API design — not just "making it work," but making deliberate choices.

**Stack:** Go · chi · JWT (HS256) · UUID · Swagger (godoc annotations)

---

## Design Decisions Worth Noting

### 1. HTTP Semantics Done Right

This API treats status codes and headers as part of the contract, not decoration.

- `POST /user` returns **201 Created** with a `Location: /user/{uuid}` header — the client immediately knows where the new resource lives, without making a second request.
- `PUT` and `DELETE` return **204 No Content** — success is communicated, no redundant body.
- **400 vs 404 are explicitly separated**: malformed JSON is a different error from "user not found," and the API says so.
- Invalid UUID format on path params resolves to **404**, not 500 — the server doesn't panic on bad input.

### 2. Domain Modeling: `UserInfo` vs `User`

The project distinguishes between two structs:

- `UserInfo` — the shape of what the client submits (username + password)
- `User` — the full entity, which includes a system-assigned UUID

This reflects a real design principle: request DTOs and domain entities are not the same thing. `UserInfo.makeUser(uuid)` is the explicit handoff between the two layers.

### 3. JWT Implementation

Rather than wiring up an auth middleware blindly, JWT is implemented with genuine understanding of the internals:

- Custom `Claims` struct embedding `jwt.RegisteredClaims`
- HS256 signing with a shared secret
- 24-hour expiry baked into the token at issuance
- `ParseID()` validates signature and extracts the user ID — one clean function, no leaking of JWT internals into the handler layer

The token binds to a user ID, meaning auth and identity travel together in the token payload.

### 4. Self-Documenting API

All handlers carry Swagger-compatible godoc annotations (`@Summary`, `@Param`, `@Success`, `@Failure`, `@Router`), making the API self-describing and ready for tooling like Swaggo.

---

## Endpoints

| Method   | Path         | Description              |
|----------|--------------|--------------------------|
| `POST`   | `/user`      | Create a new user        |
| `GET`    | `/user/{id}` | Read user by UUID        |
| `PUT`    | `/user/{id}` | Update user info         |
| `DELETE` | `/user/{id}` | Delete user              |

### Request / Response Examples

**Create user**
```
POST /user
Content-Type: application/json

{ "username": "alice", "password": "secret" }

→ 201 Created
   Location: /user/550e8400-e29b-41d4-a716-446655440000
```

**Read user**
```
GET /user/550e8400-e29b-41d4-a716-446655440000

→ 200 OK
   { "user_id": "550e8400-...", "username": "alice", "password": "******" }
```

**Update user**
```
PUT /user/550e8400-e29b-41d4-a716-446655440000
Content-Type: application/json

{ "username": "alice2", "password": "newSecret" }

→ 204 No Content
```

**Delete user**
```
DELETE /user/550e8400-e29b-41d4-a716-446655440000

→ 204 No Content
```

---

## Project Structure (Overview)

```
.
├── handler.go      # HTTP handlers: Create, Read, Update, Delete
├── jwt.go          # Token issuance and parsing (IssueToken, ParseID)
├── model.go        # UserInfo, User structs and makeUser()
├── db.go           # DB layer: create, read, update, delete
└── main.go         # Router setup (chi) and server entry point
```

---

## Running Locally

```bash
go mod tidy
go run .
```

API will be available at `http://localhost:8080`.

---

---

# 用户 CRUD 服务（中文说明）

一个用 Go 编写的 RESTful HTTP API，实现了完整的用户生命周期管理，并集成了 JWT 鉴权。这是一个求职用途的作品集项目，重点展示的不是"能跑起来"，而是背后的设计判断力。

**技术栈：** Go · chi · JWT (HS256) · UUID · Swagger（godoc 注解）

---

## 值得关注的设计决策

### 1. HTTP 语义的正确使用

本项目把状态码和响应头当成 API 契约的一部分，而不是摆设。

- `POST /user` 返回 **201 Created**，同时附带 `Location: /user/{uuid}` 响应头——客户端不需要二次请求就能知道新资源在哪。
- `PUT` 和 `DELETE` 成功后返回 **204 No Content**——不多说废话，但清楚表达了成功。
- **400 和 404 明确区分**：JSON 格式错误和"用户不存在"是两种不同的错误，API 如实反映。
- 路径参数里的非法 UUID 会返回 **404**，而不是让服务器崩成 500。

### 2. 领域建模：`UserInfo` 与 `User` 的区分

项目里有两个结构体：

- `UserInfo`——客户端提交的数据形状（用户名 + 密码）
- `User`——完整的实体，包含系统分配的 UUID

这体现了一个真实的设计原则：请求 DTO 和领域实体不是同一个东西。`UserInfo.makeUser(uuid)` 是两层之间明确的转换点。

### 3. JWT 的实现

没有黑盒接入，而是真正理解了 JWT 的内部机制：

- 自定义 `Claims` 结构体，内嵌 `jwt.RegisteredClaims`
- 使用 HS256 算法 + 共享密钥签名
- 签发时写入 24 小时过期时间
- `ParseID()` 负责验签并提取用户 ID——一个干净的函数，JWT 细节不泄露到 handler 层

Token 与用户 ID 绑定，鉴权和身份信息在 payload 里一起走。

### 4. API 自文档化

所有 handler 都带有 Swagger 兼容的 godoc 注解（`@Summary`、`@Param`、`@Success`、`@Failure`、`@Router`），可以直接配合 Swaggo 等工具生成接口文档。

---

## 接口列表

| 方法     | 路径         | 说明         |
|----------|--------------|--------------|
| `POST`   | `/user`      | 创建用户     |
| `GET`    | `/user/{id}` | 按 UUID 查询 |
| `PUT`    | `/user/{id}` | 更新用户信息 |
| `DELETE` | `/user/{id}` | 删除用户     |

---

## 本地运行

```bash
go mod tidy
go run .
```

服务默认运行在 `http://localhost:8080`。
