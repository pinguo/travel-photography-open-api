# 服务端通知文档

旅拍开放平台 OPEN API 服务端通知文档

消息通知为json格式，包含以下字段

| 通知类型名   | 数据类型   | 说明                  |
|---------|--------|---------------------|
| type   | string | 当前通知类型              |
| data   | string | JSON字符串，当前通知类型的具体数据 |


**所有可用通知类型**

| 通知类型名   | 说明       |
|---------|----------|
| order   | 订单状态变更通知 |

**通知示例数据**

```json
{
"type":"order",
"data": "{\"id\":\"\",\"tradeNo\":\"\",\"outOrderId\":\"\",\"status\":\"\",\"qrCode\":\"http://xxxx/a/b\",\"storeId\":\"\"}"
}
```

-------

## 订单状态变更通知

通知数据结构

| 参数名        | 类型     | 说明     |
|------------|--------|--------|
| id         | string | 旅拍订单ID |
| tradeNo    | string | 业务订单号  |
| outOrderId | string | 外部订单号  |
| status     | string | 订单状态   |
| qrCode     | string | 订单二维码  |
| storeId    | string | 旅拍门店ID |

**通知数据示例**

```json
{
  "id": "",
  "tradeNo": "",
  "outOrderId": "",
  "status": "",
  "qrCode": "http://xxxx/a/b",
  "storeId": ""
}
```