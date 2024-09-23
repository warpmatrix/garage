<!-- omit in toc -->
# Python3 的一些语法

<!-- omit in toc -->
## Table of Contents

- [1. 类型和运算](#1-类型和运算)
  - [1.1. 基本类型](#11-基本类型)
  - [1.2. 容器类型](#12-容器类型)
- [2. 库相关](#2-库相关)
  - [2.1. 官方标准库](#21-官方标准库)
  - [2.2. sortedcontainers](#22-sortedcontainers)
  - [2.3. 命令应用构建](#23-命令应用构建)
  - [2.4. 抽象基类](#24-抽象基类)
  - [2.5. 变量序列化和反序列化](#25-变量序列化和反序列化)
  - [2.6. python 对象的信息查看](#26-python-对象的信息查看)
- [3. 开发相关](#3-开发相关)
  - [3.1. 类型注释](#31-类型注释)
  - [3.2. 依赖管理](#32-依赖管理)
  - [3.3. `import` 机制](#33-import-机制)
  - [3.4. pdb 调试](#34-pdb-调试)
  - [3.5. 多线程编程](#35-多线程编程)
  - [3.6. 程序构建与分发](#36-程序构建与分发)

## 1. 类型和运算

### 1.1. 基本类型

python 的整数范围没有限制，maxint 可以用无穷大浮点数代替：`float("inf")`

逻辑运算：

- 与和或的关键字：`and` 和 `or`
- 真和假的关键字需要大写：`True`、`False`
  - 被当成内置常量，首字母大写
- else if 在 python 中使用关键字 `elif`

字符串：是不可修改的，单个字符也被视作一个字符串

- 修改某个字符需要转化为列表：`s = list(s)`
- 最后将列表转回字符串：`"".join(s)`

在 python 中函数间传递对象均使用引用方式，可以使用 `id` 函数查看对象的地址。

- 对象不变，对象的地址不会发生变化；对象发生变化即使使用同一个变量接收，对象的地址也发生变化
- [关于对象地址的示例代码](./addr-demo.py)

### 1.2. 容器类型

列表操作：

- `num1, num2, ..., numn = arr` 可以直接获取列表对应元素
  - 常见于 pair 数组组成的二维矩阵，使用 `for key, val in pairArr` 进行遍历
- `len([1, 2, 3])` 得到列表长度
- `[1, 2] + [3, 4]` 列表拼接
- `[1, 2] * 4` 列表重复，指针指向同一地址空间
- `3 in [1, 2, 3]` 检查元素是否在列表中
- `enumerate` 和 `in` 遍历列表
- `[] for _ in range(3)` 区别列表重复创建了新的数组空间
- `l.reverse()` 反转整个列表，没有返回值
  - 切片使用值传递，使用切片翻转不能影响原列表
  - 翻转部分子列表需要使用内置函数 `reversed`：`l[i:i+c] = reversed(l[i:i+c])`
- 末尾添加 `list.append(obj)`，插入数据 `list.insert(idx, obj)`
- 删除特定位置元素 `del list[idx]`，删除一个匹配元素 `list.remove(obj)`

`enumerate` 相当于 range，用于遍历支持迭代的对象

```py
for idx, elem in enumerate([1, 2, 3]):
    print(idx, elem)
```

推导式：

- `<expr> for <var> in <list>`
- `<expr> for <var> in <list> [if <cond>]`
- `<expr> [if <cond> else <expr>] for <var> in <list>`
- 不同类型的推导式：
  - 列表推导式：使用 `[]` 包围
  - 集合推导式：使用 `{}` 包围，增加去重功能
  - 字典推导式：使用 `{}` 包围，键和值使用 `:` 隔开，不同的键值使用 `,` 隔开
  - 生成器：使用 `()` 包围

生成器：

- 调用 `next` 一次产生一个元素，没有元素抛出异常
- 函数调用直接写函数名不需再写 `()`
- 函数中通过 `yield` 作为生成器生成数据
  - 运行到 `yield` 返回结果暂停运行
- 通过 `iterator = iter(iterable)` 生成迭代器
  - 判断可迭代：`isinstance(obj, Iterable)`

哈希表：字典 `dict`

- 键必须唯一且不可变（字符串、数字），值可以为任意类型
- 键值不在字典里报错
- `defaultdict`：不存在的键值，提供一个默认的字段
  - 接受类型或工厂函数作为参数：`m = defaultdict(int)`
- 可以使用 `items()` 方法枚举字典的键值对

无序集合：

- 使用 `{}` 包围，空集合声明使用 `set()`
- 使用面向对象得方法增加元素或删除元素

## 2. 库相关

### 2.1. 官方标准库

[标准库文档](https://docs.python.org/3/library/index.html)

堆：标准库模块 [heapq](https://docs.python.org/3/library/heapq.html)

- `heapq.heappush(heap, item)`, `heapq.heappop(heap)`
- heapq 实现的是小顶堆，使用大顶堆需要取相反数
- 判空操作：`if not heap:`
  - 区分 `heap is None`，`None` 表示空（类型），不是同一个类型

二分搜索模块：`bisect`

- `bisect_left`, `bisect_right`

官方日志库格式：https://docs.python.org/3/library/logging.html#logrecord-attributes

### 2.2. sortedcontainers

一个 Apache2 授权的排序集合库，[github 地址](https://github.com/grantjenks/python-sortedcontainers)、[官网地址](http://www.grantjenks.com/docs/sortedcontainers/#)

有序数组 `sortedList`：

- `from sortedcontainers import SortedList`

[有序映射 `SortedDict`](http://www.grantjenks.com/docs/sortedcontainers/sorteddict.html)：

- `from sortedcontainers import SortedDict`

### 2.3. 命令应用构建

构建命令行应用：[`typer`](https://github.com/tiangolo/typer) 库通过装饰器修饰。

命令行参数解析：使用 `argparse` 模块，一个简单的 [demo](./argparse-demo.py)。

终端美化：`Rich` 终端应用提供富文本格式。

### 2.4. 抽象基类

`abc` 库

### 2.5. 变量序列化和反序列化

`pickle` 库实现变量对象序列化和反序列化

### 2.6. python 对象的信息查看

[`inspect`](https://docs.python.org/3.11/library/inspect.html) 模块可以用于类、方法、函数等元信息的查看

## 3. 开发相关

python 变量对于数组进行引用传递，一个变量的操作可能导致另一个变量的改变

### 3.1. 类型注释

在函数中进行参数和返回值的类型声明：

```py
def func(var: int) -> int:
    return var
```

- 容器类型：python 里面的容器类型包括 `list, dict, tuple`，在 `python3.9+` 的版本可以直接使用内置的容器完成声明
  - 在 `python3.8-` 的版本使用容器类型需要导入 `typing` 模块：`from typing import List, Dict, Tuple`
- 函数声明：`Callable[[<args-type>], <ret-type>]`，需要导入 `typing` 模块
- 类型别名：使用自定义的类型名称
  - 与原始类型等价的声明方式：`<typename> = <type-expression>`
  - 创建新子类的声明方式：`<typename> = NewType('<typename>', <type-expression>)`
- [function overloading](https://adamj.eu/tech/2021/05/29/python-type-hints-how-to-use-overload/)

### 3.2. 依赖管理

导出依赖：`pipreqs . --encoding utf-8`

安装依赖：`pip install -r requestment.txt`

### 3.3. `import` 机制

开发过程中需要避免入口函数所在目录的文件和标准库重名。

`import` 的过程中，与标准库中的模块相比，当前目录的模块优先级更高。

python 存在两种运行方式，因此 `import` 也有[不同的讨论](https://zhuanlan.zhihu.com/p/63143493)：

### 3.4. pdb 调试

python 可以通过 [pdb](https://docs.python.org/3/library/pdb.html) 进行调试

- 并且 `.pdbrc` 可以实现与 `.gdbinit` 相同的功能

### 3.5. 多线程编程

多线程编程实例：[multi-threads.py](./multi-threads.py)

### 3.6. 程序构建与分发

- [setuptools](https://setuptools.pypa.io/en/latest/index.html)
- [pip install](https://pip.pypa.io/en/stable/cli/pip_install/)
