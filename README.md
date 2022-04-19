<p align="center">
  <a href="https://ishkong.github.io/go-cqhttp-docs/"><img src="https://s2.loli.net/2022/04/19/VxkCUtHl5Sa2zgb.png" width="200" height="200" alt="go-cqhttp"></a>
</p>

<div align="center">

# 澪酱-你的生活小助手

_✨ 基于 [MiraiGo](https://github.com/Mrs4s/MiraiGo)，使用[MiraiGo-Template](https://github.com/Logiase/MiraiGo-Template)模块化开发 ✨_

</div>

<p align="center">
  <a href="#">
    <img src="https://img.shields.io/badge/go-v1.17.8-blue" alt="license">
  </a>
  <a href="#">
    <img src="https://img.shields.io/badge/release-v0.0.1-orange" alt="release">
  </a>
  <a href="#">
    <img src="https://img.shields.io/badge/LICENSE-AGPL--3.0-red" alt="action">
  </a>
</p>

---
MiraiGo-MioBot是对MiraiGo-Template的模块化编程的一次拓展实践，项目拥有**部署简单**，**拓展性强**等特点

## 项目部署

请在[Release](https://github.com/t0hka1/MiraiGo-MioBot/releases/tag/v0.0.1)处自行下载对应平台的二进制文件

并在二进制文件目录下新建如下配置文件(后续版本考虑将配置文件的生成一起打包)

**application.yaml** -登录

```yaml
bot:
  loginmethod: qrcode
  account: ""
  password: ""
```

**homework.yaml** - 作业记录

```yaml
高数 : '学习通第八周作业(04-20 22:30)'
离散 : 'Dont care!'
英语听说 : '配音(04-17 12:00)'
英语精读 : '1.平台4单元作业
          2.批改网作文
          3.我爱记单词(04-17 12:00)'
大物 : '学习通 电场与高斯定理(05-05 16:48)'
数据结构 : '学习通 树与二叉树 (04-18 23:55)'
工程经济学 : 'None'
数字电路 : 'None'
```

**device.json** - 设备信息

```json
{"display":"MIRAI.156311.001","product":"mirai","device":"mirai","board":"mirai","model":"mirai","finger_print":"mamoe/mirai/mirai:10/MIRAI.200122.001/6611562:user/release-keys","boot_id":"11afc12f-e067-95ab-03d0-7a59f3c25218","proc_version":"Linux version 3.0.31-IAD5IDEb (android-build@xxx.xxx.xxx.xxx.com)","protocol":5,"imei":"250346797735787","brand":"mamoe","bootloader":"unknown","base_band":"","version":{"incremental":"5891938","release":"10","codename":"REL","sdk":29},"sim_info":"T-Mobile","os_type":"android","mac_address":"00:50:56:C0:00:08","ip_address":[10,0,1,3],"wifi_bssid":"00:50:56:C0:00:08","wifi_ssid":"\u003cunknown ssid\u003e","imsi_md5":"19773e595e92e49b21707c08cfa4e565","android_id":"65684d724991464f","apn":"wifi","vendor_name":"MIUI","vendor_os_name":"mirai"}
```

运行二进制文件即可

## 现有功能

-  支持关键字学习功能
-  支持对作业的简易通知（定时模块未完善）

## 现有指令说明

```go
//关键字设置
learn `keyword` `value`
```

```go
//关键字回复
keyword
```

## 模块拓展开发

module参考[log.go](./modules/logging/log.go)

```go
package mymodule

import (
    "aaa"
    "bbb"
    "MiraiGo-Template/bot"
)

var instance *Logging

func init() {
	instance = &Logging{}
	bot.RegisterModule(instance)
}

type Logging struct {
}

// ...
```

编写自己的Module后在[app.go](./app.go)中启用Module 

```go
package main

import (
    // ...
    
    _ "modules/mymodule"
)

// ...
```


## 未来将会加上的功能

-  作业通知与超星泛雅对接
-  支持上课啦签到的功能
-  迁移杭电助手现有服务的推送
-  自动健康打卡(存在风控)
-  我爱记单词自动答题
-  座位预约功能

## 后续将支持docker
