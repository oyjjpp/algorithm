### 题目

50. Pow(x, n)

### 题目描述

实现Pow(x,n)即计算x的n次幂函数（即，x^n）。

### 案例

#### 示例一

```golang
输入： x = 2.00000, n = 10
输出： 1024.00000
```

#### 示例二

```golang
输入： x = 2.10000, n = 3
输出： 9.26100
```

#### 示例三

```golang
输入： x = 2.00000, n = -2
输出： 0.25000
解释： 2-2 = 1/22 = 1/4 = 0.25
```

### 思路

### 代码

```golang
// moveZeroes
// 283. 移动零
func moveZeroes(nums []int) {
 left, right, n := 0, 0, len(nums)
 for right < n {
  if nums[right] != 0 {
   nums[left], nums[right] = nums[right], nums[left]
   left++
  }
  right++
 }
}
```

### 参考

来源：力扣（LeetCode）  
链接：<https://leetcode-cn.com/problems/powx-n>  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。  
