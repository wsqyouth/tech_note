
### 解析一次完整的HTTP请求与响应
一次完整的http请求主要包含4步：

输入域名，DNS(域名系统)进行域名解析，得到IP地址
客户端和服务器通过三次握手建立连接
连接建立之后，客户端将包装好的http请求传输到服务器，服务器处理后将数据返回给客户端
客户端和服务器之间不再有数据传输后，发起四次挥手断开连接
三次握手：为了防止服务器开启一些无用的连接

客户端发送syn=1告诉服务器需要联机，并发送一个随机生成的序列号seq=x
服务器接收到消息，回复客户端syn=1和一个确认码ack=x+1，并发送一个随机生成的序列号seq=y
客户端接受到消息，再次回复服务器一个确认码ack=y+1，并发送一个随机生成的序列号seq=z
四次挥手：由于数据传输是双向的，每个方向都必须单独进行关闭

客户端发送fin=1告诉服务器断开连接，并发送一个随机生成的序列号seq=x
服务器收到消息，回复客户端一个确认码ack=x+1，并发送一个随机生成的序列号seq=y，客户端进入等待状态
服务器发送fin=1告诉客户端断开连接，并发送一个随机生成的序列号seq=z
客户端收到消息，回复服务器一个确认码ack=z+1，并发送一个随机生成的序列号seq=r，断开连接

----
https://zhuanlan.zhihu.com/p/161560683
1、前言
当我们在浏览器栏输入：https://blog.csdn.net/dianxin113 的时候，具体发生了什么呢？这个请求是怎么到达服务器及返回结果的呢？

2、概述
浏览器进行DNS域名解析，得到对应的IP地址
根据这个IP，找到对应的服务器建立连接（三次握手）
建立TCP连接后发起HTTP请求（一个完整的http请求报文）
服务器响应HTTP请求，浏览器得到html代码（服务器如何响应）
浏览器解析html代码，并请求html代码中的资源（如js、css、图片等）
浏览器对页面进行渲染呈现给用户
服务器关闭TCP连接（四次挥手）

3、过程详解
一、DNS解析
首先会搜索浏览器自身的DNS缓存（缓存时间比较短，大概只有1分钟，且只能容纳1000条缓存）
如果浏览器自身的缓存里面没有找到，那么浏览器会搜索系统自身的DNS缓存
如果还没有找到，那么尝试从 hosts文件里面去找
在前面三个过程都没获取到的情况下，浏览器就会发起一个DNS的系统调用，就会向本地配置的首选DNS服务器（一般是电信运营商提供的，也可以使用像Google提供的DNS服务器）发起域名解析请求（通过的是UDP协议向DNS的53端口发起请求，这个请求是递归的请求，也就是运营商的DNS服务器必须得提供给我们该域名的IP地址）

二、TCP三次握手
三次握手完成之后这个TCP连接就进入Established状态，就可以发起http请求了。

【问题1】：TCP 为什么需要3次握手？

2个计算机通信是靠协议（目前流行的TCP/IP协议）来实现,如果2个计算机使用的协议不一样，那是不能进行通信的，所以这个3次握手就相当于试探一下对方是否遵循TCP/IP协议，协商完成后就可以进行通信了，当然这样理解不是那么准确。

【问题2】：为什么HTTP协议要基于TCP来实现？

目前在Internet中所有的传输都是通过TCP/IP进行的，HTTP协议作为TCP/IP模型中应用层的协议也不例外，TCP是一个端到端的可靠的面向连接的协议，所以HTTP基于传输层TCP协议不用担心数据的传输的各种问题。

三、 发起HTTP请求
HTTP是一个客户端和服务器端请求和应答的标准（TCP）。客户端是终端用户，服务器端是网站。通过使用Web浏览器、网络爬虫或者其它的工具，客户端发起一个到服务器上指定端口（默认端口为80）的HTTP请求。

通俗来讲，他就是计算机通过网络进行通信的规则，是一个基于请求与响应，无状态的，应用层的协议
一个HTTP请求报文由请求行（request line）、请求头部（header）、空行和请求数据4个部分组成
【问题3】：那什么是URL、URI、URN？

URI Uniform Resource Identifier 统一资源标识符
URL Uniform Resource Locator 统一资源定位符
URN Uniform Resource Name 统一资源名称

URL和URN 都属于 URI，为了方便就把URL和URI暂时都通指一个东西

四、服务器响应HTTP请求
接收到HTTP请求之后，就轮到负载均衡登场了，它位于网站的最前端，把短时间内较高的访问量分摊到不同机器上处理。负载均衡方案有软件、硬件两种。软件方案常见的就是NGINX了。

Nginx的作用主要有两个1，处理静态文件请求，2转发请求给后端服务器。然后后端服务器查询数据库返回数据。数据返回给客户端仍然通过HTTP协议传输。

HTTP响应报文主要由状态行、响应头部、空行以及响应数据组成。

五、浏览器解析
浏览器拿到index.html文件后，就开始解析其中的html代码，遇到js/css/image等静态资源时，就向服务器端去请求下载（会使用多线程下载，每个浏览器的线程数不一样），这个时候就用上keep-alive特性了，建立一次HTTP连接，可以请求多个资源，下载资源的顺序就是按照代码里的顺序，但是由于每个资源大小不一样，而浏览器又多线程请求请求资源，所以从下图看出，这里显示的顺序并不一定是代码里面的顺序。

浏览器在请求静态资源时（在未过期的情况下），向服务器端发起一个http请求（询问自从上一次修改时间到现在有没有对资源进行修改），如果服务器端返回304状态码（告诉浏览器服务器端没有修改），那么浏览器会直接读取本地的该资源的缓存文件。

六、浏览器进行页面渲染
最后，浏览器利用自己内部的工作机制，把请求的静态资源和html代码进行渲染，渲染之后呈现给用户，浏览器是一个边解析边渲染的过程。