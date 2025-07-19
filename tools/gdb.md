# gdb

使用 gdb 的窗口调试界面：shell 环境下执行 `gdb -tui <bin>`

执行 gdb 的脚本文件：可以在 gdb 环境下执行 `source <file>`，也可以在 shell 环境下执行 `gdb -x <file>`

[调试技巧](https://ftp.gnu.org/old-gnu/Manuals/gdb/html_node/gdb_33.html)：设置断点可以引入断点条件，如：`b func if i == 2`

gdb 也可以调试正在运行的程序（死锁场景）

## Remote debugging with GDB (gdbserver)

client: `(gdb) target remote <target>`

- `<target>`：可以是 serial-device (tty)、socket、端口
- 通过串口甚至可以调试 docker 内的程序

server: `gdbserver` 可以对外暴露 gdb 调试端口

例子：ssh 场景下使用 gdb 调试程序

```gdb
target remote | ssh -T hostname gdbserver - hello
```

## ref

- [gdb - connect](https://sourceware.org/gdb/current/onlinedocs/gdb.html/Connecting.html)
- [gdbserver](https://sourceware.org/gdb/current/onlinedocs/gdb.html/Server.html)
- [Remote debugging with GDB](https://developers.redhat.com/blog/2015/04/28/remote-debugging-with-gdb)
