package main

import "fmt"

//1 1 1 1
//1 2 3 4
//1 3 6 10
func main(){
	fmt.Println(uniquePaths(3,7))
}

func uniquePaths(m,n int)int{
	//1.定义容器。
	dp:=make([][]int,m)
	for i:=0;i<m;i++{
		dp[i]=make([]int,n)
	}
	fmt.Println(dp)

	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			//2.找到初始值。
			if i==0||j==0{
				dp[i][j]=1
				continue
			}
			//3.根据状态转移方程写代码。
			dp[i][j]=dp[i-1][j]+dp[i][j-1]
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}