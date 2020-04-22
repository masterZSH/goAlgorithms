package dp

// 价格列表  即长度0时候价格为0，长度为1的时候价格为1...
var price []int = []int{0, 1, 5, 8, 9, 17, 17, 17, 20, 24, 30}

//						0  1  2  3  4  5   6   7   8   9   10

// 返回 a,b中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 递归
func recursivelyCut(len int) int {
	if len == 0 {
		return 0
	}
	r := 0
	for i := 1; i < len; i++ {
		r = max(r, price[i]+recursivelyCut(len-i))
	}
	return r
}

// 动态规划
func recursivelyCutDp(len int, res []int) int {
	// 数组中有了就返回
	if res[len-1] > 0 {
		return res[len-1]
	}
	var r int
	if len == 0 {
		r = 0
	} else {
		r = 0
		for i := 1; i < len; i++ {
			r = max(r, price[i]+recursivelyCutDp(len-i, res))
		}
	}
	res[len-1] = r
	return r
}
