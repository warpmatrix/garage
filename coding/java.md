# java

CLI 下的 java 编译过程：

- `javac` 编译 java 文件生成对应的 class 文件
- `java` 使用 jvm 运行 class 文件

java 的打包：[jar](https://docs.oracle.com/javase/tutorial/deployment/jar/basicsindex.html) 用于无损压缩文件，通常用于分发、共享 java 软件，可以通过 `unzip` 或 `jar` 对 jar 文件进行解压。

java 的反编译：

- 通过 eclipse 可以直接查看 `class` 对应的源码
- 在命令行下可以使用相应的反编译器：[cfr](http://www.benf.org/other/cfr/) 等
  - 具体的使用方法：`java -jar <cfr.jar> <classfile> > <output-javafile>`
  - [cfr](https://www.cnblogs.com/wowcl/p/15185482.html)，[其他的反编译器](https://forum.ubuntu.org.cn/viewtopic.php?t=68007)

java 多个文件的编译：

- 同一个文件夹下多个源文件：同一个 package 名称，`javac <srcs>`
- `-d` 设置生成 class 文件的目录
- `-cp` (-classpath) 指定依赖的包的路径
- `@<file>` 可以使用文件内容指代编译源文件列表

java 的源文件名字和类名保持一致

java 的 Stream：

- 流式风格的编程，方便进行并行、延迟执行、短路等操作
- 提供 `foreach` 的 vistor 的内部迭代模式访问，无需显式在集合外部迭代
- 创建流：`array.stream()`
- `forEach`、`map`、`filter`、`limit`、`sorted`、`collect`、`summaryStatistics`

java 中的 lambda 函数：`() -> {}`
