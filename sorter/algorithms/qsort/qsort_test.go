package qsort

import (
        "testing"
       )

func sliceEqual(s1,s2 []int)bool {
    if len(s1) != len(s2) {
        return false
    }
    for i, e := range(s1) {
        if e != s2[i] {
            return false
        }
    }
    return true
}

func QsortValidate(t *testing.T) {
    var tests = []struct {
        input       []int
        output      []int
        equal       bool
    }{
        { []int{1,2,3,4,5}, []int{1,2,3,4,5}, true},
        { []int{1,2,3,5,4}, []int{1,2,3,4,5}, true},
    }

    for _, test := range(tests) {
        Qsort(test.input)
        if sliceEqual(test.input, test.output) != test.equal {
            t.Errorf("faile to quick sort")
        }
    }
}

