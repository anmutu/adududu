package main

import "fmt"

//斐波那契数列
//1,2,3,4,5,6,7
//1,1,2,3,5,8,13,21,35
//规律如上，求第n个数的值。

func main(){
  fmt.Println(fib(1))
  fmt.Println(fib(7))
}


func fib(n int) int{
	if n==1||n==2{
		return 1
	}
	//1.定义容器。
	dp:=make([]int,n)
	//2.找到初始值。
	dp[0],dp[1]=1,1
	//3.根据状态转移公式写代码。
	for i:=2;i<n;i++{
		dp[i]=dp[i-2]+dp[i-1]
	}
	return dp[n-1]
}