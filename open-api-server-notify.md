# 服务端通知文档

## 通知规则说明
- 1、回调通知仅关注响应的http状态码是否等于200，不会使用任何响应体中的数据
- 1.1、响应http状态等于200，则认为回调成功，不会再触发回调
- 1.2、响应http状态不等于200，则认为回调失败，在依次间隔15s、15s、30s、3m、10m、20m、30m、30m、30m、60m、3h、3h、3h、6h、6h进行回调重试，总共回调15次。如果所有重试都不成功，则认为回调失败，不再触发回调。
- 2、回调接口应在5s内响应，否则此次回调则认为失败回调
- 3、对于多个回调通知有顺序要求的，仅在上一个回调成功后才会触发下一个回调，如果上一个回调失败，后续回调将都不会回调。

##  文档
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
| orderPurchase  | 订单购买通知 |

**订单通知示例数据**

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
**订单购买通知示例数据**

```json
{
    "type": "orderPurchase",
    "data": "{"id":"68adfad1890e8445894ae8d0","tradeNo":"31069127","outOrderId":"4692","storeId":"6791fb6efd51813c7293002e","purchaseItems":[{"photoId":"68adfad6d5c50e919e6dc62d","photoName":"AB_00679.JPG","photoUrl":"https://a1.91lvpai.cn/photo/vip/pic/2025/08/26/4692/1/17562221200002125.JPG","photoRetouchUrl":"https://ali-private-travel-photography.c360dn.com/linNytd0gNAvvyt-2oPhpsvsDh6U?Expires=1756878474\u0026OSSAccessKeyId=LTAI5tBQAdSmFuf4vrqDTyh2\u0026Signature=O9ZyCHAUeZ5RYQCe5kdSXdHJHco%3D"},{"photoId":"68adfad7d5c50e919e6dc631","photoName":"AB_00683.JPG","photoUrl":"https://a1.91lvpai.cn/photo/vip/pic/2025/08/26/4692/1/17562221200009934.JPG","photoRetouchUrl":"https://ali-private-travel-photography.c360dn.com/ltVUjF0wlMgS3EpUz7Zazh3utAdZ?Expires=1756878474\u0026OSSAccessKeyId=LTAI5tBQAdSmFuf4vrqDTyh2\u0026Signature=75BJ8Two%2B94Mc%2BU%2FuIyZ%2Fxy4kRY%3D"},{"photoId":"68adfad8d5c50e919e6dc635","photoName":"AB_00677.JPG","photoUrl":"https://a1.91lvpai.cn/photo/vip/pic/2025/08/26/4692/1/17562221200001256.JPG","photoRetouchUrl":"https://ali-private-travel-photography.c360dn.com/lqvnOU-eLCYPYq9BkAuBhauwjS_3?Expires=1756878474\u0026OSSAccessKeyId=LTAI5tBQAdSmFuf4vrqDTyh2\u0026Signature=AzQxL9eGdXy1y56XEVEU0GiWsXs%3D"},{"photoId":"68adfadad5c50e919e6dc63b","photoName":"AB_00675.JPG","photoUrl":"https://a1.91lvpai.cn/photo/vip/pic/2025/08/26/4692/1/17562221200001663.JPG","photoRetouchUrl":"https://ali-private-travel-photography.c360dn.com/luAkWbE4_YKN2MteG1osSLXhIf0f?Expires=1756878474\u0026OSSAccessKeyId=LTAI5tBQAdSmFuf4vrqDTyh2\u0026Signature=QOLoc%2F8l3mb6Qm2rLOXzIMYGPX0%3D"},{"photoId":"68adfaded5c50e919e6dc64f","photoName":"AB_00664.JPG","photoUrl":"https://a1.91lvpai.cn/photo/vip/pic/2025/08/26/4692/1/17562221200009863.JPG","photoRetouchUrl":"https://ali-private-travel-photography.c360dn.com/lkVNsdVbErx3sXXDT7J6cCsZPC9x?Expires=1756878474\u0026OSSAccessKeyId=LTAI5tBQAdSmFuf4vrqDTyh2\u0026Signature=1GV9q9dC5ryrnoTUVxvm4N2e38A%3D"},{"photoId":"68adfadfd5c50e919e6dc653","photoName":"AB_00663.JPG","photoUrl":"https://a1.91lvpai.cn/photo/vip/pic/2025/08/26/4692/1/17562221200005899.JPG","photoRetouchUrl":"https://ali-private-travel-photography.c360dn.com/lqdmzTFplHZuSAwZaUpfdk4gOE2r?Expires=1756878474\u0026OSSAccessKeyId=LTAI5tBQAdSmFuf4vrqDTyh2\u0026Signature=v5SNGLn5nRbgMWb7nanI2Vwnxes%3D"},{"photoId":"68adfae0d5c50e919e6dc657","photoName":"AB_00660.JPG","photoUrl":"https://a1.91lvpai.cn/photo/vip/pic/2025/08/26/4692/1/17562221200007893.JPG","photoRetouchUrl":"https://ali-private-travel-photography.c360dn.com/lhyOvnHiVoPzZcNOLblX0QeCF6aH?Expires=1756878474\u0026OSSAccessKeyId=LTAI5tBQAdSmFuf4vrqDTyh2\u0026Signature=ECiSJ%2Bk7zubKbY8PQYLXfjO6Hts%3D"},{"photoId":"68adfae1d5c50e919e6dc65b","photoName":"AB_00658.JPG","photoUrl":"https://a1.91lvpai.cn/photo/vip/pic/2025/08/26/4692/1/17562221200001548.JPG","photoRetouchUrl":"https://ali-private-travel-photography.c360dn.com/lq9DnWDJmEGLapVqNlsroG6nK5pz?Expires=1756878474\u0026OSSAccessKeyId=LTAI5tBQAdSmFuf4vrqDTyh2\u0026Signature=cFJdJ8r8NHhdWnidcNCdwZkZQn4%3D"},{"photoId":"68adfae2d5c50e919e6dc65f","photoName":"AB_00655.JPG","photoUrl":"https://a1.91lvpai.cn/photo/vip/pic/2025/08/26/4692/1/17562221200001853.JPG","photoRetouchUrl":"https://ali-private-travel-photography.c360dn.com/lq0rVYxMqQ4YoREB5K3EQLE18d5J?Expires=1756878474\u0026OSSAccessKeyId=LTAI5tBQAdSmFuf4vrqDTyh2\u0026Signature=cIy0%2FtKnqJOJaYTVhNXc9Jd%2BMc4%3D"},{"photoId":"68adfae5d5c50e919e6dc66b","photoName":"AB_00650.JPG","photoUrl":"https://a1.91lvpai.cn/photo/vip/pic/2025/08/26/4692/1/17562221200005584.JPG","photoRetouchUrl":"https://ali-private-travel-photography.c360dn.com/lr4iN8_GV6IKAOJzu-72t_KnS9Ut?Expires=1756878474\u0026OSSAccessKeyId=LTAI5tBQAdSmFuf4vrqDTyh2\u0026Signature=ZC7TX54TZqCspXtSzIOUGvfHzfQ%3D"}]}"
}
```

-------

## 订单购买通知

通知数据结构

| 参数名        | 类型     | 说明     |
|------------|--------|--------|
| id         | string | 旅拍订单ID |
| tradeNo    | string | 业务订单号  |
| outOrderId | string | 外部订单号  |
| storeId    | string | 旅拍门店ID |
| purchaseItems | array | 订单已购买内容列表 |
| purchaseItems.photoId | string | 已购图片id |
| purchaseItems.photoName | string | 已购图片名称 |
| purchaseItems.photoUrl | string | 已购图片url |
| purchaseItems.photoRetouchUrl | string | 已购图片精修url |

**通知数据示例**

```json
{
    "id": "68adfad1890e8445894ae8d0",
    "tradeNo": "31069127",
    "outOrderId": "4692",
    "storeId": "6791fb6efd51813c7293002e",
    "purchaseItems": [
        {
            "photoId": "68adfad6d5c50e919e6dc62d",
            "photoName": "AB_00679.JPG",
            "photoUrl": "https://a1.91lvpai.cn/photo/vip/pic/2025/08/26/4692/1/17562221200002125.JPG",
            "photoRetouchUrl": "https://ali-private-travel-photography.c360dn.com/linNytd0gNAvvyt-2oPhpsvsDh6U?Expires=1756878474&OSSAccessKeyId=LTAI5tBQAdSmFuf4vrqDTyh2&Signature=O9ZyCHAUeZ5RYQCe5kdSXdHJHco%3D"
        }
    ]
}
```