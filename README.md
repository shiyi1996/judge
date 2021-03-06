# 基于docker虚拟化的云评测系统
## 功能
根据用户提交的代码和测试样例，进行评测，根据运行结果、时空消耗来评判是否正确。

## 难点
##### 安全性
评测系统的功能，相当于将**代码上传权限、代码执行权限**开放给用户，那么，安全是绕不过的一个问题，也是本项目最重要的一个问题。
1. 编译期
很多恶意的代码，都可以在编译器间干掉我们的系统，例如：linux上 `#include "/dev/random"`，系统会直接卡死。
还有用户写特殊代码让编译过程延长，占用系统资源等等，举不胜举，都充满着风险。
2. 运行期
运行期间的问题就更多了，最容易想到的就是`fork炸弹`了，还有，使用一些危险的系统调用，或是恶意删改系统服务器的文件，等等。

##### 伸缩性
评测系统经常遇到一些高峰期，如，发布一个比赛，那么在比赛期间，评测系统的压力远大于正常期间。因此，一个完善的评测系统需要在架构上，具有良好的伸缩性，能够方便的针对业务需求进行机器扩容操作。

##### 扩展性
互联网发展迅速，各种语言层出不穷，各种判题模式也五花八门，评测系统要在架构上，满足`语言支持`、`判题模式`等的方便扩展。


#### 技术
Go并发编程
Docker容器技术
消息队列
Redis
Mysql

