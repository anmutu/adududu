package main

import "fmt"

func main(){
	var grid =[][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	//1.定义容器。就用原来的grid.

	//2.找到初始值。
	m:=len(grid)
	for i:=1;i<m;i++{//第一列。
		grid[i][0]+=grid[i-1][0]
	}
	fmt.Println(grid)

	n:=len(grid[0])
	for j:=1;j<n;j++{//第一行
		grid[0][j]+=grid[0][j-1]
	}
	fmt.Println(grid)

	//3.根据状态转移方程写代码。
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			grid[i][j]+=min(grid[i-1][j],grid[i][j-1])
		}
	}
	fmt.Println(grid)
	
	return grid[m-1][n-1]
}

func min(x,y int)int{
	if x<y{
		return x
	}else{
		return y
	}
}

