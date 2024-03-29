# 线性代数

2022 年只考一道大题：二次型、特征值、特征向量、相似理论

- 综合考察：行列式、矩阵化简、向量组、线性方程组
- 一些题目可以进行简单的验证：求矩阵的逆、解线性方程组

数乘向量表示对向量的放缩

- scaling 放缩
- scalars 标量

高维到一维的线性变换等价于高维向量之间的点积（对偶性）

线性变换与基变换：

- 线性变换可以用基变换表示
- 基变换实际上是一个线性变换
- $A^{-1}MA$ 可以看作是进行转换、变换、逆转换

经过正交变换后，向量与基向量的点积不变：

- $T(x) \cdot T_i = x \cdot I_i = x_i$
  - $T_i, I_i$ 表示对第 $i$ 个基向量进行的变换

快速求解二阶矩阵行列式：$m\pm\sqrt{m^2-p}$

## 一些基础概念的明晰

矩阵和左乘和右乘、行向量和列向量：

- 矩阵左乘可以看作对列空间的基向量进行变换后对列向量进行批处理
  - 可以看作对行向量的线性组合
  - 行变换不改变列向量的线性表示方式
- 矩阵右乘可以看作对列向量的线性组合
  - 可以看作对行空间的基向量进行变换后对行向量进行批处理

区别将基变为 $A$ 和使用 $B$ 作为基表示向量：

- 基变为 $A$：变换时原向量看作的坐标不变、基发生变化，在原参考系下向量变为 $A\beta$
- 使用 $B$ 作为基表示向量：在原坐标系下观察向量不发生变化，在新坐标系下满足 $Bx = E\beta$
  - 因此，对应的变换为 $P = B^{-1}$
