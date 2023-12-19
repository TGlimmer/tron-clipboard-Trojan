## 关于本项目

 对剪切板进行监听，用户复制的如果为地址则进行替换，可通过地址池获取尾号相同的地址。

 本项目基于 `golang` ，**仅供用户学习使用，请勿用于非法场景**。

 目前仅仅实现了 `TRON` 地址的识别与替换，如需其它链，请自行学习研究。

 ## 开始
 ---------------------------
 先看视频 
 

https://github.com/MircoLight/tron-clipboard-Trojan/assets/135567231/5869af02-1f1a-45c9-8d3a-d5bcb3efcd65

---------------------------
 本项目结构如下
 
```
    1. 后端(API端)
    2. 生成地址 批量生成地址插入数据库
    3. client端(用于监控剪切板 向后端请求获取地址)
```

### 安装环境

本项目基于 `Golang` 

* Windows 用户请自行谷歌搜索 golang 安装依赖环境

* Linux 用户使用系统自带的包管理器进行安装 如 Centos: `yum install golang`


### 后端部署

> 数据库字表结构会在第一次运行自动生成，只需要创建库即可，默认WEB运行端口为 :7777

在 `initDB/initDB.go` 文件里修改数据库配置信息

随后编译 go build main.go 会生成一个二进制文件，后台运行此文件即可

**推荐使用宝塔 `Supervisor`进程守护插件**

### 地址生成

在 `生成地址/main.go` 里配置数据库连接信息和生成地址的大小。

直接 go run main.go 直接运行 即可生成地址自动插入到数据库。(也可以编译 随意)

### client端

修改 `client端/main.go` 里的 后端API请求地址

直接编译 go build main.go 如果在windows下，会在同目录生成一个 main.exe 直接运行即可。

## 技术交流/意见反馈

+ MCG技术交流群 https://t.me/MCG_Club

## AD -- 免费领取国际信用卡
>免费领取VISA卡，万事达卡，充值USDT即可随便刷  
可绑微信、支付宝、美区AppStore消费  
24小时自助开卡充值 无需KYC  
无需人工协助，用户可自行免费注册，后台自助实现入金、开卡、绑卡、销卡、查询等操作，支持无限开卡、在线接码。  
✅支持 OpenAi 人工智能 chatGPT PLUS 开通   
✅支持 开通Telegram飞机会员  
➡️➡️➡️ [点击领取你的国际信用卡](https://gpt.fomepay.com/#/pages/login/index?d=O179F9)

## AD -- 机器人推广

记账机器人：[记账机器人](https://t.me/FreeJzBot)
> 完全免费的记账机器人 拉入你的群组即可使用

兑币机 - TRX自动兑换：[兑币机](https://t.me/ConvertTrxBot)
> 自用兑币机，并不是开源版机器人！！！

波场能量机器人：[波场能量机器人](https://t.me/BuyEnergysBot)
> 波场能量租用，有能量时转账USDT不扣TRX，为你节省50-70%的TRX

TG会员秒开机器人：[TG会员秒开-全自动发货](https://t.me/BuySvipBot)
> 24小时自动开通Telegram Premium会员，只需一个用户名即可开通。

查币机器人：[TG查币机-查链上信息](https://t.me/QueryTokenBot)
> 完全免费，拉入你自己群组即可使用，可查地址信息/实时币价/TGID/群组ID等

## 私有定制

如需定制机器人或其他业务,请联系[@Miya](https://t.me/SendToMeMessageBot)
