# Python3 的一些语法

函数参数和返回值的类型声明：

```py
def func(var: int) -> int:
    return var
```

python 的列表操作：

- `len([1, 2, 3])` 得到列表长度
- `[1, 2] + [3, 4]` 列表拼接
- `[1, 2] * 4` 列表重复
- `3 in [1, 2, 3]` 检查元素是否在列表中
- `enumerate` 和 `in` 遍历列表

`enumerate` 相当于 range，用于遍历支持迭代的对象

```py
for idx, elem in enumerate([1, 2, 3]):
    print(idx, elem)
```

python 里面的哈希表：字典 `dict`

- 键值不在字典里报错
- `defaultdict`：不存在的键值，提供一个默认的字段
  - 接受工厂函数作为参数

```py
m = defaultdict(int)
```

python 逻辑运算中，与和或的关键字：`and` 和 `or`
