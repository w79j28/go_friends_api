package util

func QuickSort(nums []int64) {
	recursionSort(nums, 0, int64(len(nums)-1))
}

func recursionSort(nums []int64, left int64, right int64) {
	if left < right {
		pivot := partition(nums, left, right)
		recursionSort(nums, left, pivot-1)
		recursionSort(nums, pivot+1, right)
	}
}

func partition(nums []int64, left int64, right int64) int64 {
	for left < right {
		for left < right && nums[left] <= nums[right] {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		for left < right && nums[left] <= nums[right] {
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}
	return left
}
