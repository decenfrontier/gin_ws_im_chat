# gin_ws_im_chat


> 这个项目是基于WebSocket + MongoDB + MySQL + Redis。 业务逻辑很简单，只是两人的聊天。

- MySQL 用来存储用户基本信息
- MongoDB 用来存放用户聊天信息
- Redis 用来存储处理过期信息

## 1 项目结构
```bash
gin-chat-demo/
├── api
├── conf
├── model
├── ret
├── router
├── serializer
└── service
```
- api : 接口处理函数
- conf : 放置配置文件 
- model : 数据库模型类
- ret : 返回码和返回消息
- router ： 路由模块
- serializer : 序列化与反序列化
- service ：服务模块


# 2 项目功能

- 两人通信
- 在线、不在线应答
- 查看历史聊天记录


# 3 搭建开发环境

## (1) 下载依赖

```bash
go mod tidy
```

## (2) mysql容器
```bash
docker pull mysql
docker run -d --name mysql8 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -v ~/DockDir/mysql8/conf:/etc/mysql/conf.d mysql
docker exec -it mysql8 mysql -uroot -p123456
```

其中映射目录下的配置文件my.cnf内容如下:
```bash
# mysql下的是MySQL客户端的配置
[mysql]


# mysqld下的是MySQL服务端的配置
[mysqld]
user		= mysql
port		= 3306
pid-file	= /var/run/mysqld/mysqld.pid
socket	= /var/run/mysqld/mysqld.sock

# 允许连接的IP地址, 注释掉之后所有IP的机器都可以连接本MySQL服务端
# bind-address		= 127.0.0.1

# 指定单个查询能够使用的缓冲区大小
key_buffer_size		= 16M
```

## (3) redis容器
```bash
docker pull redis
docker run -d --name=redis -p 6379:6379 -v ~/DockerDir/redis:/etc/redis redis
docker exec -it xxx redis-cli
```

其中映射目录下的配置文件redis.conf内容如下:
```bash
bind 127.0.0.1

protected-mode no

port 6379

tcp-backlog 511

timeout 0

tcp-keepalive 300

daemonize no

pidfile /var/run/redis_6379.pid
```

## (4) mongodb容器
```bash
docker pull mongo
docker run -d --name mongo -p 27017:27017 mongo
docker exec -it mongo mongosh
```

- 执行

```bash
go run main.go
```

# 4 功能演示
## (1) 1 -> 2
[![7DgmBF.png](https://s4.ax1x.com/2022/01/19/7DgmBF.png)](https://imgtu.com/i/7DgmBF)

## (2) 2 -> 1
[![7DggHg.png](https://s4.ax1x.com/2022/01/19/7DggHg.png)](https://imgtu.com/i/7DggHg)