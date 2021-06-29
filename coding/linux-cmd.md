# Linux 常用软件的命令

## bash

```term
code // 打开 vscode
google-chrome // 打开 chrome 浏览器
chromium // 打开 chromium 浏览器
netease-cloud-music // 打开网易云音乐

ln -s src_addr link_addr // 创建文件夹快捷方式
showkey -a // 输出按键的 ascii 码
nautilus // 打开文件浏览器
xrandr -o <left/right/inverted/normal> // 旋转屏幕

// 通过修改设备文件，控制硬件
echo 593 | sudo tee /sys/class/backlight/intel_backlight/brightness

// 不挂起后台执行指令，并且将标准输出和错误输出进行重定向
nohup command > myout.file 2>&1 &

// 系统设置：swap 分区和防火墙设置
swapoff -a // 关闭 swap 分区
swapon -a // 开启 swap 分区
ufw status // 查看防火墙信息

// apt 包管理
apt-mark <auto/manual> <pkgname>

// 将terminal显示的内容保存到txt
$ script -a <filename>
$ exit

// 用户操作
adduser <username> // 创建用户
deluser --remove-all-files <username> // 删除用户
su <username> // 切换用户

// 某用户对一个应用添加sudo权限
sudo groupadd <group-name>
sudo chgrp <group-name> /usr/bin/<app-name>
sudo chmod 4755 /usr/bin/<app-name>
// -rwxr-xr-x -> -rwsr-xr-x
sudo gpasswd -a <user> <group-name>

echo !# // 获得上一个命令名
echo !$ // 获得上一个指令的最后一个参数
echo !:n // 获得上一个指令的第 n 个参数
!! // 执行上一条指令

whatis // 查看命令用途
// 设置默认的编辑器
sudo update-alternatives --config editor
iwconfig // 配置无线网络接口
ifconfig // 配置网络接口
```

## gdb

```gdb
// 执行 gdb 的脚本文件
// 亦可在 shell 环境下使用 -x <file> 参数
source <file>
```

## curl

```shell
# -d 发送 POST 请求的数据体，-X 指定方法
curl -d 'key1=value1&key2=value2' -X Post
```
