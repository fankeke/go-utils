package  sort


func partition(nums []int, left, right int) int{
    pivot := nums[left]
    for left < right {
        for left < right && nums[right] >= pivot {
            right--
        }
        nums[left]=nums[right]
        for left < right && nums[left] < pivot {
            left++
        }
        nums[right]=nums[left]
    }
    nums[left]=pivot
    return left
}    


func quickSort(nums []int, left, right int) {
    if left < right {
        pos := partition(nums, left, right)
        quickSort(nums, left, pos-1)
        quickSort(nums, pos+1, right)
    }
}    

func Sort(nums []int) {
    quickSort(nums, 0, len(nums)-1)
}



////find k th max elemnt

func find(nums []int, left, right, k int) int{
    if left <= right {    //和全排序不一样的是，此处如果只剩一个元素，也许要进入流程
        pos := partition(nums, left, right)
        if pos == k {
            return nums[k]
        } else if pos < k {
            return find(nums, pos+1, right, k)
        }else {
            return find(nums, left, pos-1, k)
        }
    }
    return -1
}

func FindKthMax(nums []int, k int) int {
    return find(nums, 0, len(nums)-1, len(nums)-k)
}

func FindKthMin(nums []int, k int) int {
    return find(nums, 0, len(nums)-1, k-1)
}
