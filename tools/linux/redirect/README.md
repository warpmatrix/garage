# Redirection

`./test.sh >test.log 2>&1` 的运行结果：

- 标准输出和标准错误的输出一起输出到日志文件
- `&` 表示引用，避免二次打开日志文件，造成输出结果的覆盖

```plaintext
./test.sh: 2: t: not found
Sat 17 Apr 2021 04:29:23 PM CST
```

`./test.sh >test.log 2>test.log` 的运行结果：

- 将日志文件打开了两次，标准输出将标准错误的输出覆盖

```plaintext
Sat 17 Apr 2021 04:31:42 PM CST
```
