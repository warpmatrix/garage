<!-- omit in toc -->
# c 的一些冷门易忘的知识点

<!-- omit in toc -->
## Table of Contents

- [语言机制](#语言机制)
  - [`scanf` 读到回车结束输入](#scanf-读到回车结束输入)
  - [c的字符串操作](#c的字符串操作)
  - [c中的布尔类型 (C99)](#c中的布尔类型-c99)
  - ['\\r'的作用](#r的作用)
  - [表示浮点数的方法](#表示浮点数的方法)
  - [关于序列点](#关于序列点)
  - [函数原型和函数签名](#函数原型和函数签名)
  - [数组的初始化](#数组的初始化)
  - [数组与指针](#数组与指针)
  - [c 的多维变长数组 (C99)](#c-的多维变长数组-c99)
  - [关于 typedef](#关于-typedef)
  - [复合字面量](#复合字面量)
  - [关闭缓冲区和取消回显](#关闭缓冲区和取消回显)
  - [调用子程序](#调用子程序)
  - [定义二进制数](#定义二进制数)
  - [位字段](#位字段)
  - [linux-fork 函数](#linux-fork-函数)
  - [`malloc` 与 `free` 的行为定义](#malloc-与-free-的行为定义)
  - [全局变量初始化](#全局变量初始化)
  - [预定义](#预定义)
  - [隐式转换优先级](#隐式转换优先级)
  - [`while` 的两种汇编方式](#while-的两种汇编方式)
  - [浮点数内存布局和表示范围的探索](#浮点数内存布局和表示范围的探索)
  - [获得系统变量](#获得系统变量)
- [linux 与 c](#linux-与-c)
  - [编译运行过程](#编译运行过程)
  - [链接器的使用](#链接器的使用)
  - [跨架构运行二进制代码](#跨架构运行二进制代码)
  - [系统调用](#系统调用)
  - [`glibc` 的调试](#glibc-的调试)
  - [查看库版本](#查看库版本)
  - [获取符号对应的函数](#获取符号对应的函数)
- [MPI](#mpi)

## 语言机制

### `scanf` 读到回车结束输入

```c
#include <stdio.h>
#define MAXLEN 100

int main(int argc, char const *argv[]) {
    char str[MAXLEN];
    scanf("%[^\n]%*c", str);
    printf("%s\n", str);
    return 0;
}
```

### c的字符串操作

c中大多数自带的字符串操作都自带 `'\0'`。

例外：strncpy（若没有复制到 ```'\0```，则不会带上空字符）

### c中的布尔类型 (C99)

```c
#include <stdio.h>

int main(int argc, char const *argv[]) {
    _Bool tf = 2;
    printf("%d", tf);
    return 0;
}
```

在里`<stdbool.h>`，将 `bool` 作为 `_Bool` 的别名，并且可以使用 `true` 和 `false`。

```c
#include <stdbool.h>

int main(int argc, char const *argv[]) {
    bool tf = true;
    return 0;
}
```

### '\r'的作用

```c
#include <stdio.h>

int main(int argc, char const *argv[]) {
    printf("1234567890\r");
    printf("abc\n");
    return 0;
}
```

### 表示浮点数的方法

1. 固定小数点形式：e.g. ```123.4```
2. 指数形式：e.g. ```1.234e2```
3. 十六进制数和2的幂形式：e.g. ```0x1.ap1```

### 关于序列点

用于区分语句，所有的副作用都在下一个序列点之前发生。

**四种序列点：**

1. ```;```
2. ```,```
3. ```&&```
4. ```||```

### 函数原型和函数签名

- 函数原型可在函数内声明，但定义一定要在函数外。
- 函数签名指函数的参数和返回值。
- c 中不允许函数名相同、参数列表不同的函数。

```c
#include <stdio.h>

int main(int argc, char const *argv[]) {
    void fun();
    fun();
    return 0;
}
void fun() {
    printf("hahaha\n");
}
```

### 数组的初始化

若不对数组进行初始化，则数组内存储 junk value；若进行初始化（可以部分初始化），则剩下未初始化的元素被初始化为 0。

```c
#include <stdio.h>

int main(int argc, char const *argv[]) {
    const int arr[4] = {1};
    for (size_t i = 0; i < 4; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
    return 0;
}
```

特别地，可以使用 ```int arr[SIZE] = {};``` 进行初始化为0。

C99 中可以使用指定初始化器，指定初始化某些元素。剩下未初始化的元素同样被初始化为 0，且多次初始化同一元素，使用最后初始化的值。

```c
#include <stdio.h>

int main(int argc, char const *argv[]) {
    const int arr[] = {1, 3, [4] = 5, 6, [1] = 2};
    for (size_t i = 0; i < sizeof arr / sizeof arr[0]; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
    return 0;
}
```

另，**变长数组不能被初始化**。

### 数组与指针

指针与数组是不同的东西。例，```sizeof ptr``` 不能读取数组的大小，数组的地址没有 ```++``` 运算符。

```c
#include <stdio.h>
#define SIZE 100

int main(int argc, char const *argv[]) {
    const int arr[SIZE] = {};
    const int *ptr = arr;
    printf("size: %ld\n", sizeof ptr / sizeof ptr[0]);
    return 0;
}
```

且向函数传递参数时，```int arr[]``` 与 ```int *arr``` 等价，即 ```arr``` 已退化为指针。

```c
#include <stdio.h>
#define SIZE 100

void fun(const int arr[]) {
    printf("size: %ld\n", sizeof arr / sizeof arr[0]);
}

int main(int argc, char const *argv[]) {
    const int arr[SIZE] = {};
    fun(arr);
    return 0;
}
```

另外，可以在 ```for``` 循环中使用头尾指针遍历数组。

### c 的多维变长数组 (C99)

c 在函数传参的过程中，可以直接传递多维变长数组。这一点与 c++ 不同。函数声明可以用 * 代替数组的长度变量。同样要强调，传递参数后数组的首地址退化为指针。

```c
#include <stdio.h>

void fun(int , int , int ptr[][*]);

int main(int argc, char const *argv[]) {
    const int rows = 2, cols = 3;
    int arr[rows][cols];
    for (size_t i = 0; i < rows * cols; i++) (*arr)[i] = i;
    fun(rows, cols, arr);
    return 0;
}

void fun(int rows, int cols, int ptr[][cols]) {
    for (size_t i = 0; i < rows * cols; i++) {
        printf("%d ", (*ptr)[i]);
    }
    printf("\n");
}
```

### 关于 typedef

```c
#include <stdio.h>
typedef int arr4[4];
typedef arr4 arr3x4[3];

int main(int argc, char const *argv[]) {
    arr3x4 mat;
    for (size_t i = 0; i < 3 * 4; i++) {
        (*mat)[i] = i;
    }
    return 0;
}
```

### 复合字面量

复合字面量可以创建一个匿名数组，供临时使用。

```c
#include <stdio.h>

void fun(int arr[], int len) {
    for (size_t i = 0; i < len; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

int main(int argc, char const *argv[]) {
    fun((int[]){1, 2, 3}, 3);
    // int *ptr = (int[]){1, 2, 3};
    // fun(ptr, 3);
    return 0;
}
```

### 关闭缓冲区和取消回显

```c
system("stty -icanon");
system("stty -echo");
system("stty echo");
system("stty icanon");
```

### 调用子程序

使用 system 函数，可以调用其它应用程序。（windows 下？）

### 定义二进制数

```c
#include <stdio.h>

int main(int argc, char const *argv[]) {
    int num = 0b0101;
    printf("%d", num);
    return 0;
}
```

### 位字段

```c
#include <limits.h>
#include <stdio.h>
const static int SIZE = CHAR_BIT * sizeof(int);

char *itobs(int num, char *ps);

typedef struct St {
    unsigned int num : 8;
    unsigned int : 4;
    unsigned int num2 : 16;
    unsigned int : 0;  // 自动对齐
} St;

typedef union Views {
    St st_view;
    unsigned int ui_view;
} Views;

int main(int argc, char const *argv[]) {
    char binStr[SIZE + 1];
    Views inst = {{0xff, 0xffff}};
    printf("%s\n", itobs(inst.ui_view, binStr));
    return 0;
}

char *itobs(int num, char *ps) {
    for (int i = SIZE - 1; i >= 0; num >>= 1, i--) {
        ps[i] = (1 & num) + '0';
    }
    ps[SIZE] = '\0';
    return ps;
}
```

### linux-fork 函数

```c
#include <pthread.h>
#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>
#include <wait.h>

int value = 0;
void *runner(void *param);

int main(int argc, const char *argv[]) {
    int pid;
    pthread_t tid;
    pthread_attr_t attr;
    pid = fork();
    if (pid == 0) {
        pthread_attr_init(&attr);
        pthread_create(&tid, &attr, runner, NULL);
        pthread_join(tid, NULL);
        printf("CIDLD: value = %d\n", value);
    } else if (pid > 0) {
        wait(NULL);
        printf("PARENT: value = %d\n", value);
    }
    printf("value: %d\n", value);
}

void *runner(void *param) {
    value = 5;
    pthread_exit(0);
}
```

### `malloc` 与 `free` 的行为定义

- 针对 `malloc` 无法申请内存的 `NULL` 情况，`free` 无行为
  - 写 `NULL` 操作导致 seg fault
- `free` 接受的指针和 `malloc` 获得的内存指针必须相同，否则产生 RE
- 避免动态内存遗漏回收

### 全局变量初始化

全局变量默认零值初始化，对全局变量进行初始化只能使用常量，哪怕是常量变量也不行

```c
const int num = 12;
// compile error
// int copy = num;
```

### 预定义

输出对应的名称，通常用于 debug 等环境

- `__FUNCTION__`：显示函数名
- `__FILE__`：显示文件名
- `__LINE__`：显示当前文件的行数
- `__DATE__`：显示预处理的日期
- `__TIME__`：显示编译的时间
- `__STDC__`：要求遵循标准 C 的标志

### 隐式转换优先级

隐式转换：优先转换为无符号数字

```c
#include <stdio.h>

int main(int argc, char const *argv[]) {
    if (-1 > 0u) {
        printf("surprise\n");
    }
    return 0;
}
```

整数的比较，有符号整数优先转换为无符号整数，常见于 `sizeof`：

```c
for (int i = 10; i >= 0u; i--) {
    // dead loop
}
// never reach
```

count down with unsigned int:

```c
for (unsigned i = cnt - 1; i < cnt; i--) {
    // loop execution
}
// normal exit
```

### `while` 的两种汇编方式

```c
    goto test;
loop: // body
test: if (cond) goto loop;
```

```c
if (!cond) goto done;
loop: // body
test: if (cond) goto loop;
done:
```

### 浮点数内存布局和表示范围的探索

```c
#include <stdint.h>
#include <stdio.h>

void showBytes(uint8_t *addr, size_t size) {
    for (int i = 0; i < size; i++) {
        printf("%p %02x\n", addr + i, addr[i]);
    }
}

int main(int argc, char const *argv[]) {
    float num = 0;
    num = 1 / num;
    printf("%e\n", num * 0.0);
    printf("%f\n", num * 0.0);
    showBytes((uint8_t *)&num, sizeof(num));

    uint8_t *ptr = (uint8_t *)(&num);
    *ptr = 0xff;
    // ptr[1] = 0xff;

    showBytes((uint8_t *)&num, sizeof(num));
    printf("%e\n", num);
    printf("%f\n", num);
    return 0;
}
```

### 获得系统变量

```c
#include <stdlib.h>
#include <stdio.h>

int main() {
    const char *env = getenv("CLASSPATH");
    printf("%s\n", env);
    return 0;
}
```

```sh
export CLASSPATH=abc
./test
```

## linux 与 c

### 编译运行过程

生成可执行文件的过程：

1. 预处理：`-E` 宏替换、头文件展开，生成 `.i` 文件
2. 编译：`-S` 汇编语言文件，生成 `.s` 文件
3. 汇编：`-c` 二进制文件，生成 `.o` 文件
4. 链接：`-o` 可执行文件

[动态库的链接](https://www.cnblogs.com/dpf-learn/p/6202432.html)

可以使用 `bear` 等工具生成 ide 所需的编译指令文件 `compile_commands.json`

`cmake` 也可以通过参数 `-DCMAKE_EXPORT_COMPILE_COMMANDS=YES` 导出 `compile_commands.json`

### 链接器的使用

链接器 `ld` 可以通过 `-T` 参数自定义链接脚本，从而指定链接时的重定向地址定义可执行文件的内存布局

通过 `ld --verbose` 可以查看默认的链接脚本

### 跨架构运行二进制代码

如果需要跨平台运行二进制代码，可以使用 `qemu` 模拟器

如 `mips` 架构下的 `qemu-mips`（可以通过 `apt install qemu-user` 安装）

### 系统调用

- 通过 `man syscall` 和 `man syscalls` 可以查看 linux 的系统调用
- 通过 `man 7 signal` 可以查看 linux 的信号
- 并且在程序中也可以进行调用。简单示例：[syscall](./syscall.c)

### `glibc` 的调试

源代码的下载：[glibc homepage](https://www.gnu.org/software/libc/sources.html)

```bash
cd /usr/src/
git clone https://sourceware.org/git/glibc.git
```

content of `~/.gdbinit`:

```gdb
directory /usr/src/glibc
```

正常情况下，`gdb` 能够递归查找 `directory` 的内容，否则需要[结合 `find` 指令和 `-d` 参数](https://stackoverflow.com/questions/1103161/searching-for-source-directories-in-gdb)

### 查看库版本

```bash
ldd <executable>
/path/to/lib
```

### 获取符号对应的函数

- [demangler](https://demangler.com/)
- `nm` 可以用于查看二进制文件（.so 文件、可执行文件）的符号表

## MPI

- [compile: mpicc/mpic++](https://docs.open-mpi.org/en/v5.0.x/building-apps/quickstart.html)
- [run: mpirun](https://docs.open-mpi.org/en/v5.0.x/launching-apps/quickstart.html#using-mpirun-to-launch-applications)
