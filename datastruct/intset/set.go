package intset

//集合数据结构，集合为一个数组，每个元素为64位，每位为0/1，1表示该位
//的顺序的那个数字存在在此集合,从左到右的顺序

type IntSet struct {
    words    []uint64
}

func (s *IntSet)Has(x int)bool {
    word, bit := x / 64, uint(x%64)
    return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func(s *IntSet)Add(x int) {
    word, bit := x/64, uint(x%64)
    for word >= len(s.words) {
        s.words = append(s.words, 0)
    }
    s.words[word] |= 1 << bit
}

func(s *IntSet)Del(x int) {
    word, bit := x/64, uint(x%64)

    if word > len(s.words) {
        return
    }
    s.words[word] &= ^(1<<bit)
}
    

func (s *IntSet)UnionWith(t *IntSet) {
    for i, tword := range(t.words) {
        if i < len(s.words) {
            s.words[i] |= tword
        } else {
            s.words = append(s.words, tword)
        }
    }
}

//func CreateSet(sl []uint64)(*IntSet) {
//    return &IntSet{sl}
//}
