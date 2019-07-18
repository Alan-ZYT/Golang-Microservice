# 初识gRPC：gRPC及Protobuf相关介绍

作为开篇章，将会介绍 gRPC 相关的一些知识。简单来讲 gRPC 是一个 基于 HTTP/2 协议设计的 RPC 框架，它采用了 Protobuf 作为 IDL（接口定义语言Interface Define Language）

你是否有过疑惑，它们都是些什么？本文将会介绍一些常用的知识和概念，更详细的会给出手册地址去深入

## 一、RPC

### 什么是 RPC

RPC 代指远程过程调用（Remote Procedure Call），它的调用包含了传输协议和编码（对象序列号）协议等等。允许运行于一台计算机的程序调用另一台计算机的子程序，而开发人员无需额外地为这个交互作用编程

#### 实际场景：

有两台服务器，分别是A、B。在 A 上的应用 C 想要调用 B 服务器上的应用 D，它们可以直接本地调用吗？
答案是不能的，但走 RPC 的话，十分方便。因此常有人称使用 RPC，就跟本地调用一个函数一样简单

### RPC 框架

我认为，一个完整的 RPC 框架，应包含负载均衡、服务注册和发现、服务治理等功能，并具有可拓展性便于流量监控系统等接入
那么它才算完整的，当然了。有些较单一的 RPC 框架，通过组合多组件也能达到这个标准

你认为呢？

### 常见 RPC 框架

- [gRPC](https://grpc.io/)
- [Thrift](https://github.com/apache/thrift)
- [Rpcx](https://github.com/smallnest/rpcx)
- [Dubbo](https://github.com/apache/incubator-dubbo)

### 比较一下

| \      | 跨语言 | 多 IDL | 服务治理 | 注册中心 | 服务管理 |
| ------ | ------ | ------ | -------- | -------- | -------- |
| gRPC   | √      | ×      | ×        | ×        | ×        |
| Thrift | √      | ×      | ×        | ×        | ×        |
| Rpcx   | ×      | √      | √        | √        | √        |
| Dubbo  | ×      | √      | √        | √        | √        |

### 为什么要 RPC

简单、通用、安全、效率

### RPC 可以基于 HTTP 吗

RPC 是代指远程过程调用，是可以基于 HTTP 协议的

肯定会有人说效率优势，我可以告诉你，那是基于 HTTP/1.1 来讲的，HTTP/2 优化了许多问题（当然也存在新的问题），所以你看到了本文的主题 gRPC

## 二、Protobuf

### 介绍

Protocol Buffers 是一种与语言、平台无关，可扩展的序列化结构化数据的方法，常用于通信协议，数据存储等等。相较于 JSON、XML，它更小、更快、更简单，因此也更受开发人员的青眯

### 语法

```
syntax = "proto3";

service SearchService {
    rpc Search (SearchRequest) returns (SearchResponse);
}

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}

message SearchResponse {
    ...
}
```

1、第一行（非空的非注释行）声明使用 `proto3` 语法。如果不声明，将默认使用 `proto2` 语法。同时我建议用 v2 还是 v3，都应当声明其使用的版本

2、定义 `SearchService` RPC 服务，其包含 RPC 方法 `Search`，入参为 `SearchRequest` 消息，出参为 `SearchResponse` 消息

3、定义 `SearchRequest`、`SearchResponse` 消息，前者定义了三个字段，每一个字段包含三个属性：类型、字段名称、字段编号

4、Protobuf 编译器会根据选择的语言不同，生成相应语言的 Service Interface Code 和 Stubs

最后，这里只是简单的语法介绍，详细的请右拐 [Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3)

### 数据类型

| .proto Type | C++ Type | Java Type  | Go Type | PHP Type       |
| ----------- | -------- | ---------- | ------- | -------------- |
| double      | double   | double     | float64 | float          |
| float       | float    | float      | float32 | float          |
| int32       | int32    | int        | int32   | integer        |
| int64       | int64    | long       | int64   | integer/string |
| uint32      | uint32   | int        | uint32  | integer        |
| uint64      | uint64   | long       | uint64  | integer/string |
| sint32      | int32    | int        | int32   | integer        |
| sint64      | int64    | long       | int64   | integer/string |
| fixed32     | uint32   | int        | uint32  | integer        |
| fixed64     | uint64   | long       | uint64  | integer/string |
| sfixed32    | int32    | int        | int32   | integer        |
| sfixed64    | int64    | long       | int64   | integer/string |
| bool        | bool     | boolean    | bool    | boolean        |
| string      | string   | String     | string  | string         |
| bytes       | string   | ByteString | []byte  | string         |

### v2 和 v3 主要区别

- 删除原始值字段的字段存在逻辑
- 删除 required 字段
- 删除 optional 字段，默认就是
- 删除 default 字段
- 删除扩展特性，新增 Any 类型来替代它
- 删除 unknown 字段的支持
- 新增 [JSON Mapping](https://developers.google.com/protocol-buffers/docs/proto3#json)
- 新增 Map 类型的支持
- 修复 enum 的 unknown 类型
- repeated 默认使用 packed 编码
- 引入了新的语言实现（C＃，JavaScript，Ruby，Objective-C）

以上是日常涉及的常见功能，如果还想详细了解可阅读 [Protobuf Version 3.0.0](https://github.com/protocolbuffers/protobuf/releases?after=v3.2.1)

### 相较 Protobuf，为什么不使用XML？

- 更简单
- 数据描述文件只需原来的1/10至1/3
- 解析速度是原来的20倍至100倍
- 减少了二义性
- 生成了更易使用的数据访问类

## 三、gRPC

### 介绍

gRPC 是一个高性能、开源和通用的 RPC 框架，面向移动和 HTTP/2 设计

#### 多语言

- C++
- C#
- Dart
- Go
- Java
- Node.js
- Objective-C
- PHP
- Python
- Ruby

#### 特点

1、HTTP/2

2、Protobuf

3、客户端、服务端基于同一份 IDL

4、移动网络的良好支持

5、支持多语言

### 概览

[![image](https://camo.githubusercontent.com/d36773abdc460e28e4b30f194d7858c8f30bd86f/68747470733a2f2f677270632e696f2f696d672f6c616e64696e672d322e737667)](https://camo.githubusercontent.com/d36773abdc460e28e4b30f194d7858c8f30bd86f/68747470733a2f2f677270632e696f2f696d672f6c616e64696e672d322e737667)

### 讲解

1、Ruby 或者 Java 语言编写的客户端（gRPC Sub）调用 A 方法，发起 RPC 调用

2、对请求信息使用 Protobuf 进行对象序列化压缩（IDL）

3、服务端（gRPC Server）接收到请求后，解码请求体，进行业务逻辑处理并返回

4、对响应结果使用 Protobuf 进行对象序列化压缩（IDL）

5、客户端接受到服务端响应，解码请求体。回调被调用的 A 方法，唤醒正在等待响应（阻塞）的客户端调用并返回响应结果

### 示例

在这一小节，将简单的给大家展示 gRPC 的客户端和服务端的示例代码，希望大家先有一个基础的印象，将会在下一章节详细介绍 

#### 构建和启动服务端

```
lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
if err != nil {
        log.Fatalf("failed to listen: %v", err)
}

grpcServer := grpc.NewServer()
...
pb.RegisterSearchServer(grpcServer, &SearchServer{})
grpcServer.Serve(lis)
```

1、监听指定 TCP 端口，用于接受客户端请求

2、创建 gRPC Server 的实例对象

3、gRPC Server 内部服务和路由的注册

4、Serve() 调用服务器以执行阻塞等待，直到进程被终止或被 Stop() 调用

#### 创建客户端

```
var opts []grpc.DialOption
...
conn, err := grpc.Dial(*serverAddr, opts...)
if err != nil {
    log.Fatalf("fail to dial: %v", err)
}

defer conn.Close()
client := pb.NewSearchClient(conn)
...
```

1、创建 gRPC Channel 与 gRPC Server 进行通信（需服务器地址和端口作为参数）

2、设置 DialOptions 凭证（例如，TLS，GCE凭据，JWT凭证）

3、创建 Search Client Stub

4、调用对应的服务方法

## 思考题

1、什么场景下不适合使用 Protobuf，而适合使用 JSON、XML？

2、Protobuf 一节中提到的 packed 编码，是什么？

## 总结

在开篇内容中，我利用了尽量简短的描述给你介绍了接下来所必须、必要的知识点 希望你能够有所收获，建议能到我给的参考资料处进行深入学习，是最好的了

## 参考资料

- [Protocol Buffers](https://developers.google.com/protocol-buffers/docs/proto3)
- [gRPC](https://grpc.io/docs/)
