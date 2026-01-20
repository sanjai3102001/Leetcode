package main

import(
	"fmt"
)

func main(){
	nums :=[]int{1,2,3,4}
	n:=len(nums)
	result:=0
	for i:=1;i<=len(nums);i++{
		if n%i==0{
			fmt.Println("result is ", result, "and the nums is", nums[i-1])
			result=nums[i-1]*nums[i-1]+result
		}
	}
	fmt.Println(result)
}