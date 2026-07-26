# 项目迭代文档

## 定位

`go-grpc-demo` 是 `go-order-management-system` 的旁路实验项目，目标是用 **gRPC + protobuf** 暴露简单的订单能力。  
当前阶段以**只读订单查询**为主，写路径（创建/支付/取消）仍由主系统负责。

## 1. 订单模块

### 已完成

- [x] 对齐主系统 `orders` / `order_items` 数据模型
- [x] 完成简单的订单查询（DAO → Service）
- [x] HTTP 查询：`GET /api/v1/orders/:id`
- [x] gRPC 查询：`order.OrderService/GetOrder`
- [x] 双入口启动：gRPC `:50051` + HTTP `:8083`

### 进行中 / 待办

- [ ] 补全 gRPC `GetOrder` 响应字段（`user_id`、状态文案、`order_no` 等）
- [ ] gRPC 错误语义（`NotFound` 等），避免仅靠业务 `status: FAIL`
- [ ] 完成订单创建模块（建议仍委托主系统，本服务不做库存/幂等）
- [ ] 通过 gRPC 调用订单创建接口（或明确标记为不实现）
- [ ] 开启 gRPC Server Reflection，方便 `grpcurl` 无 proto 调用
- [ ] 列表查询 `ListOrders`（分页 + `user_id`）

## 2. 架构约定

```text
Client
  ├─ gRPC :50051  → grpc_server → service → dao → MySQL(orders)
  └─ HTTP :8083   → handler     → service → dao → MySQL(orders)
```

- 不 import 主系统 `internal` 包
- 查询共享库 `go_order_management_system`，业务表以只读为主
- 不在本项目对主库执行破坏性 migration / AutoMigrate

## 3. 里程碑记录

| 阶段 | 内容 | 状态 |
|------|------|------|
| M0 | 项目脚手架、配置、MySQL 连接 | 完成 |
| M1 | proto `OrderService`、gRPC 服务注册 | 完成 |
| M2 | 订单模型对齐 + GetOrder 查询链路 | 完成 |
| M3 | 查询响应完善 + 测试/文档固化 | 进行中 |
| M4 | 创建链路策略（委托主系统 or 本地 stub） | 未开始 |
