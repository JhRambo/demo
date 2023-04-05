1.go垃圾回收机制
答：go语言中的垃圾回收是一种自动内存管理管理机制，它可以自动地释放不再被程序使用的内存。垃圾回收器会周期性的扫描程序中所有已分配的对象，并标记那些不再被使用的对象，然后将它们的内存释放回系统。go语言的垃圾回收器使用一种叫做“三色标记”的算法，它可以在不暂停程序执行的情况下完成垃圾回收。

2.内存逃逸
答：内存逃逸是指在 Go 语言中，一个对象被分配在堆上，但在它本应该被回收之前，它的引用被传递到了其他函数中。这就可能导致这个对象的生命周期被延长，从而增加了垃圾回收器的工作量。

有几种方法可以避免内存逃逸：

使用局部变量：如果一个对象只在函数内部使用，那么就应该将它定义为局部变量。这样就可以确保这个对象只能在函数内部访问，避免了内存逃逸。
使用结构体字段：如果一个对象需要在多个函数之间共享，那么可以将它定义为结构体的字段。这样就可以通过传递结构体的指针来访问这个对象，而不是直接传递它的引用。
使用指针：如果一个对象的大小不固定，那么可以使用指针来避免内存逃逸。因为指针只占用 4 个字节（或 8 个字节）的内存，所以在函数中传递指针比传递对象要高效得多。
使用垃圾回收器优化：Go 语言有一个自带的垃圾回收器，可以帮助我们回收内存。如果你的程序中存在内存逃逸，那么垃圾回收器会帮助你回收这些内存，但是这会增加程序的开销。为了优化程序的性能，应该尽量避免内存逃逸，而不是依靠垃圾回收器来解决问题。
最后，注意 Go 语言的编译器有一些优化技巧可以帮助我们避免内存逃逸。例如，如果编译器发现一个对象在函数调用结束后就不再被使用，那么它会自动将这个对象分配在栈上，而不是堆上。这样就可以避免内存逃逸，提高程序的效率。

3.为什么要 GC
答：手动管理内存挺麻烦，管错或者管漏内存也很糟糕，将会直接导致程序不稳定（持续泄露）甚至直接崩溃。

4.GC 触发场景
答：GC 触发的场景主要分为两大类，分别是：
系统触发：运行时自行根据内置的条件，检查、发现到，则进行 GC 处理，维护整个应用程序的可用性。
手动触发：开发者在业务代码中自行调用 runtime.GC 方法来触发 GC 行为。

5.监控线程
答：实质上在 Go 运行时（runtime）初始化时，会启动一个 goroutine，用于处理 GC 机制的相关事项。

代码如下：

func init() {
 go forcegchelper()
}
 
func forcegchelper() {
 forcegc.g = getg()
 lockInit(&forcegc.lock, lockRankForcegc)
 for {
  lock(&forcegc.lock)
  if forcegc.idle != 0 {
   throw("forcegc: phase error")
  }
  atomic.Store(&forcegc.idle, 1)
  goparkunlock(&forcegc.lock, waitReasonForceGCIdle, traceEvGoBlock, 1)
    // this goroutine is explicitly resumed by sysmon
  if debug.gctrace > 0 {
   println("GC forced")
  }
 
  gcStart(gcTrigger{kind: gcTriggerTime, now: nanotime()})
 }
}
注：在这段程序中，需要特别关注的是在 forcegchelper 方法中，会调用 goparkunlock 方法让该 goroutine 陷入休眠等待状态，以减少不必要的资源开销。
在休眠后，会由 sysmon 这一个系统监控线程来进行监控、唤醒等行为：

func sysmon() {
 ...
 for {
  ...
  // check if we need to force a GC
  if t := (gcTrigger{kind: gcTriggerTime, now: now}); t.test() && atomic.Load(&forcegc.idle) != 0 {
   lock(&forcegc.lock)
   forcegc.idle = 0
   var list gList
   list.push(forcegc.g)
   injectglist(&list)
   unlock(&forcegc.lock)
  }
  if debug.schedtrace > 0 && lasttrace+int64(debug.schedtrace)*1000000 <= now {
   lasttrace = now
   schedtrace(debug.scheddetail > 0)
  }
  unlock(&sched.sysmonlock)
 }
}
注：这段代码核心的行为就是不断地在 for 循环中，对 gcTriggerTime 和 now 变量进行比较，判断是否达到一定的时间（默认为 2 分钟）。
若达到意味着满足条件，会将 forcegc.g 放到全局队列中接受新的一轮调度，再进行对上面 forcegchelper 的唤醒。

6.能介绍下 rune 类型吗？
相当int32
golang中的字符串底层实现是通过byte类型实现的，中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8
byte 等同于int8，常用来处理ascii字符
rune 等同于int32，常用来处理unicode或utf-8字符
ASCII码的值如下
48～57为0到9十个阿拉伯数字；
65～90为26个大写英文字母；
97～122号为26个小写英文字母.
所以应该写：
数字、大写英文字母、小写英文字母