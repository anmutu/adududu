package main

import "fmt"

func main(){
	var nums=[]int{2,7,9,3,1}
	fmt.Println(rob(nums))
	var nums1=[]int{2,7}
	fmt.Println(rob(nums1))
}

// 2,7,       9,         3,              1
// 2,7,    9+2,7=>11   3+2,11=>11     1+11,11=>12
//dp[i]=max(nums[i]+dp[i-2],dp[i-1])
func rob(nums []int) int {
	var res=nums[0]
	if len(nums)==1{
		return nums[0]
	}
	if len(nums)==2{
		return max(nums[0],nums[1])
	}

	//1.定义容器。
	n:=len(nums)
	dp:=make([]int,n)

	//2.找到初始值。
	dp[0],dp[1]=nums[0],max(nums[0],nums[1])

	//3.根据状态转移方程写代码。
	for i:=2;i<n;i++{
		dp[i]=max(nums[i]+dp[i-2],dp[i-1])
		if dp[i]>res{
			res=dp[i]
		}
	}
	return res
}

func max(x,y int)int{
	if x>y{
		return x
	}else{
		return y
	}
}

