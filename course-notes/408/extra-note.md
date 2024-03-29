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
- EA (Effective Address)
- SP (Stack Pointer)
- LIFO (Last In First Out)
- CM (Control Memory)
- CMAR (Control Memory Address Register) 微地址寄存器
- CIR, $\mu$IR (Control Instruction Register) 微指令寄存器
- RAW (Read After Write) 写后读
- WAR (Write After Read) 读后写
- WAW (Write After Write) 写后写
- DMA (Direct Memory Access)
- VRAM (Video Random Access Memory)
- RAID (Redunant Array of Inexpensive Disks)
- INTR (INTerrupt Register)
- IF (Interrupt Flag)
- NMI (Non Maskable Interrupt)
- FCFS (First Come First Served)
- SSTC (Shortest Seek Time First)
- SCAN 扫描算法
- C-SCAN (Circular SCAN)
- FCB (File Control Block)
- MFD (Master File Directory)
- UFD (User File Directory)
- ACL (Access Control List)
- SPOOL (Simultaneous Peripheral Operation On-Line)
- DCT (Device Control Table)
- COCT (COntroller Control Table)
- CHCT (CHannel Control Table)
- SDT (System Device Table)
- LUT (Logical Unit Table)
- PCB (Process Control Block)
- TCB (Thread Control Block)
- ULT (User Level Thread)
- KLT (Kernel Level Thread)
- SJF (Shortest Job First)
- SPF (Shortest Process First)
- WPL (Weighted Path Length)
- BFS (Breadth First Search)
- DFS (Depth First Search)
- MST (Minimum Spanning Tree)
- AOV (Activity On Vertex network)
- AOE (Activity On Edge network)
- ASL (Average Search Length)
- WAN (Wide Area Network)
- MAN (Metropolitan Area Network)
- LAN (Local Area Network)
- PAN (Personal Area Network)
- WPAN (Wireless Personal Area Network)
- RTT (Round-Trip Time)
- SDU (Service Data Unit)
- PCI (Protocol Control Infomation)
- PDU (Protocol Data Unit)
- SAP (Service Access Point)
- DTE (Data Terminal Equipment)
- DTE (Data Communication Equipment)
- TCP (Transmission Control Protocol)
- UDP (User Datagram Protocol)
- IP (Internet Protocol)
- PCM (Pulse Code Modulation)
- RZ (Return Zero code)
- NRZ (Non Return Zero code)
- NRZI (Non Return Zero Inverse code)

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

数据通路、数据总线：

- 数据通路：数据通过数据总线连接形成的数据传输路径
- 数据总线：数据传输的媒介

接口 (interface)、端口 (port)：

- 端口：接口电路中可以进行读写的寄存器
- 接口：若干端口加上对应的控制逻辑

保存断点、保存现场：

- 保存断点：保存原程序的 PC 和 PSWR 的内容
- 保存现场：将（中断/子）程序使用到的寄存器的信息保存到栈中

## 疑惑问题

- 芯片中片选端有两位的原因？（便于扩展？方便粗粒度和细粒度的两级管理？）
- 一个 TLB 对应一个进程？只有一个 TLB（失效问题）？
- `goto` 语句违背了局部性原理？
- 对应用程序员而言虚拟存储器是透明的（指的是内存？）
- 段号是否需要存储
- 进程映像和进程之间的区别
- 同一个服务程序被不同的用户调用，被创建为不同的进程？
- 多对一模型的用处
