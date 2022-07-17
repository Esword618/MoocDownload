# MoocDownload

## 声明

本项目`仅仅用于自己购买课程后`，想保存下来后面观看。不可用于下载售卖。如果侵犯慕课权益，售卖者将自己承担所有，与本程序无关。如果被发现在售卖，将考虑某些措施限制下载。

望大家遵守规则。

`同时现在本项目也支持公开课的下载，包括视频以及课件。`

欢迎大家关注公众号，有什么问题可以直接跟我反馈，我会更新。记得点一个小星星。

<img src="https://cdn.jsdelivr.net/gh/Esword56/blogImg@main/vx/902a72e0543e80d43177a2b7ddc3806.7absp7tl8d00.png" alt="902a72e0543e80d43177a2b7ddc3806" style="zoom: 80%;" />

交流群

<img src="https://cdn.jsdelivr.net/gh/Esword56/blogImg@main/vx/微信图片_20220404193340.7jarc0vlns40.jpg" alt="微信图片_20220404193340" style="zoom: 50%;" />

## 教程

如果在你的电脑上运行有问题，请不要慌，可以通过我公众号加我微信，与我联系进行交流，提供一流的免费售后服务，前提是我有空。

如果你是编程高手，建议你在自己的电脑上编译，这样适配性更好。

## 唠嗑

本项目是本人在之前用`python`构造的基础上用`go`语言搭建，与之前的项目，在下载速度与稳定性来说都提升了好几个档次。这个是本人第一次完完整整用`go`搭建的第一个项目，有许多不足之处望担待，咱们不喜就喷。哈哈哈哈

## 更新内容

- 修复协程下载的不稳定
- 修复部分视频不能下载问题
- 更新了项目的结构，更加结构化，为一会加入其它平台做准备
- 不再支持登录（直接去浏览器复制cookie，简单的不要不要的！）
- 加入检测更新，再有新的版本会在使用的时候通知（其实很简单）

## ToDo

| 1     | ~~mp4 or flv 文件的分段协程下载~~ |
| ----- | ------------------------ |
| **2** | **可选择性下载文件**             |
| **3** | **其它平台的也可以下载 如imooc等平台** |
| **4** | **~~解密文件本地运行~~**         |
| **5** | **.....还不知道....**        |

## 依赖项目

下一步一类项目(做可选择性下载)
https://github.com/rivo/tview

## 参考文章

[mp4/flv文件分片下载](https://polarisxu.studygolang.com/posts/go/action/build-a-concurrent-file-downloader/)

./download/课程名/单元名/data.pdf或视频.mp4

## 感谢

[goja](https://github.com/dop251/goja)