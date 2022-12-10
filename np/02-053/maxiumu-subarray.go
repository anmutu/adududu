package main

import "fmt"

//-2,         1,           -3,                4,            -1,            2,             1,          -5,              4
//-2,     -2+1,1=>1    1+(-3),-3=> -2    4+(-2),4=>4    -1+(4),-1=>3     3+2,2=>5      5+1,1=>6   6+(-5),-5=>1      1+4，4=>5
//难点：找关系。
//第一个位置，就是取第一个位置的值。
//从第二个数开始就是取[自己]和[自己与其前面的数相加的结果]中的较大值。
func main(){
	nums:=[]int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubarray(nums))
	nums1:=[]int{1}
	fmt.Println(maxSubarray(nums1))
}

func maxSubarray(nums []int)int{
	var res=nums[0]
	//1.定义容器。里面的每个元素是到这个元素的最大的子数和的值。
	var n=len(nums)
	dp:=make([]int,n)
	//2.找到初始值。
	dp[0]=nums[0]
	//3，根据状态转移方程写相关代码。
	for i:=1;i<n;i++{
		dp[i]=max(nums[i],nums[i]+dp[i-1])
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

