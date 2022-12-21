package main

import "fmt"
// import "sort"

func main(){
	//  coins:=[]int{1,2,5}
	//  amount:=11
	// fmt.Println(coinChange(coins,amount))

	// coins1:=[]int{1,2,5,10}
	// amount1:=27
	// fmt.Println(coinChange(coins1,amount1))

	// coins2:=[]int{2,5,10,1}
	// amount2:=27
	// fmt.Println(coinChange(coins2,amount2))

	coins:=[]int{1,2,5}
	amount:=11
    fmt.Println(coinChange(coins,amount))
}

func coinChange(coins []int, amount int) int {
	//1.定义容器。
	n:=amount+1
	dp:=make([]int,amount+1)

	//2.初始化。
	dp[0]=0

	//3.根据状态转移方程写代码。 
	for i:=1;i<n;i++{
		dp[i]=-1
		fmt.Println("*****")
		fmt.Println(dp)
		for j:=0;j<len(coins);j++{
			if coins[j]<=i{
				if dp[i]==-1||dp[i]>dp[i-coins[j]]+1{
					dp[i]=dp[i-coins[j]]+1
				}
			}
		}
		fmt.Println(dp)
		fmt.Println("*****")
	}
	// fmt.Println(dp)
	return dp[amount]
}
