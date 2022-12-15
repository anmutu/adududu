package main

import "fmt"

func main(){
	var nums=[]int{2,7,9,3,1}
	fmt.Println(rob(nums))
}

// 2,7,       9,         3,              1
// 2,7,    9+2,7=>11   3+2,11=>11     1+11,11=>12
//dp[i]=max(nums[i]+dp[i-2],dp[i-1])
func rob(nums []int) int {
	
}