[TOC]

# getXlsx() - 读写xlsx表格

```go
func main() {
	// 不存在则新建, 会读取所有内容到内存
	xlsx := getXlsx("Book1.xlsx")
	// 不存在则新建, 表名字的第一个字母会大写
	sheet1 := xlsx.getSheet("sheet4")
	// 设置B列第14行的值为value, 和C列第3行的值为key
	sheet1.set("B14", "value").set("C3", "key")
	// 获取值
	fmt.Println(sheet1.get("C3"))
	// 关闭
	xlsx.close()
}
```

# getSelenium() - 使用selenium做web的测试

查xpath元素的时候要使用selenium打开的窗口去查, 这样才准

```go
func main() {
	// 需要chromedriver在PATH环境变量所在的目录里面, 会先起一个服务, 然后连接, 返回一个客户端
	sn := getSelenium("https://example.com/auth/login")
	defer sn.close() // 关闭服务端和客户端

	// 登录
	lg.trace("选语言")
	// 直接写select的下拉菜单的option的xpath
	sn.first(`/html/body/div/div[1]/div[1]/div[2]/form/div[3]/div/select/option[2]`).click()
	lg.trace("输入用户名")
	sn.first(`//*[@id="login"]`).clear().input("user") // 先清空再输入
	lg.trace("输入密码")
	sn.first(`//*[@id="password"]`).clear().input("pass")
	lg.trace("点击登录")
	sn.first(`/html/body/div[1]/div[1]/div[1]/div[2]/form/center/div/input`).click()

	// 查找单个会员信息
	lg.trace("搜索单个会员帐号")
	sn.first(`//*[@id="gotomemberinfo"]`).input("ui147147").pressEnter() // 输入并回车
	vipLevel := sn.first(`/html/body/div[2]/div/div[1]/div[1]/div[3]/div/div/div/div[1]/fieldset[2]/div/div[1]/div[1]`).text()
	lg.trace("VIP等级:", is.VipLevel)

	lg.trace("直接查询接口")
	url := sn.url() // 获取当前的url
	uid := strSplit(url, "/")[len(strSplit(url, "/"))-1]
	j := httpGet("https://example.com/player_management/getBigWalletDetails/"+uid, httpHeader{
		"cookie": sn.cookie(), // cookie的字符串
	}).content
	//lg.trace(j)
	jj := getXPathJson(j)
	balance = jj.first("//bigWallet/total").text()
	lg.trace("总余额:", is.balance)


	select {}
}
```

# getWebSocket() - websocket的客户端

```go
// 连接 
ws := getWebSocket("wss://wsapi.example.com:8001/socket/?user=example&password=example")

// 发送文本
ws.send(jsonDumps(j))

// 遍历收到的数据
for msg := range ws.recvMsgChan {
  lg.trace(msg)
}
```

# getMatrix() - Matrix服务器的Bot

```go
  // 设置homeserver的url, 以及需要发送到的会话的id, 这个房间的消息不能是加密的
  cli = getMatrix(args["server"]).setRoomID(args["roomID"])

  // 可以使用帐号密码登录, 它会先获取token, 然后使用token去交互, 每次登录获取一次token
  token := cli.login(args["username"], args["password"]) // 登录成功会返回一个token
  // 也可以直接使用token, 默认synapse的token是没有过期时间的
  cli.setToken(args["userid"], args["token"])
  
	cli.send(msg)
}
```

# argparser() - 解析命令行参数

* -b在后台执行
* -c指定配置文件, 如果没有指定, 会寻找二进制目录下的配置文件, 以及当前工作目录下的配置文件, 查找规则例如二进制文件名字为app, 则查找app.ini文件
* -h查看帮助, 可以直接在命令行指定
* 可以通过环境变量指定, 指定的名字参考命令行方式
* 参数读取优先级, 首先命令行指定, 其次环境变量, 然后读配置文件, 如果都没有, 就使用内置的默认值

```go
type argStruct struct {
	InCluster      bool
	ConfigFile     string
	Namespace      string
	TelegramAPIKey string
	TelegramChatID int64
}

func main() {
	args := new(argStruct)
	a := argparser("kubernetes的pod程序崩溃通知程序")
	args.InCluster = a.getBool("", "InCluster", "false", "是否在集群内部, 如果不在集群内部需要指定config文件")
	args.ConfigFile = a.get("", "ConfigFile", "", "如果不在集群内部, 需要指定配置文件")
	args.Namespace = a.get("", "Namespace", "", "需要监听事件的namespace, 逗号分割, 默认为空, 即监听所有")
	args.TelegramAPIKey = a.get("", "TelegramAPIKey", "", "telegram bot的api key")
	args.TelegramChatID = a.getInt64("", "TelegramChatID", "0", "telegram bot需要发送通知到的group或者channel的id")
	a.parseArgs()
}
```

# zlibCompress() - 压缩字符串

# zlibDecompress() - 解压字符串

# getLanguage() - 获取文本的语言

注意: 需要测试一下, 看看具体会输出什么, 再做判断, 例如中文就不是Chinese

# en2zh() - 翻译英文到中文

调用百度, 百度有一个从github拿到的appkey, 1秒一次请求, 没有限制使用(要做try-except, 不保证每次都能成功)

# strRemoveHtmlTag() - 把html标签去掉, 留下文本

# getRabbitMQ() - 使用RabbitMQ队列

```go
func main() {
	go func() {
		rb := getRabbitMQ("amqp://guest:guest@rabbitmq-svc:5672/", "default")
		rb.send(map[string]string{"data": "Test Message"})
	}()

	go func() {
		rb := getRabbitMQ("amqp://guest:guest@rabbitmq-svc:5672/", "default")
		msg := <-rb.recv()
		lg.debug(msg)
	}()

	select {}
}
```

# getTelegraph() - 往Telegraph上面Post文章

```go
func main(){
  tp := getTelegraph("Author Name") // 会有速率限制
  pg = tp.post("Article Title", "Plain text or HTML document") // 会有速率限制
  lg.trace("url:", pg.url)
}
```

# getPrometheus() - 查询prometheus

如果聚合成一个值, 没有label的, 应该是不支持(暂时没这个需求, 没做适配)

```go
func main() {
	p := getPrometheus("http://localhost:9090")
	pr := p.query("sum_over_time(channel_register_count_in_5_minutes{channel=\"1\"}[1h]) / sum_over_time(channel_inpour_count_in_5_minutes[1h]) < 100")
	lg.debug(pr)
}
```

上例输出

```go
[]main.prometheusResultStruct{
  main.prometheusResultStruct{
    Label: map[string]string{
      "instance":  "10.0.0.1:9100",
      "job":       "my-service-svc",
      "namespace": "default",
      "pod":       "my-service-332332234-1231232",
      "service":   "my-service-svc",
      "channel":   "1",
      "endpoint":  "my-endpoint",
    },
    Value: 44.651376,
  },
}
```

# nslookup() - 查询dns的各种记录

```go

func main() {
	fmt.Println(nslookup("www.facebook.com", "a"))     // [[www.facebook.com. CNAME star-mini.c10r.facebook.com.] [star-mini.c10r.facebook.com. A 157.240.7.35]]
	fmt.Println(nslookup("facebook.com", "ns"))        // [[facebook.com. NS d.ns.facebook.com.] [facebook.com. NS c.ns.facebook.com.] [facebook.com. NS b.ns.facebook.com.] [facebook.com. NS a.ns.facebook.com.]]
	fmt.Println(nslookup("facebook.com", "mx"))        // [[facebook.com. MX 10 smtpin.vvv.facebook.com.]]
	fmt.Println(nslookup("google.com", "aaaa"))        // [[google.com. AAAA 2a00:1450:4019:80d::200e]]
	fmt.Println(nslookup("www.facebook.com", "cname")) // [www.facebook.com. CNAME star-mini.c10r.facebook.com.]]
  fmt.Println(nslookup("_acme-challenge.example.com", "txt")) // [[_acme-challenge.example.com. TXT "97szgn1OZluTy7_70WDRW_x4nJ1TTlHC1BHRwNAMoCs"] [_acme-challenge.example.com. TXT "UwD4UhNf8S8s8cBGOuT3kipMYL1nZgpktnX4RoZYItM"]]
}
```

# pathIsSymlink() - 检查给定的路径是否是符号链接, 如果是则返回true, 否则返回false

# getTTLCache() - 返回一个k/v带ttl的cache

```go
func main() {
	cache := getTTLCache(1) // 1秒超时
	lg.debug(cache)
	cache.set("k1", "v1")
	lg.trace(cache.exists("k1")) // true
	sleep(1)
	cache.set("k2", "v2")
	lg.trace(cache.exists("k1")) // false
	lg.trace(cache.get("k2"))    // v2
	lg.trace(cache.count())      // 1
	sleep(1)
	lg.trace(cache.get("k2")) // Key k2 not found in cache
}
```

每次set都会刷新时间

```go
func main() {
	cache := getTTLCache(1) // 1秒超时

	go func() {
		for {
			cache.set("k2", "v2")
			sleep(0.1)
		}
	}()

	cache.set("k2", "v2")

	for {
		lg.trace(cache.get("k2"))
		sleep(1)
	}
}
```

# getGin() - 返回一个gin的结构体

# getXPathJson() - 用xpath去解析json

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `{
		"name": "John",
		"age"      : 26,
		"address"  : {
		  "streetAddress": "naist street",
		  "city"         : "Nara",
		  "postalCode"   : "630-0192"
		},
		"phoneNumbers": [
		  {
			"type"  : "iPhone",
			"number": "0123-4567-8888"
		  },
		  {
			"type"  : "home",
			"number": "0123-4567-8910"
		  }
    ],
    "nullvalue": null
	}`
	x := getXPathJson(s)
	name := x.first("//name").text()
	fmt.Printf("Name: %s\n", name)

	var a []string
	for _, n := range x.find("//phoneNumbers/*/number") {
		a = append(a, n.text())
	}
	fmt.Printf("All phone number: %s\n", strings.Join(a, ","))

	if n := x.first("//address/streetAddress"); n != nil {
		fmt.Printf("address: %s\n", n.text())
	}

  fmt.Println("First phone number:", x.first("//phoneNumbers[1]/*/number").text()) // 这里有个*, 因为数组的话会额外加一个元素element在xml里面, 去包含
  
  lg.debug(x.first("//nullvalue").text()) // ""
}
```

输出

```
Name: John
All phone number: 0123-4567-8888,0123-4567-8910
address: naist street
First phone number: 0123-4567-8888
```

# getXPath() - 使用xpath去解析html(xml)

准备个xml实例

```xml
<bookstore>

<book category="cooking">
  <title lang="en">Everyday Italian</title>
  <author>Giada De Laurentiis</author>
  <year>2005</year>
  <price>30.00</price>
</book>

<book category="children">
  <title lang="zh-cn">Harry Potter</title>
  <author>J K. Rowling</author>
  <year>2005</year>
  <price>29.99</price>
</book>

<book category="web">
  <title lang="zh-tw">XQuery Kick Start</title>
  <author>James McGovern</author>
  <author>Per Bothner</author>
  <author>Kurt Cagle</author>
  <author>James Linn</author>
  <author>Vaidyanathan Nagarajan</author>
  <year>2003</year>
  <price>49.99</price>
</book>

<book category="web">
  <title lang="zh-hk">Learning XML</title>
  <author>Erik T. Ray</author>
  <year>2003</year>
  <price>39.95</price>
</book>

</bookstore> 
```

如下代码

```go
package main

func main() {
	content := open("i.html").read()
	doc := getXPath(content)
	for _, title := range doc.find("//title") {
		lg.trace("获取lang属性: " + title.getAttr("lang") + ". 获取title标签的文字: " + title.text())
	}

	book := doc.find("//bookstore/book[1]")[0]
	lg.trace("只包含子节点的html: ", book.childHTML())

	lg.trace("包含book标签本身的html: ", book.html())

	author := doc.find("//bookstore/book[1]/author[2]")
	lg.trace("在第一个book找不到第二个author:", author)
}
```

输出

```s
02-10 00:54:10   1 [TRAC] (main.go:7) 获取lang属性: en. 获取title标签的文字: Everyday Italian
02-10 00:54:10   1 [TRAC] (main.go:7) 获取lang属性: zh-cn. 获取title标签的文字: Harry Potter
02-10 00:54:10   1 [TRAC] (main.go:7) 获取lang属性: zh-tw. 获取title标签的文字: XQuery Kick Start
02-10 00:54:10   1 [TRAC] (main.go:7) 获取lang属性: zh-hk. 获取title标签的文字: Learning XML
02-10 00:54:10   1 [TRAC] (main.go:11) 只包含子节点的html:  
                        <title lang="en">Everyday Italian</title>
                        <author>Giada De Laurentiis</author>
                        <year>2005</year>
                        <price>30.00</price>
02-10 00:54:10   1 [TRAC] (main.go:13) 包含book标签本身的html:  <book category="cooking">
                        <title lang="en">Everyday Italian</title>
                        <author>Giada De Laurentiis</author>
                        <year>2005</year>
                        <price>30.00</price>
                      </book>
02-10 00:54:10   1 [TRAC] (main.go:16) 在第一个book找不到第二个author: []
```


# getJavascriptVM() - 执行javascript代码

```go
func main() {
  s := "a = 1;console.log(b);"
  vm := getJavascriptVM()
  vm.set("b", "2")
  vm.run(s)
  print(vm.get("a"))
}
```

# getHightLightHTML() - 获取代码的高亮的html

```go
func main() {
	fd := open("tmp/i.html", "w")
	buf := getHightLightHTML(open("main.go").read(), "go")
	fd.write(buf)
	fd.close()
}
```

# inotify() - 监听文件和目录的改变, 支持递归监控所有子目录

如果被监控的文件或者目录不存在了, 就退出

```go
func main() {
	for ev := range inotify(".") {
		lg.debug(ev)
	}
}
```

# smuxServerWrapper() - 多路复用, 包装kcp或者tcp为一个服务端
# smuxClientWrapper() - 多路复用, 包装kcp或者tcp为一个客户端

以下代码, 服务端kcp监听64321端口, 客户端tcp监听9999端口
启动之后, 客户端跟服务端建立kcp链接, 并做好端口复用的准备
当客户端收到tcp链接, 则对服务端发起一个smux连接, 服务端服务这个连接为socks5代理

```go
package main

import (
	"io"

	"github.com/armon/go-socks5"
)

func main() {
	go func() {
		lg.trace("监听kcp")
		kl := kcpRawListen("0.0.0.0", 64321, "demo key", "demo salt")
		for kc := range kl.accept() {
			ss := smuxServerWrapper(kc)
			for sc := range ss.accept() {
				go func(sc *smuxServerSideConnection) {
					conf := &socks5.Config{}
					server, err := socks5.New(conf)
					if err != nil {
						lg.trace("创建Socks5服务器的时候发生错误", err)
					}
					server.ServeConn(sc.stream)
					sc.close()
				}(sc)
			}
		}
	}()
	lg.trace("监听tcp")
	tl := tcpListen("0.0.0.0", 9999)
	kc := kcpRawConnect("127.0.0.1", 64321, "demo key", "demo salt")
	ss := smuxClientWrapper(kc)
	for tc := range tl.accept() {
		sc := ss.connect()
		go func(tc *tcpServerSideConn) {
			io.Copy(sc.stream, tc.conn)
		}(tc)

		io.Copy(tc.conn, sc.stream)

		sc.close()
		tc.close()
	}
}
```

# map2bin() - 把字符串map转为二进制
# bin2map() - 把二进制转为字符串map

```go
func main() {
	a := map2bin(map[string]string{
		"a": "b",
		"c": "d",
	})
	print(a)

	b := bin2map(a)
	print(b)
}
```

# cmdExists() - 检查指定的命令是否存在

# setFileTime() - 设置文件的mtime和atime

# getArgparseIni() - 处理命令行参数和配置文件

1. 不指定配置文件，只指定参数，则没指定的参数使用默认值
2. 指定配置文件，但是配置文件不存在，则根据默认值生成配置文件，保存退出
3. 指定配置文件，并配置文件存在，读取配置文件的值，如果命令行参数也有指定，则返回命令行参数指定的值。
4. 指定-h或者--help，打印帮助然后退出
5. 制定配置文件，并配置文件存在，格式化配置文件保存。

```go
func main() {
  // 自动读取-c参数指定的配置文件
  ini := getArgparser("一个测试的程序")                      
  // 返回的配置顺序为
  // 1. 如果配置文件的值为空, 检查当前可执行文件目录下面的ini文件, 例如二进制文件名为run, 则查找run.ini
	// 2. 如果配置文件存在, 返回配置文件的值
  // 3. 如果配置文件不存在, 返回默认值
	// 3. 如果有设定环境变量，section.key=value的方式，则不返回默认值，返回环境变量的值，测试k8s可用
  // 4. 如果有设定命令行参数，--section.key的方式，则不返回环境变量的值，返回命令行参数
  ini.getInt("", "env", "product", "环境类型") // 第一个参数, 区块名称可以省略
	ini.getInt("base", "bindPort", "8080", "HTTP服务监听的端口") 
	ini.getInt("base", "JUMP_LIMIT_PER_IP", "1", "每个IP每天跳转的次数")
	ini.getInt("base", "ACCESS_COUNT_TO_JUMP", "1", "每个IP访问多少次之后跳转")
	ini.getInt("base", "PRECENT", "50", "如果要跳转，跳转的概率百分比")
	ini.get("base", "DESTINATION", "http://baidu.com", "跳转到的地址")
	ini.getInt("base", "JUMP_DELAY", "5", "跳转之前等待的秒数")
	ini.get("log", "enableLogFile", "false", "是否启用日志文件")
	ini.get("log", "file", "name.log", "留空则不启用日志")
	ini.getInt("log", "count", "10", "每天一个日志，留几个日志")
	ini.get("log", "level", "trace", "可选trace, debug, info, warn, error")
  p := ini.getBool("log", "displayOnTerminal", "true", "是否打印到屏幕")
  // 处理-h和--help， 如果有指定就打印帮助并退出。
  // 处理-b参数，如果有指定就后台运行（只支持Linux）。
  // 如果初始配置文件不存在，保存模板配置文件并退出。
	ini.parseArgs() 

	print(p)
}
```

# getGoroutineID() - 获取当前Go Routine的ID

# str() - 转换任意类型到字符串类型

# parseUserAgent() - 解析User-Agent

```go
func main() {
	str := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36"
	print(parseUserAgent(str))
}
```

输出

```
ua.UserAgent{
  Name:      "Chrome",
  Version:   "85.0.4183.102",
  OS:        "Linux",
  OSVersion: "x86_64",
  Device:    "",
  Mobile:    false,
  Tablet:    false,
  Desktop:   true,
  Bot:       false,
  URL:       "",
  String:    "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36",
}
```

# fakeName() - 获取假名字

# getTempFilePath() - 返回临时文件的路径

备注：文件不会存在

# getTempDirPath() - 返回临时目录的路径

备注：会创建目录

# getSelfDir() - 返回当前可执行文件所在的目录的绝对路径

# getPinYin() - 中文转拼音

```go
func main() {
	hans := "中文-123-yingwen 空格 many blank 符号 ff --- 8&^^^f fff"
	p := getPinYin(hans)
	fmt.Println(strJoin("-", p)) // zhong-wen-123-yingwen-kong-ge-many-blank-fu-hao-ff-8-f-fff
}
```

# getSnowflakeID() - 获取集群内的唯一有序id

```go
func main() {
	print(getSnowflakeID())   // node1  ==> 1322148295358812160
	print(getSnowflakeID(23)) // node23 ==> 1322148295358902272
}
```

# zipDir() - 压缩文件夹为zip

# unzipDir() - 解压缩zip到文件夹

# md2html() - 转换markdown到html

# fileType() - 根据文件头判断文件类型

备注：文件名和文件扩展名什么的不影响判断

# resizeImg() - 调整图片尺寸大小

备注: height可以指定可以不指定

# sshConnect() - ssh到其他服务器并执行命令、上传下载文件

```go
func main() {
	s := sshConnect("root", "root", "192.168.152.19", 22)
	print(s.exec("id"))
	s.pullFile("anaconda-ks.cfg", "tmp.file")
	s.pushFile("main.go", "main.go")
}
```

# httpHead() - http head 请求

```go

func main() {
	print(httpHead("http://google.com/bzip2"))
}
```

# httpGet() - http get 请求

```go
func main() {
	url := "http://localhost:8888"
	resp := httpGet(url, httpHeader{"abc": "def"}, httpParam{"arg1": "val1"}, httpConfig{timeout: 5}) // 顺序无关，类型判断参数. 设置的是http头，参数，和超时时间5秒
	fmt.Println(resp.statusCode)
	fmt.Println(resp.content)
	fmt.Println(resp.headers)
	fmt.Println(resp.url) // 如果有重定向，为重定向之后的url
}
```

# httpPost() - http Post 请求

```go
func main() {
	url := "http://localhost:8888"
	resp := httpPost(url, httpHeader{"abc": "def"}, httpParam{"arg1": "val1"}) // 顺序无关，类型判断参数
	fmt.Println(resp.statusCode)
	fmt.Println(resp.content)
	fmt.Println(resp.headers)
}
```

# httpPostRaw() - http Post 请求，直接指定body字符串

```go
func main() {
	url := "http://localhost:8888"
	resp := httpPostRaw(url, "a=b", httpHeader{"abc": "def"})
	fmt.Println(resp.statusCode)
	fmt.Println(resp.content)
	fmt.Println(resp.headers)
}
```

# httpPostJSON() - http Post Json

```go
func main() {
	url := "http://localhost:8888"
	post := map[string]interface{}{
		"labels": map[string]interface{}{
			"alertname": "urlJumper ERROR: ",
			"types":     []float64{1, 2, 3.4},
		},
		"annotations": map[string]interface{}{
			"description": "this is a description",
		},
		"startsAt": strftime("%Y-%m-%d %H:%M:%S", now()),
	}
	resp := httpPostJSON(url, &post, httpHeader{"abc": "def"}, httpParam{"arg1": "val1"})
	fmt.Println(resp.statusCode)
	fmt.Println(resp.content)
	fmt.Println(resp.headers)
}
```

# httpPostFile() - http Post 文件

Post的文件在表单里面名字为media

```go
func main() {
	url := "http://localhost:8888"
	resp := httpPostFile(url, "main.go", httpHeader{"abc": "def"}, httpParam{"arg1": "val1"})
	fmt.Println(resp.statusCode)
	fmt.Println(resp.content)
	fmt.Println(resp.headers)
}
```

# kcpListen() - kcp客户端和服务端

注意: 

1. 没有TCP的连接这个概念, 所以需要手动维护连接. 
  1.1 客户端的Connect实际上不会发送任何数据到服务端, 需要随便发送一个东西, 服务端才会收到连接, 类似于TCP的SYN包
  1.2 NAT网关上面有个UDP的超时时间, 如果超时了, 那服务端发送的数据就到不了客户端了, 这里实现了心跳，20秒一次，客户端发给服务端。如果服务端3次20秒都没收到心跳，则关闭连接。如果客户端3次20秒没有收到心跳的回复，也关闭连接。
  1.3 如果发送数据包， 发送之后程序立刻就退出或者关闭连接了，那么这个发送是没有成功的，sleep一下，等它发送完，最好等对端有回应确认
  1.4 任意一端关闭了连接，另一端是不知道的，另一端发送或者读取会timeout，timeout时间是120秒。（应该120秒内还没timeout，就会被心跳goroutine关掉连接）
2. 其它的
  1.1 写入一个关闭的链接就抛异常， 调用send，在另一个goroutine关闭链接也是这样

```go
package main

var key string = "demo key keykeykeykeykeykeykey"
var salt string = "demo salt saltsaltsaltsaltsaltsalt"

var lg *logStruct

func main() {
	args := argparser("test kcp")
	side := args.get("", "side", "s", "\"c\" for client, \"s\" for server")
	addr := args.get("", "addr", "127.0.0.1", "address for listen or connect to")
	port := args.getInt("", "port", "12345", "port for listen or connect to")
	args.parseArgs()

  // 客户端
	if side == "c" {
		c := kcpConnect(addr, port, key, salt)
		c.send("1", "2", "3")
    sleep(1) // 等待1秒， 让数据都发出去， 再退出
  // 服务端
	} else if side == "s" {
		k := kcpListen(addr, port, key, salt)
		c := <-k.accept()
    print(c.recv(10)) // 收到[]string{"1", "2", "3"}, 如果10秒内没收到, 得到nil
	}
}
```

# tcpConnect() - tcp客户端

```go
func main() {
	c := tcpConnect("localhost", 8888)
	defer c.close()
	c.send("GET / HTTP/1.1\r\n\r\n")
	fmt.Println(c.recv(1024))
}
```

# tcpListen() - tcp服务端

```go
func main() {
	l := tcpListen("0.0.0.0", 8899)
	defer l.close()

	for c := range l.accept() {
		fmt.Println(c.recv(1024))
		c.send("HTTP/1.1 200 OK\r\n\r\n")
		c.close()
	}
}
```

# sslConnect() - ssl客户端

```go
func main() {
	c := sslConnect("google.com", 443)
	defer c.close()
	c.send("GET / HTTP/1.1\r\n\r\n")
	fmt.Println(c.recv(1024))
	
}
```

# sslListen() - ssl服务端

```go
func main() {
	l := sslListen("0.0.0.0", 443, "google.com.key", "google.com.crt")
	defer l.close()

	for c := range l.accept() {
		fmt.Println(c.recv(1024))
		c.send("HTTP/1.1 200 OK\r\n\r\n")
		c.close()
	}
}
```

# udpConnect() - udp客户端

```go
func main() {
	c := udpConnect("localhost", 8899)
	defer c.close()
	c.send("Hello World!")
	fmt.Println(c.recv(1024))
}
```

# udpListen() - udp服务端

```go
func main() {
	c := udpListen("0.0.0.0", 8899)
	defer c.close()
	data, addr := c.recvfrom(1024)
	fmt.Println(data)
	c.sendto("You are welcome!", addr)
}
```

# getTelegram() - 使用telegram的bot去发送消息

```go
func main() {
	tg := getTelegram("123234:abcdefg").setChatID(123456)
	tg.send("test message")
	tg.sendFile("a.doc")
	tg.sendImage("j.png")
  // 发送可以点击的链接
  tg.send("<a href=\"http://www.example.com/\">inline URL</a>", tgMsgConfig{"parseMode": "html"}) 
	tg.send("[abc](http://www.example.com/)", tgMsgConfig{"parseMode": "markdown", "DisableWebPagePreview": true}) // 禁用预览
}
```

# getTotp() - 基于时间的一次性密码

```go

func main() {
	print(getTotp("GRIXYK3EV2GGNUUM").validate("055819"))
}
```

# pexpect() - 操控子进程

准备个py，做交互的子进程

```python
try:
    while True:
        a = input("Enter something: ")
        print("You Enter: ", a)
except:
	pass
```

注意：并不是每次都能获取到所有屏幕的输出，需要具体场景具体测试。例如yum就拿不到那个安装程序的时候要输入y的提示，而certbot可以。

```go
func main() {
	i := 0
	p := pexpect("python3 test.py")
	defer p.close()
	p.logToStdout = true
	for p.isAlive {
		sleep(1)
		if strings.Contains(p.buf, "Enter something:") {
			p.sendline(toString(i))
		}
		i++
		if i >= 5 {
			p.close()
			break
		}
	}
	print("Exit code:", p.exitCode)
}
```

# getAES() - aes加密解密

```go
func main() {
	try(func() {
		a := getAES("%SVi9Y]4U@<I[)iO")
		e := a.encrypt("1234567890")
		log.Debug("encrypted: " + e)
		log.Debug("decrypt: " + a.decrypt(e))
	}).except(func(e eee) {
		log.Error(e)
	})
}
```

# strFilter() - 把不符合要求的字符都过滤掉，只剩下符合要求的

注意：如果不指定字符串，则从以下挑选：1234567890_qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM.

# keyInMap() - key是否在map里面

# randomStr() - 生成指定长度的随机字符串

注意：如果不指定字符串，则从以下挑选：abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789

# getRedis() - 读写redis

```go
func main() {
	rdb := getRedis("redis-svc", 6379, "", 1)
	rdb.set("key", "value")
	fmt.Println(*rdb.get("key")) // 如果key不存在返回nil, 存在则返回value的类型为string的指针
	rdb.set("ttl", "delete after 1 second", 1)
	rdb.set("ttl2", "delete after 0.5 second", 0.5)
	rdb.del("key")
}
```

或者, 使用redis的分布式锁

```go
func main() {
	go func() {
		r := getRedis("redis-svc", 6379, "", 0)
		l := r.getLock("tg", 5) // key为tg, 通过这个可以获取同一把锁, 5秒之后锁自动释放
		c := 1
		for {
			l.acquire()
			lg.trace(c)
			c++
      l.release()
		}
	}()

	r := getRedis("redis-svc", 6379, "", 0)
	l := r.getLock("tg", 5)
	c := 1
	for {
		l.acquire()
		lg.info(c)
		c++
    // 没有release, 5秒后自动release
	}
}
```

# getLock() - 获取锁

```go
func main() {
	l := getLock()
	l.acquire()
	l.release()
}
```

# numToBHex() - 10进制转16进制和36进制，十进制转十六进制和三十六进制

```go
func main() {
	print(numToBHex(122342, 16)) // 十进制转十六进制
	print(numToBHex(122342, 36)) // 十进制转三十六进制
	print(bhex2Num("1dde6", 16)) // 十六进制转十进制
	print(bhex2Num("2mee", 36))  // 三十六进制转十进制
}
```

# strStartsWith() - 检查字符串是否以子字符串开始的

# strEndsWith() - 检查字符串是否以子字符串结束的

# fmtTimeDuration() - 把秒的时间转换成年月日时分秒

```go
 func main() {
	fmt.Println("3600 seconds : ", fmtTimeDuration(3600))
	fmt.Println("9999 seconds : ", fmtTimeDuration(9999))
	fmt.Println("8888888888 seconds : ", fmtTimeDuration(8888888888))
}
```

输出

```
3600 seconds : 1 hour 0 minute 0 second
9999 seconds : 2 hours 46 minutes 39 seconds
8888888888 seconds : 40 years 9 months 27 weeks 1 day 15 hours 48 minutes 8 seconds
```

# getChinaCityRank() - 获取中国城市排名

注意：如果大于100则返回">100",参考2019年数据

# sniffer() - 抓网卡的包，返回udp和tcp协议的数据

```go
func main() {
  // for i := range sniffer("any", "port 53 or port 8888") { // 不设置网卡为混杂模式
  for i := range sniffer("any", "port 53 or port 8888", true) { // 设置网卡为混杂模式
		print(i)
	}
}
```

输出：

如果没有mac地址, 可能就是没有mac地址, 例如发送给自己. 
如果网卡名字指定为any, 则没有二层协议, 而是一个linux cooked什么的协议, 只有源的mac, 没有目标的mac
windows无法指定网卡为any
以下的包, 这个mac地址是在一个linux桥接网卡上面抓的, 虚拟机桥接到这个网卡上, 抓到的包是这个网卡的mac地址和物理路由器网关的mac地址

```go
main.networkPacketStruct{
  data:  "\v\xf2\x01\x00\x00\x01\x00\x00\x00\x00\x00\x00\baccounts\x06google\x03com\x00\x00\x01\x00\x01",
  sport: 60928,
  dport: 53,
  proto: "udp",
  ipv:   4,
  sip:   "192.168.0.129",
  dip:   "192.168.0.1",
  smac:  "55:34:bb:56:23:89",
  dmac:  "30:2a:46:45:86:a4", 
}
```

# drawBarChartWithSeries() - 以名字为x轴，数据为y轴，作柱状图

注意：数据的数量会决定图片的宽度

# drawPieChartWithSeries() - 以x为名字，y的百分比为数据，做饼图

注意：数量不能太多，太多的话名字会超出区块范围，10个差不多

# drawLineChartWithNumberSeries() - 以数据为x和y轴，做折线图

# drawLineChartWithTimeSeries() - 以时间为x轴，数据为y轴，做折线图

```go
func main() {
	x := "2020-04-21,2020-05-09,2020-05-11,2020-05-14,2020-05-15,2020-05-17,2020-05-19,2020-05-20,2020-05-21,2020-05-22,2020-05-24,2020-05-25,2020-05-26,2020-05-27,2020-05-28,2020-05-29,2020-05-30,2020-05-31,2020-06-01,2020-06-02,2020-06-03,2020-06-04,2020-06-05,2020-06-06,2020-06-07,2020-06-08,2020-06-09,2020-06-10,2020-06-11,2020-06-12,2020-06-13,2020-06-14,2020-06-15,2020-06-16,2020-06-17,2020-06-18,2020-06-19,2020-06-20,2020-06-21,2020-06-22,2020-06-23,2020-06-25,2020-06-26,2020-06-27,2020-06-28,2020-06-29,2020-06-30,2020-07-01,2020-07-02,2020-07-03,2020-07-04,2020-07-05,2020-07-06,2020-07-07,2020-07-08,2020-07-09,2020-07-10,2020-07-11,2020-07-12,2020-07-13,2020-07-14,2020-07-15,2020-07-16,2020-07-17,2020-07-18,2020-07-19,2020-07-20,2020-07-21,2020-07-22,2020-07-23,2020-07-24,2020-07-25,2020-07-26,2020-07-27,2020-07-28,2020-07-29,2020-07-30,2020-07-31,2020-08-01,2020-08-02,2020-08-03,2020-08-04,2020-08-05,2020-08-06,2020-08-07,2020-08-08,2020-08-09,2020-08-10,2020-08-11,2020-08-12,2020-08-13"
	var xx []int64
	for _, i := range strSplit(x, ",") {
		xx = append(xx, strptime("%Y-%m-%d", i))
	}

	y := "100,100,500,100,100,100,200,700,200,700,300,400,900,1100,1400,900,3004,908,1460,4400,1500,2000,2950,2150,2750,7150,3850,4050,3900,4800,4200,7400,6700,6150,7400,7250,7550,9800,8900,5300,1700,1000,800,1500,1150,1300,2060,3820,4852,4320,4960,5160,2610,2640,3300,1770,2690,2020,2360,2050,1580,1410,1080,850,1540,1410,1460,1540,1620,1370,3328,3898,2218,2238,2398,2038,1700,750,1100,1700,1650,1340,950,2270,540,890,1390,1900,1580,2450,1680"
	var yy []float64
	for _, i := range strSplit(y, ",") {
		yy = append(yy, toFloat64(i))
	}

	drawLineChartWithTimeSeries(xx, yy, "时期", "金额", "每日充值", "output.png") // 不论文件名，都输出png格式图片
}
```

# getCrontab() - 定时任务

注意：不论主机设置的时区，遵从香港市区，即GMT+8

```go
func func1(arg string) {
	print(arg)
}

func main() {
	c := getCrontab()
	
	c.add("*/1 * * * *", func() {
		print(now())
	})

	c.add("*/1 * * * *", func(param1 string, param2 string) {
		print(now(), "with param: "+param1+" and "+param2)
	}, "paramValue1", "paramValue2")

	c.add("00 16 * * *", func1, "args1")

	select {}
}
```

关于时间格式

```
*     *     *     *     *        

^     ^     ^     ^     ^
|     |     |     |     |
|     |     |     |     +----- day of week (0-6) (Sunday=0)
|     |     |     +------- month (1-12)
|     |     +--------- day of month (1-31)
|     +----------- hour (0-23)
+------------- min (0-59)
```

举例子

1. `* * * * *` run on every minute
2. `10 * * * *` run at 0:10, 1:10 etc
4. `10 15 * * *` run at 15:10 every day
5. `* * 1 * *` run on every minute on 1st day of month
6. `0 0 1 1 *` Happy new year schedule
7. `0 0 * * 1` Run at midnight on every Monday
8. `* 10,15,19 * * *` run at 10:00, 15:00 and 19:00
9. `1-15 * * * *` run at 1, 2, 3...15 minute of each hour
10. `0 0-5,10 * * *` run on every hour from 0-5 and in 10 oclock
11. `*/2 * * * *` run every two minutes
12. `10 */3 * * *` run every 3 hours on 10th min
13. `0 12 */2 * *` run at noon on every two days
14. `1-59/2 * * *` * run every two minutes, but on odd minutes

# getIni() - 操作INI配置文件

最佳实践

```go
func main() {
	i := getIni("c.ini")                               // 直接读，不存在就返回空的ini结构体
	print(i.get("section", "key", "value", "comment")) // 直接读，如果存在就返回配置项目，如果不存在就写入配置项目到结构体，并返回默认值
	print(i.save())                                    // 保存文件，如果文件存在返回true，不存在返回false
}
```

其它详细操作

```go
func main() {
	i := getIni("config.ini")                    // 加载config.ini
	print(i.get("paths", "data"))                // 读取paths区块的配置data
	print(i.getInt("threads", "worker"))         // 读取值并转换为Int
	print(i.getFloat64("level", "precent"))      // 读取值并转换为Float64
	i.set("app_mode", "production")              // 设置顶层的配置app_mode
	i.set("paths", "data", "/data/git/grafana")  // 设置paths区块的配置data
	i.set("paths", "config", "/data/git/config") // 新增paths区块的配置config
	i.set("net", "interface", "eth0")            // 新增区块net和配置interface
	i.save("config.ini.local")                   // 保存到文件，会覆盖

	i = getIni() // 建立一个空的ini
	i.set("paths", "data", "/data/git/grafana")
	i.set("paths", "config", "/data/git/config")
	i.set("net", "interface", "eth0")
	i.save("config.ini.local.new")
}
```

测试的config.ini内容

```ini
[paths]
# Path to where grafana can store temp files, sessions, and the sqlite3 db (if that is used)
data = /home/git/grafana
[threads]
worker = 1
[level]
precent = 3.5
```

使用get方法来保存配置

```go
func main() {
	// 初始配置文件
	// [existsSection]
	// existsKey = existsValue
	i := getIni("c.ini")
	print(i.get("existsSection", "existsKey", "defaultValue"))                      // 打印existsValue， 因为值不为空，所以不会覆盖配置文件已有的配置
	print(i.get("notExistsSection", "notExistsKey", "defaultValue"))                // 打印defaultValue, 因为值为空，返回defaultValue，并在save的时候保存到配置文件
	print(i.get("notExistsSection1", "notExistsKey1", "defaultValue1", "Comment1")) // 打印defaultValue, 因为值为空，返回defaultValue，并在save的时候保存到配置文件，并设置注释
	print(i.save())                                                                 // true， 因为c.ini已经存在。
	// 结果配置文件
	// [existsSection]
	// existsKey = existsValue
	// [notExistsSection]
	// notExistsKey = defaultValue
	// [notExistsSection1]
	// ; Comment1
	// notExistsKey1 = defaultValue1
}
```



# getSQLite() - 操作SQLite

其他操作跟MySQL的一样, 但是内存数据库, `:memory:`这种方式有坑, 会说找不到表

# getMySQL() - 操作MySQL

操作表

```go
func main() {
	db := getMySQL("mysql-svc", 3306, "root", "", "test")
	//db := getSQLite("data.db")
	// 建表
	db.createTable("tbName")
	// 添加列
	db.table("tbName").addColumn("intType", "int")      // bigint
	db.table("tbName").addColumn("floatType", "float")  // double
	db.table("tbName").addColumn("vcharType", "string") // VARCHAR(512)
	db.table("tbName").addColumn("textType", "text")    // LONGTEXT
	// 删除列
	db.table("tbName").dropColumn("intType") // SQLite不支持
	// 添加索引
	db.table("tbName").addIndex("floatType")
	db.table("tbName").addIndex("floatType", "vcharType")
	// 删除索引
	db.table("tbName").dropIndex("floatType")
	db.table("tbName").dropIndex("floatType", "vcharType")

	// 也可以链式操作
	db.createTable("usercodes").
		addColumn("usercode", "string").
		addColumn("start", "int").
		addColumn("duration", "int").
		addIndex("usercode")

	// 如果不存在才创建
	db.createTableIfNotExists("usercodes"). // 表存在，后面的也会执行
		addColumnIfNotExists("usercode", "string"). // 列存在，后面的也会执行
		addColumnIfNotExists("start", "int").
		addColumnIfNotExists("duration", "int").
		addIndexIfNotExists("start"). // 索引存在，后面的也会执行
		addIndexIfNotExists("duration")

	// 临时使用的表
	pg := getSQLite(":memory:").
		createTable("progress").
		addColumn("pid", "float").
		addColumn("name", "string").
		addColumn("cpu", "float").
		addColumn("cmd", "string").
		addColumn("start", "int").
		addColumn("end", "int").
		addColumn("notified", "string").
		addIndex("cpu")
}
```

数据查询

```go
func main() {
	db := getMySQL("mysql-svc", 3306, "root", "", "test")

	u := db.table("user") // 之后使用u这个变量去操作的话，线程安全，会上锁

	// select
	reses := db.table("user").fields("id", "name", "age").where("age", ">", 0).orderby("id desc").limit(2).get()
	fmt.Println(reses) // [map[age:6 id:5 name:cat ] map[age:5 id:4 name:monkey]]
	// 获取一个行
	fmt.Println(reses[0]) // map[age:6 id:5 name:cat ]
	// 便利所有行获取某个字段
	for _, r := range reses {
		fmt.Println(r["name"])
	}

	// 获取第一条记录
	res := db.table("user").where("age", ">", 0).orderby("id", "desc").first()
	fmt.Println(res) // map[age:6 id:5 name:cat ]
	print(len(res)) // 0, 如果没数据 

	count := db.table("user").where("age", ">", 0).count()
	fmt.Println(count) // 5

	// 便利数据库
	db.table("user").fields("id", "name", "age").where("age", ">", 0).orderby("id").chunk(2, func(data []gorose.Data) error {
		fmt.Println("In Chunk: ", data)
		// In Chunk:  [map[age:1 id:1 name:cookie] map[age:2 id:2 name:ares]]
		// In Chunk:  [map[age:3 id:3 name:div] map[age:5 id:4 name:monkey]]
		// In Chunk:  [map[age:6 id:5 name:cat ]]
		return nil
	})

	// 插入一条数据
	var data = map[string]interface{}{"age": 17, "name": "it3"}
	id := db.table("user").data(data).insertGetID()
	fmt.Println(id) // 6， 为新数据的id

	// 插入多条数据
	var multiData = []map[string]interface{}{{"age": 18, "name": "it4"}, {"age": 19, "name": "it5"}}
	re := db.table("user").data(multiData).insert()
	fmt.Println(re) // 2 ， 插入的数据条数

	// 更新数据
	re = db.table("user").where("id", "=", 1).orWhere("age", ">", 5).data(map[string]interface{}{"age": 29, "name": "new Name"}).update()
	fmt.Println(re) // 5, 更新的数据条数

	// 删除数据
	re = db.table("user").where("id", "=", 1).delete()
	fmt.Println(re) // 1, 删除的数据条数

	rese := db.query("select count(id) as `count`, `age` from `user` group by `age` order by `count` desc")
	fmt.Println(rese) // [map[age:29 count:4] map[age:2 count:1] map[age:3 count:1] map[age:5 count:1]]

	ress := db.execute("delete from `user` where `age` = 29")
	fmt.Println(ress) // 4

	sql, param := db.table("user").fields("id", "name", "age").where("age", ">", 0).orderby("id desc").limit(2).buildSQL()
	fmt.Println(sql, param) // SELECT `id`,`name`,`age` FROM `user` WHERE `age` > ? ORDER BY id desc LIMIT 2 [0]
}
```

上面范例用到的SQL

```sql
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `age` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

INSERT INTO `user` VALUES (1, 'cookie', 1);
INSERT INTO `user` VALUES (2, 'ares', 2);
INSERT INTO `user` VALUES (3, 'div', 3);
INSERT INTO `user` VALUES (4, 'monkey', 5);
INSERT INTO `user` VALUES (5, 'cat ', 6)
```

# jsonDumps() - 序列化类型到json字符串

```go
func main() {
	a := jsonMap{
		"a": "b",
		"c": "d",
		"e": jsonMap{"f": "g"},
		"h": jsonArr{1, "k"},
	}
	j := jsonDumps(a) // {"a":"b","c":"d","e":{"f":"g"},"h":[1,"k"]}
	print(j)
	k := jsonLoads(j)
	print(k)      // map[a:b c:d e:map[f:g] h:[1 k]]
	print(k["a"]) // b
}
```

# systemWithShell() - 调用shell命令， 通过-c参数， 执行系统命令

支持管道, 退出状态码是命令的退出状态码

# system() - 执行系统命令

注意：不支持管道，要的话要自己接: bash -c '命令'

```go
func main() {
	//status := system("id"， 2) // 超时2秒，如果超时，status为-1
	status := system("id")
	fmt.Println(status)
}
```

# getOutputWithShell() - 获取命令输出

支持管道

# getStatusOutputWithShell() - 获取命令输出和退出状态码

支持管道, 退出状态码是命令的退出状态码

# getOutput() - 获取命令输出
# getStatusOutput() - 获取命令输出和退出状态码

注意：不支持管道，要的话要自己接: bash -c '命令'

```go
func main() {
	// status, output := getStatusOutput("id", 2) // 超时2秒，如果超时，status为-1
	status, output := getStatusOutput("id")
	fmt.Println(status)
	fmt.Println(output)
}
```

# getGodaddy() - 操作Godaddy的DNS记录

```go
func main() {
	gd := getGodaddy("333", "222")
	print(gd.list())                                  // 获取域名列表
	dm := gd.domain("yletx.com")                      // 处理单个主域名
	dm.add("googledns", "A", "8.8.8.8")               // 增
	dm.add("googledns_cname", "CNAME", "twitter.com") // 增
	dm.add("googledns_txt", "TXT", "by twitter")      // 增
	dm.delete("googledns")                            // 删，需要传入名称、类型、值，如果传入空字符串则忽略这一项目的判断
	dm.modify("googledns", "A", "1.1.1.1")            // 改
	print(dm.list())                                  // 查
}
```

# getCloudflare() - 操作Cloudflare的DNS记录

整体来说, 操作接口跟godaddy的没什么区别, 是多了一个添加域名, 以及是否开启cdn. 返回的域名的列表, 信息没有godaddy的多. 

```go
package main

func main() {
	cf := getCloudflare("ip5lwomzy87ohjuoacfzvqup591ipsqi", "example@gmail.com")

	// 添加cloudflare还没有接管的域名
	cf.add("example.com")

	lg.trace("获取域名列表")
	for _, i := range cf.list() {
		if i.Status == "active" {
			print(i)
			break
		}
	}

	dm := cf.domain("example.com")

	lg.trace("获取域名的记录列表")
	for _, i := range dm.list() {
		print(i)
	}

	lg.trace("添加dns记录")
	dm.add("@", "A", "8.8.8.8")
	dm.add("arecord", "a", "7.7.7.7")
	dm.add("cnamerecord", "cname", "google.com")
	dm.add("txtrecord", "txt", "this is a text")

	lg.trace("删除所有a记录")
	dm.delete("", "a", "")

	lg.trace("修改指定a记录")
	dm.add("@", "A", "8.8.8.8")
	sleep(5) // 连续添加, 后面这个会加不上, 不知道休眠多少秒, 就随便5秒吧
	dm.add("@", "A", "6.6.6.6")
	dm.add("arecord", "a", "7.7.7.7")
	// 修改指定a记录
	dm.modify("@", "a", "8.8.8.8", "a", "5.5.5.5")

	lg.trace("设置开启cdn")
	dm.setProxied("@", true)

	lg.trace("删除所有记录")
	dm.delete("", "", "")
}
```

# open() - 操作文件

```go
func main() {
	for line := range open("/etc/passwd").readlines() { // 返回一个chan, for循环处理
		fmt.Println(line) // 要循环完，才会close，所以不要在这里面break
	}

	fd := open("/etc/passwd") // 默认打开模式是r
	defer fd.close()
	fmt.Println(fd.readline()) // 打印一行， 需要手动调用close
	fmt.Println(fd.readline()) // 再打印一行， 需要手动调用close
	fmt.Println(fd.read(10))   // 打印10个字符， 需要手动调用close
	fmt.Println(fd.read())     // 打印所有, 会自动close

	fd = open("text.txt", "w") // 以写方式打开文件
	defer fd.close()
	fd.write("this is a test text")
	fd.close()

	fd = open("text.txt", "rb") // 二进制打开，read函数返回字节数组
	defer fd.close()
	fmt.Println(fd.read(5))

	fd = open("text.txt", "a") // 以追加写方式打开文件
	defer fd.close()
	fd.write(" append text 中文")
	fd.close()

}
```

# tailf() - 持续获取文件的追加内容

如果文件不存在就报错, 如果文件之后不存在了就退出

```go
func main() {
	for line := range tailf("/var/log/apache2/access.log") {
		fmt.Println(line)
	}
}
```

# getTableWithWordWrap() - 在终端打印表格, 如果一行太长的话, 会在设定的宽度内自动换行

如果有设置宽度, 如果一行太长的话, 会在设定的单元格宽度内自动换行

```go
func main() {
  tb := getTable("Header 1", "header 2", "header 3", "header 4")
  
  tb.setMaxCellWidth() // 默认设置30
  tb.setMaxCellWidth(30) // 设定单元格最大宽度

	tb.addRow("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis",
		"Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis",
		"zzLorem ipsum",
		" test",
		"test")

	tb.addRow("漢語，又稱中文[3]、唐話[4]、華語[5]为整个汉语族，或者其语族里的一种语言——汉语族为东亚分析语的一支家族，属汉藏语系。语言如視為單一語言，為世界使用人数最多的语言，目前世界有五分之一人口做為母語。其有多種分支，當中官話最為流行，其衍生而來的現代標準漢語，為中华人民共和国的普通話、以及中華民國的國語，同時是華人地區的通用語。此外，漢語還是聯合國正式語文[6][3]，並被上海合作组织等國際組織採用為官方語言。漢語在以其做為母語的地方會有不同的通稱，例如在臺灣[7]、香港[8]及澳門[9]通稱為「中文」，在馬來西亞及新加坡通稱為「華語」等[註 1]。 ",
		"漢語，又稱中文[3]、唐話[4]、華語[5]为整个汉语族，或者其语族里的一种语言——汉语族为东亚分析语的一支家族，属汉藏语系。语言如視為單一語言，為世界使用人数最多的语言，目前世界有五分之一人口做為母語。其有多種分支，當中官話最為流行，其衍生而來的現代標準漢語，為中华人民共和国的普通話、以及中華民國的國語，同時是華人地區的通用語。此外，漢語還是聯合國正式語文[6][3]，並被上海合作组织等國際組織採用為官方語言。漢語在以其做為母語的地方會有不同的通稱，例如在臺灣[7]、香港[8]及澳門[9]通稱為「中文」，在馬來西亞及新加坡通稱為「華語」等[註 1]。 ",
		"zzLorem ipsum",
		" test",
		"test")

	fmt.Println(tb.render())
}
```

# getTable() - 在终端以打印表格, 如果一行太长的话, 不会自动换行



# getProgressBar() - 进度条

```go
func main() {
	bar := getProgressBar("example bar", 100)
	for i := 0; i < 100; i++ {
		bar.add(1)
		time.Sleep(5 * time.Millisecond)
		if i == 80 {
			bar.setTotal(300) // 重设最大长度
		}
	}

	for i := 0; i < 200; i++ {
		bar.add(1)
		time.Sleep(5 * time.Millisecond)
	}
}
```

# try-except() - 异常处理

```go
func main() {
	try(func() {
		toBool("abc")
	}).except(func(err error) {
		fmt.Println(err)
	})
}
```

把抛异常的方式还原为golang的判断err的方式

```go
func main() {
	err := try(func() {
		toInt("abc")
  }).Error // 不论扔的是什么, 都会转为string, 然后放到Error里面, 可以err.Error()获取这个string

	if err == nil {
		print("正常")
	} else {
		print("出错")
	}
}
```

以及处理完成之后，也可以知道是否进行了异常处理

```go
func main() {
	err := try(func() {
		a := []string{"1", "2"}
		print(a[6])
	}).except(func(e error) {
		// 处理异常
	})

	if err == nil {
		print("正常")
	} else {
		print("出错， 但已经处理完毕")
	}
}
```

支持出错的时候retry的次数, 例如出错retry 3次, 那么实际执行代码会是4次

```go
func main() {
	lg.debug(try(func() {
		lg.trace("try..", now())
		mkdir("mydir")
		system("touch mydir/" + str(now()))
		panicerr("test panic")
	}, tryConfig{retry: 3, sleep: 3})) // 出错之后重试3次, 每次重试间隔3秒
}
```

# panicerr() - 抛异常

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	err := try(func() {
		_, err := strconv.ParseInt("Not Int", 0, 0)
		panicerr(err) // 可以是err
	}).Error
	fmt.Printf("%s\n", err.Error())

	err = try(func() {
		panicerr("自定义消息异常") // 可以是string
	}).Error
	fmt.Printf("%s\n", err.Error())

	err = nil
	panicerr(err) // 如果是nil则不做任何处理

	fmt.Println(newerr("test new error").Error())
}
```

输出如下

```
main.go:11 >> strconv.ParseInt: parsing "Not Int": invalid syntax >> (main.go:9 => try.go:26 => main.go:11)
main.go:16 >> 自定义消息异常 >> (main.go:15 => try.go:26 => main.go:16)
main.go:23 >> test new error >> (main.go:23)
```

# getLogger() - 日志

```go
func main() {
	lg := getLogger()

	// 设置输出为json格式
	lg.json = true
	// 设置为输出Text格式
	lg.json = false

	// 当输出为Text格式的时候
	// 关掉颜色
	lg.color = false
	// 开启颜色
	lg.color = true

	// 不要打印到终端上(可以setLogFile写日志)
	lg.displayOnTerminal = false

	// 设置日志文件
  log.setLogFile("/tmp/my.log", 10) // 每天一个文件，最多存储10个文件，不会删除当前程序没有写过的日志文件，即使名字符合，日志名字格式为/tmp/my.2020-10-10.log
  log.setLogFile("/tmp/my.log", 10, 100) // 每个文件100MB，最多存储10个文件，不会删除当前程序没有写过的日志文件，日志名字格式为/tmp/my.0.log，/tmp/my.1.log，/tmp/my.2.log

	// 设置日志等级
	lg.setLevel("error")
	lg.setLevel("warn")
	lg.setLevel("info")
	lg.setLevel("debug")
  lg.setLevel("trace")
  lg.setLevel("") // 设置为空则禁用日志, 不打印也不写文件
	// 打日志
	lg.error("error")
	lg.warn("warn")
	lg.info("info")
	lg.debug("debug") // 变量会被更详细的打印，类型，值，结构里面的内容
	lg.trace("trace")
}
```

# now() - 当前时间戳(秒)

```go
func main() {
	fmt.Println(now())
}
```

# strptime() - 转换字符串到时间戳

```go
func main() {
	fmt.Println(strptime("%Y-%m-%d %H:%M:%S", "2020-09-30 01:02:03"))
}
```

# strftime() - 转换时间戳到字符串

```go
func main() {
	fmt.Println(strftime("%Y-%m-%d %H:%M:%S", now()))
}
```

# sleep() - 休眠一段时间(秒)

```go
func main() {
	sleep(1)
	sleep(0.5)
}
```

# strIndex() - 查找子字符串在字符串中出现的位置

```go
func main() {
	strIndex("Hello World!", "World")
}
```

# strReplace() - 字符串替换

```go
func main() {
	new := strReplace("Hello World!", "World", "Trump")
	fmt.Println(new)
}
```

# strUpper() - 字符串转换为大写

```go
func main() {
	new := strUpper("Hello World!")
	fmt.Println(new)
}
```

# strLower() - 字符串转换为小写

```go
func main() {
	new := strLower("Hello World!")
	fmt.Println(new)
}
```

# strJoin() - 用指定字符连接字符数组为字符串

```go
func main() {
	new := strJoin(",", []string{"a", "b", "c"})
	fmt.Println(new)
}
```

# strStrip() - 去除字符串首尾处的空白字符（或者其他字符）

```go
func main() {
	new := strStrip("\n     abc    \ndef    \t")
	fmt.Println(new)
}
```

# strLStrip() - 去除字符串开头处的空白字符（或者其他字符）

```go
func main() {
	new := strLStrip("\n     abc    \ndef    \t")
	fmt.Println(new)
}
```

# strRStrip() - 去除字符串末尾处的空白字符（或者其他字符）

```go
func main() {
	new := strRStrip("\n     abc    \ndef    \t")
	fmt.Println(new)
}
```

# urlparse() - 解析url

```go
func main() {
	new := urlparse("https://user:pass@google.com/drive/user?id=john&group=sre")
	fmt.Println(new)
}
```

# urlEncode() - 编码 URL 字符串

```go
func main() {
	new := urlEncode("abc!@#$%^&*()")
	fmt.Println(new)
}
```

# urlDecode() -  解码已编码的 URL 字符串

```go
func main() {
	new := urlDecode("abc%21%40%23%24%25%5E%26%2A%28%29")
	fmt.Println(new)
}
```

# base64Encode() - base64编码

```go
func main() {
	new := base64Encode("Hello World!")
	fmt.Println(new)
}
```

# base64Decode() - base64解码

```go
func main() {
	new := base64Decode("SGVsbG8gV29ybGQh")
	fmt.Println(new)
}
```

# randint() - 生成随机整数

```go
func main() {
	new := randint(0, 100)
	fmt.Println(new)
}
```

# randomChoice() - 随机选取

注意：只支持内置的类型的数组

```go
func main() {
	a := []interface{}{"a", 1, 0.6}
	b := randomChoice(&a)
	fmt.Println(b)
}
```

# pathExists() - 检查文件是否存在

```go
func main() {
	fmt.Println(pathExists("main.go"))
}
```

# pathIsFile() - 路径是否是文件

```go
func main() {
	fmt.Println(pathIsFile("main.go"))
	fmt.Println(pathIsFile("."))
}
```

# pathIsDir() - 路径是否是目录

```go
func main() {
	fmt.Println(pathIsDir("main.go"))
	fmt.Println(pathIsDir("."))
}
```

# unlink() - 删除文件或者目录, 目录或者文件不存在不报错

```go
func main() {
	fmt.Println(unlink("main.go"))
}
```

# copy() - 拷贝文件

```go
func main() {
	copy("main.go", "main.go.bak")
}
```

# rangeInt() - 生成指定范围的int数组

```go
rangeInt(10) // [0,1,2,3,4,5,6,7,8,9]
rangeInt(5, 10) // [5,6,7,8,9]
```

# rename() - 重命名文件

# mkdir() - 递归新建目录, 目录存在不报错，如果路径存在又是一个文件，则报错

# getcwd() - 取得当前工作目录

# basename() - 返回路径中的文件名部分

# isdigit() - 检测变量是否为数字或数字字符串

# getStatusOutput() - 执行命令获取命令退出代码和命令输出

```go
func main() {
	fmt.Println(getStatusOutput("ping -c 5 8.8.8.8"))
	fmt.Println(getStatusOutput("ping -c 5 8.8.8.8", 2)) // 命令执行超时2秒
}
```

# print() - 打印字符串到标准输出，带换行

# printf() - 格式化打印字符串到标准输出

# system() - 执行系统命令并获取退出代码

```go
func main() {
	fmt.Println(system("ping -c 5 8.8.8.8"))
	fmt.Println(system("ping -c 5 8.8.8.8", 2)) // 命令执行超时2秒，未知原因，不是所有都可以确定超时然后终结
}
```

# gethostname() - 获取当前主机名

# gethostbyname() - 获取域名的IP

```go
func main() {
  print(gethostbyname("google.com")) // 使用系统默认的dns服务器
  print(gethostbyname("google.com", "127.0.0.1")) // 指定使用的dns服务器
	print(gethostbyname("google.com", "127.0.0.1:5354")) // 指定使用的dns服务器并指定端口
}
```

# getenv() - 获取环境变量

# dirname() - 获取路径当中的目录部分

# uuid4() - 获取一个UUID

# shortuuid4() - 获取一个短的UUID

# walk() - 遍历目录

返回包括了目录和文件路径

```go
func main() {
	for i := range walk("/etc/php") {
		fmt.Println(i) // /etc/php/7.4/mods-available/sysvsem.ini
	}
}
```

# reFindAll() - 正则表达式查找

<!-- 注意: 如果要匹配换行符, 则在开头指定匹配规则`(?s)`例如: `r := reFindAll("(?s)<script type=\"text/javascript\">(.+?)</script>", data)` -->

```go
func main() {
	for _, v := range reFindAll(`{([a-z]+),([a-z]+)}`, "{city,day}, {state,exit} {zip,ok}") {
		fmt.Println(v)
	}
}
```

输出

```
[{city,day} city day]
[{state,exit} state exit]
[{zip,ok} zip ok]
```

# typeof() - 获取变量的类型

```go
func main() {
	a := 1
	fmt.Println(typeof(a))
}
```

# now() - 返回当前时间戳

对接strftime

```go
func main() {
	fmt.Println(now())
}
```

# md5sum() - 计算字符串的 MD5 散列值

# md5File() - 计算指定文件的 MD5 散列值

# sha1sum() - 计算字符串的 sha1 散列值

# sha1File() - 计算文件的 sha1 散列值

# getBuffer() - 获取一个buffer变量

```go
func main() {
	a := getBuffer()
	a.WriteString("abc")
	fmt.Println(a.String())
}
```

# tailCmdOutput() - 持续获取命令的输出

```go
func main() {
	for line := range tailCmdOutput("dstat -smar") {
		fmt.Println(line)
	}
}
```

# getStdin() - 获取管道的标准输入

```go
func main() {
	for line := range getStdin().readlines() { // 按行读取
		fmt.Println(line)
	}

	fmt.Print(getStdin().read()) // 读取所有
}
```

# strIn() - 子字符串是否在另一个字符串里面

# strSplit() - 按子字符串分割字符串，然后把子字符串的头尾的空白字符去掉

# strSplitlines() - 按行分割字符串, 然后把子字符串的尾部的\r字符去掉

# checkDate() - 检查是否是合法的年月日

# ucfirst() - 将字符串的首字母转换为大写

# lcfirst() -  使一个字符串的第一个字符小写

# ucwords() -  将字符串中每个单词的首字母转换为大写

# substr() -  返回字符串的子串

# parseStr() -  将字符串解析成多个变量

# numberFormat() -  以千位分隔符方式格式化一个数字

# chunkSplit() -  将字符串分割成小块

# wordwrap() - 按照指定长度对字符串进行折行处理

# strlen() -  获取字符串长度

# mbStrlen() -  获取字符串的长度

# strRepeat() - 重复一个字符串

# strstr() - 查找字符串的首次出现

# strtr() - 转换指定字符

# strShuffle() -  随机打乱一个字符串

# explode() -  使用一个字符串分割另一个字符串

# chr() - 返回指定的字符

# ord() -  转换字符串第一个字节为 0-255 之间的值

# nl2br() -  在字符串所有新行之前插入 HTML 换行标记

# addslashes() -  使用反斜线引用字符串

# stripslashes() -   反引用一个引用字符串

# quotemeta() -  转义元字符集

# htmlentities() -  将字符转换为 HTML 转义字符

# htmlEntityDecode() -  将HTML 转义字符转换为 字符

# crc32sum() -  计算一个字符串的 crc32 多项式

# levenshtein() -  计算两个字符串之间的编辑距离

# SimilarText() - 计算两个字符串的相似度

# soundex() - Calculate the soundex key of a string.

# rawurlencode() - 按照 RFC 3986 对 URL 进行解码

# rawurldecode() -  按照 RFC 3986 对 URL 进行编码

# httpBuildQuery() -  生成 URL-encode 之后的请求字符串

# arrayFill() - 使用指定的键和值填充数组

# arrayFlip() - 交换数组中的键和值

# arrayKeys() -  返回数组中部分的或所有的键名

# arrayValues() -  返回数组中所有的值

# arrayMerge() - 合并一个或多个数组

# arrayChunk() -  将一个数组分割成多个

# arrayPad() -  以指定长度将一个值填充进数组

# arraySlice() -  从数组中取出一段

# ArrayRand() -  从数组中随机取出一个或多个单元

# arrayColumn() - 返回数组中指定的一列

# arrayPush() -  将一个或多个单元压入数组的末尾（入栈）

# ArrayPop() - 弹出数组最后一个单元（出栈）

# arrayUnshift() -  在数组开头插入一个或多个单元

# arrayShift() -  将数组开头的单元移出数组

# arrayKeyExists() -  检查数组里是否有指定的键名或索引

# ArrayCombine() - 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值

# arrayReverse() -  返回单元顺序相反的数组

# inArray() - 检查数组中是否存在某个值

# round() -  对浮点数进行四舍五入

# floor() -  舍去法取整. 返回不大于 value 的最接近的整数，舍去小数部分取整。

# ceil() -  进一法取整. 返回不小于 value 的下一个整数，value 如果有小数部分则进一位。

# pi() -  得到圆周率值

# max() -  找出最大值

# min() -  找出最小值

# decbin() - 十进制转换为二进制

# bindec() - 二进制转换为十进制

# hex2bin() - 转换十六进制字符串为二进制字符串

# bin2hex() - 函数把包含数据的二进制字符串转换为十六进制值

# dechex() -  十进制转换为十六进制

# hexdec() -  十六进制转换为十进制

# decoct() -  十进制转换为八进制

# octdec() -  八进制转换为十进制

# baseConvert() -  在任意进制之间转换数字

# IsNan() -  判断是否为合法数值

# stat() -  给出文件的信息

# pathinfo() -  返回文件路径的信息

# fileSize() - 取得文件大小

# filePutContents() - 将一个字符串写入文件

# fileGetContents() - 将整个文件读入一个字符串

# isReadable() -  判断给定文件名是否可读

# isWriteable() - 判断给定的文件名是否可写

# touch() - 设定文件的访问和修改时间

# abspath() -  返回规范化的绝对路径名

# chmod() -  改变文件模式

# chown() -  改变文件的所有者

# fclose() - 关闭一个已打开的文件指针

# filemtime() -  取得文件修改时间

# fgetcsv() - 从文件指针中读入并解析 CSV 字段

# glob() - 寻找与模式匹配的文件路径

# empty() - 检查一个变量是否为空

# passthru() -  执行外部程序并且显示原始输出

# gethostbynamel() -  获取互联网主机名对应的 IPv4 地址列表

# gethostbyaddr() -  获取指定的IP地址对应的主机名

# ip2long() -  将 IPV4 的字符串互联网协议转换成长整型数字

# long2ip() - 将长整型转化为字符串形式带点的互联网标准格式地址（IPV4）

# echo() -  输出一个或多个字符串

# uniqid() -  生成一个ID. 此函数不保证返回值的唯一性。

# putenv() - The setting, like "FOO=BAR"

# memoryGetUsage() - 返回分配给 golang 的内存量

# VersionCompare() -  对比两个「PHP 规范化」的版本数字字符串

# zipOpen() -  打开ZIP存档文件

# pack() - 将数据打包成二进制字符串

# unpack() -  从二进制字符串还原数据

# getTimeDuration() - 返回一个time.Duration结构

# 其他

## vscode的快捷键配置

```json
[
    {
        "key": "alt+q",
        "command": "editor.action.insertSnippet",
        "when": "editorTextFocus",
        "args": {
            "snippet": "if err != nil {\n_, fn, line, _ := runtime.Caller(0)\n panic(filepath.Base(fn) + \":\" + strconv.Itoa(line-2) + \" >> \" + err.Error() + \" >> \" + fmtDebugStack(string(debug.Stack())))\n }"
        }
    }
]
```

## Gin的模板

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/akamensky/argparse"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var ginEng *gin.Engine
var lg *logStruct
var bindPort string
var logLevel string

func main() {
	parser := argparse.NewParser("urljump", "url跳转劫持系统")
	cfg := parser.String("c", "config", &argparse.Options{Required: false, Help: "配置文件路径，不存在则生成示例然后退出"})
	background := parser.Flag("b", "background", &argparse.Options{Required: false, Help: "后台运行"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		exit(0)
	}

	if *background {
		args := os.Args[1:]
		for i := 0; i < len(args); i++ {
			if args[i] == "-b" {
				args[i] = ""
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		os.Exit(0)
	}

	lg = getLogger()

	if *cfg != "" {
		if !pathExists(*cfg) {
			fd := open(*cfg, "w")
			fd.write("[base]\n")
			fd.write("bindPort      = 8080\n")
			fd.write("\n")
			fd.write("[log]\n")
			fd.write("enableLogFile           = false # 是否启用日志文件\n")
			fd.write("file                    = name.log # 留空则不启用日志\n")
			fd.write("count                   = 10 # 每天一个日志，留几个日志\n")
			fd.write("level                   = trace # 可选trace, debug, info, warn, error\n")
			fd.write("displayOnTerminal       = true # 是否打印到屏幕\n")
			fd.close()
			exit(0)
		} else {
			ini := getIni(*cfg)
			bindPort = ini.get("base", "bindPort")
			if toBool(ini.get("log", "enableLogFile")) {
				lg.setLogFile(ini.get("log", "file"), toInt(ini.get("log", "count")))
			}
			logLevel = ini.get("log", "level")
			lg.displayOnTerminal = toBool(ini.get("log", "displayOnTerminal"))
		}
	} else {
		logLevel = "trace"
		bindPort = "8888"
	}

	lg.setLevel(logLevel)

	if !strStartsWith(getSelfDir(), "/tmp/go-build") {
		gin.SetMode(gin.ReleaseMode)
	}

	ginEng = gin.New()

	ginEng.Use(gin.Recovery())
	ginEng.Use(gzip.Gzip(gzip.DefaultCompression))
	ginEng.Use(func(c *gin.Context) {
		if logLevel == "trace" {
			nowTime := time.Now()
			c.Next()

			ipaddr := c.ClientIP()
			var ipinfo *ipLocationInfo
			try(func() {
				ipinfo = getIPLocation(ipaddr)
			}).except(func(e eee) {
				ipinfo = &ipLocationInfo{}
			})

			ipinfostr := ipinfo.Country
			if !itemInArray(ipinfo.Region, []string{"", "N/A", ipinfo.Country}) {
				ipinfostr = ipinfostr + " " + ipinfo.Region
			}
			if ipinfo.City != "" && ipinfo.City != "N/A" {
				ipinfostr = ipinfostr + " " + ipinfo.City
			}

			ua := parseUserAgent(c.Request.Header["User-Agent"][0])

			tm := toString(time.Since(nowTime))
			tmf := fmt.Sprintf("%.2f", toFloat64(reFindAll("[0-9\\.]+", tm)[0][0]))
			tms := reFindAll("[a-zµ]+", tm)[0][0]
			tm = tmf + tms

			logstr := toString(c.Writer.Status()) + " " + fmt.Sprintf("%8v", tm) + " " + c.Request.Method + " " + c.Request.URL.String() + " " + ipaddr + " " + ipinfostr + " " + ua.OS + " " + ua.OSVersion + " " + ua.Name + " " + ua.Version + " " + ua.Device
			lg.trace(logstr)
		} else {
			c.Next()
		}
	})

	statikFS, err := fs.New()
	if err != nil {
		_, fn, line, _ := runtime.Caller(0)
		panic(filepath.Base(fn) + ":" + strconv.Itoa(line-2) + " >> " + err.Error() + " >> " + fmtDebugStack(string(debug.Stack())))
	}
	ginEng.StaticFS("/static", statikFS)
	ginEng.Use(static.Serve("/uploads", static.LocalFile(uploadPath, false)))

	ginEng.NoRoute(func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, statikOpen("404.html").read())
	})

	ginEng.POST("/js/zztj.js", func(c *gin.Context) {

	})

	lg.info("服务器监听在: " + bindPort)
	ginEng.Run(":" + bindPort)
}
```
