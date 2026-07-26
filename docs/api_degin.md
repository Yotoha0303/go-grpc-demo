# API 设计文档

> 文件名历史拼写为 `api_degin.md`，内容以本文为准。

## 服务入口

| 协议 | 地址 | 说明 |
|------|------|------|
| gRPC | `localhost:50051` | 明文（无 TLS） |
| HTTP | `localhost:8083` | Gin，配置项 `server.port` |

共享数据库：`go_order_management_system`（与主系统一致）。

---

## 1. 订单模块

### 1.1 订单查询（HTTP）

| 项 | 说明 |
|----|------|
| 方法 | `GET` |
| 路径 | `/api/v1/orders/:id` |
| 路径参数 | `id`：订单主键，正整数 |
| 鉴权 | 当前无（演示用） |

**成功响应示例：**

```json
{
  "code": 0,
  "message": "ok",
  "data": {
    "id": 1,
    "user_id": 1,
    "order_no": "...",
    "total_amount_fen": 1000,
    "status": 1,
    "created_at": "...",
    "updated_at": "..."
  }
}
```

**失败：**

- `id` 非法 → 400
- 订单不存在 → 404

### 1.2 订单查询（gRPC）

| 项 | 说明 |
|----|------|
| 服务 | `order.OrderService` |
| 方法 | `GetOrder` |
| 全名 | `order.OrderService/GetOrder` |

**请求 `GetOrderRequest`：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `order_id` | int64 | 订单 ID |

**响应 `GetOrderResponse`（当前实现）：**

| 字段 | 说明 |
|------|------|
| `order_id` | 订单 ID |
| `amount` / `TotalAmountFen` | 总金额（分） |
| `status` | 业务字符串：`SUCCESS` / `FAIL`（非订单状态枚举） |

> 说明：proto 中另有 `user_id`、`product_id` 等字段，服务端尚未全部填充。  
> 查不到订单时当前返回 `status = "FAIL"` 且 RPC error 为 nil，后续应改为 gRPC `NotFound`。

**grpcurl 示例（项目根目录）：**

```bash
grpcurl -plaintext \
  -import-path ./api/proto \
  -proto order/order.proto \
  -d "{\"order_id\": 1}" \
  localhost:50051 \
  order.OrderService/GetOrder
```

### 1.3 订单创建（占位，未完成）

| 协议 | 入口 | 状态 |
|------|------|------|
| HTTP | `POST /api/v1/orders` | Handler 存在，业务未实现 |
| gRPC | `CreateOrder` | 方法存在，写库不完整，不可用于真实下单 |

真实下单请走主系统：`POST /api/v1/orders`（需 JWT + 幂等键）。

### 1.4 健康检查

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/ping` | 进程存活 |

---

## 2. 订单状态约定（对齐主系统）

| 值 | 含义 |
|----|------|
| 1 | pending（待支付） |
| 2 | paid（已支付） |
| 3 | finished（已完成） |
| 4 | cancelled（已取消） |

金额字段使用 **分**：`total_amount_fen`。
