package main

import "fmt"

func main(){
	 coins:=[]int{1,2,5}
	 amount:=11
	fmt.Println(coinChange(coins,amount))
}

func coinChange(coins []int, amount int) int {
	//1.定义容器。
	n:=amount+1
	dp:=make([]int,amount+1)
	fmt.Println(dp)

	//2.初始化。
	dp[0]=0

	//3.根据状态转移方程写代码。 
	for i:=1;i<n;i++{
		dp[i]=-1
		for j:=0;j<len(coins);j++{
			if coins[j]<=i&&dp[i-coins[j]]!=-1{
				dp[i]=dp[i-coins[j]]+1
			}
		}
	}
	fmt.Println(dp)
	return dp[amount]
}
