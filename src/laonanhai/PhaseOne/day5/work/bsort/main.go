package main

import (
	"fmt"
)

func bsort(n []int)  {
	//需要冒泡的 ishu5-8 len(a)-1
	for i :=0;i<len(n)-1;i++ {
		//从10开始，交换的测试10-4
		for j :=1;j<len(n)-i;j++{
			if(n[j]<n[j-1]){
				n[j],n[j-1] = n[j-1],n[j]
			}
		}

	}
}

func ssort(n []int)  {
	//需要冒泡的 ishu5-8 len(a)-1
	for i :=0;i<len(n)-1;i++ {
		//从10开始，交换的测试10-4
		for j :=i+1;j<len(n);j++{
			if(n[i]<n[j]){
				n[i],n[j] = n[j],n[i]
			}
		}

	}
}
func isort(a []int)  {
	for i:=1;i<len(a);i++{
		for j :=i;j>0 ;j--  {
			if a[j]>a[j-1]{
				break
			}
			a[j],a[j-1] =a[j-1],a[j]
		}
	}
}
func PartSort()  {
	
}
func main()  {
	//b := [...]int{8,2,3,6,7,8,9}
	b := [...]int{5,10,3,6,7,8,4}
	//bsort(b[:])
	ssort(b[:])
	fmt.Println(b)
}
