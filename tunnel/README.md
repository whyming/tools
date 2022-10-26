## 简单的代理通道

### client
提供本地http代理, 可以通过pac或者浏览器插件控制
`tunnel --client --port=8001 --addr=server.com:123`
`port`: 监听本地端口
`addr`: 将流量转发到该地址

### server
提供服务器tcp服务,接收来自`client`的数据
`tunnel --server --port=6001 -addr=127.0.0.1:80`
`port`: 监听服务器的端口,客户端连接到此端口
`addr`: 将客户端的全部流量转发到此地址,比如`127.0.0.1:80`,就是让客户端访问本地的http服务