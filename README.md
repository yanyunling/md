# 云文档

## 介绍
一个简单的markdown文档服务器

## 软件架构
- go1.20
- sqlite
- vue3

## 编译步骤
- 安装go环境
- 安装nodejs环境
- 进入目录，添加权限并执行 **build.bat** 或 **build.sh**
- 编译后的可执行文件在 **./md** 目录下

## 命令行参数
- `-p`：监听端口号。默认值：**9900**
- `-log`：日志目录，存放近30天的日志。默认值：**./logs**
- `-data`：数据目录，存放数据库文件和图片。默认值：**./data**
- `-reg`：是否允许注册（即使禁止注册，在没有任何用户的情况时仍可注册）。默认值：**true**

## docker镜像
- [https://hub.docker.com/r/streamerzero/md](https://hub.docker.com/r/streamerzero/md)

## 隐含的功能按钮
- 编辑模式：“云文档”标题
- 收缩边栏：“云文档”标题左边图标
- OpenApi保存：Ctrl + s
- 发布链接：文档列表蓝色倒三角