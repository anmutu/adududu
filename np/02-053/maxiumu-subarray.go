package main

import "fmt"

//-2,1,-3,4,-1,2,1,-5,4

//难点：找出关系。
//创建一个数组，里面装的是到它这个位置的最大值是多少。
//1.对于第一个元素来说就是它自己了。
//2.对于第二个元素来说就是[第一个元素+自己]和[自己]比较大的值了。
//-2,            1,           -3,              4,        -1,              2,      1,     -5,    4
//-2,      -2+1，1,=>1    1+(-3),-3=>-2     -2+4,4=>4  4+(-1),-1=>3        5      6      1     5
func main(){
	nums:=[]int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubarray(nums))
	nums1:=[]int{1}//返回值要定义为nums[0],不然是会不通过的。
	fmt.Println(maxSubarray(nums1))
}

func maxSubarray(nums []int)int{
	var res=nums[0]
	
	//1.定义容器
	n:=len(nums)
	dp:=make([]int,n)

	//2.找到初始值
	dp[0]=nums[0]
	// dp[1]=max(nums[0],nums[1])

	//3.状态转移方程
	for i:=1;i<n;i++{
		dp[i]=max(nums[i]+dp[i-1],nums[i])
		if res<dp[i]{
			res=dp[i]
			fmt.Println(res)
		}
		fmt.Println(dp)
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