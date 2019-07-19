# 何为微服务？

在传统的应用中，所有的功能都是存在于单一的代码库(Monotholic Code Base)中。在表面上看，代码库中的代码可以有几种聚合方式。可能会按照其类型分割，比如controllers, entity, factories，也有可能按照其功能拆分成几个包，比如auth, articles等等。但无论如何，整个应用是建立在一个单一代码库上的。

微服务是对于上述第二种聚合方式的拓展。我们依旧将应用按照其功能拆分成几个包，但不同的是，这些功能包现在都是一个可独立运行的代码库。

# 为何用微服务？

降低复杂性 - 将功能拆分成对应的微服务可以将你的整个代码拆分成更小，更易维护的代码库。这有点类似早期Unix的开发哲学“只做一件事，并做到做好“。在传统的单一代码库应用中，代码的耦合性往往更容易往高耦合发展。这就可能导致撰写和维护代码变得很复杂，也更容易出现漏洞。

拓展性 - 在单一代码库应用中，可能某些代码的使用频率会比其他代码高很多。当出现需要拓展我们的应用时，我们此时只能拓展整个代码库而非其中的部分代码。比如现在应用的瓶颈出现在了验证模块上，由于验证模块是和整个应用的代码库高度耦合的，那么我们只能选择拓展整个代码库来摆脱瓶颈。但如果验证模块本身是一个微服务，那么我们只需要拓展验证模块即可。

微服务的理念让你能撰写低耦合的代码，这样更容易横向拓展，这非常适合于如今云端的开发环境。

**Nginx有一系列文章来探讨了有关微服务的诸多概念，可以在此阅读。**

# 为何选用Golang?

尽管很多语言都能实现微服务（毕竟微服务只是一种概念而非具体的框架），但有些语言对于微服务的支持会更好。Golang就是其中之一。

Golang本身非常的轻量，速度飞快。最重要的是，它对并发提供了非常好的支持，这一点能更好的利用多核处理器，以及帮助我们同时在不同的机器上运行代码。

Golang的标准库对网络服务有非常好的支持。

最后，Golang有一个非常棒的微服务框架，go-mirco，我们将在以后用到他。

# 何为protobuf/gRPC

由于每个微服务对应一个独立运行的代码库，一个很自然的问题就是如何在这些微服务之间通信。

我们可以使用传统的REST，用http传输JSON或者XML。但用这种方法的一个问题在于，当两个微服务A和B之间要通信时，A要先把数据编码成JSON/XML，然后发送一个大字符串给B，然后B在将数据从JSON/XML解码。这在大型应用中可能会造成大量的开销。尽管我们在和浏览器交互时必须使用这种方法，但微服务之间可以选择其他方式。

gRPC就是这另外一种方式。gRPC是谷歌出品的一个RPC通信工具，它很轻量，且其协议是基于二进制的。让我们来仔细研究下这个定义。gRPC将二进制当作其核心的编码格式。在我们使用JSON的RESTful例子中，我们的数据会以字符串的格式通过http传输。字符串包含了相对大量的元数据，用于描述其编码格式，长度，内容格式以及其他必要数据。之所以包含这些元数据，是因为要让传统的网页浏览器知道收到的数据会是怎样的。但是在两个微服务之间通信时，我们不一定需要这么多元数据。我们可以只需要更轻量的二进制数据。gRPC支持全新的HTTP 2协议，正好可以使用二进制数据。gRPC甚至可以建立双向的流数据。HTTP 2是gRPC的基础，如果你想了解更多HTTP 2的内容，可以看[Google的这篇文章](https://developers.google.com/web/fundamentals/performance/http2/)。

