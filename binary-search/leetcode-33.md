# 搜索旋转排序数组

## 题目

整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。


## 总结
这题主要分析的是二分查找选边的问题，找到中间那个数的值很简单，怎么选择那边去找到目标结果很关键

本题中left这个位置的值，一定可以划分将数组划分成两个数组，必然存在一个数组严格递增。只要判断目标值是否在这个严格递增的数组中即可