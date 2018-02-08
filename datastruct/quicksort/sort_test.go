package sort

import ("testing")

func equal(a []int, b []int)bool{
    if len(a) != len(b) {
        return false
    }
    for i := range(a) {
        if b[i] != a[i] {
            return false
        }
    }
    return true
}

func TestSortSanity(t *testing.T) {
    var tests = []struct {
        input []int
        output []int
    }{
        { []int{1,2,3,5,4}, []int{1,2,3,4,5}},
        { []int{1}, []int{1}},
        { []int{1,2,9,5,4}, []int{1,2,4,5,9}},
    }

    for _, test := range(tests) {
        input := test.input
        Sort(input)
        if !equal(input, test.output) {
            t.Errorf("failed to fullsort")
        }
    }
}

func TestKthMaxSanity(t *testing.T) {
    var tests = []struct {
        input []int
        k       int
        res     int
    }{
        { []int{1,2,3,5,4}, 3, 3},
        { []int{1},1, 1},
        { []int{1,2,9,5,4}, 2, 5},
    }

    for _, test := range(tests) {
        input := test.input
        got := FindKthMax(input, test.k)
        if got != test.res {
            t.Errorf("failed to findKthMaxsort")
        }
    }
}

func TestKthMinSanity(t *testing.T) {
    var tests = []struct {
        input []int
        k       int
        res     int
    }{
        { []int{1,2,3,5,4}, 3, 3},
        { []int{1},1, 1},
        { []int{1,2,9,5,4}, 2, 2},
    }

    for _, test := range(tests) {
        input := test.input
        got := FindKthMin(input, test.k)
        if got != test.res {
            t.Errorf("failed to findKthMinsort")
        }
    }
}


