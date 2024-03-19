# [gdb](https://www.sourceware.org/gdb)

## gdb common flags

使用 gdb 的窗口调试界面：shell 环境下执行 `gdb -tui <bin>`

- `--tui`: Use a terminal user interface.

执行 gdb 的脚本文件：可以在 gdb 环境下执行 `source <file>`，也可以在 shell 环境下执行 `gdb -x <file>`

- `--command=FILE`, `-x`: Execute GDB commands from FILE.

[调试技巧](https://ftp.gnu.org/old-gnu/Manuals/gdb/html_node/gdb_33.html)：

- 设置断点可以引入断点条件，如：`b func if i == 2`
- 可以 attach 到正在运行的程序进行调试（死锁场景）

## `.gdbinit`

```.gdbinit
# execute gdb script
source <file>

# setting layout for tui
layout src
layout asm

# setting lock for multi-thread process
set scheduler-locking on
```

## Remote debugging with GDB (gdbserver)

client: `(gdb) target remote <target>`

- `<target>`：可以是 serial-device (tty)、socket、端口
- 通过串口甚至可以调试 docker 内的程序

server: `gdbserver` 可以对外暴露 gdb 调试端口

例子：ssh 场景下使用 gdb 调试程序

```gdb
target remote | ssh -T hostname gdbserver - hello
```

## References

- [edit-mode](https://unix.stackexchange.com/questions/748258/quick-swap-tui-mode-while-using-vi-keybindings-in-gdb)
- [gdb - connect](https://sourceware.org/gdb/current/onlinedocs/gdb.html/Connecting.html)
- [gdbserver](https://sourceware.org/gdb/current/onlinedocs/gdb.html/Server.html)
- [Remote debugging with GDB](https://developers.redhat.com/blog/2015/04/28/remote-debugging-with-gdb)
