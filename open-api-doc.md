# 接口文档

## 帐户管理

#### 获取当前帐户配置信息

>**GET** /api/v1/account

**返回数据**

| 参数名             | 类型     | 说明      |
|--------|--------|---------|
| serverNotifyUrl | string | 服务端通知地址 |

**返回示例**

```json
 {
  "serverNotifyUrl": "https://xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

#### 设置回调地址

>**PATCH** /api/v1/account

**请求参数**

| 参数名             | 类型     | 说明      |
|--------|--------|---------|
| serverNotifyUrl | string | 服务端通知地址 |

**示例数据**

```json
 {
  "serverNotifyUrl": "https://xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```
------

## 门店管理

### 获取当前关联门店信息
>**GET** /api/v1/stores

**返回数据**

| 参数名               | 类型       | 说明     |
|-------------------|----------|--------|
| stores            | object[] | 门店列表   |
| stores.id         | string   | 门店ID   |
| stores.outStoreId | string   | 外部门店ID |

**返回样例**

```json
{
  "stores": [
     {
       "id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
       "outStoreId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
     }
  ]
}
```

### 配置关联门店信息
>**PATCH** /api/v1/stores

**请求参数**

| 参数名               | 类型       | 说明     |
|-------------------|----------|--------|
| stores            | object[] | 门店列表   |
| stores.id         | string   | 门店ID   |
| stores.outStoreId | string   | 外部门店ID |

**请求示例**

```json
 {
  "stores": [
     {
       "id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
       "outStoreId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
     }
  ]
}
```

------

### 创建订单
>GET /api/v1/stores

| 参数名 | 参数描述 | --- 必填 --- |--- 默认值 ---|
| --- | --- | --- | --- |
| access_key | 访问密钥 | 是 | - |

------

