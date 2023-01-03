package main

import "fmt"

func main(){
	coins:=[]int{1,2,5}
	amount:=11
    fmt.Println(coinChange(coins,amount))

	// coins1 :=[]int{2}
	// amount1:=3
	// fmt.Println(coinChange(coins1,amount1))
}

func coinChange(coins []int, amount int) int {
	//1.定义容器。
	dp:=make([]int,amount+1)
	for i:=0;i<len(dp);i++{
		dp[i]=amount+1
	}

	//2.找出初始值。
	dp[0]=0

	//3.根据状态转移方程写代码。
	for i:=1;i<=amount;i++{
		fmt.Println("*************************start******",i)
		for j:=0;j<len(coins);j++{
			if coins[j]<=i{
				fmt.Println("dp[i-coins[j]]+1=>",dp[i-coins[j]]+1)
				dp[i]=min(dp[i],dp[i-coins[j]]+1)
			}
		}
		fmt.Println("*************************end*******",i)
	}
	fmt.Println(dp)
	if dp[amount]>amount{
		return -1
	}
	return dp[amount]
}

func min(x,y int)int{
	if x<y{
		return x
	}
	return y
}
