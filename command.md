# Linux 常用软件的命令

## bash

```shell
$ code // 打开vscode
$ google-chrome // 打开chrome浏览器
$ netease-cloud-music // 打开网易云音乐
$ ln -s src_addr link_addr // 创建文件夹快捷方式
$ showkey -a // 输出按键的 ascii 码
$ nautilus // 打开文件浏览器
$ xrandr -o <left/right/inverted/normal> // 旋转屏幕

// 将terminal显示的内容保存到txt
$ script -a <filename>
$ exit

// 某用户对一个应用添加sudo权限
sudo groupadd <group-name>
sudo chgrp <group-name> /usr/bin/<app-name>
sudo chmod 4755 /usr/bin/<app-name>
// -rwxr-xr-x -> -rwsr-xr-x
sudo gpasswd -a <user> <group-name>
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
