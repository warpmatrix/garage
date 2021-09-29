# Python3 的一些语法

函数参数和返回值的类型声明：

```py
def func(var: int) -> int:
    return var
```

python 的整数范围没有限制，maxint 可以用无穷大浮点数代替：`float("inf")`

python 的逻辑运算

- 与和或的关键字：`and` 和 `or`
- 真和假的关键字需要大写：`True`、`False`
  - 被当成内置常量，首字母大写
- else if 在 python 中使用关键字 `elif`

python 的字符串是不可修改的，单个字符也被视作一个字符串：

- 修改某个字符需要转化为列表：`s = list(s)`
- 最后将列表转回字符串：`"".join(s)`

python 的列表操作：

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

python 的列表推导式：

- `[<expr> for <var> in <list>]`
- `[<expr> for <var> in <list> [if <cond>]]`
- `[<expr> [if <cond> else <expr>] for <var> in <list>]`

python 里面的哈希表：字典 `dict`

- 键值不在字典里报错
- `defaultdict`：不存在的键值，提供一个默认的字段
  - 接受类型或工厂函数作为参数：`m = defaultdict(int)`

python 的堆：标准库模块 [heapq](https://docs.python.org/3/library/heapq.html)

- `heapq.heappush(heap, item)`, `heapq.heappop(heap)`
- heapq 实现的是小顶堆，使用大顶堆需要取相反数
- 判空操作：`if not heap:`
  - 区分 `heap is None`，`None` 表示空（类型），不是同一个类型

python 中的有序数组 `sortedList`：

- `from sortedcontainers import SortedList`

python 中的二分搜索模块：`bisect`

- `bisect_left`, `bisect_right`