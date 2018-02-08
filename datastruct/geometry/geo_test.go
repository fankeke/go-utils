package geometry

import "testing"


func TestSanity(t *testing.T) {
    var tests = []struct {
        input1 Point
        input2 Point
        want   float64
    }{
        { Point{1,2}, Point{4, 6}, 5},
        { Point{3,4}, Point{6, 8}, 5},
        { Point{5,6}, Point{8, 10}, 5},
    }

    for _, test := range(tests) {
        if got := test.input1.Distance(test.input2); got != test.want {
            t.Errorf("failed to test:%v", test)
        }
    }
}

func TestPathSanity(t *testing.T) {
    var tests = []struct {
        input Path
        want  float64
    }{
        { Path{ Point{1,1}, Point{5,1}, Point{5,4}, Point{1,1} }, 12},
    }

    for _, test := range(tests) {
        if got := test.input.Distance(); got != test.want {
            t.Errorf("failed to test:%v", test)
        }
    }
}

