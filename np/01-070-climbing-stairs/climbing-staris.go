package main

import "fmt"
//1,2,3,4,5
//1,2,3,5,8

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	if n==1{
		return 1
	}
	if n==2{
		return 2
	}
	//1.定义容器
	dp:=make([]int,n)
	//2.初始化值。
	dp[0],dp[1]=1,2
	//3.根据状态转移方程写相关代码。
	for i:=2;i<n;i++{
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n-1]
}
