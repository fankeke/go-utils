package btree

import (
        "fmt"
        "testing"
       )


func equal(a1, a2 []int)bool{
    if len(a1) != len(a2) {
        return false
    }

    for i := 0; i < len(a1); i++ {
        if a1[i] != a2[i] {
            return false
        }
    }

    return true
}


//func TestSort(t *testing.T) {
//
//    var tests = []struct {
//        input   []int
//        want    []int
//    }{
//        { []int{1,2,3,6,4}, []int{1,2,3,4,6}},
//        { []int{1,2}, []int{1,2}},
//        { []int{1}, []int{1}},
//        { []int{7,2,3,6,4}, []int{2,3,4,6,7}},
//    }
//
//    for _, test := range(tests) {
//        Sort(test.input)
//        if !equal(test.input, test.want) {
//            t.Errorf("failed to test")
//        }
//    }
//}
//


func TestCreateBST(t *testing.T) {
    var tests = []struct {
        input   []int
    }{
        { []int{1,2,3,4,5}},
    }
    for _, test := range(tests) {
        root := CreateBST(test.input)
        root.InorderTravel()
    }
}

func TestInoderTraval(t *testing.T) {
    var tests = []struct {
        input   []int
        output  []int
    }{
        { []int{1,2,3,4,5}, []int{1,2,3,4,5}},
    }
    for _, test := range(tests) {
        root := CreateBST(test.input)
        got := root.InorderTravel2()
        if equal(got, test.output) != true {
            fmt.Printf("%+v\n", got)
            t.Errorf("faild to inordertravel")
        }
    }
}

func TestPreoderTraval(t *testing.T) {
    var tests = []struct {
        input   []int
        output  []int
    }{
        { []int{1,2,3,4,5}, []int{3,1, 2, 4,5}},
    }
    for _, test := range(tests) {
        root := CreateBST(test.input)
        got := root.PreOrderTravel()
        if equal(got, test.output) != true {
            fmt.Printf("%+v\n", got)
            t.Errorf("faild to preordertravel")
        }
    }
}

func TestPostoderTraval(t *testing.T) {
    var tests = []struct {
        input   []int
        output  []int
    }{
        { []int{1,2,3,4,5}, []int{2,1,5,4,3}},
    }
    for _, test := range(tests) {
        root := CreateBST(test.input)
        got := root.PostOrderTravel()
        if equal(got, test.output) != true {
            fmt.Printf("%+v\n", got)
            t.Errorf("faild to postordertravel")
        }
    }
}
 
func TestLevelOderTraval(t *testing.T) {
    var tests = []struct {
        input   []int
        output  []int
    }{
        { []int{1,2,3,4,5}, []int{3,1,4,2,5}},
    }
    for _, test := range(tests) {
        root := CreateBST(test.input)
        got := root.LevelOrderTravel()
        if equal(got, test.output) != true {
            fmt.Printf("%+v\n", got)
            t.Errorf("faild to LevelOrdertravel")
        }
    }
}

func TestHasPathSum(t *testing.T) {
    var tests = []struct {
        input   []int
        sum       int
        got       bool
    }{
        { []int{1,2,3,4,5}, 6, true},
    }
    for _, test := range(tests) {
        root := CreateBST(test.input)
        got := HasPathSum(root, test.sum)
        if got != test.got {
            fmt.Printf("%+v\n", got)
            t.Errorf("faild to HasPathSum")
        }
    }
}

func TestAllPathSum(t *testing.T) {
    var tests = []struct {
        input   []int
        sum       int
    }{
        { []int{1,2,3,4,5}, 12},
        { []int{5,7,8,9,10,12}, 30},
    }
    for _, test := range(tests) {
        root := CreateBST(test.input)
        AllPathSum(root, test.sum)
        //if got != test.got {
        //    fmt.Printf("%+v\n", got)
        //    t.Errorf("faild to HasPathSum")
        //}
    }

}


func TestCreateBtreFromInorderAndPreOrder(t *testing.T) {
    var tests = []struct {
        pre   []int
        in    []int
        post  []int //后序访问顺序
    }{
        { []int{1,2,4,5,3,6,7}, []int{4,2,5,1,7,6,3}, []int{4,5,2,7,6,3,1}},
    }
    for _, test := range(tests) {
        root := CreateBtreFromInorderAndPreOrder(test.pre, test.in)
        got := root.PostOrderTravel()
        if equal(got, test.post) != true {
            t.Errorf("failed to create btree from inorder and preorder")
        }
    }
}

func TestCreateBtreFromInorderAndPostOrder(t *testing.T) {
    var tests = []struct {
        pre   []int
        in    []int
        post  []int //后序访问顺序
    }{
        { []int{1,2,4,5,3,6,7}, []int{4,2,5,1,7,6,3}, []int{4,5,2,7,6,3,1}},
    }
    for _, test := range(tests) {
        root := CreateBtreFromInorderAndPostOrder(test.post, test.in)
        got := root.InorderTravel2()
        if equal(got, test.in) != true {
            t.Errorf("failed to create btree from inorder and preorder")
        }
    }
}





