感谢：[这个链接](https://github.com/zhouhaibing089/Blog/issues/5)解决了我的迷惑

在学习go语言的过程当中, 你也许问过自己一个问题, `GOPATH`到底是做什么的? 或许在写过一些代码之后, 很快你就会在网上找到一些资料, 它告诉我们`GOPATH`这个环境变量它指定了一个目录, 这个目录包含了我们所有的源码, 比如这里的介绍, 但是我觉得光是理解这一点是不够的, 因为你会发现让人为难的是, 看起来似乎我们每新建一个项目, 都要往`GOPATH`再补充一条路径, 这实在是反人类的做法啊.

我自己曾经也有过那样的顾虑, 为了巩固自己的理解, 同时也为了让其他读者更清楚`GOPATH`的作用, 我希望这篇小文希望能够起到一定的作用.

基本探索
让我们先从目录结构开始说起, 在写registry-watch这个项目的时候, 编辑器经常提示我结构不正确, 于是在网上搜了搜, 找到了这个.

按照文档的说法, 作为一个workspace, 它需要包含三个目录:

```bash
workspace
  |-- src
  |-- pkg
  |-- bin
```

`src`表示我们的源码目录, 在不考虑另外两个目录的情况下, 假设我写了一个demo项目, 并新建了一个`hello.go`, 它是`main`模块, 同时它依赖于另外一个函数, 该函数位于`library/world.go`:

hello.go

```go
package main

import "fmt"
import "demo/library"

func main() {
    fmt.Println("Hello, " + library.World())
}
```

world.go

```go
package library

// World return "World"
func World() string {
    return "World"
}
```

于是目录结构变成了现在这个样子:

```bash
workspace
  |-- src
  |     |-- demo
  |     |     |-- hello.go
  |     |     |-- library
  |     |     |     |-- world.go
  |-- pkg
  |-- bin
```

现在我`cd src`目录并运行`go install demo`后, 我们发现以下变化:

```bash
workspace
  |-- src
  |     |-- demo
  |     |     |-- hello.go
  |     |     |-- library
  |     |     |     |-- world.go
  |-- pkg
  |     |-- darwin_amd64
  |     |     |-- demo
  |     |     |     |-- library.a
  |-- bin
  |     |-- demo
```
这个例子基本说明了问题, 当我们写的模块是`main`时, 它会对应到一个可执行文件, 并且编译后的文件会被复制到`bin`目录, 如果是其他模块, 它会被编译成一个库文件, 并且被复制到`pkg`目录. 这就是我们必须提供三个目录的原因, 一个放源代码, 一个放编译后的可执行文件, 另外一个放编译后的库文件.

协作
一个项目总是会由多个成员进行协作开发, 在观察我们的项目结构之后, 很自然的, 我们会发现应该被提交的代码只有`hello.go`和`library/world.go`, 也就是说我们的`.git`目录应该位于`src/demo`之下.

```bash
workspace
  |-- src
  |     |-- demo
  |     |     |-- .git
  |     |     |-- hello.go
  |     |     |-- library
  |     |     |     |-- world.go
  |-- pkg
  |-- bin
```

再考虑另外一个问题, 我们经常会运行`go get xxx`去安装一些包, 这些包会被下载到`$GOPATH/src`目录下, 当另外一个团队成员拉下我们的代码时, 如果还需要逐个去运行`go get xxx`那就太不方便了, 于是我们需要借助一个依赖管理的工具, 这就是`godep`. 当我们运行过`go get xxxx`安装过依赖之后, 只要再到项目目录下, 运行`godep save`, 便会将所有的依赖记录在一个文件当中, 并且这个时候我们的目录结构再次发生了一些小小的变化.

```bash
workspace
  |-- src
  |     |-- demo
  |     |     |-- .git
  |     |     |-- Godeps
  |     |     |     |-- _workspace
  |     |     |     |-- Godeps.json
  |     |     |-- hello.go
  |     |     |-- library
  |     |     |     |-- world.go
  |-- pkg
  |-- bin
```
现在来到需要拉代码的这一方, 对于像上面这样一个仓库. 我们需要多考虑一些事情. 我们不如以kubernetes为例来做示例:

```shell
cd workspace/src
git clone https://github.com/kubernetes/kubernetes.git k8s.io/kubernetes
godep restore
```

初看可能会奇怪, 为什么需要指定目的目录呢, 尤其是`k8s.io`看起来好生奇怪, 其实只要大概找一个源码文件看看就知道了, 比如这个文件

pkg/api/context.go

```go
package api

import (
    stderrs "errors"
    "time"

    "golang.org/x/net/context"
    "k8s.io/kubernetes/pkg/auth/user"
)
```

我们知道`import`路径都是针对`$GOPATH/src`的, 所以很自然的我们要有`k8s.io`这个目录.

总结
这个时候, 我们回到一开始提到的顾虑, 我们真的需要对每一个项目都添加一条`GOPATH`路径吗? 不然, 我们所说的项目不过是以一个文件夹的形式存在于`$GOPATH/src`中, 我们只需要配置一个`GOPATH`, 并把项目都建在`src`目录下就可以了.