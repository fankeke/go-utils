package qsort


func partition(values []int, left, right int)int {
    pivot := values[left]
    for left < right {
        for left < right && values[right] >= pivot {
            right--
        }

        values[left]=values[right]
         
        for left < right && values[left] <= pivot {
            left++
        }

        values[right] = values[left]
    }
    values[left] = pivot
    return left
}

func quicksort(values []int, left, right int ){
    if left < right {
        pos := partition(values, left, right) 
        quicksort(values, left, pos-1)
        quicksort(values, pos+1, right)
    }
}

func Qsort(values []int){
    quicksort(values, 0, len(values)-1)
}
