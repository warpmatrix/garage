<!-- omit in toc -->
# c++ 的一些冷门易忘的点

<!-- omit in toc -->
## Table of Contents

- [引用的用法](#引用的用法)
- [在函数中传递函数形参](#在函数中传递函数形参)
- [名字空间的使用方法 \& 关键字 `enum`](#名字空间的使用方法--关键字-enum)
- [c++11 枚举类 `enum class`](#c11-枚举类-enum-class)
- [构造函数初始化列表](#构造函数初始化列表)
- [使用花括号初始化变量](#使用花括号初始化变量)
- [delete 每次调用删除整块内存](#delete-每次调用删除整块内存)
- [c++函数引用](#c函数引用)
- [c++ 的变长数组](#c-的变长数组)
- [STL sort](#stl-sort)
- [尾置返回类型 (c++11)](#尾置返回类型-c11)
- [`typeid` 获取类型信息](#typeid-获取类型信息)
- [lambda 函数实现递归](#lambda-函数实现递归)
- [cpp 的多值返回 - Structured Bindings](#cpp-的多值返回---structured-bindings)
- [字面量中的数字分隔符 (c++14)](#字面量中的数字分隔符-c14)
- [类成员函数指针](#类成员函数指针)

## 引用的用法

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

## 在函数中传递函数形参

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

## 名字空间的使用方法 & 关键字 `enum`

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

## c++11 枚举类 `enum class`

[c++11-enum class](https://blog.csdn.net/weixin_42817477/article/details/109029172)

```cpp
enum class TransportControllerType : uint8_t {
    NONE = 0
};
```

## 构造函数初始化列表

```cpp
// Program Designing Lesson/The Date class(version 3)
class Date {
public:
    Date(int year=0, int month=0, int day=0);
};
Date::Date(int year, int month, int day)
    :year(year), month(month), day(day) {}
```

## 使用花括号初始化变量

在C++11中，自动变量和全局变量的初始化方式包括：

- 等号 `=` 加上赋值表达式（assignment-expression），例如：`int a=2+3`;
- 等号 `=` 加上花括号表达式的初始化列表，例如：`int a = {3+4}`;
- 圆括号式的表达式列表（expression-lit），例如：`int a(6+8)`;
- 花括号式的初始化列表，例如：`int a{6+8}`;

另外，目前也支持使用 initialize_list 进行初始化 [link](https://www.cnblogs.com/zyk1113/p/13452493.html)

```cpp
struct TransportDownloadTaskInfo {
    fw::ID m_rid;
    uint64_t m_filelength{ std::numeric_limits<uint64_t>::max() };
    uint32_t m_byterate{ std::numeric_limits<uint32_t>::max() };
};
```

## delete 每次调用删除整块内存

```cpp
int main(int argc, char const *argv[]) {
    int *arr = new int[10];
    delete (arr+9);
    delete []arr;
    //false, 2 allocs, 3 frees
    return 0;
}
```

## c++函数引用

```cpp
// Program Designing Lesson/FunctionTemplate(eden)
inline bool myCmp(const int & a, const int & b) {
    return a > b;
}

bool (&fun) (const int &a, const int &b) = myCmp;
```

## c++ 的变长数组

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

## STL sort

`std::sort` 封装了快排，不是一个稳定的排序，保证稳定性需要使用 `std::stable_sort`

## 尾置返回类型 (c++11)

通过关键字 `auto` 和 `->` 尾置返回类型 (trailing return types)：类似 golang

- 提高可读性，使用模板时可以进行类型推导

```cpp
template <typename X, typename Y>
auto f(X x, Y y) -> decltype(x + y) {
    return x + y + 1.5;
}

// template <typename X, typename Y>
// decltype(*(X*)0 + *(Y*)0) f(X x, Y y) {
//     return x + y + 1.5;
// }
```

没有类型推断的情况下，只能手动推导具体的模板类型，或者使用 auto?

```cpp
template<typename X, typename Y, typename Z>
Z f(X x, Y y) {
    return Z(x + y);
}

// template <typename X, typename Y>
// auto f(X x, Y y) {
//     return x + y;
// }
```

## `typeid` 获取类型信息

头文件 `<typeinfo>`，`typeid` 是一个运算符类似于 `sizeof`，除了模板和引用在编译时进行推导

- 可以对类型名称、常量、变量、模板变量进行推导
- 存在方法 `name` 输出类型名称，不同编译器可能不一样

```cpp
template <typename X, typename Y, typename Z>
Z f(X x, Y y) {
    Z z = Z(x + y + 1.5);
    std::cout << (typeid(int) == typeid(1)) << '\n';
    std::cout << typeid(x).name() << ' ' << typeid(y).name() << ' ' << typeid(z).name() << '\n';
    return z;
}
```

## lambda 函数实现递归

```cpp
auto f = [&](auto &&self, int n) -> {
    self(self, n - 1);
}
```

## cpp 的多值返回 - [Structured Bindings](./structured_bindings.cpp)

在 c++17 之前使用 `std::tie` 接收多值返回，c++17 之后可以直接使用结构化绑定 (structured bindings)。不过一般情况下，为了保证代码的可读性还是尽量声明结构体比较合适。

ref: [C++17 结构化绑定(Structured Bindings)初探](https://zhuanlan.zhihu.com/p/139140185)、[c++函数返回多值](https://www.zhihu.com/question/57540006)

## 字面量中的数字分隔符 (c++14)

c++14 中引入了数字分隔符可以使用在整数（包含其他进制）、浮点数：

```cpp
int bin = 0b1001'1101;
int oct = 0123'456;
int dec = 123'456;
int hex = 0x7C00'7FFF;
float f = 123.456'789;
```

## 类成员函数指针

ref: [类成员函数指针](https://www.runoob.com/w3cnote/cpp-func-pointer.html)
