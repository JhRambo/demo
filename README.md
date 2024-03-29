1.go垃圾回收机制
答：go语言中的垃圾回收是一种自动内存管理管理机制，它可以自动地释放不再被程序使用的内存。垃圾回收器会周期性的扫描程序中所有已分配的对象，并标记那些不再被使用的对象，然后将它们的内存释放回系统。go语言的垃圾回收器使用一种叫做“三色标记”的算法，它可以在不暂停程序执行的情况下完成垃圾回收。

2.什么是内存逃逸
答：内存逃逸是指在 Go 语言中，一个对象被分配在堆上，但在它本应该被回收之前，它的引用被传递到了其他函数中。这就可能导致这个对象的生命周期被延长，从而增加了垃圾回收器的工作量。

有几种方法可以避免内存逃逸：

使用局部变量：如果一个对象只在函数内部使用，那么就应该将它定义为局部变量。这样就可以确保这个对象只能在函数内部访问，避免了内存逃逸。
使用结构体字段：如果一个对象需要在多个函数之间共享，那么可以将它定义为结构体的字段，这样就可以通过传递结构体的指针来访问这个对象。
使用指针：如果一个对象的大小不固定，那么可以使用指针来避免内存逃逸。因为指针只占用 4 个字节（或 8 个字节）的内存，所以在函数中传递指针比传递对象要高效得多。
使用垃圾回收器优化：Go 语言有一个自带的垃圾回收器，可以帮助我们回收内存。如果你的程序中存在内存逃逸，那么垃圾回收器会帮助你回收这些内存，但是这会增加程序的开销。为了优化程序的性能，应该尽量避免内存逃逸，而不是依靠垃圾回收器来解决问题。
最后，注意 Go 语言的编译器有一些优化技巧可以帮助我们避免内存逃逸。例如，如果编译器发现一个对象在函数调用结束后就不再被使用，那么它会自动将这个对象分配在栈上，而不是堆上。这样就可以避免内存逃逸，提高程序的效率。

3.为什么要 GC
答：手动管理内存挺麻烦，管错或者管漏内存也很糟糕，将会直接导致程序不稳定（持续泄露）甚至直接崩溃。

4.GC 触发场景
答：GC 触发的场景主要分为两大类，分别是：
系统触发：运行时自行根据内置的条件，检查、发现到，则进行 GC 处理，维护整个应用程序的可用性。
手动触发：开发者在业务代码中自行调用 runtime.GC 方法来触发 GC 行为。

5.rune 是什么类型
答：相当于int32
golang中的字符串底层实现是通过byte类型实现的，中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码是utf-8
byte 等同于int8，常用来处理ascii字符
rune 等同于int32，常用来处理unicode或utf-8字符
ASCII码的值如下：
48～57为0到9十个阿拉伯数字
65～90为26个大写英文字母
97～122号为26个小写英文字母

6.什么是Goroutine
答：Goroutine可以理解为一种Go语言的协程。Gorotine可以运行在一个或多个线程上。
Goroutine是一个轻量级的执行线程，多个Goroutine比一个线程轻量，所以管理Goroutine消耗的资源相对更少。
Goroutine是Go中最基本的执行单元，每一个Go程序至少有一个Goroutine：主Goroutine。程序启动时会自动创建。

7.Context的理解
答：context包可以提供一个请求从API请求边界到各goroutine的请求域数据传递、取消信号及截止时间等能力。
Go 语言中的每一个请求的都是通过一个单独的 Goroutine 进行处理的。
HTTP/RPC 请求的处理器往往都会启动新的 Goroutine 访问数据库和 RPC 服务，我们可能会创建多个 Goroutine 来处理一次请求。
而 Context 的主要作用就是在不同的 Goroutine 之间同步请求特定的数据、取消信号以及处理请求的截止日期。
每一个 Context 都会从最顶层的 Goroutine 一层一层传递到最下层，这也是 Golang 中上下文最常见的使用方式。
如果没有 Context，当上层执行的操作出现错误时，下层其实不会收到错误而是会继续执行下去。
在不同 Goroutine 之间对信号进行同步避免对计算资源的浪费，与此同时 Context 还能携带以请求为作用域的键值对信息。
接口：
Deadline：返回 context.Context 被取消的时间，也就是完成工作的截止日期。
Done：    返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消之后关闭，多次调用 Done 方法会返回同一个 Channel。
Err：     返回 context.Context 结束的原因，它只会在 Done 返回的 Channel 被关闭时才会返回非空的值，
          如果 context.Context 被取消，会返回 Canceled 错误；如果 context.Context 超时，会返回 DeadlineExceeded 错误。
Value：   从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据。

8.context的作用和使用场景
答：context 是 Go 语言标准库中提供的一种用于在多个Goroutine之间共享上下文信息的机制，它对控制goroutine运行的生命周期和取消很有用。context 可以用于在一个请求的多个goroutine之间传递请求的元数据、控制goroutine的超时操作和取消操作。在 Go 语言中一个goroutine所运行的都是同一个request请求，每一个goroutine都可能需要访问该请求的元数据、控制goroutine的超时操作和取消操作，如果每个goroutine都需要独立传递这些数据和控制，就需要在多个goroutine之间传递相同的数据，这样会非常麻烦，而 context 可以使用上下文信息来更方便地传递数据。
context 在网络编程中的某些情景下非常有用，例如在 HTTP 服务器中处理请求时，多个处理程序可能需要访问一些请求参数，如请求头或URL参数，而 context 可以让一个处理程序将这些参数存储在上下文中，以便其他处理程序可以方便地访问它们。
context 的几个使用场景：
1.在 HTTP 服务器中处理请求时，可以将请求作为一个上下文，它可以用于在多个处理程序之间传递参数和元数据。
2.在长时间运行的 Goroutine 中控制请求的生命周期。例如，在运行一个耗时的请求时，如果用户提前取消该请求，则可以使用 context 将取消信号传递给 Goroutine。
3.在多个 Goroutine 之间共享取消信号。例如，如果你的应用程序由多个 Goroutine 组成，其中一个 Goroutine（如主 Goroutine）负责收到系统终止信号，那么你可以使用 context 来通知其他 Goroutine 系统正在关闭，以便让它们可以完美停止并释放资源。
总之，context 为多个 Goroutine 之间共享数据提供了便利，并且在某些场景下，使用 context 能让代码更加简洁和易于维护。


99.proto
切到当前文件所在的目录下执行
protoc --go_out=./ hello.proto						生成：hello.pb.go         包含 Protocol Buffers 生成的代码，定义您的消息类型和服务接口。
protoc --go-grpc_out=./ hello.proto					生成：hello_grpc.pb.go    包含 gRPC 服务实现的代码，实现由 .proto 文件定义的服务方法。
protoc --grpc-gateway_out=./ hello.proto		    生成：hello.pb.gw.go      包含 gRPC-Gateway 的代码，它允许您通过 RESTful API 访问您的 gRPC 服务。

protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. hello.proto  一次性生成3个文件

--go_out 		        生成 .pb.go			    需要安装：go install github.com/golang/protobuf/protoc-gen-go@latest
--go-grpc_out 	        生成 _grpc.pb.go		需要安装：go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
--grpc-gateway_out      生成 pb.gw.go			需要安装：go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest


创建容器命令：在根目录执行：
docker-compose -f docker-compose-win.yml up -d --build