# [gdb](https://www.sourceware.org/gdb)

## gdb common flags

- `--tui`: Use a terminal user interface.
- `--command=FILE`, `-x`: Execute GDB commands from FILE.

## `.gdbinit` related

```.gdbinit
# execute gdb script
source <file>

# setting layout for tui
layout src
layout asm

# setting lock for multi-thread process
set scheduler-locking on
```

## other

- [调试技巧](https://blog.csdn.net/robinblog/article/details/17652541)：设置断点可以引入断点条件，如：`b fun if i == 2`
- [edit-mode](https://unix.stackexchange.com/questions/748258/quick-swap-tui-mode-while-using-vi-keybindings-in-gdb)
