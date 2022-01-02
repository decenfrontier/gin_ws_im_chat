# gin_ws_im_chat


> 这个项目是基于WebSocket + MongoDB + MySQL + Redis。 业务逻辑很简单，只是两人的聊天。

- MySQL 用来存储用户基本信息
- MongoDB 用来存放用户聊天信息
- Redis 用来存储处理过期信息

## 项目结构
```bash
gin-chat-demo/
├── api
├── cache
├── conf
├── model
├── router
├── serializer
└── service
```
- cache : 放置redis配置
- conf : 放置配置文件 
- model : 数据库模型
- router ： 路由模块
- service ：服务模块

# 项目功能

- 两人通信
- 在线、不在线应答
- 查看历史聊天记录

# 项目运行

- 下载依赖

```bash
go mod tidy
```

- 执行

```bash
go run main.go
```