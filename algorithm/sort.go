package algorithm

// 冒泡排序（排序10000个随机整数，用时约145ms）
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums) - i; j++ {
			if nums[j] < nums[j - 1] {
				nums[j], nums[j - 1] = nums[j - 1], nums[j]
			}
		}
	}
}

//选择排序（排序10000个随机整数，用时约45ms）
func SelectSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		maxIndex := 0
		//寻找最大的一个数，保存索引值
		for j := 1; j < length - i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[length - i - 1], nums[maxIndex] = nums[maxIndex], nums[length - i - 1]
	}
}

//快速排序（排序10000个随机整数，用时约0.9ms）
func QuickSort(nums []int) {
	recursionSort(nums, 0, len(nums) - 1)
}

func recursionSort(nums []int, left int, right int) {
	if left < right {
		pivot := partition(nums, left, right)
		recursionSort(nums, left, pivot - 1)
		recursionSort(nums, pivot + 1, right)
	}
}

func partition(nums []int, left int, right int) int {
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

//插入排序（排序10000个整数，用时约30ms）
func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i - 1] {
			j := i - 1
			temp := nums[i]
			for j >= 0 && nums[j] > temp {
				nums[j + 1] = nums[j]
				j--
			}
			nums[j + 1] = temp
		}
	}
}
