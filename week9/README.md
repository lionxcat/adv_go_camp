## 第九周 网络编程
### 问题
1. 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用。
2. 实现一个从 socket connection 中解码出 goim 协议的解码器。
### 回答
1. 
- fix length: 固定长度，可采用固定长度缓冲区，每次接收消息至缓冲区满再解析，短于固定长度的消息可以用\0等byte填充但会造成浪费。
适用于特定消息体的频繁交互，如上报心跳、保活、传递整形数字（例如自增ID、授时）或作为消息head等。
因为缓冲区长度固定，每连接只需要一个缓冲区，因此也适用要求内存较小性能较高的场景。
- length field based: 消息体分为head和body，head内包含body长度信息，head可以是fixed length，解析出body length后再获取相应长度消息体。
适用于各种场景，如纯文本传输，二进制的文件和音视频传输。
- delimiter based: 采用特定分隔符来作为消息切分，例如\r\n，^D。
一般用于纯文本传输，适用方便实现简单。
*其实上面三种方法在设计一种传输协议时可能会组合适用*

2. 
依据[官方文档](https://github.com/Terry-Mao/goim/blob/master/docs/proto.md)，goim protocol设计如下图。
![goim protocol](https://github.com/Terry-Mao/goim/blob/master/docs/protocol.png?raw=true)
详见代码，protos.go > func readSock() 是解码器，goimProtoDemo.go是测试程序。
代码没经优化还请见谅。