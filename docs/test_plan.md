# 项目测试文档

## 前置条件

1. MySQL 可连，库名 `go_order_management_system`
2. 环境变量或配置中提供 `MYSQL_PASSWORD`（`config.yml` 中 password 为空时）
3. `orders` 表中至少有一条可查询订单（可用主系统创建）
4. 在项目根目录启动：

```bash
go run ./cmd
```

确认日志包含：

- `mysql connected: ...`
- `gRPC Server is running on :50051...`
- HTTP 监听 `:8083`（或 `config.yml` 中 `server.port`）

---

## 1. 订单模块测试

### 1.1 HTTP 查询订单

| 用例 | 步骤 | 期望 |
|------|------|------|
| 查询成功 | `GET http://localhost:8083/api/v1/orders/{存在的id}` | 200，body 含订单字段 |
| ID 非法 | `GET .../orders/abc` 或 `.../orders/0` | 400 |
| 不存在 | `GET .../orders/999999999` | 404 |

```bash
curl -s http://localhost:8083/api/v1/orders/1
curl -s http://localhost:8083/ping
```

### 1.2 gRPC 查询订单

| 用例 | 步骤 | 期望 |
|------|------|------|
| 查询成功 | `GetOrder` + 存在的 `order_id` | 返回 `order_id`、金额；`status` 为 `SUCCESS` |
| 查询失败 | `GetOrder` + 不存在的 `order_id` | `status` 为 `FAIL`（当前实现；后续可改为 NotFound） |

```bash
grpcurl -plaintext \
  -import-path ./api/proto \
  -proto order/order.proto \
  -d "{\"order_id\": 1}" \
  localhost:50051 \
  order.OrderService/GetOrder
```

### 1.3 回归检查清单

- [ ] 服务可启动且 MySQL 连接成功
- [ ] HTTP `/ping` 返回 success
- [ ] HTTP 按 id 查询已有订单成功
- [ ] gRPC `GetOrder` 查询已有订单成功
- [ ] 非法 id / 不存在 id 行为符合上表
- [ ] 不通过本服务创建真实业务订单（创建仍为占位）

---

## 2. 已知限制（验收时注意）

- gRPC 响应字段未对齐完整订单模型（如 `user_id`、订单状态枚举文案）
- 未开启 gRPC reflection 时，`grpcurl` 必须带 `-proto`
- `CreateOrder`（HTTP/gRPC）不作为本阶段验收项
