# Makefile

[csdn-命令前提条件](https://blog.csdn.net/Decisiveness/article/details/52199708?utm_source=blogkpcl12)：

> Type II： 命令前提目标（order-only Prerequisites）
> 偶尔的，我们会遇到这样的情况：我们需要执行某个或某些规则，但不能引起生成目标被重新生成。
> 此时你就需要使用命令前提目标。命令前提目标由一个管道符号即竖线“|”指示，位于前提目标列表中。
> 竖线左边的目标就是正常前提目标，竖线右边的目标就是命令前提目标，形式如下：
>
> [java] view plain copy print?
> target : normal-prerequisites | order-only-prerequisites
>
> 竖线左边的正常前提目标列表可以是空。
>
> 注意：如果前提目标中同时存在正常前提目标和命令前提目标，则正常前提目标优先生成。

makefile 空格的替换：常见于将 makefile 的变量列表转换为系统变量

```makefile
nullstring :=
space := $(nullstring) #end
delim := :
env := $(subst $(space),$(delim),$(vars))
```

makefile 默认使用 `/bin/sh` 作为执行指令的 shell

- 可以通过修改 `SHELL` 变量来修改默认 shell，如 `make all SHELL=$(which zsh)`
- 上述执行指令不会引入用户空间的环境变量，如果需要在用户空间下执行指令可以在 `makefile` 中使用 `$(shell <cmd>)`
- 或者通过 `$$<VAR>` 查看用户空间下的环境变量
