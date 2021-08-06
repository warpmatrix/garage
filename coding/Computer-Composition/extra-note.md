# 各种缩写

- MAR (Memory Address Register)
- MDR (Memory Data Register)
- ALU (Arithmetic and Logical Unit)
- PSW (Program State Word) 程序状态字，程序状态寄存器
- PC (Program Counter)
- IR (Instruction Register)
- CU (Control Unit)
- GPR (General Propose Register)
- ROM (Read Only Memory)
- ACC (ACCumulator)
- MQ (Multiple—Quotient Register)
- X 操作数寄存器
- IX (IndeX register) 变址寄存器
- BR (Base Register)
- BCD (Binary-Coded Decimal)
- FA (Full Adder)
- CLA (Carry Lookahead Adder)
- BCLA (Block Carry Lookahead Adder) 成组先行进位电路
- RAM (Random Access Memory)
- ROM (Read Only Memory)
- DRAM (Dynamic Random Access Memory)
- SRAM (Static Random Access Memory)
- MROM (Mask Read Only Memory)
- PROM (Programable Read Only Memory)
- EPROM (Erasable Programable Read Only Memory)
- SSD (Solid State Drive)
- MM (Main Memory)
- RAND (RANDom)
- FIFO (First In First Out)
- LRU (Least Recently Used)
- LFU (Least Frequently Used)
- TLB (Translation Lookaside Buffer) 转译后备缓冲区，快表
- PTR (Page Table Register) 页表寄存器

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

## 疑惑问题

- 芯片中片选端有两位的原因？（便于扩展？方便粗粒度和细粒度的两级管理？）
- 一个 TLB 对应一个进程？只有一个 TLB（失效问题）？
