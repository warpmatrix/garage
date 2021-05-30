<!-- omit in toc -->
# c 和 c++ 的一些冷门易忘的点

<!-- omit in toc -->
## Content of Table

- [C语言篇](#c语言篇)
  - [scanf()读到回车结束输入](#scanf读到回车结束输入)
  - [c的字符串操作](#c的字符串操作)
  - [c中的布尔类型 (C99)](#c中的布尔类型-c99)
  - ['\r'的作用](#r的作用)
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
  - [隐式转换的优先级](#隐式转换的优先级)
- [C plusplus篇](#c-plusplus篇)
  - [引用的用法](#引用的用法)
  - [在函数中传递函数形参](#在函数中传递函数形参)
  - [名字空间的使用方法 & 关键字 `enum`](#名字空间的使用方法--关键字-enum)
  - [构造函数初始化列表](#构造函数初始化列表)
  - [delete不方便之处](#delete不方便之处)
  - [c++函数引用](#c函数引用)
  - [c++ 的变长数组](#c-的变长数组)

## C语言篇

### scanf()读到回车结束输入

```c
#include <stdio.h>
#define MAXLEN 100

int main(int argc, char const *argv[]) {
    char str[MAXLEN];
    scanf("%[^\n]", str);
    printf("%s\n", str);
    return 0;
}
```

### c的字符串操作

c中大多数自带的字符串操作都自带 ```'\0'```。

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

在里```<stdbool.h>```，将 ```bool``` 作为 ```_Bool``` 的别名，并且可以使用 ```true``` 和 ```false```。

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

### 隐式转换的优先级

整数的比较，有符号整数优先转换为无符号整数：

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

## C plusplus篇

### 引用的用法

```cpp
#include <iostream>

int &min(int &num1, int &num2) {
    return num1 < num2 ? num1 : num2;
}

int main(int argc, char const *argv[]) {
    int num1 = 2, num2 = 4;
    int &lesser = min(num1, num2);
    // int &lesser = num1 < num2 ? num1 : num2;
    lesser++;
    std::cout << num1 << num2 << '\n';
    return 0;
}
```

也可以嵌套使用条件运算符。

```c
#include <stdio.h>

int main(int argc, char const *argv[]) {
    int num1 = 4, num2 = 2, num3 = 1;
    int &min = (num1 < num2) ? (num1 < num3 ? num1 : num3)
                             : (num2 < num3 ? num2 : num3);
    printf("%d\n", min);
    return 0;
}
```

### 在函数中传递函数形参

```cpp
// Program Designing Lesson/Double Linked List
list &list::remove_if(bool (*condition)(listPointer p) ) {
    listPointer p = head;
    for(int i=0; i<_size; i++) {
        if(condition(p) ) {
            erase(i);
            i--;
        }
        p = p->next;
    }
}
```

### 名字空间的使用方法 & 关键字 `enum`

```cpp
// Program Designing Lesson/Alipay System 1 user
namespace alipay {
namespace Gender {
    enum Gender { Female = 0, Male = 1, Unknown = 2 };
}
class User {
public:
    User() {
      this->gender = Gender::Unknown;
    }
    Gender::Gender getGender(void);
    bool setGender(Gender::Gender gender);
private:
    Gender::Gender gender;
}
}
namespace alipay {
Gender::Gender User::getGender(void) {
    return gender;
}
bool User::setGender(Gender::Gender gender){
    this->gender = gender;
    return true;
}
}
```

### 构造函数初始化列表

```cpp
// Program Designing Lesson/The Date class(version 3)
class Date {
public:
    Date(int year=0, int month=0, int day=0);
};
Date::Date(int year, int month, int day)
    :year(year), month(month), day(day) {}
```

### delete不方便之处

```cpp
int main(int argc, char const *argv[]) {
    int *arr = new int[10];
    delete (arr+9);
    delete []arr;
    //false, 2 allocs, 3 frees
    return 0;
}
```

### c++函数引用

```cpp
// Program Designing Lesson/FunctionTemplate(eden)
inline bool myCmp(const int & a, const int & b) {
    return a > b;
}

bool (&fun) (const int &a, const int &b) = myCmp;
```

### c++ 的变长数组

c++ 中不能直接传递变长多维数组，只能通过函数模板传递非类型参数。

```cpp
#include <iostream>

template<int rows, int cols>
void fun(int (&ptr)[rows][cols]) {
    for (size_t i = 0; i < rows*cols; i++)
        std::cout << (*ptr)[i] << ' ';
    std::cout << '\n';
}

int main(int argc, char const *argv[]) {
    const int rows = 2, cols = 3;
    int arr[rows][cols];
    for (size_t i = 0; i < rows * cols; i++)
        (*arr)[i] = i;
    fun<rows, cols>(arr);
    return 0;
}
```
