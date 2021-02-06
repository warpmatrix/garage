<!-- omit in toc -->
# 数学必背结论

<!-- omit in toc -->
## Table of Contents

- [1. 基础篇](#1-基础篇)
- [2. 奇淫巧计篇](#2-奇淫巧计篇)
  - [2.1. 恒等变化](#21-恒等变化)
  - [2.2. 微积分技巧](#22-微积分技巧)
  - [2.3. 组合数技巧](#23-组合数技巧)
  - [2.4. 复数技巧](#24-复数技巧)
  - [2.5. 三角函数技巧](#25-三角函数技巧)
- [3. 结论篇](#3-结论篇)
- [4. 傻子才犯的错误](#4-傻子才犯的错误)

## 1. 基础篇

函数：

- 研究函数时，需要讨论**定义域**和依赖关系
- 复合函数 $f(g(x))$ 需满足：$g(x)$ 的值域与 $f(x)$ 的定义域有交集

极限：

- 数列 $\left \{ x_n \right\}$ 的极限与前面有限项无关
- $\lim_{x \to x_0}f(x)$ 只与 $f(x)$ 中 $x = x_0$ 的去心邻域相关，与 $f(x_0)$ 无关
- 讨论左右极限的两种情况：
  - 分段函数分界点两端
  - 形如：$e^\infty, \arctan\infty$ 的函数
- 有界性：收敛 $\Rightarrow$ 有界，有界 $\nRightarrow$ 收敛，无界 $\Rightarrow$ 发散，发散 $\nRightarrow$ 收敛
- 保号性（正负判断）：极限 $\Rightarrow$ 极限过程，使用严格不等号；极限过程 $\Rightarrow$ 极限，使用非严格不等号

## 2. 奇淫巧计篇

### 2.1. 恒等变化

- $x - 1 < [x] \leq x$
- 分子有理化：$\sqrt{a} - \sqrt{b} = \frac{a - b}{\sqrt{a} + \sqrt{b}}$
  - $\sqrt{1 + x^2} - x = \frac{1}{x + \sqrt{1 + x^2}}$

### 2.2. 微积分技巧

涉及三角的积分，多留意函数的奇偶性。

1. $\int^{+\infty}_0e^{-x^2}dx = \frac{\sqrt\pi}2$

### 2.3. 组合数技巧

1. $\Sigma_{k=0}^nC_n^kx^k = (1+x)^n$

2. 形如 $\Sigma_{k=0}^nk^mC^k_n$ 的和：对上式求导，乘 $x$，重复 k 次。

   - $\Sigma_{k=0}^nkC^k_n=n2^{n-1}$（或者使用倒序相加）
   - $\Sigma_{k=0}^nk^2C^k_n=n(n+1)2^{n-2}$

3. $\Sigma_{k=0}^n(C^k_n)^2 = C^n_{2n}$

### 2.4. 复数技巧

以下形式的式子都可以通过欧拉公式进行转换：

- $\cos x \pm jsinx$
- $j\cos x \pm sinx$

### 2.5. 三角函数技巧

$\sin$ 与 $\cos$ 的求和公式，上下同除 $\sin\frac{x}{2}$ 构造和差化积公式得到：

- $\sum_{k=1}^n \sin kx = \frac{\sin\frac{nx}{2}\sin\frac{(n+1)x}{2}}{\sin \frac{x}{2}}, x \neq 0$
- $\sum_{k=1}^n \cos kx = \frac{\sin\frac{nx}{2}\cos\frac{(n+1)x}{2}}{\sin \frac{x}{2}}, x \neq 0$

## 3. 结论篇

- $\int^{+\infty}_0e^{-x^2}dx = \frac{\sqrt\pi}2$
- $\Sigma_{k=0}^nC_n^kx^k = (1+x)^n$
- $\Sigma_{k=0}^nkC^k_n=n2^{n-1}$
- $\Sigma_{k=0}^nk^2C^k_n=n(n+1)2^{n-2}$

## 4. 傻子才犯的错误

- $x \neq \sqrt{x^2} = |x|$
  - $\frac{\sqrt{x^2+1}}{x} \neq \sqrt{1 + \frac{1}{x^2}}$
