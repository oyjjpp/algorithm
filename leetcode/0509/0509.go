package leetcode

// fib
// 斐波那契数列
func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fib(N-1) + fib(N-2)
}

// fibV2
// 斐波那契数列 通过备忘录解决重叠子问题
// 自顶向下计算的实现
// 从一个规模较大的原问题比如说f(20)，向下逐渐分解规模，直到f(1)和f(2) 这两个base case
func fibV2(N int) int {
	data := map[int]int{
		0: 0,
		1: 1,
	}

	// 使用递归思想解决问题
	var fib func(n int) int
	fib = func(n int) int {
		if rs, ok := data[n]; ok {
			return rs
		}
		if n == 0 || n == 1 {
			return 1
		}
		data[n] = fib(n-1) + fib(n-2)
		return data[n]
	}
	return fib(N)
}

// fibV3
// 斐波那契数列 通过DP table解决重叠子问题
// 自底向上计算的实现
// 我们直接从最底下，最简单，问题规模最小的f(1)和f(2)开始往上推，直到推到我们想要的答案f(20)，
// 这就是动态规划的思路，这也是为什么动态规划一般都脱离了递归，而是由循环迭代完成计算
func fibV3(N int) int {
	data := map[int]int{
		0: 0,
		1: 1,
	}
	if _, ok := data[N]; ok {
		return data[N]
	}
	for i := 2; i <= N; i++ {
		data[i] = data[i-1] + data[i-2]
	}
	return data[N]
}

// fibV3
// 斐波那契数列 通过DP table解决重叠子问题
// 自底向上计算的实现
// 我们直接从最底下，最简单，问题规模最小的f(1)和f(2)开始往上推，直到推到我们想要的答案f(20)，
// 算法中状态转移f(n) = f(n-1)+f(n-2)，计算中仅需要使用两个状态，上面都是使用map保存了所有状态
// 此类问题我们可以通过“状态压缩”，减少内存的浪费
func fibV4(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return 1
	}
	// 状态压缩，当前算法仅需要前两个值，不需要将所有状态保存下来，可以节省内存空间
	pre, cur := 1, 1
	for i := 3; i <= N; i++ {
		sum := pre + cur
		pre = cur
		cur = sum
	}
	return cur
}
