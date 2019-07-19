# 何为微服务？

在传统的应用中，所有的功能都是存在于单一的代码库(Monotholic Code Base)中。在表面上看，代码库中的代码可以有几种聚合方式。可能会按照其类型分割，比如controllers, entity, factories，也有可能按照其功能拆分成几个包，比如auth, articles等等。但无论如何，整个应用是建立在一个单一代码库上的。

微服务是对于上述第二种聚合方式的拓展。我们依旧将应用按照其功能拆分成几个包，但不同的是，这些功能包现在都是一个可独立运行的代码库。

# 为何用微服务？

降低复杂性 - 将功能拆分成对应的微服务可以将你的整个代码拆分成更小，更易维护的代码库。这有点类似早期Unix的开发哲学“只做一件事，并做到做好“。在传统的单一代码库应用中，代码的耦合性往往更容易往高耦合发展。这就可能导致撰写和维护代码变得很复杂，也更容易出现漏洞。

拓展性 - 在单一代码库应用中，可能某些代码的使用频率会比其他代码高很多。当出现需要拓展我们的应用时，我们此时只能拓展整个代码库而非其中的部分代码。比如现在应用的瓶颈出现在了验证模块上，由于验证模块是和整个应用的代码库高度耦合的，那么我们只能选择拓展整个代码库来摆脱瓶颈。但如果验证模块本身是一个微服务，那么我们只需要拓展验证模块即可。

微服务的理念让你能撰写低耦合的代码，这样更容易横向拓展，这非常适合于如今云端的开发环境。

**Nginx有一系列文章来探讨了有关微服务的诸多概念，可以[在此阅读](https://www.nginx.com/blog/introduction-to-microservices/)。**

# 为何选用Golang?

尽管很多语言都能实现微服务（毕竟微服务只是一种概念而非具体的框架），但有些语言对于微服务的支持会更好。Golang就是其中之一。

Golang本身非常的轻量，速度飞快。最重要的是，它对并发提供了非常好的支持，这一点能更好的利用多核处理器，以及帮助我们同时在不同的机器上运行代码。

Golang的标准库对网络服务有非常好的支持。

最后，Golang有一个非常棒的微服务框架，go-mirco，我们将在以后用到他。

# 何为protobuf/gRPC

由于每个微服务对应一个独立运行的代码库，一个很自然的问题就是如何在这些微服务之间通信。

我们可以使用传统的REST，用http传输JSON或者XML。但用这种方法的一个问题在于，当两个微服务A和B之间要通信时，A要先把数据编码成JSON/XML，然后发送一个大字符串给B，然后B在将数据从JSON/XML解码。这在大型应用中可能会造成大量的开销。尽管我们在和浏览器交互时必须使用这种方法，但微服务之间可以选择其他方式。

gRPC就是这另外一种方式。gRPC是谷歌出品的一个RPC通信工具，它很轻量，且其协议是基于二进制的。让我们来仔细研究下这个定义。gRPC将二进制当作其核心的编码格式。在我们使用JSON的RESTful例子中，我们的数据会以字符串的格式通过http传输。字符串包含了相对大量的元数据，用于描述其编码格式，长度，内容格式以及其他必要数据。之所以包含这些元数据，是因为要让传统的网页浏览器知道收到的数据会是怎样的。但是在两个微服务之间通信时，我们不一定需要这么多元数据。我们可以只需要更轻量的二进制数据。gRPC支持全新的HTTP 2协议，正好可以使用二进制数据。gRPC甚至可以建立双向的流数据。HTTP 2是gRPC的基础，如果你想了解更多HTTP 2的内容，可以看[Google的这篇文章](https://developers.google.com/web/fundamentals/performance/http2/)。

那么我们该怎么用二进制数据呢？gRPC使用protobuf来描述数据格式。使用Protobuf，你可以清晰的定义一个微服务的interface。关于gRPC，我建议你读一读[这篇文章](https://blog.gopheracademy.com/advent-2017/go-grpc-beyond-basics/)。


## Go语言实现GRPC远程调用

### 定义服务(Service)

  如果想要将消息类型用在RPC(远程方法调用)系统中，可以在.proto文件中定义一个RPC服务接口，protocol buffer编译器将会根据所选择的不同语言生成服务接口代码及存根。如，想要定义一个RPC服务并具有一个方法，该方法能够接收 SearchRequest并返回一个SearchResponse，此时可以在.proto文件中进行如下定义：

```protobuf
service GreetService {
  //rpc 服务的函数名 (传入参数)      返回    (返回参数)
    rpc SayHello (HelloRequest) returns (HelloReply) {}
}
```

  最直观的使用protocol buffer的RPC系统是gRPC一个由谷歌开发的语言和平台中的开源的RPC系统，gRPC在使用protocl buffer时非常有效，如果使用特殊的protocol buffer插件可以直接为您从.proto文件中产生相关的RPC代码。

在项目的根目录下创建一个 proto 文件`$GOPATH/src/grpc/myproto/myproto.proto`

`myproto.proto`文件内容为：

```protobuf
//$GOPATH/src/grpc/myproto/myproto.proto
syntax = "proto3";

package myproto;

service GreetService {
	// 创建接口
    rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
    string name = 1;
}

message HelloReply {
    string message = 1;
}
```

这里有几点需要注意。首先，你得定义`service`。一个`service`定义了此服务暴露给外界的交互interface。然后，你得定义`message`。宽泛的讲，`message`就是你的数据结构

这个文件里，`message`由protobuf处理，而`service`则是由protobuf的grpc插件处理。这个grpc插件使我们定义的`service`能使用`message`。

有了这个proto文件还不够，我们需要使用protobuf的工具来编译它。为了方便，让我们写一个`Makefile`来帮助我们编译文件。`grpc/Makefile`内容如下：

```Makefile
build:
        protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/grpc/ \
          myproto/myproto.proto
```

这段代码会调用protoc，它负责将我们的protobuf文件编译成代码。同时我们还指定了grpc的插件，以及最终输出文件的位置。

现在，如果你在 grpc 项目根目录运行`make build`，然后前往文件夹`myproto/`，你应该可以看到一个新的Golang文件`myproto.pb.go`。这个文件是protoc自动生成的，它将proto文件中的`service`转化成了需要我们在Golang代码中需要编写的`interface`。

### gRPC-Server编写 

让我们现在来实现这个`interface`。创建`srv/main.go`文件：

```go
package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc/myproto" // 导入生成的myproto.pb.go文件
	"log"
	"net"
)

const (
	Addr = "127.0.0.1:8889"
)

// GreetService要实现在proto中定义的所有方法。当你不确定时
// 可以去对应的*.pb.go文件里查看需要实现的方法及其定义
type GreetService struct{}

// SayHello - 在proto中，我们只给这个微服务定义了一个方法
// 就是这个SayHello方法，它接受一个context以及proto中定义的
// HelloRequest消息，这个HelloRequest是由gRPC的服务器处理后提供给你的
func (s *GreetService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
     // 返回的数据也要符合proto中定义的数据结构
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
    // 配置gRPC服务器
	lis, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// new一个服务对象
	srv := grpc.NewServer()
	// 在gRPC服务器上注册微服务，这会将我们的代码和*.pb.go中的各种interface对应起来
	pb.RegisterGreetServiceServer(srv, &GreetService{})
    // 在gRPC服务器上注册反射服务。
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
```

总的来说，我们实现了 HelloRequest 微服务所需要的方法，并建立了一个服务器监听本机的 8889 端口。如果你此时运行`go run main.go`，你肯定看不见任何输出，因为我们还没写客户端代码呢！


### gRPC-Client编写 

现在就让我们看看怎么写客户端代码；

请在项目的根目录下建立一个新的文件夹`mkdir cli/`；在这个文件夹中，我们创建`main.go`文件:

```go
package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc/myproto" //导入编译好的protobuf
	"log"
	"os"
	"time"
)

const (
	address     = "127.0.0.1:8889"
	defaultName = "World!"
)

func main() {
    // 建立到服务器的连接。
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial not connect: %v", err)
	}
    //延迟关闭连接
	defer conn.Close()
	//调用protobuf的函数创建客户端连接句柄
	cli := pb.NewGreetServiceClient(conn)
	//命令行传参
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//调用protobuf的sayhello函数
	r, err := cli.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//打印结果
	log.Printf("Greeting: %s", r.Message)
}
```

在`srv`下运行`go run main.go`, 再另一个终端中，在`cli`下运行`go run main.go`, 你应该能看到一条消息`Hello world`。

至此，我们使用protobuf和grpc创建了一个微服务以及一个客户端。





