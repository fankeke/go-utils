package intset

import (
        "testing"
       )

func TestSanity(t *testing.T) {
    var tests = []struct {
        input  int
        want    bool
    }{
         {1, false},
         {2, false},
         {3, false},
         {498, false},
         {1000, false},
    }

    set := &IntSet{}
    
    for _, test := range(tests) {
        if got := set.Has(test.input); got != test.want {
            t.Errorf("failed to test:%v %v", test.input, test.want)
        }
        set.Add(test.input)
        if got := set.Has(test.input); got != true {
            t.Errorf("failed to add")
        }
        set.Del(test.input)
        if got := set.Has(test.input); got != false {
            t.Errorf("failed to del")
        }
    }
}


func TestSanity2(t *testing.T) {
    var tests = []struct {
        input  int
        want    bool
    }{
         {1, false},
         {2, false},
         {3, false},
         {498, false},
         {1000, false},
    }

    set := &intset2{1}
    
    for _, test := range(tests) {
        if got := set.Has(test.input); got != test.want {
            t.Errorf("failed to test:%v %v", test.input, test.want)
        }
        set.Add(test.input)
        if got := set.Has(test.input); got != true {
            t.Errorf("failed to add")
        }
        (*set)[0]=2
        set.Del(test.input)
        if got := set.Has(test.input); got != false {
            t.Errorf("failed to del")
        }
    }
}

