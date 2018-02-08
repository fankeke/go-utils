package word

import "testing"


func TestPalindrome(t *testing.T) {
    if !IsPalindrome("detartrated") {
        t.Error(`"IsPalindrome("detartrated") = false`)
    }

    if !IsPalindrome("kayak") {
        t.Error(`"IsPalindrome("kayak") = false`)
    }

}

func TestNonPalindrome(t*testing.T) {
    if IsPalindrome("hello") {
        t.Error(`IsPalindrome("hello") = true`)
    }
}

func TestIspalindrome(t *testing.T){
    var tests = []struct {
        input   string
        want    bool
    }{
        { "", true},
        { "a", true},
        { "aa", true},
        { "ab", false},
        { "kayak", true},
        { "A man, a plan ,a canl", false},
    }

    for _, test := range(tests) {
        if got := IsPalindrome(test.input); got != test.want {
            t.Errorf("isPalnaince(%q)=%v", test.input, got)
        }
    }
}


func TestRandomePalindoemse(t*testing.T) {
}
