# x 的平方根

## 题目

编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

- 每行中的整数从左到右按升序排列。
- 每行的第一个整数大于前一行的最后一个整数。


## 总结
这题的话，主要有三种方法。
1. 可以使用两次二分查找，首先使用二分查找找到对应的行，然后再用二分查找找对应的位置。
2. 一次二分查找，若将矩阵每一行拼接在上一行的末尾，则会得到一个升序数组，我们可以在该数组上二分找到目标元素。 代码实现时，可以二分升序数组的下标，将其映射到原矩阵的行和列上。
3. 可以使用顺序查找，时间复杂度o(m+n)