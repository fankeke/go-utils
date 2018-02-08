package list_test

import (
        "testing"
        . "../list"
       )


func TestSanity(t *testing.T) {
    var tests = []struct {
        input1 []int
        input2 []int
        equal   bool
    }{
        { []int{1,2,3,4,5}, []int{5,4,3,2,1}, true},
        { []int{1,2}, []int{2,1}, true},
        { []int{1}, []int{2,1}, false},
    }

    for _, test := range(tests) {
        l1 := CreateList(test.input1)
        l1r := l1.Reveser()

        l2 := CreateList(test.input2)

        if got := EqualList(l1r, l2); got != test.equal {
            t.Errorf("faild to test reverse")
        }
    }
}
func TestReverseCreateList(t *testing.T){
    var tests = []struct {
        input1  []int
        input2  []int
        got     bool
    }{
        { []int{1,2,3,4,5}, []int{5,4,3,2,1}, true},
        { []int{1,2,4,3,5}, []int{5,3,4,2,1}, true},
        { []int{1,2}, []int{2,1}, true},
        { []int{1}, []int{1}, true},
    }

    for _, test := range(tests) {
        l1 := CreateReverseList(test.input1)
        l2 := CreateList(test.input2)
        if got := EqualList(l1, l2); got != test.got {
            t.Errorf("failed to test CreateListReveres")
        }
    }

}

func TestRemoveKth(t *testing.T) {
    var tests = []struct {
        input1 []int
        input2 []int
        k        int
        equal   bool
    }{
        { []int{1,2,3,4,5}, []int{1,2,3,4}, 1, true},
        { []int{1,2}, []int{2}, 2, true},
        { []int{1}, []int{2,1}, 1, false},
    }

    for _, test := range(tests) {
        l1 := CreateList(test.input1)
        l2 := CreateList(test.input2)

        l1R := l1.RemoveKthFromEnd(test.k)

        if got := EqualList(l1R, l2); got != test.equal {
            t.Errorf("faild to test remove kth from end")
        }
    }
}

func TestMergeSortListh(t *testing.T) {
    var tests = []struct {
        input1 []int
        input2 []int
        input3 []int
    }{
        { []int{1,2,3,4,5}, []int{1,2,3,4}, []int{1,1,2,2,3,3,4,4,5}},
        { []int{1,2}, []int{2}, []int{1,2,2}},
        { []int{1}, []int{1}, []int{1,1}},
    }

    for _, test := range(tests) {
        l1 := CreateList(test.input1)
        l2 := CreateList(test.input2)
        l3 := CreateList(test.input3)

        lm := MergeTwoSortList(l1, l2)

        if got := EqualList(l3, lm); got != true {
            lm.Println()
            l3.Println()
            t.Errorf("faild to test mergesortlist")
        }
    }
}


func TestSortListh(t *testing.T) {
    var tests = []struct {
        input1 []int
        input2 []int
    }{
        { []int{1,2,4,5,3}, []int{1,2,3,4, 5}},
        { []int{2,1}, []int{1, 2}},
        { []int{1}, []int{1}},
    }

    for _, test := range(tests) {
        l1 := CreateList(test.input1)
        l2 := CreateList(test.input2)

        lr := l1.MergeSortList()

        if got := EqualList(l2, lr); got != true {
            lr.Println()
            l2.Println()
            t.Errorf("faild to test sortlist")
        }
    }
}

func TestDeleteDuplicates(t *testing.T) {
    var tests = []struct {
        input1 []int
        input2 []int
    }{
        { []int{1,2,3,3,4}, []int{1,2,3,4}},
        { []int{1,2,3,4,4}, []int{1,2,3,4}},
        { []int{1,1}, []int{1}},
        { []int{1}, []int{1}},
    }

    for _, test := range(tests) {
        l1 := CreateList(test.input1)
        l2 := CreateList(test.input2)

        lr := l1.DeleteDuplicates()

        if got := EqualList(l2, lr); got != true {
            lr.Println()
            l2.Println()
            t.Errorf("faild to test deleteduplicateslist")
        }
    }
}

func TestDeleteDuplicatesAll(t *testing.T) {
    var tests = []struct {
        input1 []int
        input2 []int
    }{
        { []int{1,2,3,3,4}, []int{1,2,4}},
        { []int{1,2,3,4,4}, []int{1,2,3}},
        { []int{1,1}, []int{}},
        { []int{1,1,2,2}, []int{}},
        { []int{1}, []int{1}},
    }

    for _, test := range(tests) {
        l1 := CreateList(test.input1)
        l2 := CreateList(test.input2)

        lr := l1.DeleteDuplicatesAll()

        if got := EqualList(l2, lr); got != true {
            lr.Println()
            l2.Println()
            t.Errorf("faild to test deleteduplicatesalllist")
        }
    }
}

func TestIsPalindrome(t *testing.T) {
    var tests = []struct {
        input1 []int
        res     bool
    }{
        { []int{1,2,3,3,4}, false},
        { []int{1,2,3,2,1}, true},
        { []int{1,1}, true},
        { []int{1,1,2,2}, false},
        { []int{1}, true},
    }

    for _, test := range(tests) {
        l1 := CreateList(test.input1)
        if got := l1.IsPalindrome(); got != test.res {
            t.Errorf("faild to test palindrome")
        }
    }
}



