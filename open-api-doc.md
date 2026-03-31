# 接口文档

## 帐户管理

### 获取当前帐户配置信息

>**GET** /v1/account

**返回数据**

| 参数名             | 类型     | 说明      |
|--------|--------|---------|
| serverNotifyUrl | string | 服务端通知地址 |

**返回数据示例**

```json
 {
  "serverNotifyUrl": "https://xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

### 设置回调地址

>**PATCH** /v1/account

**请求参数**

| 参数名             | 类型     | 说明      | 默认值 |
|--------|--------|---------|-----|
| serverNotifyUrl | string | 服务端通知地址 |     |

**请求数据示例**

```json
 {
  "serverNotifyUrl": "https://xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```
------

## 门店管理

### 获取当前关联门店信息
>**GET** /v1/stores

**返回数据**

| 参数名               | 类型       | 说明     |
|-------------------|----------|--------|
| stores            | object[] | 门店列表   |
| stores.id         | string   | 旅拍门店ID |
| stores.name       | string   | 店铺名称   |
| stores.outStoreId | string   | 外部门店ID |

**返回数据示例**

```json
{
  "stores": [
     {
       "id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
       "outStoreId": "3333333",
       "name": "name"
     }
  ]
}
```

### 配置关联门店信息
>**PATCH** /v1/stores

**请求参数**

| 参数名               | 类型       | 是否必填 | 说明     | 默认值 |
|-------------------|----------|------|--------|-----|
| stores            | object[] | 是    | 门店列表   |     |
| stores.id         | string   | 是    | 门店ID   |     |
| stores.outStoreId | string   | 否    | 外部门店ID |     |


**请求数据示例**

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

## 订单管理

**订单状态枚举**

| 状态码           | 描述          |
|---------------|-------------|
| Initializing  | 修图准备中       |
| Retouching    | 修图中         |
| Retouched     | 修成完成，等待用户选片 |
| PhotoSelected | 用户选片完成      |

### 创建订单
>**POST** /v1/orders

**请求参数**

| 参数名                                   | 类型       | 是否必填 | 说明       | 默认值 |
|---------------------------------------|----------|------|----------|-----|
| storeId                               | string   | 是    | 旅拍门店ID   |     |
| outOrderId                            | string   | 是    | 外部系统订单ID |     |
| customers                             | object[] | 是    | 客户信息     |     |
| customers.mobile                      | string   | 是    | 客户手机号    |     |
| customers.age                         | string   | 否    | 年龄       |     |
| consumptionPackage                    | object   | 是    | 套餐信息     |     |
| consumptionPackage.photoNumber        | number   | 是    | 底片数量     |     |
| consumptionPackage.retouchPhotoNumber | number   | 是    | 精修数量     |     |

**请求数据示例**
```json
{
  "storeId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "outOrderId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "customers": [
    {
      "mobile": "15900000000",
      "age": "18"
    }
  ],
  "consumptionPackage": {
    "photoNumber": 1,
    "retouchPhotoNumber": 1
  }
}
```

**返回数据**

| 参数名        | 类型     | 说明     |
|------------|--------|--------|
| orderId    | string    | 旅拍订单ID |
| outOrderId | string    | 外部订单ID |
| storeId    | string    | 旅拍门店ID |
| tradeNo    | string    | 业务订单号  |
| status     | string   | 订单状态   |
| qrCode     | string   | 二维码    |

**返回数据示例**
```json
 {
  "orderId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "outOrderId": "4433455",
  "storeId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "tradeNo": "345324",
  "status": "Initializing",
  "qrCode": "http://xxxxx"
}
```

### 提交订单图片urls
>**PATCH** /v1/orders/{id}/image-urls

**请求参数**

| 参数名         | 类型       | 是否必填 | 说明   | 默认值 |
|-------------|----------|------|------|-----|
| images      | object[] | 是    | 图片列表 |     |
| images.name | string   | 否    | 图片名  |     |
| images.url  | string   | 是    | 下载地址 |     |

**请求数据示例**
```json
{
  "images": [
     {
       "name": "name",
       "url": "http://xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
     }
  ]
}
```

**返回数据**

| 参数名        | 类型     | 说明     |
|------------|--------|--------|
| orderId    | string    | 旅拍订单ID |
| outOrderId | string    | 外部订单ID |
| storeId    | string    | 旅拍门店ID |
| tradeNo    | string    | 业务订单号  |
| status     | string   | 订单状态   |
| qrCode     | string   | 二维码    |

**返回数据示例**
```json
 {
  "orderId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "outOrderId": "4433455",
  "storeId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "tradeNo": "345324",
  "status": "Initializing",
  "qrCode": "http://xxxxx"
}
```


### 获取订单详情
>**GET** /v1/orders/{orderId}

**返回数据**

| 参数名        | 类型     | 说明     |
|------------|--------|--------|
| orderId    | string    | 旅拍订单ID |
| outOrderId | string    | 外部订单ID |
| storeId    | string    | 旅拍门店ID |
| tradeNo    | string    | 业务订单号  |
| status     | string   | 订单状态   |
| qrCode     | string   | 二维码    |

**返回数据示例**
```json
 {
  "orderId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "outOrderId": "4433455",
  "storeId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "tradeNo": "345324",
  "status": "Initializing",
  "qrCode": "http://xxxxx"
}
```

### 获取订单图片和视频
>**GET** /v1/orders/{orderId}/media

**返回数据**

| 参数名                          | 类型       | 说明           |
|------------------------------|----------|--------------|
| orderId                      | string   | 订单ID         |
| status                       | string   | 订单状态         |
| urlExpiresIn                 | number   | URL有效期（秒）    |
| photos                       | object[] | 照片列表         |
| photos.id                    | string   | 照片ID         |
| photos.number                | string   | 照片编号         |
| photos.type                  | string   | 照片类型：retouch-精修照，creative-创意照 |
| photos.origin                | object   | 原图信息         |
| photos.origin.url            | string   | 原图地址         |
| photos.origin.width          | number   | 原图宽度         |
| photos.origin.height         | number   | 原图高度         |
| photos.origin.size           | number   | 原图大小（字节）     |
| photos.origin.mimeType       | string   | 原图格式         |
| photos.origin.thumbnailUrl   | string   | 原图缩略图地址      |
| photos.retouch               | object   | 精修图信息        |
| photos.retouch.url           | string   | 精修图地址        |
| photos.retouch.width         | number   | 精修图宽度        |
| photos.retouch.height        | number   | 精修图高度        |
| photos.retouch.size          | number   | 精修图大小（字节）    |
| photos.retouch.mimeType      | string   | 精修图格式        |
| photos.retouch.thumbnailUrl  | string   | 精修图缩略图地址     |
| photos.metaData              | object   | 元数据          |
| photos.shootLocation         | string   | 拍摄地点         |
| photos.shootTime             | number   | 拍摄时间（时间戳）    |
| photos.photographer          | string   | 摄影师          |
| videos                       | object[] | 视频列表         |
| videos.url                   | string   | 视频地址         |
| videos.number                | string   | 视频编号         |
| videos.description           | string   | 视频描述         |
| videos.thumbnailUrl          | string   | 视频缩略图地址      |
| videos.width                 | number   | 视频宽度         |
| videos.height                | number   | 视频高度         |

**返回数据示例**
```json
{
  "orderId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "status": "Retouched",
  "urlExpiresIn": 3600,
  "photos": [
    {
      "id": "673c392b0d987fa73241b682",
      "number": "1",
      "type": "retouch",
      "origin": {
        "url": "https://cdn-qa-all.c360dn.com/travel-photography/resource/673c3341b14c2d8b7f4ae228/FnwmjdSnZS0lVL0flj9odIXy4Zs3",
        "width": 1080,
        "height": 1445,
        "size": 151273,
        "mimeType": "JPEG",
        "thumbnailUrl": "https://cdn-qa-all.c360dn.com/travel-photography/resource/673c3341b14c2d8b7f4ae228/FnwmjdSnZS0lVL0flj9odIXy4Zs3?imageView2/1/w/200/h/200"
      },
      "retouch": {
        "url": "https://cdn-qa-all.c360dn.com/central-kitchen/173200083886282210297_retouch_673c392b0d987fa73241b682.jpeg",
        "width": 1080,
        "height": 1445,
        "size": 399420,
        "mimeType": "",
        "thumbnailUrl": "https://cdn-qa-all.c360dn.com/central-kitchen/173200083886282210297_retouch_673c392b0d987fa73241b682.jpeg?imageView2/1/w/200/h/200"
      },
      "metaData": {},
      "shootLocation": "二仙桥",
      "shootTime": 1732001102,
      "photographer": "新-不要删"
    }
  ],
  "videos": [
    {
      "url": "https://cdn-qa-all.c360dn.com/central-kitchen/404fa70a2e4d141d2005581b0b4d98b0",
      "number": "V0001",
      "description": "",
      "thumbnailUrl": "https://cdn-qa-all.c360dn.com/central-kitchen/404fa70a2e4d141d2005581b0b4d98b0-1.jpg",
      "width": 400,
      "height": 400
    }
  ]
}
```