# Linux

- [服务器管理](#服务器管理)
- [常用软件的命令](#常用软件的命令)
  - [update-alternatives](#update-alternatives)
  - [bash](#bash)
  - [gdb](#gdb)
  - [curl](#curl)
  - [ssh](#ssh)
  - [vim](#vim)
  - [apt](#apt)

## 服务器管理

- 部署 quark 时如果使用的是 proxmox 需要开启嵌套虚拟化的支持 ([ref](https://zhuanlan.zhihu.com/p/593472919))
- proxmox 虚拟机的扩容：[逻辑卷组的操作](https://einverne.github.io/post/2020/11/extend-proxmox-system-partition-and-pve-file-system.html)、[已有分区的扩展](https://cloud.tencent.com/developer/beta/article/1671893)、[扩容文件系统](https://blog.csdn.net/youjin/article/details/79137203)
- jetson 创建容器所使用的[镜像文档](https://catalog.ngc.nvidia.com/orgs/nvidia/containers/l4t-pytorch)、查看 L4T 使用的指令 `jtop`
- linux 开机自启动：可以利用 crontab 中的 `@reboot`

## 常用软件的命令

### update-alternatives

ubuntu 可以通过 `update-alternatives` 来管理软件的版本，如：通过以下指令可以指定 cuda 版本

```bash
sudo update-alternatives --config cuda
```

### bash

```bash
code # 打开 vscode
google-chrome # 打开 chrome 浏览器
chromium # 打开 chromium 浏览器
netease-cloud-music # 打开网易云音乐

ln -s src_addr link_addr # 创建文件夹快捷方式
xev # 查看按键的信息
showkey -a # 输出按键的 ascii 码
nautilus # 打开文件浏览器
xrandr -o <left/right/inverted/normal> # 旋转屏幕

# 通过修改设备文件，控制硬件
echo 593 | sudo tee /sys/class/backlight/intel_backlight/brightness

# 不挂起后台执行指令，并且将标准输出和错误输出进行重定向
nohup command > myout.file 2>&1 &

# 系统设置：swap 分区和防火墙设置
swapoff -a # 关闭 swap 分区
swapon -a # 开启 swap 分区
ufw status # 查看防火墙信息

# apt 包管理
apt-mark <auto/manual> <pkgname>

# 将terminal显示的内容保存到txt
$ script -a <filename>
$ exit

# 用户操作
adduser <username> # 创建用户
deluser --remove-all-files <username> # 删除用户
su <username> # 切换用户

# 某用户对一个应用添加sudo权限
sudo groupadd <group-name>
sudo chgrp <group-name> /usr/bin/<app-name>
sudo chmod 4755 /usr/bin/<app-name>
# -rwxr-xr-x -> -rwsr-xr-x
sudo gpasswd -a <user> <group-name>

echo !# # 获得上一个命令名
echo !$ # 获得上一个指令的最后一个参数
echo !:n # 获得上一个指令的第 n 个参数
!! # 执行上一条指令

whatis # 查看命令用途
sudo update-alternatives --config editor # 设置默认的编辑器
iwconfig # 配置无线网络接口
ifconfig # 配置网络接口

ulimit # 配置 shell 使用的资源（可以完成 coredump 等配置）

stat <file> # 查看文件的 inode 信息
df # 查看磁盘的的使用情况 (disk free)
ls -i <file> # 查看文件的 inode 编号
dumpe2fs <device> # dump file extx filesystem information
```

记录 sudoer 信息的文件：`/etc/sudoers`

### gdb

执行 gdb 的脚本文件：可以在 gdb 环境下执行 `source <file>`，也可以在 shell 环境下执行 `gdb -x <file>`

使用 gdb 的窗口调试界面：shell 环境下执行 `gdb -tui <bin>`

[调试技巧](https://blog.csdn.net/robinblog/article/details/17652541)：设置断点可以引入断点条件，如：`b fun if i == 2`

### curl

```shell
# -d 发送 POST 请求的数据体，-X 指定方法
curl url -d 'key1=value1&key2=value2' -X Post
# -b 设置 cookie 信息
curl url -b 'key=val'
# -G 指定 get 方法，--data-urlencode 使用 url 传递参数
curl -G url --data-urlencode 'key=val'
curl --resolve addr:ip:port url # 模拟 DNS 服务器将 url 中的地址解析为 ip:port
```

### ssh

远程连接：`ssh user@address`

- 缺省 `user@` 的情况下，默认使用和本地相同用户名作为登陆的用户
- `ssh-copy-id user@address` 设置免密登陆
- `-X` 参数可以进行图形的转发
- ssh 可以直接执行指令
- 在 `~/.ssh/config` 中可以配置 ssh 的别名

    ```sshconfig
    Host <alias-name>
        HostName <host-name>
        User <user>
    ```

### vim

vim 中使用 `verbose set` 可以查看具体设定，如：

- `verbose set` 查看文件类型

### apt

apt 卸载本地的 deb 包，需要通过 `dpkg -l` 查询安装的包名 ([ref](https://devimalplanet.com/how-to-uninstall-deb-packages-linux))
