<!-- omit in toc -->
# 数学必背结论

<!-- omit in toc -->
## Table of Contents

- [1. 奇淫巧计篇](#1-奇淫巧计篇)
  - [1.1. 微积分技巧](#11-微积分技巧)
  - [1.2. 组合数技巧](#12-组合数技巧)
  - [1.3. 复数技巧](#13-复数技巧)
  - [三角函数技巧](#三角函数技巧)
  - [1.4. 结论篇](#14-结论篇)
- [2. 基础篇](#2-基础篇)
  - [2.1. 微积分基础](#21-微积分基础)
  - [2.2. 复数基础](#22-复数基础)

## 1. 奇淫巧计篇

### 1.1. 微积分技巧

涉及三角的积分，多留意函数的奇偶性。

1. $\int^{+\infty}_0e^{-x^2}dx = \frac{\sqrt\pi}2$

### 1.2. 组合数技巧

1. $\Sigma_{k=0}^nC_n^kx^k = (1+x)^n$

2. 形如 $\Sigma_{k=0}^nk^mC^k_n$ 的和：对上式求导，乘 $x$，重复 k 次。

   - $\Sigma_{k=0}^nkC^k_n=n2^{n-1}$（或者使用倒序相加）
   - $\Sigma_{k=0}^nk^2C^k_n=n(n+1)2^{n-2}$

3. $\Sigma_{k=0}^n(C^k_n)^2 = C^n_{2n}$

### 1.3. 复数技巧

以下形式的式子都可以通过欧拉公式进行转换：

- $\cos x \pm jsinx$
- $j\cos x \pm sinx$

### 三角函数技巧

$\sin$ 与 $\cos$ 的求和公式，上下同除 $\sin\frac{x}{2}$ 构造和差化积公式得到：

- $\sum_{k=1}^n \sin kx = \frac{\sin\frac{nx}{2}\sin\frac{(n+1)x}{2}}{\sin \frac{x}{2}}, x \neq 0$
- $\sum_{k=1}^n \cos kx = \frac{\sin\frac{nx}{2}\cos\frac{(n+1)x}{2}}{\sin \frac{x}{2}}, x \neq 0$

### 1.4. 结论篇

- $\int^{+\infty}_0e^{-x^2}dx = \frac{\sqrt\pi}2$
- $\Sigma_{k=0}^nC_n^kx^k = (1+x)^n$
- $\Sigma_{k=0}^nkC^k_n=n2^{n-1}$
- $\Sigma_{k=0}^nk^2C^k_n=n(n+1)2^{n-2}$

## 2. 基础篇

### 2.1. 微积分基础

### 2.2. 复数基础
