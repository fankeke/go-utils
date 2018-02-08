package bubblesort

import (
        "testing"
)

func sliceEqual(sli1, sli2 []int) bool {
    if len(sli1) != len(sli2) {
        return false
    }

    for i, e := range(sli1) {
        if e != sli2[i] {
            return false
        }
    }
    return true
}

func TestBubbleSort1(t *testing.T) {
    var tests = []struct {
        input       []int
        output      []int
        equal       bool
    }{
        {[]int{1,2,3,3,3}, []int{1,2,3,3,3}, true},
        {[]int{5,4,3,2,1}, []int{1,2,3,4,5}, true},
    }

    for _, test := range(tests) {
        BubbleSort(test.input)
        if sliceEqual(test.input, test.output) != test.equal {
            t.Errorf("Failed to sort!")
        }
    }
}    

    

