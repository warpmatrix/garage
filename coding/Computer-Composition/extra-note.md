# 各种缩写

- MAR (memory address register)
- MDR (memory data register)
- ALU (arithmetic and logical unit)
- PSW (program state word)
- PC (program counter)
- IR (instruction register)
- CU (control unit)
- GPR (general propose register)
- ROM (read only memory)
- ACC (Accumulator)
- MQ (Multiple—Quotient Register)
- X 操作数寄存器
- IX 变址寄存器
- BR (base register)

## 容易混淆的概念

翻译程序、汇编程序、编译程序、解释程序：

- 翻译程序
  - 编译程序：翻译成汇编代码或机器代码
    - 汇编程序
  - 解释程序：逐条翻译执行

机器字长、指令字长、存储字长：

- 机器字长：直接处理的二进制数据的位数
- 指令字长：指令中二进制代码的位数
- 存储字长：存储单元中二进制代码的位数

## 一些 CSAPP 上的补充

- 现代 64 位的计算机，一般只有 47 位可用的内存空间大小
- 使用虚拟内存进行管理，逻辑实现 64 位的内存
