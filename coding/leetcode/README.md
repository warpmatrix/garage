<!-- omit in toc -->
# Leetcode Note

思考如何切题：先动笔，再打码

- [54 & 59. Spiral Matrix](#54--59-spiral-matrix)
- [61. Rotate List](#61-rotate-list)
- [74. Search a 2D Matrix](#74-search-a-2d-matrix)
- [80. Remove Duplicates from Sorted Array II](#80-remove-duplicates-from-sorted-array-ii)
- [82 & 83. Remove Duplicates from Sorted List](#82--83-remove-duplicates-from-sorted-list)
- [88. Merge Sorted Array](#88-merge-sorted-array)
- [90. Subsets II](#90-subsets-ii)
- [92. Reverse Linked List II](#92-reverse-linked-list-ii)
- [115. Distinct Subsequences](#115-distinct-subsequences)
- [131. Palindrome Partitioning](#131-palindrome-partitioning)
- [132. Palindrome Partitioning II](#132-palindrome-partitioning-ii)
- [150. Evaluate Reverse Polish Notation](#150-evaluate-reverse-polish-notation)
- [173. Binary Search Tree Iterator](#173-binary-search-tree-iterator)
- [190. Reverse Bits](#190-reverse-bits)
- [191. Number of 1 Bits](#191-number-of-1-bits)
- [224 & 227. Basic Calculator](#224--227-basic-calculator)
- [240. Search a 2D Matrix II](#240-search-a-2d-matrix-ii)
- [331. Verify Preorder Serialization of a Binary Tree](#331-verify-preorder-serialization-of-a-binary-tree)
- [338. Counting Bits](#338-counting-bits)
- [341. Flatten Nested List Iterator](#341-flatten-nested-list-iterator)
- [354. Russian Doll Envelopes](#354-russian-doll-envelopes)
- [456. 132 Pattern](#456-132-pattern)
- [503. Next Greater Element II](#503-next-greater-element-ii)
- [525. Contiguous Array](#525-contiguous-array)
- [781. Rabbits in Forest](#781-rabbits-in-forest)
- [1006. Clumsy Factorial](#1006-clumsy-factorial)
- [1047. Remove All Adjacent Duplicates In String](#1047-remove-all-adjacent-duplicates-in-string)
- [1143. Longest Common Subsequence](#1143-longest-common-subsequence)
- [1603. Design Parking System](#1603-design-parking-system)
- [面试题 17.21. Volume of Histogram LCCI](#面试题-1721-volume-of-histogram-lcci)

## 54 & 59. Spiral Matrix

按层遍历，可以使用四个变量定位矩形坐标

## 61. Rotate List

新链表的最后一个节点为原链表的第 (n - 1) - (k mod n) 个节点（从 00 开始计数）

最坏情况下，遍历链表两次，时间复杂度 $O(n)$

## 74. Search a 2D Matrix

将二维矩阵打散为一维数组，二分查找；使用一维下标映射到矩阵，二分查找。时间复杂度：$O(\log m + \log n)$

- 区别：[240. Search a 2D Matrix II](#240-search-a-2d-matrix-ii)

## 80. Remove Duplicates from Sorted Array II

双指针简单模拟，`nums[slow - 2] != nums[fast]` 时，添加元素

## 82 & 83. Remove Duplicates from Sorted List

- 82: 使用区间删除，因为一个区间需要删除 $n+1$ 个元素，$n$ 为区间中相等关系数量。
- 83: 一个相等的元素对应一个删除元素，对每个元素进行判断删除即可
- 另外需要注意空指针，可以使用 dump 记录头节点的信息。

## 88. Merge Sorted Array

使用逆向指针实现 $O(1)$ 的空间复杂度

## 90. Subsets II

dfs，可以统计数字出现次数，减少递归深度；也可以使用 bool 值记录是否使用了当前数字。

需要注意 go 中的切片是指针赋值，进行赋值需要使用深复制

- `ret = append(ret, append([]int(nil), cur...))`
  - 区别 `ret = append(ret, cur)`

## 92. Reverse Linked List II

记录反转链表的前驱、后继，一次遍历实现

## 115. Distinct Subsequences

动态规划，`dp[i][j]` 表示字符串 $s$ 下标 $i$ 开始和目标字符串 $t$ 下标 $j$ 开始的匹配数。若 `s[i] == t[j]`，`dp[i][j] = dp[i + 1][j + 1] + dp[i + 1][j]` 进行匹配或放弃当前匹配；若 `s[i] != t[j]` 只能放弃匹配，`dp[i][j] = dp[i + 1][j]`。时间复杂度 $O(nm)$。

或者可以进行问题转换，先确定字符串 $t$ 中每个字符在 $s$ 中所在的位置，寻找递增序列的个数。`if (locs[t[i]][j] < locs[t[i+1]][k]) dp[i][j] += dp[i+1][k]`，`ans = sum(dp[0][i])`。时间复杂度 $O(n + m\frac{n}{|\Sigma|})$

## 131. Palindrome Partitioning

划分方法数（解空间）：$\sum\limits_{i=0}^{n-1} C_{n-1}^i = 2^{n-1}$

存在最坏情况下，解空间中的每一个解都接受，时间复杂度 $O(2^n)$

判断回文方法：动态规划，双指针判断

## 132. Palindrome Partitioning II

在 [131](#131-palindrome-partitioning) 的判断字串回文的基础上，使用动态规划。若字串 [i, j] 为回文串，则 $minCut[j] = \min(minCut[i] + 1)$。枚举所有可能的前缀 i，时间复杂度为 O(n^2)。

## 150. Evaluate Reverse Polish Notation

注意负号的存在，需要进行特别处理：

- go 使用函数 `num, err := strconv.Atoi(str)`
- cpp 使用字符串流 `std::stringstream ss; ss << str; ss >> num;`

## 173. Binary Search Tree Iterator

迭代做法：使用栈和一个指针维护递归的信息

## 190. Reverse Bits

位分治算法：使用掩码对数字进行分组，按 2 的幂次进行分组交换得到数字的反转

## 191. Number of 1 Bits

相关题目：[338. Counting Bits](#338-counting-bits)

## 224 & 227. Basic Calculator

- 224 只有 '+' 和 '-' 两种运算，针对正负两种情况设计符号栈，直接展开表达式。
- 227 只有四则运算，没有括号可以使用栈，先计算乘除，加减号将操作数入栈

显式后缀转中缀表达式，大量字符串操作容易超时，需要特别留意（尤其子字符串的传递），使用双栈（数字栈、符号栈）操作无需进行后缀转中缀的操作。

## 240. Search a 2D Matrix II

从左下角出发，或从右上角出发形成一棵二叉树，时间复杂度 $O(m+n)$

## 331. Verify Preorder Serialization of a Binary Tree

对空位进行计数，本质找规律，每增加一个节点增加两个子节点。

## 338. Counting Bits

- `x = x & (x−1)`，将 x 最低位的 1 变成 0
  - 若不返回数组，空间复杂度为 O(1)
- `bits[i] = bits[i − highBit] + 1`
- `bits[x] = bits[x >> 1] + (x & 1)`
- `bits[x] = bits[x & (x − 1)] + 1`

## 341. Flatten Nested List Iterator

递归实现嵌套扁平化；或者使用栈维护递归信息

## 354. Russian Doll Envelopes

二维字典序严格递增最长序列 -> 最长递增子序列

问题转换：先对第一维进行字典序排序，再对最后一个维度求最长递增子序列

- 避免前面维数不是严格递增，最后一维使用倒序排序
- 拓展到 $d$ 维的情况，对前 (d-1) 维的数据进行升序排序，最后一维降序排序
  - 对于固定的 $i$，$\forall j$ 满足 $a_i < a_j$ 都排在 $i$ 前面
  - 时间复杂度 $O(n^2k)$，[CDQ 分治](https://oi-wiki.org/misc/cdq-divide/)可以降低时间复杂度

## 456. 132 Pattern

- 枚举 3：维护左区间最小值，维护右区间次小值，时间复杂度 $O(n\log n)$
- 枚举 1：从右向左枚举，对于每一个右区间最大值 3，取区间次小值作为 2，使用单调递减栈维护右区间的可能最大值，时间复杂度 $O(n)$
- 枚举 2：取左区间的尽可能小值作为 1，尽可能大值作为 3，其中 `&1 < &3`

## 503. Next Greater Element II

单调栈维护序列的单调性，两个循环实现循环队列的需求，也可以在第一个循环增加循环长度下标取模实现

## 525. Contiguous Array

将 [0, 1] 映射到 [-1, 1]，使用前缀和统计 0 和 1 出现的次数。前缀和相同意味着区间内 0 和 1 的个数相等。使用哈希维护区间最大值和最小值，返回区间的最大值。

## 781. Rabbits in Forest

贪心，使用 `answers[i] + 1` 对相同的 `answers[i]` 进行分组

## 1006. Clumsy Factorial

利用数学方法找规律，相邻项可以相互约掉：

$N - 3 \approx \frac{(N-4)*(N-5)}{N-6}$

## 1047. Remove All Adjacent Duplicates In String

使用栈（string 也可以作为容器）记录移出相邻重复的字符。主要看清题意，每次移出两个字符。

## 1143. Longest Common Subsequence

经典动态规划

$$
dp[i][j] = \begin{cases}
  dp[i - 1][j - 1] + 1, & text1[i] == text2[j] \\
  max(dp[i - 1][j], dp[i][j - 1]), & text1[i] != text2[j] \\
\end{cases}
$$

## 1603. Design Parking System

简单模拟

## 面试题 17.21. Volume of Histogram LCCI

`ans[i] += min(lmax[i], rmax[i]) - h[i]`

需要注意的点：

- `lmax`/`rmax` 单调递增/递减
- 若 `lmax[i] > rmax[j]` 则，$\exist nj < j, rmax[nj] \geq lmax[i]$，可以适用双指针
