package main

import "fmt"

//1 1 1 1
//1 2 3 4
//1 3 6 10
func main(){
	fmt.Println(uniquePaths(3,7))
}

func uniquePaths(m,n int)int{
	//1.定义容器。先行再列。
	dp:=make([][]int,m)
	fmt.Println(dp)
	for i:=0;i<m;i++{
		dp[i]=make([]int,n)
	}
	fmt.Println(dp)

	//2.找初始值。
	//3.根据状态转移公式写代码。
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if i==0||j==0{
			  dp[i][j]=1
			  continue
			}
			dp[i][j]=dp[i-1][j]+dp[i][j-1]
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}