# 福州大学202401软件工程实践FuliFuli组项目(后端)

## 介绍

这个项目中，**除了测试部分**以外的均由我完成，并作为福州大学202401软件工程实践团队项目的一部分。我所在组的项目是构建一个视频创作和分享平台，名为FuliFuli。

相比于之前的视频网站后端项目，这次实践出来的代码更规整一些。

~~虽然后面有点混乱~~

本项目是单体式架构而不是微服务架构。主要是因为服务器的配置不够——一台2GB的2核服务器，显然不够运行微服务架构下的多个数据库实例。

毕竟按理来说，多个主要模块的数据库应该分离，以加速数据操作，因此理应做到每一个大模块一个MySQL实例，也就会出现服务器数量不够或内存不足的情况。

即便是单体式，这个2GB服务器也崩过几次，很多情况下还仅仅只是运行除了后端程序以外的容器镜像。

你可以看到Docker Compose配置文件中的日志服务用的是Zincsearch而不是Elasticsearch，这就是退而求其次的结果（当然还有用程序内的简单代码替代消息队列中间件的操作~~毕竟是单体式架构~~）。

## 运行

推荐在Unix/Linux平台上运行，最好使用Docker，否则需自行安装运行环境。

在运行之前，需要完成[配置文件](config.yaml)的修改：
1. 应修改邮箱地址，否则无法发送邮箱验证码（以下为例子）：

``` yaml
Email:
  address: "smtp.mxhichina.com"
  port: 465
  username: "fulifuli@sophisms.cn"
  password: "_222200316Cyk"
  conn_pool_size: 4
```

2. 应修改OSS配置信息，否则无法对音视频资源进行操作（本项目采用七牛云OSS）：

``` yaml
OSS:
  bucket: "your_bucket"
  access_key: "your_access_key"
  secret_key: "your_secret_key"
  domain: "http://your.cdn.domain"
  upload_url: "https://up.xxxx.xxx"
  callback_url: "https://your.domain.com/api/v1/oss/callback"
```

运行此项目时，对于Docker而言：
1. 需要先构建Docker镜像：

```bash
make docker_build
```

2. 然后运行Docker Compose

```bash
make docker_run
```