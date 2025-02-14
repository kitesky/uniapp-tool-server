<div align="center">

```shell
如果对您有帮助，请给一个Star！
```

</div>

### 项目介绍

该项目是AI写作小程序服务端。小程序端地址：https://gitee.com/kitesky/uniapp-tool

### 技术栈

- 小程序采用Uniapp开发，是一个使用 Vue.js 开发所有前端应用的框架，开发者编写一套代码，可发布到iOS、Android、Web（响应式）、以及各种小程序（微信/支付宝/百度/头条/飞书/QQ/快手/钉钉/淘宝）、快应用等多个平台。

- 服务端采用Go语言开发，主要技术栈：Gin、GORM、Asynq、robfig/corn、Redis

### 在线体验

微信小程序(请使用微信扫一扫)

![美智合AI](https://api.idcd.com/assets/example/10001307.png)


### 目录结构

![目录结构](https://foruda.gitee.com/images/1739524221063420006/63154032_82149.png "目录结构")

- .air 热重载配置文件
- boot 系统初始化，加载配置、MySQL、Redis连接
- cmd  服务启动命令
- config 配置文件
- controllers 控制器
- crontab cron定时器
- dao 数据操作层
- jobs 队列任务
- middlewares 路由中间件
- models 模型层
- pkg 自定义包
- public 公共目录
- routes 路由
- services 服务层
- storege 数据存储目录， 上传文件、日志log...
- types 类型， 常量、结构体定义在这里
- utils 封装的工具函数
- test.sql 数据库文件

### 服务启动

1. 启动HTTP服务

```
$ air
```

2. 启动队列服务

```
目录/main asynq
```

2. 启动Cron定时任务

```
目录/main cron
```

### 使用说明

1、本项目为个人兴趣爱好而开发，开源并且提供完整代码，不保证代码中无缺陷，也不保证能够及时修复BUG。正式用于运营需谨慎使用，对于未知缺陷造成的损失，不负任何责任，推荐用于学习和测试。

2、文档不完善，但代码结构还是比较清晰，作为全栈项目，不适合新手小白。部署运行有一定的难度，不提供任何无偿咨询和指导。

3、**提问题、请求协助，请先打赏，金额随意。** 


![打赏](https://api.idcd.com/assets/example/fengzheng.jpg)

### 作者信息

开发者：风筝

微信号: kite365 （添加好友，注明来意）

邮箱：kite365@gmail.com





