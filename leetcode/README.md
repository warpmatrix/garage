<!-- omit in toc -->
# Leetcode Note

思考如何切题：先思考，再动笔，最后打码！leetcode 难度分级不科学，困难其实是中等难度

- leetcode 的难度好像崩坏了..中等也有可能只是简单题
- 说不定带薪上洗手间就可以思路就出来了

<!-- TODO: 以后再补吧，lay 了 -->
可以留意的一些题目：28、87、160、213、240、264、338、421、456、518、525、783、1109、1473、1486、1310、1787

- 思维训练的题：149、168、457、483、581、810、1787、面试题 17.10

题目 tag：

- 二分搜索：278、374、852
- 异或性质：1787
- 搜索：752、773、815、909

<!-- omit in toc -->
## Table of Contents

- [7. Reverse Integer](#7-reverse-integer)
- [12. Integer to Roman](#12-integer-to-roman)
- [13. Roman to Integer](#13-roman-to-integer)
- [26. Remove Duplicates from Sorted Array](#26-remove-duplicates-from-sorted-array)
- [27. Remove Element](#27-remove-element)
- [28. Implement strStr()](#28-implement-strstr)
- [29. Divide Two Integers](#29-divide-two-integers)
- [31. Next Permutation](#31-next-permutation)
- [34. Find First and Last Position of Element in Sorted Array](#34-find-first-and-last-position-of-element-in-sorted-array)
- [36. Valid Sudoku](#36-valid-sudoku)
- [53. Maximum Subarray](#53-maximum-subarray)
- [54 & 59. Spiral Matrix](#54--59-spiral-matrix)
- [61. Rotate List](#61-rotate-list)
- [65. Valid Number](#65-valid-number)
- [68. Text Justification](#68-text-justification)
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
- [138. Copy List with Random Pointer](#138-copy-list-with-random-pointer)
- [149. Max Points on a Line](#149-max-points-on-a-line)
- [150. Evaluate Reverse Polish Notation](#150-evaluate-reverse-polish-notation)
- [153 & 154. Find Minimum in Rotated Sorted Array](#153--154-find-minimum-in-rotated-sorted-array)
- [160. Intersection of Two Linked Lists](#160-intersection-of-two-linked-lists)
- [162. Find Peak Element](#162-find-peak-element)
- [165. Compare Version Numbers](#165-compare-version-numbers)
- [166. Fraction to Recurring Decimal](#166-fraction-to-recurring-decimal)
- [168. Excel Sheet Column Title](#168-excel-sheet-column-title)
- [171. Excel Sheet Column Number](#171-excel-sheet-column-number)
- [173. Binary Search Tree Iterator](#173-binary-search-tree-iterator)
- [179. Largest Number](#179-largest-number)
- [187. Repeated DNA Sequences](#187-repeated-dna-sequences)
- [190. Reverse Bits](#190-reverse-bits)
- [191. Number of 1 Bits](#191-number-of-1-bits)
- [203. Remove Linked List Elements](#203-remove-linked-list-elements)
- [208. Implement Trie (Prefix Tree)](#208-implement-trie-prefix-tree)
- [212. Word Search II](#212-word-search-ii)
- [213. House Robber II](#213-house-robber-ii)
- [215. Kth Largest Element in an Array](#215-kth-largest-element-in-an-array)
- [218. The Skyline Problem](#218-the-skyline-problem)
- [220. Contains Duplicate III](#220-contains-duplicate-iii)
- [224 & 227. Basic Calculator](#224--227-basic-calculator)
- [231. Power of Two](#231-power-of-two)
- [233. Number of Digit One](#233-number-of-digit-one)
- [240. Search a 2D Matrix II](#240-search-a-2d-matrix-ii)
- [263. Ugly Number](#263-ugly-number)
- [264. Ugly Number II](#264-ugly-number-ii)
- [273. Integer to English Words](#273-integer-to-english-words)
- [274. H-Index](#274-h-index)
- [275. H-Index II](#275-h-index-ii)
- [278. First Bad Version](#278-first-bad-version)
- [279. Perfect Squares](#279-perfect-squares)
- [295. Find Median from Data Stream](#295-find-median-from-data-stream)
- [297. Serialize and Deserialize Binary Tree](#297-serialize-and-deserialize-binary-tree)
- [313. Super Ugly Number](#313-super-ugly-number)
- [331. Verify Preorder Serialization of a Binary Tree](#331-verify-preorder-serialization-of-a-binary-tree)
- [338. Counting Bits](#338-counting-bits)
- [341. Flatten Nested List Iterator](#341-flatten-nested-list-iterator)
- [342. Power of Four](#342-power-of-four)
- [345. Reverse Vowels of a String](#345-reverse-vowels-of-a-string)
- [352. Data Stream as Disjoint Intervals](#352-data-stream-as-disjoint-intervals)
- [354. Russian Doll Envelopes](#354-russian-doll-envelopes)
- [374. Guess Number Higher or Lower](#374-guess-number-higher-or-lower)
- [377. Combination Sum IV](#377-combination-sum-iv)
- [401. Binary Watch](#401-binary-watch)
- [403. Frog Jump](#403-frog-jump)
- [405. Convert a Number to Hexadecimal](#405-convert-a-number-to-hexadecimal)
- [412. Fizz Buzz](#412-fizz-buzz)
- [413. Arithmetic Slices](#413-arithmetic-slices)
- [421. Maximum XOR of Two Numbers in an Array](#421-maximum-xor-of-two-numbers-in-an-array)
- [430. Flatten a Multilevel Doubly Linked List](#430-flatten-a-multilevel-doubly-linked-list)
- [434. Number of Segments in a String](#434-number-of-segments-in-a-string)
- [441. Arranging Coins](#441-arranging-coins)
- [443. String Compression](#443-string-compression)
- [446. Arithmetic Slices II - Subsequence](#446-arithmetic-slices-ii---subsequence)
- [447. Number of Boomerangs](#447-number-of-boomerangs)
- [451. Sort Characters By Frequency](#451-sort-characters-by-frequency)
- [456. 132 Pattern](#456-132-pattern)
- [457. Circular Array Loop](#457-circular-array-loop)
- [461. Hamming Distance](#461-hamming-distance)
- [470. Implement Rand10() Using Rand7()](#470-implement-rand10-using-rand7)
- [474. Ones and Zeroes](#474-ones-and-zeroes)
- [477. Total Hamming Distance](#477-total-hamming-distance)
- [483. Smallest Good Base](#483-smallest-good-base)
- [494. Target Sum](#494-target-sum)
- [502. IPO](#502-ipo)
- [503. Next Greater Element II](#503-next-greater-element-ii)
- [516. Longest Palindromic Subsequence](#516-longest-palindromic-subsequence)
- [518. Coin Change 2](#518-coin-change-2)
- [523. Continuous Subarray Sum](#523-continuous-subarray-sum)
- [524. Longest Word in Dictionary through Deleting](#524-longest-word-in-dictionary-through-deleting)
- [525. Contiguous Array](#525-contiguous-array)
- [526. Beautiful Arrangement](#526-beautiful-arrangement)
- [528. Random Pick with Weight](#528-random-pick-with-weight)
- [541. Reverse String II](#541-reverse-string-ii)
- [551. Student Attendance Record I](#551-student-attendance-record-i)
- [552. Student Attendance Record II](#552-student-attendance-record-ii)
- [554. Brick Wall](#554-brick-wall)
- [560. Subarray Sum Equals K](#560-subarray-sum-equals-k)
- [576. Out of Boundary Paths](#576-out-of-boundary-paths)
- [581. Shortest Unsorted Continuous Subarray](#581-shortest-unsorted-continuous-subarray)
- [600. Non-negative Integers without Consecutive Ones](#600-non-negative-integers-without-consecutive-ones)
- [611. Valid Triangle Number](#611-valid-triangle-number)
- [633. Sum of Square Numbers](#633-sum-of-square-numbers)
- [645. Set Mismatch](#645-set-mismatch)
- [664. Strange Printer](#664-strange-printer)
- [671. Second Minimum Node In a Binary Tree](#671-second-minimum-node-in-a-binary-tree)
- [678. Valid Parenthesis String](#678-valid-parenthesis-string)
- [690. Employee Importance](#690-employee-importance)
- [692. Top K Frequent Words](#692-top-k-frequent-words)
- [704. Binary Search](#704-binary-search)
- [725. Split Linked List in Parts](#725-split-linked-list-in-parts)
- [726. Number of Atoms](#726-number-of-atoms)
- [740. Delete and Earn](#740-delete-and-earn)
- [752. Open the Lock](#752-open-the-lock)
- [773. Sliding Puzzle](#773-sliding-puzzle)
- [781. Rabbits in Forest](#781-rabbits-in-forest)
- [783. Minimum Distance Between BST Nodes](#783-minimum-distance-between-bst-nodes)
- [787. Cheapest Flights Within K Stops](#787-cheapest-flights-within-k-stops)
- [797. All Paths From Source to Target](#797-all-paths-from-source-to-target)
- [802. Find Eventual Safe States](#802-find-eventual-safe-states)
- [810. Chalkboard XOR Game](#810-chalkboard-xor-game)
- [815. Bus Routes](#815-bus-routes)
- [847. Shortest Path Visiting All Nodes](#847-shortest-path-visiting-all-nodes)
- [852. Peak Index in a Mountain Array](#852-peak-index-in-a-mountain-array)
- [863. All Nodes Distance K in Binary Tree](#863-all-nodes-distance-k-in-binary-tree)
- [872. Leaf-Similar Trees](#872-leaf-similar-trees)
- [877. Stone Game](#877-stone-game)
- [879. Profitable Schemes](#879-profitable-schemes)
- [881. Boats to Save People](#881-boats-to-save-people)
- [897. Increasing Order Search Tree](#897-increasing-order-search-tree)
- [909. Snakes and Ladders](#909-snakes-and-ladders)
- [930. Binary Subarrays With Sum](#930-binary-subarrays-with-sum)
- [938. Range Sum of BST](#938-range-sum-of-bst)
- [981. Time Based Key-Value Store](#981-time-based-key-value-store)
- [987. Vertical Order Traversal of a Binary Tree](#987-vertical-order-traversal-of-a-binary-tree)
- [993. Cousins in Binary Tree](#993-cousins-in-binary-tree)
- [1006. Clumsy Factorial](#1006-clumsy-factorial)
- [1011. Capacity To Ship Packages Within D Days](#1011-capacity-to-ship-packages-within-d-days)
- [1035. Uncrossed Lines](#1035-uncrossed-lines)
- [1047. Remove All Adjacent Duplicates In String](#1047-remove-all-adjacent-duplicates-in-string)
- [1049. Last Stone Weight II](#1049-last-stone-weight-ii)
- [1074. Number of Submatrices That Sum to Target](#1074-number-of-submatrices-that-sum-to-target)
- [1104. Path In Zigzag Labelled Binary Tree](#1104-path-in-zigzag-labelled-binary-tree)
- [1109. Corporate Flight Bookings](#1109-corporate-flight-bookings)
- [1137. N-th Tribonacci Number](#1137-n-th-tribonacci-number)
- [1143. Longest Common Subsequence](#1143-longest-common-subsequence)
- [1190. Reverse Substrings Between Each Pair of Parentheses](#1190-reverse-substrings-between-each-pair-of-parentheses)
- [1221. Split a String in Balanced Strings](#1221-split-a-string-in-balanced-strings)
- [1239. Maximum Length of a Concatenated String with Unique Characters](#1239-maximum-length-of-a-concatenated-string-with-unique-characters)
- [1269. Number of Ways to Stay in the Same Place After Some Steps](#1269-number-of-ways-to-stay-in-the-same-place-after-some-steps)
- [1310. XOR Queries of a Subarray](#1310-xor-queries-of-a-subarray)
- [1337. The K Weakest Rows in a Matrix](#1337-the-k-weakest-rows-in-a-matrix)
- [1418. Display Table of Food Orders in a Restaurant](#1418-display-table-of-food-orders-in-a-restaurant)
- [1442. Count Triplets That Can Form Two Arrays of Equal XOR](#1442-count-triplets-that-can-form-two-arrays-of-equal-xor)
- [1449. Form Largest Integer With Digits That Add up to Target](#1449-form-largest-integer-with-digits-that-add-up-to-target)
- [1473. Paint House III](#1473-paint-house-iii)
- [1480. Running Sum of 1d Array](#1480-running-sum-of-1d-array)
- [1482. Minimum Number of Days to Make m Bouquets](#1482-minimum-number-of-days-to-make-m-bouquets)
- [1486. XOR Operation in an Array](#1486-xor-operation-in-an-array)
- [1583. Count Unhappy Friends](#1583-count-unhappy-friends)
- [1588. Sum of All Odd Length Subarrays](#1588-sum-of-all-odd-length-subarrays)
- [1600. Throne Inheritance](#1600-throne-inheritance)
- [1603. Design Parking System](#1603-design-parking-system)
- [1646. Get Maximum in Generated Array](#1646-get-maximum-in-generated-array)
- [1707. Maximum XOR With an Element From Array](#1707-maximum-xor-with-an-element-from-array)
- [1711. Count Good Meals](#1711-count-good-meals)
- [1713. Minimum Operations to Make a Subsequence](#1713-minimum-operations-to-make-a-subsequence)
- [1720. Decode XORed Array](#1720-decode-xored-array)
- [1723. Find Minimum Time to Finish All Jobs](#1723-find-minimum-time-to-finish-all-jobs)
- [1734. Decode XORed Permutation](#1734-decode-xored-permutation)
- [1736. Latest Time by Replacing Hidden Digits](#1736-latest-time-by-replacing-hidden-digits)
- [1738. Find Kth Largest XOR Coordinate Value](#1738-find-kth-largest-xor-coordinate-value)
- [1743. Restore the Array From Adjacent Pairs](#1743-restore-the-array-from-adjacent-pairs)
- [1744. Can You Eat Your Favorite Candy on Your Favorite Day?](#1744-can-you-eat-your-favorite-candy-on-your-favorite-day)
- [1787. Make the XOR of All Segments Equal to Zero](#1787-make-the-xor-of-all-segments-equal-to-zero)
- [1818. Minimum Absolute Sum Difference](#1818-minimum-absolute-sum-difference)
- [1833. Maximum Ice Cream Bars](#1833-maximum-ice-cream-bars)
- [1838. Frequency of the Most Frequent Element](#1838-frequency-of-the-most-frequent-element)
- [1846. Maximum Element After Decreasing and Rearranging](#1846-maximum-element-after-decreasing-and-rearranging)
- [1877. Minimize Maximum Pair Sum in Array](#1877-minimize-maximum-pair-sum-in-array)
- [1893. Check if All the Integers in a Range Are Covered](#1893-check-if-all-the-integers-in-a-range-are-covered)
- [1894. Find the Student that Will Replace the Chalk](#1894-find-the-student-that-will-replace-the-chalk)
- [剑指 Offer 10- I. 斐波那契数列](#剑指-offer-10--i-斐波那契数列)
- [剑指 Offer 15. 二进制中1的个数](#剑指-offer-15-二进制中1的个数)
- [剑指 Offer 22. 链表中倒数第k个节点](#剑指-offer-22-链表中倒数第k个节点)
- [剑指 Offer 37. 序列化二叉树](#剑指-offer-37-序列化二叉树)
- [剑指 Offer 38. 字符串的排列](#剑指-offer-38-字符串的排列)
- [剑指 Offer 42. 连续子数组的最大和](#剑指-offer-42-连续子数组的最大和)
- [剑指 Offer 52. 两个链表的第一个公共节点](#剑指-offer-52-两个链表的第一个公共节点)
- [剑指 Offer 53 - I. 在排序数组中查找数字 I](#剑指-offer-53---i-在排序数组中查找数字-i)
- [面试题 10.02. Group Anagrams LCCI](#面试题-1002-group-anagrams-lcci)
- [面试题 17.10. Find Majority Element LCCI](#面试题-1710-find-majority-element-lcci)
- [面试题 17.21. Volume of Histogram LCCI](#面试题-1721-volume-of-histogram-lcci)
- [LCP 07. 传递信息](#lcp-07-传递信息)

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

## 29. Divide Two Integers

中等题，求除法的商，可以使用移位或者使用递归完成

## 31. Next Permutation

获得排列的下一个排列，时间复杂度 $O(n)$

- 将数组中的低位非递减排列变为递增序列 e.g. [3, **1, 4, 2**] -> [3, 2, **1, 4**]
- 注意题目中记录排列的是数组，数组的低位元素对应排列的高位

## 34. Find First and Last Position of Element in Sorted Array

简单题，在数组中找到目标数字出现的区间，使用二分

## 36. Valid Sudoku

简单模拟

## 53. Maximum Subarray

简单题，返回数组中的最大子序和

- 使用前缀和完成，注意子序的长度至少为 1

## 54 & 59. Spiral Matrix

按层遍历，可以使用四个变量定位矩形坐标

## 61. Rotate List

新链表的最后一个节点为原链表的第 (n - 1) - (k mod n) 个节点（从 00 开始计数）

最坏情况下，遍历链表两次，时间复杂度 $O(n)$

## 65. Valid Number

判断有效数字，考虑使用正则表达式？（本质上构造一个自动状态机）

- go 的正则表达式的性能也太拉了吧，编译器就不能顺便优化一波？
- 实际过程可能需要自己构造自动状态机

## 68. Text Justification

中等题，偏工程的题目

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

## 138. Copy List with Random Pointer

带有随机指针的链表深复制：

- 可以使用哈希表记录原链表和新链表节点之间的映射关系，两次遍历可以实现
- 也可以使用哈希表记录节点的创建情况，使用回溯实现
- 还可以利用原链表的位置信息，避免哈希表的创建，再进行拆分，空间复杂度可以达到 $O(1)$

## 149. Max Points on a Line

求一条直线能通过点集的最大数量：枚举直线经过的第一个点，计算后面的点有多少个和起始点的斜率相等

- 使用浮点数计算斜率可能精度达不到题目的要求
  - 最小斜率为 $0.5*10^{-4}$，最大斜率为 $2*10^4$
- 优化的一些思路：
  - 当枚举到下标为 $i$ 的点，此时最多可以找到共线的点数为 $n - i$，当前最优解更优时，可以提前退出
  - 找到共线的点超过半数，可以提前退出
  - 根据数据范围可以将二维的哈希表映射到一维上，进一步节约空间

## 150. Evaluate Reverse Polish Notation

注意负号的存在，需要进行特别处理：

- go 使用函数 `num, err := strconv.Atoi(str)`
- cpp 使用字符串流 `std::stringstream ss; ss << str; ss >> num;`

## 153 & 154. Find Minimum in Rotated Sorted Array

- 153：旋转数组无重复元素可以直接二分查找旋转位置
- 154：左边数组的元素大于右边数组的元素，使用二分进行查找
  - 最差情况 $O(n)$ 也可以使用线性搜索

## 160. Intersection of Two Linked Lists

寻找链表相交节点，可以使用哈希进行遍历，时间和空间复杂度均为 $O(n)$

居然还可以使用双指针解决，时间复杂度 $O(n)$，空间复杂度 $O(1)$

- 两个指针将两个链表的路都走一次可以确保两个指针最终会走到一起
- 共同节点或 `null`

## 162. Find Peak Element

简单题，使用二分搜索局部判断是否为峰值

## 165. Compare Version Numbers

简单题模拟

## 166. Fraction to Recurring Decimal

简单模拟，需要注意整除，正负等情况

## 168. Excel Sheet Column Title

数值转换的变种，核心在于将权重从 1-26 映射到 0-25

- 可以将每次求得的权重重新进行映射
- `wei = num % 26 - 1`
  - if wei == -1 {wei, num = 25, num / 26 - 1}

## 171. Excel Sheet Column Number

简单题，进程转换：将 0-25 映射到 1-26。区别：

- [168. Excel Sheet Column Title](#168-excel-sheet-column-title)

## 173. Binary Search Tree Iterator

迭代做法：使用栈和一个指针维护递归的信息

## 179. Largest Number

拼接数最大，使用两个 segment 进行拼接比较，算法课已经做过一次了 orz

## 187. Repeated DNA Sequences

简单模拟，哈希记录子串计数信息

- 可以使用状态压缩，使用一个 20 位整数表示子串信息

## 190. Reverse Bits

位分治算法：使用掩码对数字进行分组，按 2 的幂次进行分组交换得到数字的反转

## 191. Number of 1 Bits

相关题目：[338. Counting Bits](#338-counting-bits)

## 203. Remove Linked List Elements

去除链表相同取值的元素：使用虚拟头节点和两个指针节点实现

## 208. Implement Trie (Prefix Tree)

简单模拟：使用前缀和表示搜索词

优化：使用循环表示递归，本题没有必要使用递归

- 发现自己非递归实现树的算法是一个弱项

## 212. Word Search II

字典树 + 搜索，实现在二维矩阵中检索单词

## 213. House Robber II

循环队列动态规划

- 第一个元素和最后一个元素不能共存，分成两个区间进行讨论：$[0, len(nums) - 1], [1, len(nums)]$
- 动态规划递推公式包含有限项元素，可以使用**滚动数组**

## 215. Kth Largest Element in an Array

使用快排的思路寻找元素，快速选择的模板

## 218. The Skyline Problem

中等题目，给定多条线段（每条线段有对应的权重），求边界上（任意一点）最大的权重

- 只需考察边界上的信息，在同一个位置权重由包含这一位置的线段的权重决定
  - 可以动态维护一个大顶堆，确定在一个位置中最大的权重，时间复杂度 $O(n \log n)$
- go 使用 `heap` 确实比较麻烦需要多练练
- TODO: 思考是否可能使用线段树完成

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

## 231. Power of Two

判断最低位的 1 是否是最高位的 1 -> 将最低位的 1 修改位 0 后数字是否为 0

- 需要注意处理非正数的情况

## 233. Number of Digit One

求 0-n 中 1 出现的次数：注意题目给定的数字是十进制数，小学奥数题

- 数据范围是 $10^9$，计算机算力通常是 $10^7/s$，需要使用 $O(\log n)$ 的方法：按照位数枚举
- 每一个位置上 1 的数量分别由高位（凑整的 1）和低位（不能凑整的 1）决定
  - 高位：第 $i$ 位 1 出现的连续出现的次数与周期之比为 $10^i / 10^{i+1}$
  - 低位：分权大于 1，等于 1，小于 1 三种情况讨论

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

## 273. Integer to English Words

简单语法分析题目，每三位进行分节

## 274. H-Index

简单题，求最大的 $h$ 满足 $cnt(num \geq h) \geq h$

- 构造 `cnt` 数组：`cnt[i]` 为数组中元素大于 `i` 的个数，时间复杂度为 $O(n)$
- 也可以对数组进行排序，时间复杂度为 $O(n \log n)$

## 275. H-Index II

数组升序的情况下，直接使用二分进行查找，时间复杂度 $O(\log n)$

- [274. H-Index](#274-h-index)

## 278. First Bad Version

二分算法，可以使用 go 的标准库 `sort.Search`，注意标准库对应下标从零开始

## 279. Perfect Squares

求解一个数最少可以分解为多少个完全平方数的和

- 可以使用动态规划，枚举 square 进行求解：`dp[n] = min(dp[n], dp[n - square] + 1)`
- 数学定理的求解方式（四平方和定理）：每一个整数至多被四个正整数的平方和表示
  - 当 $n = 4^k(8m + 7)$ 时，$n$ 只能表示为四个正整数的平方和
  - 特判答案为 1 或答案为 2 的情况

## 295. Find Median from Data Stream

维护数据流的中位数：

- 使用两个堆分别维护一半的数字
- 使用有序集合作为自动排序的数组，直接获得中间的（两个）数字

## 297. Serialize and Deserialize Binary Tree

将二叉树进行序列化和反序列化：

- 序列化：考虑对二叉树进行先序遍历
  - 一个 int 转成四个 byte，nil 考虑转为在区间以外的数字？

## 313. Super Ugly Number

简单题，新的丑数由先前计算的丑数延申得到，类似的题目：

- [264. Ugly Number II](#264-ugly-number-ii)

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

## 342. Power of Four

4 的幂：1 只出现 1 次，并且出现的位置在偶数位置上

- 对比：[231. Power of Two](#231-power-of-two)
- 使用掩码进行处理可以避免一个 32 次的循环
- 或者在判断 2 的幂的基础上，使用取模运算进行判断
  - $(4^n) \% 3 = 1$
  - $(2 \times 4^n) \% 3 = 2$

## 345. Reverse Vowels of a String

简单题，将字符串的特定字符进行反转：双指针 $O(n)$ 实现

## 352. Data Stream as Disjoint Intervals

中等简单题，考点数据结构，有序映射的应用

## 354. Russian Doll Envelopes

二维字典序严格递增最长序列 -> 最长递增子序列

问题转换：先对第一维进行字典序排序，再对最后一个维度求最长递增子序列

- 避免前面维数不是严格递增，最后一维使用倒序排序
- 拓展到 $d$ 维的情况，对前 (d-1) 维的数据进行升序排序，最后一维降序排序
  - 对于固定的 $i$，$\forall j$ 满足 $a_i < a_j$ 都排在 $i$ 前面
  - 时间复杂度 $O(n^2k)$，[CDQ 分治](https://oi-wiki.org/misc/cdq-divide/)可以降低时间复杂度

## 374. Guess Number Higher or Lower

二分搜索，类似题目：[278. First Bad Version](#278-first-bad-version)

## 377. Combination Sum IV

记忆化搜索，动态规划：$f(n) = \sum_i f(n - x_i)$

- 进阶问题：若组合数中存在负数，考虑 $a, -b$，则可能出现无限长度排列 $ab - ba = 0$

## 401. Binary Watch

固定海明权重的简单题，可以使用暴力搜索

- TODO: 高级一点的做法，考虑如何枚举固定海明权重的数字

## 403. Frog Jump

如果青蛙可以两边跳，需要使用类似于图的遍历方式，需要进行标记。

状态转移方程：`dp[loc][speed] = dp[ploc][speed - 1] || dp[ploc][speed] || dp[ploc][speed + 1]`

题目给定的要求为单边跳，上式也同样成立，可以结合性质，考虑使用动态规划：

- $speed < t$、$loc(t) > t$

由于只是沿一个方向前进，可以考虑对数组进行差分，相邻数据可以求和合并，使得新数组相邻数据大小相差 1，开头数据为 1。

## 405. Convert a Number to Hexadecimal

简单题，位运算（辗转相除法）

## 412. Fizz Buzz

简单模拟，[背后的起源](https://blog.codinghorror.com/why-cant-programmers-program/)

## 413. Arithmetic Slices

简单题，得到数组中等差序列的子数组个数

- 统计差分相等的个数，再使用求和公式即可
- 等价写法 `ret += ++cnt`

## 421. Maximum XOR of Two Numbers in an Array

求值问题转换为判断问题，二分法枚举最大的 $x$ 使得 $x = x_i \oplus x_j$，问题转换为对于 $a_i$ 数组中是否存在元素 $a_j = x \oplus a_i$

本质上，提前将数的前缀保存起来，避免枚举

- 可以使用哈希表，保存数组中的元素前 $k$ 位是什么数
- 可以使用字典树，保存数组中的元素对应前 $k$ 位的数字

## 430. Flatten a Multilevel Doubly Linked List

简单题，使用栈记录信息进行遍历

## 434. Number of Segments in a String

简单模拟，可以使用内置函数 `split`，注意空字符串不可以作为 segment

## 441. Arranging Coins

简单题，$O(1)$ 计算

## 443. String Compression

简单题，字符串原位压缩：使用双指针复用字符串的内存空间

- 创建字符串处理，空间复杂度 $O(\log n)$
- 倒转数字处理，空间复杂度 $O(1)$

## 446. Arithmetic Slices II - Subsequence

类似于 [413. Arithmetic Slices](#413-arithmetic-slices) 的第二种思想：每增加一个元素，增加当前公差相同的个数加 1

- 动态规划：`dp[i][d] += dp[j][d] + 1`

## 447. Number of Boomerangs

双重循环算距离，哈希加速

## 451. Sort Characters By Frequency

统计字符出现次数，倒序输出字符

- 使用桶进行记录（按照字符出现的次记录信息）
- go 中构建重复的字符串可以使用 `bytes.Repeat` 接口

## 456. 132 Pattern

- 枚举 3：维护左区间最小值，维护右区间次小值，时间复杂度 $O(n\log n)$
- 枚举 1：从右向左枚举，对于每一个右区间最大值 3，取区间次小值作为 2，使用单调递减栈维护右区间的可能最大值，时间复杂度 $O(n)$
- 枚举 2：取左区间的尽可能小值作为 1，尽可能大值作为 3，其中 `&1 ( &3`

## 457. Circular Array Loop

中等题，本质判断图中成环是否符号一致并且成环的距离应大于 1

- 每个节点出度为 1，必定成环
- 可以使用快慢指针判断，leetcode 141
- 也可以作两次标记，记录是当前路径访问的节点还是之前的路径访问过的节点，时间复杂度 $O(n)$

## 461. Hamming Distance

异或得到不同的位为 1，再计算海明权重

- go 中海明距离可以使用 `bits.OnesCount` 函数

## 470. Implement Rand10() Using Rand7()

可以从熵的角度进行考虑，多次调用提供更高的熵；设置一定的拒绝域，使得和 `rand10()` 的熵保持一致

## 474. Ones and Zeroes

二维容量的背包问题：动态规划，时间复杂度：$O(mnl)$

## 477. Total Hamming Distance

计算所有数字不同的位数（海明距离），按位计算：

- 统计所有数字每一位上的零或一的出现的次数
- 海明距离等于出现的零一次数相乘（零一相加为数组长度）

## 483. Smallest Good Base

已知等比数列的和，求等比数列最小的底：$n = \sum_{i=0}^m b^i$

- m 和 b 呈负相关，可以从大到小枚举 m，求出对应的 b 是否符合条件
- b 可以通过二分求出，也可以通过数学方式得到
- 计算乘方避免引入精度误差，不使用 pow 函数

## 494. Target Sum

> 可以直接使用动态规划枚举前面的数可以构成的表达式和，时间复杂度和空间复杂度较高

问题转换：枚举选 $i$ 个数看能否达到题目要求的负数和（$neg = \frac{sum - target}{2}$）

- 要求 $sum - target$ 为非负偶数
- $dp[i][j] = \begin{cases}
  dp[i - 1][j], & j < nums[i] \\
  dp[i - 1][j] + dp[i - 1][j - nums[i]], & j \geq nums[i] \\
\end{cases}$

## 502. IPO

中等题，银行家算法。使用堆动态维护可以选择的项目，快速得到可执行的最大利润项目。

## 503. Next Greater Element II

单调栈维护序列的单调性，两个循环实现循环队列的需求，也可以在第一个循环增加循环长度下标取模实现

## 516. Longest Palindromic Subsequence

简单题，求最长的回文子序列：动态规划

- 首尾相等：`dp[idx][offset] = 2 + dp[idx+1][offset]`
- 首尾不等：`dp[idx][offset] = max(dp[idx][offset-1], dp[idx+1][offset-1])`

## 518. Coin Change 2

零钱兑换问题（背包问题）：硬币的种类不定，注意到答案只和不同种类的硬币相关

- 区别跳台阶问题，动态规划要特别注意循环的内外层关系
  - 不同的嵌套关系，对应不同的子问题
  - 或者引入硬币的种类状态，时间复杂度会增加

## 523. Continuous Subarray Sum

区间问题：判断区间和为 k 的倍数

- 区间问题转化为前缀和的余数问题
- 并且使用线性表或者哈希表记录之前是否存在余数为 $k_1$ 的情况
  - k 较大时使用线性表消耗的空间较大

## 524. Longest Word in Dictionary through Deleting

在字符串中找到多个匹配字符串中最长的一个，双指针进行求解

- <!-- TODO: -->可以进一步使用动态规划进行优化

## 525. Contiguous Array

将 [0, 1] 映射到 [-1, 1]，使用前缀和统计 0 和 1 出现的次数。前缀和相同意味着区间内 0 和 1 的个数相等。使用哈希维护区间最大值和最小值，返回区间的最大值。

## 526. Beautiful Arrangement

中等题，求满足一定条件的排列：动态规划

- 观察数据范围考虑做题方向，状态压缩使用二进制枚举状态，依次数字可以放置的所有位置

## 528. Random Pick with Weight

简单题，使用权重进行随机选择：前缀和 + 二分搜索

## 541. Reverse String II

简单题，翻转字符串中一些连续的子串

## 551. Student Attendance Record I

简单题，模拟

## 552. Student Attendance Record II

中等题，在 [551. Student Attendance Record I](#551-student-attendance-record-i) 的规则下，给定长度求复合条件的情况数：

- 动态规划，分成几种情况进行讨论：
  - 'A' 已经出现过、没有出现过、'L' 在最近出现的次数，共 6 种情况
  - 编写代码进行递推时分 3 种情况：A, L, P
- 通过上面的递推式，可以使用矩阵快速幂进行加速

## 554. Brick Wall

简单题，确定每一行砖块边界的位置，统计哪一个位置边界最多

## 560. Subarray Sum Equals K

一维前缀和 + 哈希表（判定所需要的前缀和出现的次数）

- 哈希表用于加速检索元素的有无

## 576. Out of Boundary Paths

中等题，计算路径数量：动态规划，时间复杂度 $O(move \times row \times col)$，空间复杂度 $O(row \times col)$

## 581. Shortest Unsorted Continuous Subarray

求数组排序后，和原数组不同的区间长度

- 使用双指针：求两端升序数组的边界，满足相应的单调性

## 600. Non-negative Integers without Consecutive Ones

中等题，动态规划：

- 斐波拉契数列得到不连续为 1 序列的数量，从高位往低位尽可能填 1
- 填 1 后后一位必须为 0，最大整数有连续为 1 的情况可以提前退出

## 611. Valid Triangle Number

给定边集，求能够组成三角形的个数：遍历一条边求满足三角形条件的区间

- 思考问题的时候可以将相同长度的边合在一起思考，但是编码的时候会涉及组合数、前缀和等十分不方便处理

## 633. Sum of Square Numbers

暴力枚举挺方便的，但双指针是真的秀，可以应该可以降低时间复杂度的常数，时间复杂度为 $O(\sqrt{c})$：

- $l^2 + r^2 < c \to l = l + 1$
- $l^2 + r^2 > c \to r = r - 1$

## 645. Set Mismatch

在数集中找到重复和缺失的数字，简单题目，时间复杂度和空间复杂度分别可以达到 $O(n)$ 和 $O(1)$

- 简单想法可以使用数组记录数字是否出现过
  - 还可以使用取相反数的方法，使用取值为负的另外一半空间，实现 $O(1)$ 的空间复杂度
- 另外可以使用异或的方法求出重复和缺失的数据，空间复杂度可以达到 $O(1)$

## 664. Strange Printer

- 左边字符和右边字符相等，右边字符不影响：$f[i][j] = f[i][j - 1]$
- 左边字符和右边字符不相等，分段讨论最大值：$min(f[i][k] + f[k+1][j])$

## 671. Second Minimum Node In a Binary Tree

简单题，寻找类似小顶堆的二叉树的第二小节点，直接遍历适当剪枝

## 678. Valid Parenthesis String

带有通配符的括号匹配问题。简单题，可以使用整数记录栈的情况

- 由于 `*` 的匹配和括号的位置相关，匹配为右括号必须在左括号的右边
- 因此 `*` 需要使用两个变量记录两种匹配情况的信息

## 690. Employee Importance

简单遍历，广度优先或深度优先

## 692. Top K Frequent Words

- 哈希统计单词次数，排序或者堆排序得到前 k 个频率最高的单词
- go 中需要自己实现堆排序，直接使用排序可能写起来比较方便

## 704. Binary Search

二分搜索，大多数的语言库可以直接调用

## 725. Split Linked List in Parts

简单题，遍历链表

## 726. Number of Atoms

简单题目，解析化学物质的表达式：使用栈和哈希表计数完成，主要是工作量比较大

- 主要注意在栈中增加数据可以直接在栈里创建匿名对象，无需深复制

## 740. Delete and Earn

简单题：排序，统计每一个数字出现的次数（方便查看对应的数字是什么，可以计算其前缀和），使用动态规划处理，转换为类似于 [213](#213-house-robber-ii) 的题目。

## 752. Open the Lock

打开转盘锁，类似于八数码问题，考虑搜索的方法：

- 生成邻居时特别注意 `(ch - '0' + 9) % 10 + '0'` 的情况
- 广度优先搜索
- 双向广度优先搜索：需要特别注意初始化的地方
- TODO: 启发式搜索

## 773. Sliding Puzzle

八数码问题，棋盘大小为 $2*3$：广度优先搜索或 A* 算法

- 记录切片的信息可以考虑使用 `string`，直接使用切片实在是太麻烦了，特别还要考虑切片的复制机制..
- map 无法映射切片...

<!-- FIXME: 重构中 -->

## 781. Rabbits in Forest

贪心，使用 `answers[i] + 1` 对相同的 `answers[i]` 进行分组

## 783. Minimum Distance Between BST Nodes

中序遍历寻找相邻数字的最小差：

- 实现过程：实现中序遍历、中序遍历完成相邻元素的比较（分开实现，各司其职），特别是用栈实现的方法
- 使用值域以外的值区分第一次遍历
- 使用栈完成中序遍历

## 787. Cheapest Flights Within K Stops

简单中等题，动态规划：复用空间

## 797. All Paths From Source to Target

简单题，得到从起点到终点的所有路径：深度优先搜索

## 802. Find Eventual Safe States

中等题，检测图中的节点出发是否可以达到环：

- 求反图使用拓扑排序去除出度为 0 （即不成环）的节点
- 注意题目要求得到的节点升序

## 810. Chalkboard XOR Game

需要分成数组长度的奇偶性进行讨论：

- $S = \oplus_{i=0}^{n-1} nums[i], S_i = S \oplus nums[i]$
- 若失败，则 $\forall i, S_i = 0$
  - $t = \oplus_{i=0}^{n-1} S_i = \oplus_{i=0}^{n} S = 0$
  - 若 $n$ 为偶数，$t = S = 0$ 与 $S \neq 0$ 矛盾，不会失败
  - A 先手为偶数不会输，B 后手为偶数不会输

## 815. Bus Routes

给定图的邻接关系，求节点之间的最短距离：

- 巴士路线决定了一个等价类，类内成员相互邻接
- 类内成员数据规模较大，主要需要设计数据结构记录邻接信息
  - 考虑存储节点和路线的对应信息，记录路径和路径之间是否相连
  - 使用路线信息进行枚举
  - 使用点进行枚举，需要对最大数据规模的样例进行特判

## 847. Shortest Path Visiting All Nodes

中等题，广度优先搜索，利用一个整数记录当前节点以及状态压缩记录遍历过的节点

## 852. Peak Index in a Mountain Array

二分搜索，数组中寻找唯一的峰值。类似题目：

- [278. First Bad Version](#278-first-bad-version)
- [374. Guess Number Higher or Lower](#374-guess-number-higher-or-lower)

## 863. All Nodes Distance K in Binary Tree

中等题，在二叉树中找到目标节点距离为 k 的节点：

- 使用深度优先搜索，记录父亲节点；使用回溯的距离信息寻找相应距离的子孙结点
- 注意不能走回头路，如果在左节点的方向找到目标节点，应在右节点的方向找距离为 k 的节点，还要特判当前节点是否距离为 k
- 也可以先使用哈希表记录节点的信息

## 872. Leaf-Similar Trees

遍历二叉树的叶子节点进行比较：两棵树同时 dfs 得到叶子节点进行比较

- 关于中序遍历的又一深入理解（有状态的中序遍历）

## 877. Stone Game

两个博弈者轮流在数组首尾选择数字，数字大的人获胜：

- 先手拥有必胜策略，对数组进行奇偶分组，选择奇数组或偶数组
- 如果要求最多能胜出多少分使用动态规划：`dp[i][j] = max(arr[i] - dp[i+1][j], arr[j] - dp[i][j-1])`
  - 可以进行空间优化

## 879. Profitable Schemes

给定任务序列，求在一定的 cost 内达到最小 goal 的方法数。

使用动态规划，转换为背包问题（二维容量）：

- `dp[task][cost][goal] = dp[i-1][cost][goal] + dp[i-1][cost - cost[i]][max(0, goal - goal[i])]`
- 注意上面的式子的种类数为成本正好为 cost 的情况，在一定的 cost 内，需要求 $\sum_c$`dp[n][c][goal]`

## 881. Boats to Save People

弱化的装箱问题，简单题，每个箱子只能装两个物品，使用双指针解决

## 897. Increasing Order Search Tree

中序遍历直接改变指针的指向

- 反问题：线性树 -> 平衡树
- 可以考虑使用旋转解决，比较麻烦

## 909. Snakes and Ladders

走带有传送门的迷宫，广度优先搜索，主要处理迷宫之字形的问题

- 提前将二维转一维处理应该好一些
- 写转换函数也可以

## 930. Binary Subarrays With Sum

计算二进制数组中和为定值的子串数量，对于一个满足条件的局部子串，同样模式的子串数量与首尾的 1 和它们各自的相邻 1 的距离相关：

- `ret += (oneLocs[head] - oneLocs[head - 1]) * (oneLocs[tail + 1] - oneLocs[tail])`
- 需要特别处理 `goal == 0` 的情况：`ret += tmp * (tmp - 1) / 2`

## 938. Range Sum of BST

简单递归，不必过度纠结中序遍历。发现自己会受别人的思路影响，神 tm jm 说了一句中序遍历...

- 也可以使用其它遍历二叉树的方法

## 981. Time Based Key-Value Store

基于时间戳查找对应的键值对，使用二分查找的方法查询

- 注意题目要求找到不大于特定时间戳的最大时间戳对应的键值对，而标准库模板为满足给定函数情况的最小下标
- 使用时间戳大于给定时间戳的函数求下标，然后减 1，需要注意边界

## 987. Vertical Order Traversal of a Binary Tree

中等题，纵向顺序遍历，遍历结果可能和节点的键值相关，需要记录信息再进行排序，不能直接遍历得到结果

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

## 1035. Uncrossed Lines

动态规划，最长相同子序列

## 1047. Remove All Adjacent Duplicates In String

使用栈（string 也可以作为容器）记录移出相邻重复的字符。主要看清题意，每次移出两个字符。

## 1049. Last Stone Weight II

问题转换：两个数相消 -> 对每一个取正负号求和

- 动态规划：背包问题，负数集的和尽可能接近总数的一半

## 1074. Number of Submatrices That Sum to Target

可以将列（行）进行压缩，枚举行（列）的上界和下界，将问题转换为[一维序列的区间和为 k 的问题](#560-subarray-sum-equals-k)

其它类似的问题：363. Max Sum of Rectangle No Larger Than K

<!-- TODO: 考虑做法？二维前缀和 + 哈希表 -->

## 1104. Path In Zigzag Labelled Binary Tree

简单题，经过奇偶反转的二叉树节点定位问题

- 利用二叉树的父节点和子节点的位置关系，在需要反转的位置上进行转换

## 1109. Corporate Flight Bookings

中等题，给定线段和求值：

- 线段问题可以使用线段树解决
- 线段和的问题可以简化使用差分的方式解决

## 1137. N-th Tribonacci Number

3-斐波拉契数列：简单模拟

- 可以使用矩阵运算，使用 $\log n$ 的时间复杂度解决：$\begin{bmatrix}
  0 & 1 & 0 \\
  0 & 0 & 1 \\
  1 & 1 & 1
\end{bmatrix}$

## 1143. Longest Common Subsequence

经典动态规划

$$dp[i][j] = \begin{cases}
  dp[i - 1][j - 1] + 1, & text1[i] == text2[j] \\
  max(dp[i - 1][j], dp[i][j - 1]), & text1[i] != text2[j] \\
\end{cases}$$

## 1190. Reverse Substrings Between Each Pair of Parentheses

遇到括号，逆序遍历，注意括号可能是嵌套的也可能是非嵌套的（题目给的样例是特地的吧 orz...

## 1221. Split a String in Balanced Strings

简单题，整个串是一个 balanced strings，贪心选择长度最短的 balanced strings

## 1239. Maximum Length of a Concatenated String with Unique Characters

- 总共只有 26 个字母，可以使用一个 32 位的整数进行状态压缩
- 使用哈希进行枚举，一次循环动态规划求解问题
- 由于哈希表只需记录有或无，go 中的哈希表可以使用 `struct{}{}` 以节约空间

## 1269. Number of Ways to Stay in the Same Place After Some Steps

动态规划：`dp[step][loc] = dp[step - 1][loc - 1] + dp[step - 1][loc] + dp[step - 1][loc + 1]`

## 1310. XOR Queries of a Subarray

类比为区间和的问题：可以使用前缀异或数组求解

## 1337. The K Weakest Rows in a Matrix

简单题，完成统计后找出数组前 k 小的数：

- 二分查找 0 的位置
- 使用堆完成，堆化时间复杂度 $O(n)$
- 使用快速选择得到前 k 小的数，时间复杂度 $O(n)$

## 1418. Display Table of Food Orders in a Restaurant

简单题 ims，使用哈希表进行计数，最后格式化输出

## 1442. Count Triplets That Can Form Two Arrays of Equal XOR

区间问题转换为前缀问题，通过下面的性质进行问题转换使用二重循环可以解决：

- $\oplus_{idx=i}^{j-1} arr[idx] = \oplus_{idx=j}^{k}arr[idx]$
- $\oplus_{idx=i}^k arr[idx] = 0$
- $\oplus_{idx=0}^{i-1} arr[idx] = \oplus_{idx=0}^k arr[idx]$
  - 考虑构造 `preXors[i] = ^arr[0..i-1]`
  - 可以使用哈希进行加速，一重循环解决
  - `ret += cnt[preXor] * idx + sumIdx[preXor]`

## 1449. Form Largest Integer With Digits That Add up to Target

完全背包问题：不同的数字对应各自的 `cost`，求总 `cost` 为 `target` 时，最大的数

- 注意到固定的 cost 对应的是每一个数字选取的个数固定
- 为了避免对字符串修改要使用循环进行增量修改，可以提前计算好各个数字的位数
  - 而非传统完全背包问题，直接枚举不同数字的位数
- 即最大的数字可以让所有数字按从大到小依次排列对应的次数

## 1473. Paint House III

困难动态规划，其实也不是很难，三维动态规划：

- 考虑和第 i-1 间房子进行比较
  - 若 $c_i = c_{i-1} = c$，则 $dp[i][c][neiNum] = dp[i-1][c][neiNum]$
  - 若 $c_i \neq c_{i-1}$，则 $dp[i][c_i][neiNum] = dp[i-1][c_{i-1}][neiNum - 1]$
- 未涂色，加 `cost[i][c]`

注意初始化：`i == 0` 和 `neiNum == 0` 的情况

- 可以使用滚动数组节省空间
<!-- TODO: 更高级的做法，日后再看看 -->

## 1480. Running Sum of 1d Array

简单题，原地计算前缀和

## 1482. Minimum Number of Days to Make m Bouquets

- 求解问题转换为判断问题，使用二分法进行求解
- go 中可以使用 `sort.Search` 函数实现二分查找
- 一些性能优化：排序后进行二分查找，时间复杂度减小，但仍不能弥补排序的开销，只能说数据出的...

## 1486. XOR Operation in an Array

简单模拟，初始值可以设为 0（`0 ^ num = num`）

看这种出题方法也可以估计出存在 $O(1)$ 时间复杂度的解法，利用异或的性质：

- 连续数字进行异或可以使用的性质：$4i \oplus (4i+1) \oplus (4i+2) \oplus (4i+3) = 0$
- 对于线性变化的数字，若倍数是 2 次幂可以对低位数字提出来特别讨论

## 1583. Count Unhappy Friends

简单模拟，使用哈希（顺序数组）记录配对信息和优先度信息

## 1588. Sum of All Odd Length Subarrays

中等题，考察元素在长度为奇数的子串中出现的次数

- 分成左右两端和两种情况考虑：奇+1+奇，偶+1+偶

## 1600. Throne Inheritance

简单模拟，多叉树的前序遍历

## 1603. Design Parking System

简单模拟

## 1646. Get Maximum in Generated Array

简单题，模拟

## 1707. Maximum XOR With an Element From Array

通过前缀树记录数组元素的每一位，寻找最大的异或值

- 对于每一个节点，记录子树对应的最小值，确保可以进行转换

## 1711. Count Good Meals

简单题目，给定数组挑选两个数，和为二的幂。枚举所有可能的幂 ($pow < \max num * 2$)，使用哈希表判断满足条件的数字个数

- 判断二的幂：[231. Power of Two](#231-power-of-two)

## 1713. Minimum Operations to Make a Subsequence

求两个数组的最长公共序列长度，将数组的元素转换为下标，求最长递增子序列长度

- 可以映射为目标数组的下标，也可以映射为原数组的下标

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

## 1736. Latest Time by Replacing Hidden Digits

简单题，暴力打表

## 1738. Find Kth Largest XOR Coordinate Value

二维前缀异或 + 快速选择，二维前缀的构造类比二维前缀和

- go 里面传递数组好像是写时复制？递归传递数组比较消耗时间
- 尽量将递归改成循环

## 1743. Restore the Array From Adjacent Pairs

简单题，用邻接数对还原数组：抓头抓尾

- 每个元素唯一比较简单；存在重复元素使用回溯的方法，双向遍历？

## 1744. Can You Eat Your Favorite Candy on Your Favorite Day?

前缀和，确定吃到特定种类的糖，需要吃多少颗糖

- 判断特定天数能否的吃糖区间是否满足条件

## 1787. Make the XOR of All Segments Equal to Zero

对于每一个长度为 $k$ 的区间异或为 $0$，则 $nums[i] == nums[i+k]$。因此，数组满足的性质为：

- 数组以 $k$ 为周期
- 数组的前 $k$ 个数的异或为 0

将数组分成 k 组，考虑不同的异或前缀路径，将 k 个数异或结果变为 0

- $dp[mask] = min(prev[mask \oplus num] - cnts[num], min(prev))$
- 主要需要遍历每一组的数据需要变成哪一个数字，分成两类进行讨论：
  - 将所有数字变成没有出现过的数字
  - 将所有数字变成一组中的一个，变化的次数为 `size - cnts[num]`

## 1818. Minimum Absolute Sum Difference

使用数组中的一个数字替换另一个数字，使得该数组的一范数尽可能接近另一个数组

- 使用二分法进行枚举，时间复杂度 $O(n \log n)$

## 1833. Maximum Ice Cream Bars

简单题，给定预算和物品的成本、数量，求可以得到的最大物品数量：排序贪心，时间复杂度 $O(n \log n)$

- 可以使用堆进行优化，堆化的时间复杂度为 $O(n)$

## 1838. Frequency of the Most Frequent Element

给出一定的补偿，求将数组进行补偿后可能的最大频数

- 滑动窗口求解，除去排序的时间，时间复杂度为 $O(n)$，全部时间复杂度为 $O(n \log n)$
- 也可以先计数，然后排序加滑动窗口求解

## 1846. Maximum Element After Decreasing and Rearranging

不知怎么描述，反正就简单题

## 1877. Minimize Maximum Pair Sum in Array

简单题，求数组中最小的最大数对和

## 1893. Check if All the Integers in a Range Are Covered

判断范围是否被线段覆盖：

- 使用贪心的方法，对线段进行排序，更新左端点，时间复杂度 $O(n \log n)$
- 差分的思想，记录覆盖的信息，使用前缀和可以得到覆盖的数量，时间复杂度 $O(n + l)$
- 线段树的做法

## 1894. Find the Student that Will Replace the Chalk

简单题，前缀和加二分搜索

## 剑指 Offer 10- I. 斐波那契数列

简单模拟

## 剑指 Offer 15. 二进制中1的个数

计算海明距离，go 使用 `bits.OnesCount` 函数

- [461. Hamming Distance](#461-hamming-distance)
- [477. Total Hamming Distance](#477-total-hamming-distance)

## 剑指 Offer 22. 链表中倒数第k个节点

简单题：使用双指针可以只遍历一次

## 剑指 Offer 37. 序列化二叉树

相同的题目：[297. Serialize and Deserialize Binary Tree](#297-serialize-and-deserialize-binary-tree)

## 剑指 Offer 38. 字符串的排列

获得所有的排列组合，获得最小的排列，然后依次遍历增加下一个递增的排列

- [31. Next Permutation](#31-next-permutation)

## 剑指 Offer 42. 连续子数组的最大和

- [53. Maximum Subarray](#53-maximum-subarray)

## 剑指 Offer 52. 两个链表的第一个公共节点

- [160. Intersection of Two Linked Lists](#160-intersection-of-two-linked-lists)

## 剑指 Offer 53 - I. 在排序数组中查找数字 I

- [34. Find First and Last Position of Element in Sorted Array](#34-find-first-and-last-position-of-element-in-sorted-array)

## 面试题 10.02. Group Anagrams LCCI

简单题，使用字母出现的次数进行哈希，go 中哈希表只能通过数组进行映射，不能通过切片映射？

## 面试题 17.10. Find Majority Element LCCI

思维训练简单题，寻找数组中的元素个数大于数组元素个数一半的元素

- 即 `cnt(num) > len(nums) - cnt(num)`
- 可以使用两个变量，一个变量记录当前认为的主要元素，另一个变量记录主要元素被部分其他元素抵消后出现的次数

## 面试题 17.21. Volume of Histogram LCCI

`ans[i] += min(lmax[i], rmax[i]) - h[i]`

需要注意的点：

- `lmax`/`rmax` 单调递增/递减
- 若 `lmax[i] > rmax[j]` 则，$\exist nj < j, rmax[nj] \geq lmax[i]$，可以适用双指针

## LCP 07. 传递信息

给定有向图，求经过 k 条边，从起点到终点的方案数

- 简单动态规划：`newdp[dest] += dp[src]`，时间复杂度 $O(kn)$
