<!-- omit in toc -->
# Leetcode Note

思考如何切题：先动笔，再打码！leetcode 难度分级不科学，困难其实是中等难度

<!-- TODO: 以后再补吧，lay 了 -->
可以留意的一些题目：28、87、213、240、264、338、421、456、525、783、1473、1486、1310

- [7. Reverse Integer](#7-reverse-integer)
- [12. Integer to Roman](#12-integer-to-roman)
- [13. Roman to Integer](#13-roman-to-integer)
- [26. Remove Duplicates from Sorted Array](#26-remove-duplicates-from-sorted-array)
- [27. Remove Element](#27-remove-element)
- [28. Implement strStr()](#28-implement-strstr)
- [54 & 59. Spiral Matrix](#54--59-spiral-matrix)
- [61. Rotate List](#61-rotate-list)
- [74. Search a 2D Matrix](#74-search-a-2d-matrix)
- [80. Remove Duplicates from Sorted Array II](#80-remove-duplicates-from-sorted-array-ii)
- [81. Search in Rotated Sorted Array II](#81-search-in-rotated-sorted-array-ii)
- [82 & 83. Remove Duplicates from Sorted List](#82--83-remove-duplicates-from-sorted-list)
- [87. Scramble String](#87-scramble-string)
- [88. Merge Sorted Array](#88-merge-sorted-array)
- [90. Subsets II](#90-subsets-ii)
- [91. Decode Ways](#91-decode-ways)
- [92. Reverse Linked List II](#92-reverse-linked-list-ii)
- [115. Distinct Subsequences](#115-distinct-subsequences)
- [131. Palindrome Partitioning](#131-palindrome-partitioning)
- [132. Palindrome Partitioning II](#132-palindrome-partitioning-ii)
- [137. Single Number II](#137-single-number-ii)
- [150. Evaluate Reverse Polish Notation](#150-evaluate-reverse-polish-notation)
- [153 & 154. Find Minimum in Rotated Sorted Array](#153--154-find-minimum-in-rotated-sorted-array)
- [173. Binary Search Tree Iterator](#173-binary-search-tree-iterator)
- [179. Largest Number](#179-largest-number)
- [190. Reverse Bits](#190-reverse-bits)
- [191. Number of 1 Bits](#191-number-of-1-bits)
- [208. Implement Trie (Prefix Tree)](#208-implement-trie-prefix-tree)
- [213. House Robber II](#213-house-robber-ii)
- [215. Kth Largest Element in an Array](#215-kth-largest-element-in-an-array)
- [220. Contains Duplicate III](#220-contains-duplicate-iii)
- [224 & 227. Basic Calculator](#224--227-basic-calculator)
- [240. Search a 2D Matrix II](#240-search-a-2d-matrix-ii)
- [263. Ugly Number](#263-ugly-number)
- [264. Ugly Number II](#264-ugly-number-ii)
- [331. Verify Preorder Serialization of a Binary Tree](#331-verify-preorder-serialization-of-a-binary-tree)
- [338. Counting Bits](#338-counting-bits)
- [341. Flatten Nested List Iterator](#341-flatten-nested-list-iterator)
- [354. Russian Doll Envelopes](#354-russian-doll-envelopes)
- [377. Combination Sum IV](#377-combination-sum-iv)
- [403. Frog Jump](#403-frog-jump)
- [421. Maximum XOR of Two Numbers in an Array](#421-maximum-xor-of-two-numbers-in-an-array)
- [456. 132 Pattern](#456-132-pattern)
- [503. Next Greater Element II](#503-next-greater-element-ii)
- [525. Contiguous Array](#525-contiguous-array)
- [554. Brick Wall](#554-brick-wall)
- [633. Sum of Square Numbers](#633-sum-of-square-numbers)
- [690. Employee Importance](#690-employee-importance)
- [740. Delete and Earn](#740-delete-and-earn)
- [781. Rabbits in Forest](#781-rabbits-in-forest)
- [783. Minimum Distance Between BST Nodes](#783-minimum-distance-between-bst-nodes)
- [872. Leaf-Similar Trees](#872-leaf-similar-trees)
- [897. Increasing Order Search Tree](#897-increasing-order-search-tree)
- [938. Range Sum of BST](#938-range-sum-of-bst)
- [993. Cousins in Binary Tree](#993-cousins-in-binary-tree)
- [1006. Clumsy Factorial](#1006-clumsy-factorial)
- [1011. Capacity To Ship Packages Within D Days](#1011-capacity-to-ship-packages-within-d-days)
- [1047. Remove All Adjacent Duplicates In String](#1047-remove-all-adjacent-duplicates-in-string)
- [1143. Longest Common Subsequence](#1143-longest-common-subsequence)
- [1269. Number of Ways to Stay in the Same Place After Some Steps](#1269-number-of-ways-to-stay-in-the-same-place-after-some-steps)
- [1310. XOR Queries of a Subarray](#1310-xor-queries-of-a-subarray)
- [1442. Count Triplets That Can Form Two Arrays of Equal XOR](#1442-count-triplets-that-can-form-two-arrays-of-equal-xor)
- [1473. Paint House III](#1473-paint-house-iii)
- [1482. Minimum Number of Days to Make m Bouquets](#1482-minimum-number-of-days-to-make-m-bouquets)
- [1486. XOR Operation in an Array](#1486-xor-operation-in-an-array)
- [1603. Design Parking System](#1603-design-parking-system)
- [1720. Decode XORed Array](#1720-decode-xored-array)
- [1723. Find Minimum Time to Finish All Jobs](#1723-find-minimum-time-to-finish-all-jobs)
- [1734. Decode XORed Permutation](#1734-decode-xored-permutation)
- [1738. Find Kth Largest XOR Coordinate Value](#1738-find-kth-largest-xor-coordinate-value)
- [面试题 17.21. Volume of Histogram LCCI](#面试题-1721-volume-of-histogram-lcci)

## 7. Reverse Integer

简单题。与数学的定义不同，求余的结果可以是负数，有负数也可以直接进行反转；另外题目要求超过表示范围的数，返回零需要自己额外处理

## 12. Integer to Roman

- 模拟，类似于词法解析，关键从低位开始解析
- 数据量较小时，硬编码枚举所有的情况，计算速度更快

## 13. Roman to Integer

简单模拟：将每个罗马字符看作一个数字，依据其位置加减对应的值

- go 中没有 char 类型，字符使用 byte 表示

## 26. Remove Duplicates from Sorted Array

简单模拟，双指针，跳过重复的元素

## 27. Remove Element

简单模拟，可以使用双指针优化一个变量

## 28. Implement strStr()

实现 KMP 算法，本质上求子串中前缀和后缀的公共子串长度

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

## 81. Search in Rotated Sorted Array II

存在重复元素不能使用二分法找到旋转数组的旋转位置，只能线性寻找然后使用二分，倒不如直接线性寻找

- 对比 [153 没有重复元素的情况](#153-find-minimum-in-rotated-sorted-array)

## 82 & 83. Remove Duplicates from Sorted List

- 82: 使用区间删除，因为一个区间需要删除 $n+1$ 个元素，$n$ 为区间中相等关系数量。
- 83: 一个相等的元素对应一个删除元素，对每个元素进行判断删除即可
- 另外需要注意空指针，可以使用 dump 记录头节点的信息。

## 87. Scramble String

字符串划分区间后位置不再改变，原问题只和两个子问题相关，可以使用动态规划，记忆化搜索：

- isScramble(s1, s2) = $\vee_{i=1}^{len-1}$ [isScramble($s1_l$, $s2_l$) $\wedge$ isScramble($s1_r$, $s2_r$)] $\vee$ [isScramble($s1_l$, $s2_r$) $\wedge$ isScramble($s1_r$, $s2_l$)]
- 还需要注意 go 语言的**二维 map** 如何声明，如何进行初始化

## 88. Merge Sorted Array

使用逆向指针实现 $O(1)$ 的空间复杂度

## 90. Subsets II

dfs，可以统计数字出现次数，减少递归深度；也可以使用 bool 值记录是否使用了当前数字。

需要注意 go 中的切片是指针赋值，进行赋值需要使用深复制

- `ret = append(ret, append([]int(nil), cur...))`
  - 区别 `ret = append(ret, cur)`

## 91. Decode Ways

简单动态规划

- $f_i = f_{i-1} + f_{i-2}$，加 $f_{i-1}$ 和 $f_{i-2}$ 需要满足对应的条件
- go 中使用 `strconv.Atoi` 必须使用两个变量接收返回值

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

## 137. Single Number II

简单做法：使用哈希完成，时间复杂度 $O(n)$，空间复杂度 $O(n)$

使用性质求解可以节约空间，达到 $O(1)$ 的空间复杂度

- 由于重复的数字，均出现三次，则重复数字对应的每一位数字求和得到的结果为 3 的倍数
- 使用数字电路的方法（真值表进行状态转移），对每一位同时进行求和操作

## 150. Evaluate Reverse Polish Notation

注意负号的存在，需要进行特别处理：

- go 使用函数 `num, err := strconv.Atoi(str)`
- cpp 使用字符串流 `std::stringstream ss; ss << str; ss >> num;`

## 153 & 154. Find Minimum in Rotated Sorted Array

- 153：旋转数组无重复元素可以直接二分查找旋转位置
- 154：左边数组的元素大于右边数组的元素，使用二分进行查找
  - 最差情况 $O(n)$ 也可以使用线性搜索

## 173. Binary Search Tree Iterator

迭代做法：使用栈和一个指针维护递归的信息

## 179. Largest Number

拼接数最大，使用两个 segment 进行拼接比较，算法课已经做过一次了 orz

## 190. Reverse Bits

位分治算法：使用掩码对数字进行分组，按 2 的幂次进行分组交换得到数字的反转

## 191. Number of 1 Bits

相关题目：[338. Counting Bits](#338-counting-bits)

## 208. Implement Trie (Prefix Tree)

简单模拟：使用前缀和表示搜索词

优化：使用循环表示递归，本题没有必要使用递归

- 发现自己非递归实现树的算法是一个弱项

## 213. House Robber II

循环队列动态规划

- 第一个元素和最后一个元素不能共存，分成两个区间进行讨论：$[0, len(nums) - 1], [1, len(nums)]$
- 动态规划递推公式包含有限项元素，可以使用**滚动数组**

## 215. Kth Largest Element in an Array

使用快排的思路寻找元素，快速选择的模板

## 220. Contains Duplicate III

检查滑动窗口内两个元素的差小于阈值，使用集合和 `lower_bound` 完成：

- 注意边界条件：`max(num[i] - t, INT_MIN)` -> `max(nums[i], INT_MIN + t) - t`
- go 中没有实现有序集合需要自己完成
- 也可以对每一个滑动窗口维护桶，每增加一个元素检查对应的区间是否有符合条件的元素
  - 每一个桶的大小为 t + 1，桶的编号为 (x + 1) / w - 1

## 224 & 227. Basic Calculator

- 224 只有 '+' 和 '-' 两种运算，针对正负两种情况设计符号栈，直接展开表达式。
- 227 只有四则运算，没有括号可以使用栈，先计算乘除，加减号将操作数入栈

显式后缀转中缀表达式，大量字符串操作容易超时，需要特别留意（尤其子字符串的传递），使用双栈（数字栈、符号栈）操作无需进行后缀转中缀的操作。

## 240. Search a 2D Matrix II

从左下角出发，或从右上角出发形成一棵二叉树，时间复杂度 $O(m+n)$

## 263. Ugly Number

简单模拟，使用循环去除因子，判断是否为 1

## 264. Ugly Number II

动态规划，当前最小的丑数由可能的三个丑数乘对应的因子得到

```go
p2, p3, p5 := 0, 0, 0
uglyNums[i] = min(uglyNums[p2]*2, uglyNums[p3]*3, uglyNums[p5]*5)
if uglyNums[i] == uglyNums[p] { p++ }
```

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

## 377. Combination Sum IV

记忆化搜索，动态规划：$f(n) = \sum_i f(n - x_i)$

- 进阶问题：若组合数中存在负数，考虑 $a, -b$，则可能出现无限长度排列 $ab - ba = 0$

## 403. Frog Jump

如果青蛙可以两边跳，需要使用类似于图的遍历方式，需要进行标记。

状态转移方程：`dp[loc][speed] = dp[ploc][speed - 1] || dp[ploc][speed] || dp[ploc][speed + 1]`

题目给定的要求为单边跳，上式也同样成立，可以结合性质，考虑使用动态规划：

- $speed < t$、$loc(t) > t$

由于只是沿一个方向前进，可以考虑对数组进行差分，相邻数据可以求和合并，使得新数组相邻数据大小相差 1，开头数据为 1。

## 421. Maximum XOR of Two Numbers in an Array

求值问题转换为判断问题，二分法枚举最大的 $x$ 使得 $x = x_i \oplus x_j$，问题转换为对于 $a_i$ 数组中是否存在元素 $a_j = x \oplus a_i$

本质上，提前将数的前缀保存起来，避免枚举

- 可以使用哈希表，保存数组中的元素前 $k$ 位是什么数
- 可以使用字典树，保存数组中的元素对应前 $k$ 位的数字

## 456. 132 Pattern

- 枚举 3：维护左区间最小值，维护右区间次小值，时间复杂度 $O(n\log n)$
- 枚举 1：从右向左枚举，对于每一个右区间最大值 3，取区间次小值作为 2，使用单调递减栈维护右区间的可能最大值，时间复杂度 $O(n)$
- 枚举 2：取左区间的尽可能小值作为 1，尽可能大值作为 3，其中 `&1 < &3`

## 503. Next Greater Element II

单调栈维护序列的单调性，两个循环实现循环队列的需求，也可以在第一个循环增加循环长度下标取模实现

## 525. Contiguous Array

将 [0, 1] 映射到 [-1, 1]，使用前缀和统计 0 和 1 出现的次数。前缀和相同意味着区间内 0 和 1 的个数相等。使用哈希维护区间最大值和最小值，返回区间的最大值。

## 554. Brick Wall

简单题，确定每一行砖块边界的位置，统计哪一个位置边界最多

## 633. Sum of Square Numbers

暴力枚举挺方便的，但双指针是真的秀，可以应该可以降低时间复杂度的常数，时间复杂度为 $O(\sqrt{c})$：

- $l^2 + r^2 < c \to l = l + 1$
- $l^2 + r^2 > c \to r = r - 1$

## 690. Employee Importance

简单遍历，广度优先或深度优先

## 740. Delete and Earn

简单题：排序，统计每一个数字出现的次数（方便查看对应的数字是什么，可以计算其前缀和），使用动态规划处理，转换为类似于 [213](#213-house-robber-ii) 的题目。

## 781. Rabbits in Forest

贪心，使用 `answers[i] + 1` 对相同的 `answers[i]` 进行分组

## 783. Minimum Distance Between BST Nodes

中序遍历寻找相邻数字的最小差：

- 实现过程：实现中序遍历、中序遍历完成相邻元素的比较（分开实现，各司其职），特别是用栈实现的方法
- 使用值域以外的值区分第一次遍历
- 使用栈完成中序遍历

## 872. Leaf-Similar Trees

遍历二叉树的叶子节点进行比较：两棵树同时 dfs 得到叶子节点进行比较

- 关于中序遍历的又一深入理解（有状态的中序遍历）

## 897. Increasing Order Search Tree

中序遍历直接改变指针的指向

- 反问题：线性树 -> 平衡树
- 可以考虑使用旋转解决，比较麻烦

## 938. Range Sum of BST

简单递归，不必过度纠结中序遍历。发现自己会受别人的思路影响，神 tm jm 说了一句中序遍历...

- 也可以使用其它遍历二叉树的方法

## 993. Cousins in Binary Tree

判断堂兄节点，需要记录深度信息和父节点信息

- 使用 dfs 进行遍历，比较好写，但需要找到两个节点
- 使用 bfs 进行遍历，可以提前退出

## 1006. Clumsy Factorial

利用数学方法找规律，相邻项可以相互约掉：

$N - 3 \approx \frac{(N-4)*(N-5)}{N-6}$

## 1011. Capacity To Ship Packages Within D Days

整数规划问题无多项式时间解法，将整数规划问题转换为判断问题，使用二分查找运载能力，得到所需天数 $d$

- 下界：运载能力为最重的物品（一般的（小数）规划问题中，最优装载量为物品重量的平均值，最大值比平均值要大）
- 上界：无法确定 $d$ 的范围，只能假设 $d$ 最小的情况（$d=1$），装载量为所有物品重量的总和

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

## 1269. Number of Ways to Stay in the Same Place After Some Steps

动态规划：`dp[step][loc] = dp[step - 1][loc - 1] + dp[step - 1][loc] + dp[step - 1][loc + 1]`

## 1310. XOR Queries of a Subarray

类比为区间和的问题：可以使用前缀异或数组求解

## 1442. Count Triplets That Can Form Two Arrays of Equal XOR

区间问题转换为前缀问题，通过下面的性质进行问题转换使用二重循环可以解决：

- $\oplus_{idx=i}^{j-1} arr[idx] = \oplus_{idx=j}^{k}arr[idx]$
- $\oplus_{idx=i}^k arr[idx] = 0$
- $\oplus_{idx=0}^{i-1} arr[idx] = \oplus_{idx=0}^k arr[idx]$
  - 考虑构造 `preXors[i] = ^arr[0..i-1]`
  - 可以使用哈希进行加速，一重循环解决
  - `ret += cnt[preXor] * idx + sumIdx[preXor]`

## 1473. Paint House III

困难动态规划，其实也不是很难，三维动态规划：

- 考虑和第 i-1 间房子进行比较
  - 若 $c_i = c_{i-1} = c$，则 $dp[i][c][neiNum] = dp[i-1][c][neiNum]$
  - 若 $c_i \neq c_{i-1}$，则 $dp[i][c_i][neiNum] = dp[i-1][c_{i-1}][neiNum - 1]$
- 未涂色，加 `cost[i][c]`

注意初始化：`i == 0` 和 `neiNum == 0` 的情况

- 可以使用滚动数组节省空间
<!-- TODO: 更高级的做法，日后再看看 -->

## 1482. Minimum Number of Days to Make m Bouquets

- 求解问题转换为判断问题，使用二分法进行求解
- go 中可以使用 `sort.Search` 函数实现二分查找
- 一些性能优化：排序后进行二分查找，时间复杂度减小，但仍不能弥补排序的开销，只能说数据出的...

## 1486. XOR Operation in an Array

简单模拟，初始值可以设为 0（`0 ^ num = num`）

看这种出题方法也可以估计出存在 $O(1)$ 时间复杂度的解法，利用异或的性质：

- 连续数字进行异或可以使用的性质：$4i \oplus (4i+1) \oplus (4i+2) \oplus (4i+3) = 0$
- 对于线性变化的数字，若倍数是 2 次幂可以对低位数字提出来特别讨论

## 1603. Design Parking System

简单模拟

## 1720. Decode XORed Array

简单题，利用两条性质完成：异或满足交换律，两次异或得到原来的操作数

- `encoded[i] = arr[i] xor arr[i + 1]`
- `arr[i + 1] = encoded[i] xor arr[i]`

## 1723. Find Minimum Time to Finish All Jobs

ref: [1011](#1011-capacity-to-ship-packages-within-d-days) 给定顺序进行规划？

1723：需要自己设定顺序

<!-- TODO: 有空看看二分的做法 -->

可以考虑动态规划与状态压缩：`dp[wNum][jobs] = min(max(dp[wNum-1][jobs-subJobs], sum(calcTime(subJobs))))`

## 1734. Decode XORed Permutation

留意到和 [1720](#1720-decode-xored-array) 的类似情景，考虑构造出 $p[0] = (\oplus_{i=1}^np[i]) \oplus (\oplus_{i=0}^np[i])$，可以有两种方式构造出来：

- $\oplus_{i=0}^{j-1}(p[i] \oplus p[i+1]) = p[0] \oplus p[j]$
  - $\oplus_{i=1}^{2n+1}(p[0] \oplus p[i]) = \oplus_{i=1}^{2n+1}p[i]$
- $\oplus_{i=2k+1}(p[i] \oplus p[i+1]) = \oplus_{i=1}^{2n+1}p[i],\ (k = 0, 1, 2, \cdots)$

## 1738. Find Kth Largest XOR Coordinate Value

二维前缀异或 + 快速选择，二维前缀的构造类比二维前缀和

- go 里面传递数组好像是写时复制？递归传递数组比较消耗时间
- 尽量将递归改成循环

## 面试题 17.21. Volume of Histogram LCCI

`ans[i] += min(lmax[i], rmax[i]) - h[i]`

需要注意的点：

- `lmax`/`rmax` 单调递增/递减
- 若 `lmax[i] > rmax[j]` 则，$\exist nj < j, rmax[nj] \geq lmax[i]$，可以适用双指针
