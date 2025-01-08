---
title: api
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# api

Base URLs:

* <a href="http://localhost:8888">测试环境: http://localhost:8888</a>

# Authentication

# 用户模块

## POST 用户注册

POST /api/v1/user/register

- 用户注册接口

> Body 请求参数

```yaml
username: fulifuli
password: Fulifuli123456
email: 3185743133@qq.com
code: "841471"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» username|body|string| 是 |用户名|
|» password|body|string| 是 |密码（满足二级强度，即包含字母、数字、特殊符号三者之二，且密码长度不少于8位）|
|» email|body|string| 是 |邮箱|
|» code|body|string| 是 |邮箱验证码|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## POST MFA绑定

POST /api/v1/user/mfa/bind

> Body 请求参数

```yaml
code: "271411"
secret: UXVK3NUPKVCS222PCYG7Y3ETPO2S2BXN

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|
|body|body|object| 否 |none|
|» code|body|string| 是 |需要用验证器（Authenticator）扫描给出的二维码，在验证器内获得验证码|
|» secret|body|string| 是 |获取二维码的接口会给出密钥|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## POST 用户登录

POST /api/v1/user/login

- 用户登录接口

> Body 请求参数

```yaml
username: fulifuli
password: Fulifuli123456
mfa_code: ""

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» username|body|string| 是 |none|
|» password|body|string| 是 |none|
|» mfa_code|body|string| 否 |MFA验证码，需要用户下载验证器（Authenticator）|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": "111431859426033664",
    "username": "fulifuli",
    "avatar_url": "",
    "created_at": 1730634625,
    "updated_at": 1730634625,
    "deleted_at": 0,
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW5fZmllbGQiOnsiUGF5bG9hZCI6IjExMTQzMTg1OTQyNjAzMzY2NCJ9LCJleHAiOjE3MzA2NDkzMzgsIm9yaWdfaWF0IjoxNzMwNjM0OTM4fQ.RrF2VyPM9oluH6DhlEj2dBMKqS9DRswGODbiyJP2euc",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA4OTQxMzgsIm9yaWdfaWF0IjoxNzMwNjM0OTM4LCJyZWZyZXNoX3Rva2VuX2ZpZWxkIjp7IlBheWxvYWQiOiIxMTE0MzE4NTk0MjYwMzM2NjQifX0.Sq46kLRPdpSCr63IPL2fRHAE5YJ3q-P3-bfxPYd3jDA",
    "role": ""
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» id|string|true|none||none|
|»» username|string|true|none||none|
|»» avatar_url|string|true|none||none|
|»» created_at|integer|true|none||none|
|»» updated_at|integer|true|none||none|
|»» deleted_at|integer|true|none||none|
|»» access_token|string|true|none||none|
|»» refresh_token|string|true|none||none|
|»» role|string|true|none||none|

## GET 用户信息

GET /api/v1/user/info

- 获取用户信息

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |用户ID|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": "111431859426033664",
    "username": "fulifuli",
    "avatar_url": "http://cdn.sophisms.cn/?e=1730638368&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:r_TMTSJZZEGK8UwC0ib5PHAKOJE=",
    "created_at": 1730634625,
    "updated_at": 1730634625,
    "deleted_at": 0
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» id|string|true|none||none|
|»» username|string|true|none||none|
|»» avatar_url|string|true|none||none|
|»» created_at|integer|true|none||none|
|»» updated_at|integer|true|none||none|
|»» deleted_at|integer|true|none||none|

## GET 用户粉丝数

GET /api/v1/user/follower_count

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |用户ID|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": "111431859426033664",
    "follower_count": 1
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» id|string|true|none||none|
|»» follower_count|integer|true|none||none|

## GET 用户关注数

GET /api/v1/user/following_count

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |用户ID|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": "111431859426033664",
    "following_count": 1
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» id|string|true|none||none|
|»» following_count|integer|true|none||none|

## GET 上传头像

GET /api/v1/user/avatar/upload

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|
|body|body|object| 否 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "upload_url": "http://up-z2.qiniup.com/",
    "uptoken": "5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:4TEu640EUly0arR81azWfDDBZ4o=:eyJjYWxsYmFja0JvZHkiOiJ7XG5cdFx0XCJrZXlcIjogXCIkKGtleSlcIixcblx0XHRcImhhc2hcIjogXCIkKGV0YWcpXCIsXG5cdFx0XCJmc2l6ZVwiOiAkKGZzaXplKSxcblx0XHRcImJ1Y2tldFwiOiBcIiQoYnVja2V0KVwiLFxuXHRcdFwibmFtZVwiOiBcIiQoeDpuYW1lKVwiLFxuXHRcdFwib3R5cGVcIjogXCJhdmF0YXJcIixcblx0XHRcIm9pZFwiOiBcIjExMTQzMTg1OTQyNjAzMzY2NFwiXG5cdH0iLCJjYWxsYmFja1VybCI6Imh0dHBzOi8vdGVzdC5zb3BoaXNtcy5jbi9hcGkvdjEvb3NzL2NhbGxiYWNrL2F2YXRhciIsImRlYWRsaW5lIjoxNzMwNjM4NTYwLCJwZXJzaXN0ZW50Tm90aWZ5VXJsIjoiaHR0cHM6Ly90ZXN0LnNvcGhpc21zLmNuL2FwaS92MS9vc3MvY2FsbGJhY2svZm9wIiwicGVyc2lzdGVudE9wcyI6ImltYWdlTW9ncjIvdGh1bWJuYWlsLzI1NngyNTYvZm9ybWF0L3dlYnAvYmx1ci8xeDAvcXVhbGl0eS83NXxzYXZlYXMvWW1sc2FYUnZhenBoZG1GMFlYSXZNVEV4TkRNeE9EVTVOREkyTURNek5qWTBMbmRsWW5BPSIsInBlcnNpc3RlbnRUeXBlIjowLCJzY29wZSI6ImJpbGl0b2s6YXZhdGFyLzExMTQzMTg1OTQyNjAzMzY2NC53ZWJwIn0=",
    "upload_key": "avatar/111431859426033664.webp"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» uptoken|string|true|none||none|
|»» upload_url|string|true|none||none|
|»» upload_key|string|true|none||none|

## GET 获取MFA二维码

GET /api/v1/user/mfa/qrcode

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "secret": "UXVK3NUPKVCS222PCYG7Y3ETPO2S2BXN",
    "qrcode": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAACs0lEQVR4nOyYQY60IBSEy7hgyRG4iVzMtHa8GN6EI7BkQaw/9eyZ7j9zADFpVqPzdSJQvFcFvuM7bjgcSe4FeGAhcwuIPDDobbkRkAG365kr5sBM1mi/GHsCSLfbh69AGLNLdXJ7Wbh1B5CrZx41B61+j0DCBGAOLSAw8Y7AKZiyaC/QQoXtxR9FXQqcR2/yzzKHLbdwCubP2bwUOIcnD/tj1H+Ov7XuWsARkcnrsxtG6uhhYPLP9zSvB4TUSaL1LWyZtEobyR/R3gJwuUIz8tK1CYbucHt5YEZHgEu2G4ff8phHVsBKcfk9m9cDCBU8/F4ettRAjbbQ74PXA4Aa68DDP1W3CJB1qHYay42AUOGSV9fbSCJI13Uow+deXA24zFRjiWVhw4wx1+iSk+NZegIg4zVhKDAf5tLZiH1DN4CWGlrZpcyYA4KjvVj4oYerAcca3auP6egFe5gg/j6A3tg0lRik6xrNmX9UuesBR5eoo0cTzPgTH979ogPAFDLwsCDWzvigqvYsS+kHsJkUE3IeqbT40vV4J0BWHD6pPQRlNUR3nLUZ/QCo0e0l8vAtrAFwydLlf13vakBODPG11Mq8CmI1yk6gG8A++/XRoaGhRh6WzNAREGoE/M7Dr5ZwnNUH8kO0/QNO5mu3AqKuN8pRAqooXQE1yjvILgaa6VXGKfBbP4BF88PvfHLLc9ALuVo+31WuDwCDXXE0NTUlh2Sz4I0Aa81yNQ+MbGB2SZr+FMz1gA0JhlzDRmbKlyd8VLnrAbtNUuJdMeYlN8iom8lZSj+A3cuVKFVTkN3LJXwk9x4A0u1lsqVu4SwYFsx/rzjuAiSfAJVidQyp+vBPtt4ATBbO1mCqTlBlXkpHgAnmxz/g5Wn96tkRYJfJPG+TqDYsJ6ZJzbgP8B3f0dX4FwAA///XIwGFJA6HWAAAAABJRU5ErkJggg=="
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» secret|string|true|none||none|
|»» qrcode|string|true|none||none|

## GET 搜索用户

GET /api/v1/user/search

> Body 请求参数

```yaml
keyword: i
page_size: 10
page_num: 0

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» keyword|body|string| 是 |none|
|» page_size|body|integer| 是 |与上次不同，页大小和页码不可缺少|
|» page_num|body|integer| 是 |与上次不同，页大小和页码不可缺少|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "1",
        "username": "admin",
        "avatar_url": "",
        "created_at": 1,
        "updated_at": 1,
        "deleted_at": 0
      },
      {
        "id": "111431859426033664",
        "username": "fulifuli",
        "avatar_url": "http://cdn.sophisms.cn/avatar/111431859426033664.webp?e=1730638769&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:HHTQqNDJdaXZMUldE01x0soEW80=",
        "created_at": 1730634625,
        "updated_at": 1730635132,
        "deleted_at": 0
      }
    ],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 2
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[[user](#schemauser)]|true|none||none|
|»»» id|string|true|none||none|
|»»» username|string|true|none||none|
|»»» avatar_url|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»» is_end|boolean|true|none||none|
|»» page_size|integer|true|none||none|
|»» page_num|integer|true|none||none|
|»» total|integer|true|none||none|

## GET 用户获赞数（仅视频）

GET /api/v1/user/like_count

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |用户ID|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": "108845644633870336",
    "like_count": 1
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» id|string|true|none||none|
|»» like_count|integer|true|none||none|

## POST 忘记密码（邮箱）

POST /api/v1/user/security/password/retrieve/email

> Body 请求参数

```yaml
email: 3185743133@qq.com

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» email|body|string| 是 |邮箱|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## POST 重置密码（邮箱）

POST /api/v1/user/security/password/reset/email

> Body 请求参数

```yaml
code: ""
email: ""
password: ""

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» code|body|string| 是 |邮箱验证码|
|» email|body|string| 是 |邮箱|
|» password|body|string| 是 |密码|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## POST 请求注册验证码

POST /api/v1/user/security/email/code

> Body 请求参数

```yaml
email: 3185743133@qq.com

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» email|body|string| 是 |邮箱|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## POST 忘记密码（用户名）

POST /api/v1/user/security/password/retrieve/username

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|username|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## POST 重置密码（用户名）

POST /api/v1/user/security/password/reset/username

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|username|query|string| 否 |none|
|password|query|string| 否 |none|
|code|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

# 视频模块

## GET 视频流

GET /api/v1/video/feed

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|category|query|string| 否 |分区|
|offset|query|integer| 是 |必须设置偏移|
|n|query|integer| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "111435433388281856",
        "user": {
          "id": "111431859426033664",
          "username": "fulifuli",
          "avatar_url": "http://cdn.sophisms.cn/avatar/111431859426033664.webp?e=1731562870&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:BfdUnyWfzdpWLOXZx8pkC03lAnI=",
          "created_at": 1730634625,
          "updated_at": 1730635320,
          "deleted_at": 0,
          "is_followed": false
        },
        "video_url": "http://cdn.sophisms.cn/video/111435433388281856/video.mp4?e=1731562870&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:33zaOtyZpQFPdfQYqHKyygMj7ak=",
        "cover_url": "http://cdn.sophisms.cn/video/111435433388281856/cover.jpg?e=1731562870&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:c31rcuwu_xDXltt6jet3RR3fIlA=",
        "title": "视频标题",
        "description": "视频简介",
        "visit_count": 0,
        "like_count": 2,
        "comment_count": 0,
        "category": "影音",
        "labels": [
          "标签测试",
          "测试标签"
        ],
        "status": "passed",
        "created_at": 1730635784,
        "updated_at": 1730635784,
        "deleted_at": 0,
        "is_liked": false
      }
    ],
    "page_num": 0,
    "page_size": 0
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[[video](#schemavideo)]|true|none||none|
|»»» id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» video_url|string|true|none||none|
|»»» cover_url|string|true|none||none|
|»»» title|string|true|none||none|
|»»» description|string|true|none||none|
|»»» visit_count|integer|true|none||none|
|»»» like_count|integer|true|none||none|
|»»» comment_count|integer|true|none||none|
|»»» category|string|true|none||none|
|»»» labels|[string]|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|string|true|none||none|

## POST 投稿

POST /api/v1/video/publish

- data仅用于上传视频，而不是**封面**

> Body 请求参数

```yaml
title: 视频标题
description: 视频简介
category: 影音
labels:
  - 测试标签
  - 标签测试

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» description|body|string| 是 |none|
|» category|body|string| 是 |分组|
|» labels|body|[string]| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": {
    "upload_url": "string",
    "upload_key": "string",
    "uptoken": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» upload_url|string|true|none||none|
|»» upload_key|string|true|none||none|
|»» uptoken|string|true|none||none|

## GET 发布列表

GET /api/v1/video/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |none|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Access-Token|header|string| 否 |非必需|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "111435433388281856",
        "user_id": "111431859426033664",
        "video_url": "http://cdn.sophisms.cn/video/111435433388281856/video.mp4?e=1730714329&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:CuDqMX6LFgfRW4YUso-lu-Rr-hg=",
        "cover_url": "http://cdn.sophisms.cn/video/111435433388281856/cover.jpg?e=1730714329&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:3tZqgoMfghOdj3LlryizVt68oPQ=",
        "title": "视频标题",
        "description": "视频简介",
        "visit_count": 0,
        "like_count": 1,
        "comment_count": 0,
        "category": "影音",
        "labels": [
          "标签测试",
          "测试标签"
        ],
        "status": "passed",
        "created_at": 1730635784,
        "updated_at": 1730635784,
        "deleted_at": 0,
        "is_liked": true
      },
      {
        "id": "111435433388281999",
        "user_id": "111431859426033664",
        "video_url": "http://cdn.sophisms.cn/video/111435433388281856/video.mp4?e=1730714329&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:CuDqMX6LFgfRW4YUso-lu-Rr-hg=",
        "cover_url": "http://cdn.sophisms.cn/video/111435433388281856/cover.jpg?e=1730714329&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:3tZqgoMfghOdj3LlryizVt68oPQ=",
        "title": "desc",
        "description": "title",
        "visit_count": 0,
        "like_count": 0,
        "comment_count": 1,
        "category": "影音",
        "labels": [],
        "status": "passed",
        "created_at": 1730423423,
        "updated_at": 1730423423,
        "deleted_at": 0,
        "is_liked": false
      }
    ],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 2
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[object]|true|none||none|
|»»» id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» video_url|string|true|none||none|
|»»» cover_url|string|true|none||none|
|»»» title|string|true|none||none|
|»»» description|string|true|none||none|
|»»» visit_count|integer|true|none||none|
|»»» like_count|integer|true|none||none|
|»»» comment_count|integer|true|none||none|
|»»» category|string|true|none||none|
|»»» labels|[string]|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|boolean|true|none||none|
|»» is_end|boolean|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» total|integer|true|none||none|

## GET 查看全部投稿

GET /api/v1/video/submit/all

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "111435433388281856",
        "user_id": "111431859426033664",
        "video_url": "http://cdn.sophisms.cn/video/111435433388281856/video.mp4?e=1730714349&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:347uzR3q417bYfzopTxQe7jmR60=",
        "cover_url": "http://cdn.sophisms.cn/video/111435433388281856/cover.jpg?e=1730714349&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:DhSzysWYJMFHnSmb44A9ugIseLo=",
        "title": "视频标题",
        "description": "视频简介",
        "visit_count": 0,
        "like_count": 1,
        "comment_count": 0,
        "category": "影音",
        "labels": [
          "标签测试",
          "测试标签"
        ],
        "status": "passed",
        "created_at": 1730635784,
        "updated_at": 1730635784,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "111435433388281999",
        "user_id": "111431859426033664",
        "video_url": "http://cdn.sophisms.cn/video/111435433388281856/video.mp4?e=1730714349&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:347uzR3q417bYfzopTxQe7jmR60=",
        "cover_url": "http://cdn.sophisms.cn/video/111435433388281856/cover.jpg?e=1730714349&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:DhSzysWYJMFHnSmb44A9ugIseLo=",
        "title": "desc",
        "description": "title",
        "visit_count": 0,
        "like_count": 0,
        "comment_count": 1,
        "category": "影音",
        "labels": [],
        "status": "passed",
        "created_at": 1730423423,
        "updated_at": 1730423423,
        "deleted_at": 0,
        "is_liked": false
      }
    ],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 2
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[[video](#schemavideo)]|true|none||none|
|»»» id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» video_url|string|true|none||none|
|»»» cover_url|string|true|none||none|
|»»» title|string|true|none||none|
|»»» description|string|true|none||none|
|»»» visit_count|integer|true|none||none|
|»»» like_count|integer|true|none||none|
|»»» comment_count|integer|true|none||none|
|»»» category|string|true|none||none|
|»»» labels|[string]|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|string|true|none||none|
|»» is_end|boolean|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» total|integer|true|none||none|

## GET 查看审核中的投稿

GET /api/v1/video/submit/review

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "111435433388281856",
        "user_id": "111431859426033664",
        "video_url": "http://cdn.sophisms.cn/video/111435433388281856/video.mp4?e=1730639744&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:4iJcjT7i_jvcBvg3ZLQBLtQzkCA=",
        "cover_url": "http://cdn.sophisms.cn/video/111435433388281856/cover.jpg?e=1730639744&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7N5vfclMg5LYM2KG9p55TLnGFpE=",
        "title": "视频标题",
        "description": "视频简介",
        "visit_count": 0,
        "like_count": 0,
        "comment_count": 0,
        "category": "影音",
        "labels": [
          "标签测试",
          "测试标签"
        ],
        "status": "review",
        "created_at": 1730635784,
        "updated_at": 1730635784,
        "deleted_at": 0
      }
    ],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 1
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[object]|true|none||none|
|»»» id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» video_url|string|true|none||none|
|»»» cover_url|string|true|none||none|
|»»» title|string|true|none||none|
|»»» description|string|true|none||none|
|»»» visit_count|integer|true|none||none|
|»»» like_count|integer|true|none||none|
|»»» comment_count|integer|true|none||none|
|»»» category|string|true|none||none|
|»»» labels|[string]|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|boolean|true|none||none|
|»» is_end|boolean|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» total|integer|true|none||none|

## GET 查看被锁定的投稿

GET /api/v1/video/submit/locked

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 0
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[[video](#schemavideo)]|true|none||none|
|»»» id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» video_url|string|true|none||none|
|»»» cover_url|string|true|none||none|
|»»» title|string|true|none||none|
|»»» description|string|true|none||none|
|»»» visit_count|integer|true|none||none|
|»»» like_count|integer|true|none||none|
|»»» comment_count|integer|true|none||none|
|»»» category|string|true|none||none|
|»»» labels|[string]|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|string|true|none||none|
|»» is_end|boolean|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» total|integer|true|none||none|

## GET 查看通过的投稿

GET /api/v1/video/submit/passed

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 0
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[[video](#schemavideo)]|true|none||none|
|»»» id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» video_url|string|true|none||none|
|»»» cover_url|string|true|none||none|
|»»» title|string|true|none||none|
|»»» description|string|true|none||none|
|»»» visit_count|integer|true|none||none|
|»»» like_count|integer|true|none||none|
|»»» comment_count|integer|true|none||none|
|»»» category|string|true|none||none|
|»»» labels|[string]|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|string|true|none||none|
|»» is_end|boolean|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» total|integer|true|none||none|

## GET 视频信息（视频播放页调用）

GET /api/v1/video/info

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|video_id|query|string| 是 |none|
|Access-Token|header|string| 否 |非必需|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": "109262262673350656",
    "user": {
      "id": "108845644633870336",
      "username": "cyk",
      "avatar_url": "/avatar/108845644633870336?e=1730951388&token=:ljeyi0JdV8SWXOXE1m6ybvYrDWc=",
      "created_at": 1730018023,
      "updated_at": 1730810406,
      "deleted_at": 0,
      "is_followed": true
    },
    "video_url": "/video/109262262673350656/video.mp4?e=1730951388&token=:Ylup5kpXdWXE-65aWCqUtRhdTzk=",
    "cover_url": "/video/109262262673350656/cover.jpg?e=1730951388&token=:2tr-vArBI7I20_tYcgZyCLlakZs=",
    "title": "title",
    "description": "desc",
    "visit_count": 1,
    "like_count": 1,
    "comment_count": 0,
    "category": "影音",
    "labels": [
      "测试1",
      "测试2"
    ],
    "status": "passed",
    "created_at": 1730117483,
    "updated_at": 1730637996,
    "deleted_at": 0,
    "is_liked": false
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» id|string|true|none||none|
|»» user|object|true|none||none|
|»»» id|string|true|none||none|
|»»» username|string|true|none||none|
|»»» avatar_url|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_followed|boolean|true|none||none|
|»» video_url|string|true|none||none|
|»» cover_url|string|true|none||none|
|»» title|string|true|none||none|
|»» description|string|true|none||none|
|»» visit_count|integer|true|none||none|
|»» like_count|integer|true|none||none|
|»» comment_count|integer|true|none||none|
|»» category|string|true|none||none|
|»» labels|[string]|true|none||none|
|»» status|string|true|none||none|
|»» created_at|integer|true|none||none|
|»» updated_at|integer|true|none||none|
|»» deleted_at|integer|true|none||none|
|»» is_liked|boolean|true|none||none|

## GET 热门排行榜

GET /api/v1/video/popular

- 如果用户没有登录，首页视频流就调用这个接口

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_size|query|integer| 是 |none|
|page_num|query|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": {
    "items": {
      "id": "string",
      "user_id": "string",
      "video_url": "string",
      "cover_url": "string",
      "title": "string",
      "description": "string",
      "visit_count": 0,
      "like_count": 0,
      "comment_count": 0,
      "category": "string",
      "labels": [
        "string"
      ],
      "status": "string",
      "created_at": 0,
      "updated_at": 0,
      "deleted_at": 0,
      "is_liked": "string"
    }
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[video](#schemavideo)|true|none||none|
|»»» id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» video_url|string|true|none||none|
|»»» cover_url|string|true|none||none|
|»»» title|string|true|none||none|
|»»» description|string|true|none||none|
|»»» visit_count|integer|true|none||none|
|»»» like_count|integer|true|none||none|
|»»» comment_count|integer|true|none||none|
|»»» category|string|true|none||none|
|»»» labels|[string]|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|string|true|none||none|

## GET 搜索视频

GET /api/v1/video/search

> Body 请求参数

```yaml
keyword: ""
page_size: 0
page_num: 0
from_date: 0
to_date: 0

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 否 |非必需|
|body|body|object| 否 |none|
|» keyword|body|string| 是 |none|
|» page_size|body|integer| 是 |none|
|» page_num|body|integer| 是 |none|
|» from_date|body|integer| 否 |none|
|» to_date|body|integer| 否 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "1",
        "user_id": "114514",
        "video_url": "/video/109262262673350656/video.mp4?e=1730813881&token=:l4sngxF-2IzAC9UBHsuGYwqIXKg=",
        "cover_url": "/video/109262262673350656/cover.jpg?e=1730813881&token=:Jl6gRdcWiDdfg-8lWhBC3wiSWZs=",
        "title": "标题",
        "description": "简介",
        "visit_count": 0,
        "like_count": 0,
        "comment_count": 0,
        "category": "影音",
        "labels": [],
        "status": "passed",
        "created_at": 1,
        "updated_at": 1730621505,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "111435433388281856",
        "user_id": "111431859426033664",
        "video_url": "/video/111435433388281856/video.mp4?e=1730813881&token=:GEjhVllcQTQLq2yZqQ8QWAja0jY=",
        "cover_url": "/video/111435433388281856/cover.jpg?e=1730813881&token=:gN9CrX3pqHN2axrGjf5anppSSuw=",
        "title": "视频标题",
        "description": "视频简介",
        "visit_count": 0,
        "like_count": 1,
        "comment_count": 0,
        "category": "影音",
        "labels": [
          "标签测试",
          "测试标签"
        ],
        "status": "passed",
        "created_at": 1730635784,
        "updated_at": 1730635784,
        "deleted_at": 0,
        "is_liked": false
      }
    ],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 2
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[[video](#schemavideo)]|true|none||none|
|»»» id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» video_url|string|true|none||none|
|»»» cover_url|string|true|none||none|
|»»» title|string|true|none||none|
|»»» description|string|true|none||none|
|»»» visit_count|integer|true|none||none|
|»»» like_count|integer|true|none||none|
|»»» comment_count|integer|true|none||none|
|»»» category|string|true|none||none|
|»»» labels|[string]|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|string|true|none||none|
|»» is_end|boolean|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» total|integer|true|none||none|

## GET 获取分区列表

GET /api/v1/video/categories

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      "游戏",
      "知识",
      "生活",
      "军事",
      "影音",
      "新闻"
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[string]|true|none||none|

## POST 视频封面上传

POST /api/v1/video/cover/upload

- 必须在视频上传完成之后调用

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|video_id|query|string| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "upload_url": "http://up-z2.qiniup.com/",
    "upload_key": "video/111435433388281856/cover.jpg",
    "uptoken": "5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:j1vip3Yu2o1x_RipMk9HDgy8NQM=:eyJjYWxsYmFja0JvZHkiOiJ7XG5cdFx0XCJrZXlcIjogXCIkKGtleSlcIixcblx0XHRcImhhc2hcIjogXCIkKGV0YWcpXCIsXG5cdFx0XCJmc2l6ZVwiOiAkKGZzaXplKSxcblx0XHRcImJ1Y2tldFwiOiBcIiQoYnVja2V0KVwiLFxuXHRcdFwibmFtZVwiOiBcIiQoeDpuYW1lKVwiLFxuXHRcdFwib3R5cGVcIjogXCJjb3ZlclwiLFxuXHRcdFwib2lkXCI6IFwiMTExNDM1NDMzMzg4MjgxODU2XCJcblx0fSIsImNhbGxiYWNrVXJsIjoiaHR0cHM6Ly90ZXN0LnNvcGhpc21zLmNuL2FwaS92MS9vc3MvY2FsbGJhY2svY292ZXIiLCJkZWFkbGluZSI6MTczMDYzOTU3Niwic2NvcGUiOiJiaWxpdG9rOnZpZGVvLzExMTQzNTQzMzM4ODI4MTg1Ni9jb3Zlci5qcGcifQ=="
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» upload_url|string|true|none||none|
|»» upload_key|string|true|none||none|
|»» uptoken|string|true|none||none|

## GET 个性化视频流（必须有token）

GET /api/v1/video/custom/feed

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|category|query|string| 否 |none|
|offset|query|integer| 是 |none|
|n|query|integer| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "111435433388281856",
        "user": {
          "id": "111431859426033664",
          "username": "fulifuli",
          "avatar_url": "/avatar/111431859426033664.webp?e=1730950426&token=:GxUceF7du6q4w2gkQWcHmk3YnV0=",
          "created_at": 1730634625,
          "updated_at": 1730635320,
          "deleted_at": 0,
          "is_followed": false
        },
        "video_url": "/video/111435433388281856/video.mp4?e=1730950426&token=:wzDo0bwoKXuWE8OLUXKUGRyNpek=",
        "cover_url": "/video/111435433388281856/cover.jpg?e=1730950426&token=:8UObKofu-8PPYZslY78r0ufg-Zo=",
        "title": "视频标题",
        "description": "视频简介",
        "visit_count": 0,
        "like_count": 2,
        "comment_count": 0,
        "category": "影音",
        "labels": [
          "标签测试",
          "测试标签"
        ],
        "status": "passed",
        "created_at": 1730635784,
        "updated_at": 1730635784,
        "deleted_at": 0,
        "is_liked": true
      },
      {
        "id": "109262262673350656",
        "user": {
          "id": "108845644633870336",
          "username": "cyk",
          "avatar_url": "/avatar/108845644633870336?e=1730950426&token=:9ToRtvDdvSVeI3VpDMjeN5K8a3g=",
          "created_at": 1730018023,
          "updated_at": 1730810406,
          "deleted_at": 0,
          "is_followed": false
        },
        "video_url": "/video/109262262673350656/video.mp4?e=1730950426&token=:1wLHw96mMTdPb3J3mx-a-946S40=",
        "cover_url": "/video/109262262673350656/cover.jpg?e=1730950426&token=:8sEbpkzDsn2R3E-uD5XeEm26-VI=",
        "title": "title",
        "description": "desc",
        "visit_count": 1,
        "like_count": 1,
        "comment_count": 0,
        "category": "影音",
        "labels": [
          "测试1",
          "测试2"
        ],
        "status": "passed",
        "created_at": 1730117483,
        "updated_at": 1730637996,
        "deleted_at": 0,
        "is_liked": false
      }
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[object]|true|none||none|
|»»» id|string|true|none||none|
|»»» user|object|true|none||none|
|»»»» id|string|true|none||none|
|»»»» username|string|true|none||none|
|»»»» avatar_url|string|true|none||none|
|»»»» created_at|integer|true|none||none|
|»»»» updated_at|integer|true|none||none|
|»»»» deleted_at|integer|true|none||none|
|»»»» is_followed|boolean|true|none||none|
|»»» video_url|string|true|none||none|
|»»» cover_url|string|true|none||none|
|»»» title|string|true|none||none|
|»»» description|string|true|none||none|
|»»» visit_count|integer|true|none||none|
|»»» like_count|integer|true|none||none|
|»»» comment_count|integer|true|none||none|
|»»» category|string|true|none||none|
|»»» labels|[string]|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|boolean|true|none||none|

## GET 获取相似视频

GET /api/v1/video/neighbour/feed

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|video_id|query|string| 是 |none|
|offset|query|integer| 是 |none|
|n|query|integer| 是 |none|
|Access-Token|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 社交模块

## POST 关注操作

POST /api/v1/relation/follow/action

> Body 请求参数

```yaml
to_user_id: "114514"
action_type: 1

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|
|body|body|object| 否 |none|
|» to_user_id|body|string| 是 |none|
|» action_type|body|integer| 是 |0取关，1关注|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## GET 关注列表

GET /api/v1/relation/follow/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |none|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "108845644633870336",
        "username": "cyk",
        "avatar_url": "http://cdn.sophisms.cn/avatar/108845644633870336?e=1730641759&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:09GYP2ntNToI9JU5klEg8KtgoQ8=",
        "created_at": 1730018023,
        "updated_at": 1730103254,
        "deleted_at": 0
      },
      {
        "id": "111431859426033664",
        "username": "fulifuli",
        "avatar_url": "http://cdn.sophisms.cn/avatar/111431859426033664.webp?e=1730641759&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:oSHwoQMm1FxyF_NhfKL-VL29Zv8=",
        "created_at": 1730634625,
        "updated_at": 1730635320,
        "deleted_at": 0
      }
    ],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 2
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[[user](#schemauser)]|true|none||none|
|»»» id|string|true|none||none|
|»»» username|string|true|none||none|
|»»» avatar_url|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|

## GET 粉丝列表

GET /api/v1/relation/follower/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |none|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "111431859426033664",
        "username": "fulifuli",
        "avatar_url": "http://cdn.sophisms.cn/avatar/111431859426033664.webp?e=1730641780&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:Xe0DSN0loAv6grhdqICbkdEQlws=",
        "created_at": 1730634625,
        "updated_at": 1730635320,
        "deleted_at": 0
      }
    ],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 1
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[[user](#schemauser)]|true|none||none|
|»»» id|string|true|none||none|
|»»» username|string|true|none||none|
|»»» avatar_url|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|

## GET 好友列表

GET /api/v1/relation/friend/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "114514",
        "username": "test",
        "avatar_url": "",
        "created_at": 1,
        "updated_at": 1,
        "deleted_at": 0
      }
    ],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 1
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[[user](#schemauser)]|true|none||none|
|»»» id|string|true|none||none|
|»»» username|string|true|none||none|
|»»» avatar_url|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|

# 动态模块

## POST 发布动态

POST /api/v1/activity/publish

> Body 请求参数

```yaml
content: 即将样杨德
image:
  - "111351437572726784"
ref_video: "109262262673350656"
ref_activity: ""

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|
|body|body|object| 否 |none|
|» content|body|string| 是 |none|
|» image|body|array| 是 |none|
|» ref_video|body|string| 否 |none|
|» ref_activity|body|string| 否 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## GET 动态列表

GET /api/v1/activity/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |none|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Access-Token|header|string| 否 |非必需|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "111447245135552512",
        "user_id": "114514",
        "content": "即将样杨德",
        "image": [
          "http://cdn.sophisms.cn/image/111351437572726784.webp?e=1730714476&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:UsOE8B82g2ADfYfMxZNWlO0TTPo="
        ],
        "ref_video": "109262262673350656",
        "ref_activity": "",
        "like_count": 1,
        "created_at": 1730638293,
        "updated_at": 1730638293,
        "deleted_at": 0,
        "is_liked": true
      }
    ],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 1
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[object]|true|none||none|
|»»» id|string|false|none||none|
|»»» user_id|string|false|none||none|
|»»» content|string|false|none||none|
|»»» image|[string]|false|none||none|
|»»» ref_video|string|false|none||none|
|»»» ref_activity|string|false|none||none|
|»»» created_at|integer|false|none||none|
|»»» updated_at|integer|false|none||none|
|»»» deleted_at|integer|false|none||none|
|»» is_end|boolean|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» total|integer|true|none||none|

## GET 动态信息

GET /api/v1/activity/info

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|activity_id|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": {
    "id": "string",
    "user_id": "string",
    "title": "string",
    "text": "string",
    "image": [
      "string"
    ],
    "ref_video": "string",
    "ref_activity": "string",
    "created_at": "string",
    "updated_at": "string",
    "deleted_at": "string",
    "is_liked": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|[activity](#schemaactivity)|true|none||none|
|»» id|string|true|none||none|
|»» user_id|string|true|none||none|
|»» title|string|true|none||none|
|»» text|string|true|none||none|
|»» image|[string]|false|none||none|
|»» ref_video|string|false|none||none|
|»» ref_activity|string|false|none||none|
|»» created_at|string|true|none||none|
|»» updated_at|string|true|none||none|
|»» deleted_at|string|true|none||none|
|»» is_liked|string|true|none||none|

## GET 动态流

GET /api/v1/activity/feed

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_size|query|integer| 是 |none|
|page_num|query|integer| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "111447245135552512",
        "user_id": "114514",
        "content": "即将样杨德",
        "image": [
          "http://cdn.sophisms.cn/image/111351437572726784.webp?e=1730714442&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:I6AMy14CH0exCDS5BdfUunaeIuw="
        ],
        "ref_video": "109262262673350656",
        "ref_activity": "",
        "like_count": 1,
        "created_at": 1730638293,
        "updated_at": 1730638293,
        "deleted_at": 0,
        "is_liked": true
      }
    ],
    "is_end": false,
    "page_num": 0,
    "page_size": 10,
    "total": 1
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[object]|true|none||none|
|»»» id|string|false|none||none|
|»»» user_id|string|false|none||none|
|»»» content|string|false|none||none|
|»»» image|[string]|false|none||none|
|»»» ref_video|string|false|none||none|
|»»» ref_activity|string|false|none||none|
|»»» like_count|integer|false|none||none|
|»»» created_at|integer|false|none||none|
|»»» updated_at|integer|false|none||none|
|»»» deleted_at|integer|false|none||none|
|»»» is_liked|boolean|false|none||none|
|»» is_end|boolean|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» total|integer|true|none||none|

# 互动模块

## POST 点赞视频

POST /api/v1/interact/like/video/action

> Body 请求参数

```yaml
video_id: "111435433388281856"
action_type: 1

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|
|body|body|object| 否 |none|
|» video_id|body|string| 是 |视频ID|
|» action_type|body|integer| 是 |0取消，1点赞|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## GET 点赞视频列表

GET /api/v1/interact/like/video/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |none|
|page_size|query|integer| 是 |none|
|page_num|query|integer| 是 |none|
|Access-Token|header|string| 否 |非必需|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "109262262673350656",
        "user_id": "108845644633870336",
        "video_url": "http://cdn.sophisms.cn/video/109262262673350656/video.mp4?e=1730642388&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:W8dQwBPrNgJXdpI_KjXRpRpmqz8=",
        "cover_url": "http://cdn.sophisms.cn/video/109262262673350656/cover.jpg?e=1730642388&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:wUDE3hyN5Ckgnu2isiH_OoQyN_M=",
        "title": "title",
        "description": "desc",
        "visit_count": 1,
        "like_count": 1,
        "comment_count": 3,
        "category": "影音",
        "labels": [
          "测试1",
          "测试2"
        ],
        "status": "passed",
        "created_at": 1730117483,
        "updated_at": 1730637996,
        "deleted_at": 0,
        "is_liked": false
      }
    ],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 1
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[[video](#schemavideo)]|true|none||none|
|»»» id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» video_url|string|true|none||none|
|»»» cover_url|string|true|none||none|
|»»» title|string|true|none||none|
|»»» description|string|true|none||none|
|»»» visit_count|integer|true|none||none|
|»»» like_count|integer|true|none||none|
|»»» comment_count|integer|true|none||none|
|»»» category|string|true|none||none|
|»»» labels|[string]|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|string|true|none||none|
|»» is_end|boolean|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» total|integer|true|none||none|

## POST 视频下评论

POST /api/v1/interact/comment/video/publish

> Body 请求参数

```yaml
video_id: "111435433388281999"
root_id: ""
parent_id: ""
content: test_parent

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|
|body|body|object| 否 |none|
|» video_id|body|string| 是 |none|
|» root_id|body|string| 否 |根评论id（不是评论的评论就不填）|
|» parent_id|body|string| 否 |回复评论id（不是评论的评论就不填）|
|» content|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## GET 视频根评论列表

GET /api/v1/interact/comment/video/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|video_id|query|string| 是 |none|
|page_size|query|integer| 是 |none|
|page_num|query|integer| 是 |none|
|Access-Token|header|string| 否 |非必需|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "117981305177649152",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "1",
        "created_at": 1732196135,
        "updated_at": 1732196135,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117615849140924416",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 1,
        "child_count": 0,
        "content": "111",
        "created_at": 1732109003,
        "updated_at": 1732109003,
        "deleted_at": 0,
        "is_liked": true
      },
      {
        "id": "117615844086788096",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "111",
        "created_at": 1732109002,
        "updated_at": 1732109002,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117615835748511744",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "123",
        "created_at": 1732109000,
        "updated_at": 1732109000,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117615829788405760",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "123",
        "created_at": 1732108999,
        "updated_at": 1732108999,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117609710957375488",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "1111",
        "created_at": 1732107540,
        "updated_at": 1732107540,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117609658100756480",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "O.o",
        "created_at": 1732107527,
        "updated_at": 1732107527,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117609627297787904",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "qaq",
        "created_at": 1732107520,
        "updated_at": 1732107520,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117609604006817792",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "呜呜呜",
        "created_at": 1732107514,
        "updated_at": 1732107514,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117609583354064896",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "太感动了",
        "created_at": 1732107509,
        "updated_at": 1732107509,
        "deleted_at": 0,
        "is_liked": false
      }
    ],
    "is_end": false,
    "page_num": 0,
    "page_size": 10,
    "total": 14
  }
}
```

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "117981305177649152",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "1",
        "created_at": 1732196135,
        "updated_at": 1732196135,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117615849140924416",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 1,
        "child_count": 0,
        "content": "111",
        "created_at": 1732109003,
        "updated_at": 1732109003,
        "deleted_at": 0,
        "is_liked": true
      },
      {
        "id": "117615844086788096",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "111",
        "created_at": 1732109002,
        "updated_at": 1732109002,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117615835748511744",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "123",
        "created_at": 1732109000,
        "updated_at": 1732109000,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117615829788405760",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "123",
        "created_at": 1732108999,
        "updated_at": 1732108999,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117609710957375488",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "1111",
        "created_at": 1732107540,
        "updated_at": 1732107540,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117609658100756480",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "O.o",
        "created_at": 1732107527,
        "updated_at": 1732107527,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117609627297787904",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "qaq",
        "created_at": 1732107520,
        "updated_at": 1732107520,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117609604006817792",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "呜呜呜",
        "created_at": 1732107514,
        "updated_at": 1732107514,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "117609583354064896",
        "user": {
          "id": "1",
          "username": "admin",
          "avatar_url": "http://cdn.sophisms.cn/avatar/1.webp?e=1732201425&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7VveMDlmu4QbaZd0NHI9uMAQBs8=",
          "created_at": 0,
          "updated_at": 1732108099,
          "deleted_at": 0,
          "is_followed": false
        },
        "otype": "video",
        "oid": "116896812492656640",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "太感动了",
        "created_at": 1732107509,
        "updated_at": 1732107509,
        "deleted_at": 0,
        "is_liked": false
      }
    ],
    "is_end": false,
    "page_num": 0,
    "page_size": 10,
    "total": 14
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[object]|true|none||none|
|»»» id|string|true|none||none|
|»»» user|object|true|none||none|
|»»»» id|string|true|none||none|
|»»»» username|string|true|none||none|
|»»»» avatar_url|string|true|none||none|
|»»»» created_at|integer|true|none||none|
|»»»» updated_at|integer|true|none||none|
|»»»» deleted_at|integer|true|none||none|
|»»»» is_followed|boolean|true|none||none|
|»»» otype|string|true|none||none|
|»»» oid|string|true|none||none|
|»»» root_id|string|true|none||none|
|»»» parent_id|string|true|none||none|
|»»» like_count|integer|true|none||none|
|»»» child_count|integer|true|none||none|
|»»» content|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|boolean|true|none||none|
|»» is_end|boolean|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» total|integer|true|none||none|

## POST 发出私信

POST /api/v1/interact/message/send

> Body 请求参数

```yaml
to_user_id: ""
content: ""

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|
|body|body|object| 否 |none|
|» to_user_id|body|string| 是 |none|
|» content|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## POST 点赞动态

POST /api/v1/interact/like/activity/action

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|activity_id|query|string| 是 |none|
|action_type|query|integer| 是 |0取消，1点赞|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## POST 点赞评论

POST /api/v1/interact/like/comment/action

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|comment_type|query|string| 是 |video, activity|
|from_media_id|query|string| 是 |视频ID或动态ID|
|comment_id|query|string| 是 |none|
|action_type|query|integer| 是 |0取消，1点赞|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## POST 动态下评论

POST /api/v1/interact/comment/activity/publish

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|activity_id|query|string| 是 |none|
|root_id|query|string| 否 |none|
|parent_id|query|string| 否 |none|
|content|query|string| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## GET 动态根评论列表

GET /api/v1/interact/comment/activity/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|activity_id|query|string| 是 |none|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Access-Token|header|string| 否 |非必需|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "111451650316574720",
        "user_id": "111431859426033664",
        "otype": "activity",
        "oid": "111021519437574144",
        "root_id": "0",
        "parent_id": "0",
        "like_count": 0,
        "child_count": 0,
        "content": "test_activity_parent",
        "created_at": 1730639344,
        "updated_at": 1730639344,
        "deleted_at": 0,
        "is_liked": false
      }
    ],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 1
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[[comment](#schemacomment)]|true|none||none|
|»»» id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» otype|string|true|none||none|
|»»» oid|string|true|none||none|
|»»» root_id|integer|true|none||none|
|»»» parent_id|string|true|none||none|
|»»» like_count|string|true|none||none|
|»»» child_count|integer|true|none||none|
|»»» content|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|string|true|none||none|
|»» is_end|boolean|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» total|integer|true|none||none|

## GET 视频子评论列表

GET /api/v1/interact/video/child_comment/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|comment_id|query|string| 是 |none|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Access-Token|header|string| 否 |非必需|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "111089319908753408",
        "user_id": "108845644633870336",
        "otype": "video",
        "oid": "109262262673350656",
        "root_id": "111088921890275328",
        "parent_id": "111089271577788416",
        "like_count": 0,
        "child_count": 0,
        "content": "test_parent",
        "created_at": 1730552957,
        "updated_at": 1730552957,
        "deleted_at": 0,
        "is_liked": false
      },
      {
        "id": "111089271577788416",
        "user_id": "108845644633870336",
        "otype": "video",
        "oid": "109262262673350656",
        "root_id": "111088921890275328",
        "parent_id": "111088921890275328",
        "like_count": 0,
        "child_count": 0,
        "content": "test_root",
        "created_at": 1730552946,
        "updated_at": 1730552946,
        "deleted_at": 0,
        "is_liked": false
      }
    ],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 2
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[[comment](#schemacomment)]|true|none||none|
|»»» id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» otype|string|true|none||none|
|»»» oid|string|true|none||none|
|»»» root_id|integer|true|none||none|
|»»» parent_id|string|true|none||none|
|»»» like_count|string|true|none||none|
|»»» child_count|integer|true|none||none|
|»»» content|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|string|true|none||none|
|»» is_end|boolean|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» total|integer|true|none||none|

## GET 动态子评论列表

GET /api/v1/interact/activity/child_comment/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|comment_id|query|string| 是 |none|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Access-Token|header|string| 否 |非必需|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [],
    "is_end": true,
    "page_num": 0,
    "page_size": 10,
    "total": 0
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[[comment](#schemacomment)]|true|none||none|
|»»» id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» otype|string|true|none||none|
|»»» oid|string|true|none||none|
|»»» root_id|integer|true|none||none|
|»»» parent_id|string|true|none||none|
|»»» like_count|string|true|none||none|
|»»» child_count|integer|true|none||none|
|»»» content|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» updated_at|integer|true|none||none|
|»»» deleted_at|integer|true|none||none|
|»»» is_liked|string|true|none||none|
|»» is_end|boolean|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» total|integer|true|none||none|

## POST 减少该类视频的推荐

POST /api/v1/interact/video/dislike

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|video_id|query|string| 是 |none|
|Access-Token|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

# 工具接口

## DELETE 删除视频

DELETE /api/v1/tool/delete/video

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|video_id|query|string| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## GET 获取上传图片URL及引用id（用于动态）

GET /api/v1/tool/upload/image

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|
|body|body|object| 否 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "image_id": "111351437572726784",
    "upload_url": "http://up-z2.qiniup.com/",
    "upload_key": "image/111351437572726784.webp",
    "uptoken": "5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:e6MtUtoCM7r3czltQj71I75y5ZQ=:eyJjYWxsYmFja0JvZHkiOiJ7XG5cdFx0XCJrZXlcIjogXCIkKGtleSlcIixcblx0XHRcImhhc2hcIjogXCIkKGV0YWcpXCIsXG5cdFx0XCJmc2l6ZVwiOiAkKGZzaXplKSxcblx0XHRcImJ1Y2tldFwiOiBcIiQoYnVja2V0KVwiLFxuXHRcdFwibmFtZVwiOiBcIiQoeDpuYW1lKVwiLFxuXHRcdFwib3R5cGVcIjogXCJpbWFnZVwiLFxuXHRcdFwib2lkXCI6IFwiMTExMzUxNDM3NTcyNzI2Nzg0XCIsXG5cdFx0XCJ1c2VyX2lkXCI6IFwiMTE0NTE0XCJcblx0fSIsImNhbGxiYWNrVXJsIjoiaHR0cHM6Ly90ZXN0LnNvcGhpc21zLmNuL2FwaS92MS9vc3MvY2FsbGJhY2svaW1hZ2UiLCJkZWFkbGluZSI6MTczMDYxOTA1MSwicGVyc2lzdGVudE5vdGlmeVVybCI6Imh0dHBzOi8vdGVzdC5zb3BoaXNtcy5jbi9hcGkvdjEvb3NzL2NhbGxiYWNrL2ZvcCIsInBlcnNpc3RlbnRPcHMiOiJpbWFnZU1vZ3IyL2Zvcm1hdC93ZWJwL2JsdXIvMXgwL3F1YWxpdHkvNzV8c2F2ZWFzL1ltbHNhWFJ2YXpwcGJXRm5aUzh4TVRFek5URTBNemMxTnpJM01qWTNPRFF1ZDJWaWNBPT0iLCJwZXJzaXN0ZW50VHlwZSI6MCwic2NvcGUiOiJiaWxpdG9rOmltYWdlLzExMTM1MTQzNzU3MjcyNjc4NC53ZWJwIn0="
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» image_id|string|true|none||none|
|»» upload_url|string|true|none||none|
|»» upload_key|string|true|none||none|
|»» uptoken|string|true|none||none|

## GET 刷新access-token

GET /api/v1/tool/token/refresh

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Refresh-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": "114514",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW5fZmllbGQiOnsiUGF5bG9hZCI6IjExNDUxNCJ9LCJleHAiOjE3MzA2NTM4NjMsIm9yaWdfaWF0IjoxNzMwNjM5NDYzfQ.Y0SI1DfabnH6Riy5rwQwb2J4wD7GUBZL3hMKGRnTQYQ"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» id|string|true|none||none|
|»» access_token|string|true|none||none|

## DELETE 删除视频（管理员）

DELETE /api/v1/admin/tool/delete/video

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|video_id|query|string| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## GET 请求引用图片URL（动态图片）

GET /api/v1/tool/get/image

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|image_id|query|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "url": "http://cdn.sophisms.cn/image/111351437572726784.webp?e=1730643074&token=5daXu_5RGQn5477EP4F3DYidrpB4RO8zB63aFlV7:7RFyCUKgqgrT6o0swQXuxmndHUQ="
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» url|string|true|none||none|

## DELETE 删除动态

DELETE /api/v1/tool/delete/activity

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|activity_id|query|string| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## DELETE 删除评论

DELETE /api/v1/tool/delete/comment

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|comment_type|query|string| 是 |video, activity|
|from_media_id|query|string| 是 |视频ID 动态ID|
|comment_id|query|string| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## DELETE 删除动态（管理员）

DELETE /api/v1/admin/tool/delete/activity

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|activity_id|query|string| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## DELETE 删除评论（管理员）

DELETE /api/v1/admin/tool/delete/comment

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|comment_type|query|string| 是 |none|
|from_media_id|query|string| 是 |none|
|comment_id|query|string| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 续期refresh-token

GET /api/v1/tool/refresh_token/refresh

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Refresh-Token|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": {
    "id": "string",
    "refresh_token": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» id|string|true|none||none|
|»» refresh_token|string|true|none||none|

# 风纪模块

## POST 举报视频

POST /api/v1/report/video

> Body 请求参数

```yaml
video_id: "109262262673350656"
content: content
label: 血腥

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|
|body|body|object| 否 |none|
|» video_id|body|string| 是 |none|
|» content|body|string| 是 |none|
|» label|body|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## POST 处理视频举报（管理员）

POST /api/v1/admin/video/report/handle

> Body 请求参数

```yaml
report_id: "10001"
action_type: 1

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|
|body|body|object| 否 |none|
|» report_id|body|string| 是 |none|
|» action_type|body|integer| 是 |1受理，0拒绝|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## POST 处理待审视频（管理员）

POST /api/v1/admin/video/handle

> Body 请求参数

```yaml
video_id: "1"
action_type: 1

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|
|body|body|object| 否 |none|
|» video_id|body|string| 是 |none|
|» action_type|body|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## GET 查看待审视频列表（管理员）

GET /api/v1/admin/video/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|category|query|string| 否 |分区|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": {
    "items": [
      {
        "id": "string",
        "user_id": "string",
        "video_url": "string",
        "cover_url": "string",
        "title": "string",
        "description": "string",
        "visit_count": 0,
        "like_count": 0,
        "comment_count": 0,
        "category": "string",
        "labels": [
          "string"
        ],
        "status": "string",
        "created_at": 0,
        "updated_at": 0,
        "deleted_at": 0
      }
    ],
    "total": 0,
    "page_num": 0,
    "page_size": 0,
    "is_end": true
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[object]|true|none||none|
|»»» id|string|false|none||none|
|»»» user_id|string|false|none||none|
|»»» video_url|string|false|none||none|
|»»» cover_url|string|false|none||none|
|»»» title|string|false|none||none|
|»»» description|string|false|none||none|
|»»» visit_count|integer|false|none||none|
|»»» like_count|integer|false|none||none|
|»»» comment_count|integer|false|none||none|
|»»» category|string|false|none||none|
|»»» labels|[string]|false|none||none|
|»»» status|string|false|none||none|
|»»» created_at|integer|false|none||none|
|»»» updated_at|integer|false|none||none|
|»»» deleted_at|integer|false|none||none|
|»» total|integer|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» is_end|boolean|true|none||none|

## GET 视频举报列表（管理员）

GET /api/v1/admin/video/report/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_size|query|integer| 是 |none|
|page_num|query|integer| 是 |none|
|keyword|query|string| 否 |none|
|status|query|string| 否 |none|
|user_id|query|string| 否 |none|
|label|query|string| 否 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "10001",
        "video_id": "109262262673350656",
        "user_id": "114514",
        "reason": "content",
        "label": "血腥",
        "status": "unsolved",
        "created_at": 1730617202,
        "resolved_at": 0,
        "admin_id": "0"
      },
      {
        "id": "111362247627923456",
        "video_id": "109262262673350656",
        "user_id": "114514",
        "reason": "content",
        "label": "血腥",
        "status": "unsolved",
        "created_at": 1730618028,
        "resolved_at": 0,
        "admin_id": "0"
      }
    ],
    "total": 2,
    "page_num": 0,
    "page_size": 10,
    "is_end": true
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[object]|true|none||none|
|»»» id|string|true|none||none|
|»»» video_id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» reason|string|true|none||none|
|»»» label|string|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» resolved_at|integer|true|none||none|
|»»» admin_id|string|true|none||none|
|»» total|integer|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» is_end|boolean|true|none||none|

## POST 举报动态

POST /api/v1/report/activity

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|activity_id|query|string| 是 |none|
|content|query|string| 是 |none|
|label|query|string| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## POST 举报评论

POST /api/v1/report/comment

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|comment_type|query|string| 是 |video, activity|
|from_media_id|query|string| 是 |none|
|comment_id|query|string| 是 |none|
|content|query|string| 是 |none|
|label|query|string| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## GET 动态举报列表（管理员）

GET /api/v1/admin/activity/report/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|status|query|string| 否 |none|
|keyword|query|string| 否 |none|
|user_id|query|string| 否 |none|
|label|query|string| 否 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "10000",
        "activity_id": "111021519437574144",
        "user_id": "114514",
        "reason": "content",
        "label": "血腥",
        "status": "solved",
        "created_at": 1730617258,
        "resolved_at": 0,
        "admin_id": "1"
      },
      {
        "id": "10001",
        "activity_id": "111021519437574144",
        "user_id": "114514",
        "reason": "content",
        "label": "血腥",
        "status": "solved",
        "created_at": 1730617272,
        "resolved_at": 0,
        "admin_id": "1"
      },
      {
        "id": "10002",
        "activity_id": "111021519437574144",
        "user_id": "114514",
        "reason": "content",
        "label": "血腥",
        "status": "unsolved",
        "created_at": 1730617281,
        "resolved_at": 0,
        "admin_id": "0"
      },
      {
        "id": "10003",
        "activity_id": "111021519437574144",
        "user_id": "114514",
        "reason": "content",
        "label": "血腥",
        "status": "unsolved",
        "created_at": 1730617282,
        "resolved_at": 0,
        "admin_id": "0"
      }
    ],
    "total": 4,
    "page_num": 0,
    "page_size": 10,
    "is_end": true
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[object]|true|none||none|
|»»» id|string|true|none||none|
|»»» activity_id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» reason|string|true|none||none|
|»»» label|string|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» resolved_at|integer|true|none||none|
|»»» admin_id|string|true|none||none|
|»» total|integer|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» is_end|boolean|true|none||none|

## GET 评论举报列表（管理员）

GET /api/v1/admin/comment/report/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|comment_type|query|string| 是 |video, activity|
|status|query|string| 否 |none|
|keyword|query|string| 否 |none|
|user_id|query|string| 否 |none|
|label|query|string| 否 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "items": [
      {
        "id": "10000",
        "comment_type": "video",
        "comment_id": "111089271577788416",
        "user_id": "114514",
        "reason": "2323",
        "label": "血腥",
        "status": "solved",
        "created_at": 1730617487,
        "resolved_at": 0,
        "admin_id": "1"
      },
      {
        "id": "10001",
        "comment_type": "video",
        "comment_id": "111089271577788416",
        "user_id": "114514",
        "reason": "2323",
        "label": "血腥",
        "status": "unsolved",
        "created_at": 1730617496,
        "resolved_at": 0,
        "admin_id": "0"
      },
      {
        "id": "10002",
        "comment_type": "video",
        "comment_id": "111089271577788416",
        "user_id": "114514",
        "reason": "2323",
        "label": "血腥",
        "status": "unsolved",
        "created_at": 1730617497,
        "resolved_at": 0,
        "admin_id": "0"
      },
      {
        "id": "10003",
        "comment_type": "video",
        "comment_id": "111089271577788416",
        "user_id": "114514",
        "reason": "2323",
        "label": "血腥",
        "status": "unsolved",
        "created_at": 1730617510,
        "resolved_at": 0,
        "admin_id": "0"
      }
    ],
    "total": 4,
    "page_num": 0,
    "page_size": 10,
    "is_end": true
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» items|[object]|true|none||none|
|»»» id|string|true|none||none|
|»»» comment_type|string|true|none||none|
|»»» comment_id|string|true|none||none|
|»»» user_id|string|true|none||none|
|»»» reason|string|true|none||none|
|»»» label|string|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|integer|true|none||none|
|»»» resolved_at|integer|true|none||none|
|»»» admin_id|string|true|none||none|
|»» total|integer|true|none||none|
|»» page_num|integer|true|none||none|
|»» page_size|integer|true|none||none|
|»» is_end|boolean|true|none||none|

## POST 处理动态举报（管理员）

POST /api/v1/admin/activity/report/handle

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|report_id|query|string| 是 |none|
|action_type|query|integer| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "msg": "ok"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

## POST 处理评论举报（管理员）

POST /api/v1/admin/comment/report/handle

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|report_id|query|string| 是 |none|
|comment_type|query|string| 是 |video, activity|
|action_type|query|integer| 是 |none|
|Access-Token|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|

# OSS对接

## POST 对象存储服务器回调（头像）

POST /api/v1/oss/callback/avatar

> Body 请求参数

```yaml
key: ""
bucket: ""
name: ""
fsize: 0
hash: ""
otype: ""
oid: ""

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |none|
|body|body|object| 否 |none|
|» key|body|string| 是 |none|
|» bucket|body|string| 是 |none|
|» name|body|string| 是 |none|
|» fsize|body|integer| 是 |none|
|» hash|body|string| 是 |none|
|» otype|body|string| 是 |none|
|» oid|body|string| 是 |none|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 对象存储服务器回调（文件处理）

POST /api/v1/oss/callback/fop

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 对象存储服务器回调（视频）

POST /api/v1/oss/callback/video

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 对象存储服务器回调（图片）

POST /api/v1/oss/callback/image

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 数据模型

<h2 id="tocS_video">video</h2>

<a id="schemavideo"></a>
<a id="schema_video"></a>
<a id="tocSvideo"></a>
<a id="tocsvideo"></a>

```json
{
  "id": "string",
  "user_id": "string",
  "video_url": "string",
  "cover_url": "string",
  "title": "string",
  "description": "string",
  "visit_count": 0,
  "like_count": 0,
  "comment_count": 0,
  "category": "string",
  "labels": [
    "string"
  ],
  "status": "string",
  "created_at": 0,
  "updated_at": 0,
  "deleted_at": 0,
  "is_liked": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||none|
|user_id|string|true|none||none|
|video_url|string|true|none||none|
|cover_url|string|true|none||none|
|title|string|true|none||none|
|description|string|true|none||none|
|visit_count|integer|true|none||none|
|like_count|integer|true|none||none|
|comment_count|integer|true|none||none|
|category|string|true|none||none|
|labels|[string]|true|none||none|
|status|string|true|none||none|
|created_at|integer|true|none||none|
|updated_at|integer|true|none||none|
|deleted_at|integer|true|none||none|
|is_liked|string|true|none||none|

<h2 id="tocS_user">user</h2>

<a id="schemauser"></a>
<a id="schema_user"></a>
<a id="tocSuser"></a>
<a id="tocsuser"></a>

```json
{
  "id": "string",
  "username": "string",
  "avatar_url": "string",
  "created_at": 0,
  "updated_at": 0,
  "deleted_at": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||none|
|username|string|true|none||none|
|avatar_url|string|true|none||none|
|created_at|integer|true|none||none|
|updated_at|integer|true|none||none|
|deleted_at|integer|true|none||none|

<h2 id="tocS_comment">comment</h2>

<a id="schemacomment"></a>
<a id="schema_comment"></a>
<a id="tocScomment"></a>
<a id="tocscomment"></a>

```json
{
  "id": "string",
  "user_id": "string",
  "otype": "string",
  "oid": "string",
  "root_id": 0,
  "parent_id": "string",
  "like_count": "string",
  "child_count": 0,
  "content": "string",
  "created_at": 0,
  "updated_at": 0,
  "deleted_at": 0,
  "is_liked": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||none|
|user_id|string|true|none||none|
|otype|string|true|none||none|
|oid|string|true|none||none|
|root_id|integer|true|none||none|
|parent_id|string|true|none||none|
|like_count|string|true|none||none|
|child_count|integer|true|none||none|
|content|string|true|none||none|
|created_at|integer|true|none||none|
|updated_at|integer|true|none||none|
|deleted_at|integer|true|none||none|
|is_liked|string|true|none||none|

<h2 id="tocS_activity">activity</h2>

<a id="schemaactivity"></a>
<a id="schema_activity"></a>
<a id="tocSactivity"></a>
<a id="tocsactivity"></a>

```json
{
  "id": "string",
  "user_id": "string",
  "title": "string",
  "text": "string",
  "image": [
    "string"
  ],
  "ref_video": "string",
  "ref_activity": "string",
  "created_at": "string",
  "updated_at": "string",
  "deleted_at": "string",
  "is_liked": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||none|
|user_id|string|true|none||none|
|title|string|true|none||none|
|text|string|true|none||none|
|image|[string]|false|none||none|
|ref_video|string|false|none||none|
|ref_activity|string|false|none||none|
|created_at|string|true|none||none|
|updated_at|string|true|none||none|
|deleted_at|string|true|none||none|
|is_liked|string|true|none||none|

<h2 id="tocS_report">report</h2>

<a id="schemareport"></a>
<a id="schema_report"></a>
<a id="tocSreport"></a>
<a id="tocsreport"></a>

```json
{
  "id": "string",
  "otype": "string",
  "oid": "string",
  "user_id": "string",
  "content": "string",
  "lables": "string",
  "created_at": 0,
  "updated_at": 0,
  "deleted_at": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||none|
|otype|string|true|none||none|
|oid|string|true|none||none|
|user_id|string|true|none||none|
|content|string|true|none||none|
|lables|string|true|none||none|
|created_at|integer|true|none||none|
|updated_at|integer|true|none||none|
|deleted_at|integer|true|none||none|

