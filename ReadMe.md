# StylishPubServer

一个简单的可配置的本地服务器
开发的目的是为了使Google Chrome浏览器的插件`Stylish`可以使用本地的图片等其它资源

## 使用方式
+ 使用源码构建
+ 使用编译好了的应用程序

### 使用源码构建
#### 安装依赖
```
go get gopkg.in/yaml.v2
```
#### 编译
```
git clone https://github.com/0lightbee0/StylishResourceServer
cd StylishResourceServer
go build ./src
```

### 使用编译好了的应用程序
直接下载`StylishPubResource.exe`就可以了

## 配置文件
需要在`StylishPubResource.exe`的同级文件夹下建立`config`文件夹，所有的配置文件存放在这里
+ dir.yaml 指定资源文件路径(`.`表示在当前文件夹下)
+ resType.yaml 指定文件类型与对应的文件夹
+ server.yaml 指定服务的端口和地址等等(不填写地址意味着将在本地开启一个服务)

### 文件夹指定
该文件夹下需要有以下文件夹:
+ Actions
+ Styles
+ Images
+ Fonts
+ Pages

## log
放在`StylishPubResource.exe`的同级文件夹`log`下的`main.log`文件