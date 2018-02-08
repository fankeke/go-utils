package intset

//type IntSet struct {
//    words    []uint64
//}
type IntSet2 []uint64


func (s *IntSet2)Has(x int)bool {
    word, bit := x / 64, uint(x%64)
    return word < len(*s) && (*s)[word]&(1<<bit) != 0
}

func(s *IntSet2)Add(x int) {
    word, bit := x/64, uint(x%64)
    for word >= len(*s) {
        *s = append(*s, 0)
    }
    (*s)[word] |= 1 << bit
}

func(s *IntSet2)Del(x int) {
    word, bit := x/64, uint(x%64)

    if word > len(*s) {
        return
    }
    (*s)[word] &= ^(1<<bit)
}
    

//func (s *IntSet2)UnionWith(t *IntSet2) {
//    for i, tword := range(t) {
//        if i < len(s) {
//            s[i] |= tword
//        } else {
//            s = append(s,tword)
//        }
//    }
//}

//func CreateSet(sl []uint64)(*IntSet) {
//    return &IntSet{sl}
//}
