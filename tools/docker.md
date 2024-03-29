# Docker

Docker 主要解决环境配置问题，它是一种虚拟化技术，对进程进行隔离，被隔离的进程**独立于宿主操作系统**和其它隔离的进程。docker 基于 go 进行开发。

与虚拟机的区别：虚拟机也是一种虚拟化技术。通过模拟硬件，并在硬件上安装操作系统来实现。

![docker-vm](imgs/docker-vm.png)

Docker 把应用程序及其依赖，打包在 image 文件里面。只有通过这个文件，才能生成 Docker 容器。image 文件可以看作是容器的模板。Docker 根据 image 文件生成容器的实例。同一个 image 文件，可以生成多个同时运行的容器实例。

image 文件生成的容器实例，本身也是一个文件，称为容器文件。也就是说，一旦容器生成，就会同时存在两个文件： image 文件和容器文件。而且关闭容器并不会删除容器文件，只是容器停止运行而已。

创建 docker 涉及到的文件：`.dockerignore`、`Dockerfile`

## dockerfile & docker-compose

docker 中不同的指令：

- FROM：定制镜像的基镜像
- RUN：构建定制镜像所提前运行的指令
- CMD：镜像运行时执行的指令

docker-compose 后台执行 bash 需要加入 `tty` 和 `stdin_true` 字段

## 镜像的操作方法

`ctrl P + ctrl Q` detach 容器而不删除容器

## docker 构建镜像的优化方式

[docker 的多阶段构建](https://blog.csdn.net/Michaelwubo/article/details/91872076)

- copy --from 甚至可以从其他镜像复制文件到构建的镜像当中

[构建 alpine 镜像以及相应的优化 (go, java, python, rust)](https://cloud.tencent.com/developer/article/1632733)
