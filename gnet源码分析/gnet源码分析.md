# 1、简介

[gnet](https://github.com/panjf2000/gnet) 是一个基于事件驱动的高性能和轻量级网络框架。基本特性：

- 直接使用了epoll和kqueue系统调用而非标准Go网络包；
- 类似netty或者libuv网络库，使得gnet到达了一个远超Go net的性能表现，不是为了取代Go的网络标准库，而是为了创造出高效处理网络包的Go语言网络服务器框架。

功能：

- 高性能的基于多线程/Go程网络模型的 event-loop 事件驱动；
- 内置 bytes 内存池，由开源库 bytebufferpool 提供支持；
- 整个生命周期是无锁的；
- 高效、可重用而且自动伸缩的环形内存 buffer；
-  灵活的事件定时器；
- 内置多种编解码器，支持对 TCP 数据流分包：LineBasedFrameCodec, DelimiterBasedFrameCodec, FixedLengthFrameCodec 和 LengthFieldBasedFrameCodec，参考自 [netty codec](https://netty.io/4.1/api/io/netty/handler/codec/package-summary.html)，而且支持自定制编解码器。

# 2、核心设计

## 2.1 多线程模型



## 2.2 可重用且自动扩容的 Ring-Buffer

gnet内置了inbound 和 outbound 两个 buffers，基于 Ring-Buffer 原理实现，分别用来缓冲输入输出的网络数据以及管理内存，gnet 里面的 ring buffer 能够重用内存以及按需扩容。

对于 TCP 协议的流数据，使用 gnet 不需要业务方为了解析应用层协议而自己维护和管理 buffers，gnet 会替业务方完成缓冲和管理网络数据的任务，降低业务代码的复杂性以及降低开发者的心智负担，使得开发者能够专注于业务逻辑而非一些底层实现。



# 3、开始使用





