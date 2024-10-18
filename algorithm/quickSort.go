package algorithm


import (
	
)

func QuickSort(nums []int, left, right int) {
	if left < right{
		mid := left+(right-left)/2
		p := Partition(nums, left, right, mid)
		QuickSort(nums, left, p-1)
		QuickSort(nums, right, p+1)	
	}
}

func Partition(nums []int, left, right, mid int) int{
	pivot := nums[mid]
	nums[mid], nums[right] = nums[right], nums[mid]
	temp := left
	for left < right{
		if nums[temp] < pivot{
			nums[left], nums[temp] = nums[temp], nums[left]
			temp++
		}
		left++
	}
	nums[right], nums[temp] = nums[temp], nums[right]
	return temp
}