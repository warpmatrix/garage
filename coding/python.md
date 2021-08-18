# Python3 的一些语法

函数参数和返回值的类型声明：

```py
def func(var: int) -> int:
    return var
```

python 的列表操作：

- `len([1, 2, 3])` 得到列表长度
- `[1, 2] + [3, 4]` 列表拼接
- `[1, 2] * 4` 列表重复，指针指向同一地址空间
- `3 in [1, 2, 3]` 检查元素是否在列表中
- `enumerate` 和 `in` 遍历列表
- `[] for _ in range(3)` 区别列表重复创建了新的数组空间

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
  - 接受工厂函数作为参数

```py
m = defaultdict(int)
```

python 的逻辑运算

- 与和或的关键字：`and` 和 `or`
- 真和假的关键字需要大写：`True`、`False`
  - 被当成内置常量，首字母大写
- else if 在 python 中使用关键字 `elif`
