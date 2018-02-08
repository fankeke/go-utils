package list

import (
        "fmt"
       )


type IntList struct {
    Value int
    Next  *IntList
}

func CreateList(sl []int) *IntList {
    if len(sl) < 0 {
        return nil
    }

    h := &IntList{}
    p := h
    for i := 0; i < len(sl); i++ {
        e := &IntList{sl[i], nil}
        h.Next = e
        h = e 
    }
    return p.Next
}

func CreateReverseList(nums []int)*IntList{
    if len(nums) < 0 {
        return nil
    }
    var r *IntList
    for i := 0; i < len(nums); i++ {
        p := &IntList{nums[i], nil}
        p.Next = r 
        r = p
    }
    return r
}

func (list*IntList) RemoveKthFromEnd (k int)(*IntList) {
    var p, q *IntList
    p, q = list, list
    for i := 0; i < k; i++ {
        p = p.Next
    }

    if p == nil {
        return q.Next
    }

    for p.Next != nil {
        p, q = p.Next, q.Next
    }
    q.Next = q.Next.Next

    return list
}





func (l *IntList)Reveser() *IntList{
    if l == nil || l.Next == nil {
        return l
    }

    var r *IntList
    for l != nil {
        q := l.Next
        l.Next = r
        r = l
        l = q
    }
    return r
}

//func(list *IntList)Reveser()*IntList{
//    if list == nil || list.Next == nil {
//        return list
//    }
//
//    var rear, p, q *IntList
//
//    p = list
//    for p != nil {
//        q = p.Next
//        p.Next = rear
//        rear = p
//        p = q
//    }
//    return rear
//}

func (list *IntList) Println() {
    if list == nil {
        return
    }

    p := list
    for p != nil {
        fmt.Printf("%d\t", p.Value)
        p = p.Next
    }
    fmt.Printf("\n")
}

func(self *IntList)Sum()int{
    if self == nil {
        return 0
    }
    return self.Value + self.Next.Sum()
}

func EqualList(l1 *IntList, l2 *IntList)bool {
    for l1 != nil && l2 != nil {
        if l1.Value != l2.Value {
            return false
        }
        l1, l2 = l1.Next, l2.Next
    }
    if l1 != nil || l2 != nil {
        return false
    }

    return true
}

func MergeTwoSortList(l1, l2 *IntList) *IntList {
    head := &IntList{}
    p := head

    for l1 != nil && l2 != nil {
        if l1.Value < l2.Value {
            p.Next = l1
            l1 = l1.Next
        }else{
            p.Next = l2
            l2 = l2.Next
        }
        p = p.Next
    }
    if l1 != nil {
        p.Next = l1
    } else {
        p.Next = l2
    }

    return head.Next
        
}

func (l *IntList)MergeSortList()*IntList{
    if l == nil  || l.Next == nil {
        return l
    }

    fast, slow := l, l
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    //从中间分割出来
    fast = slow.Next
    slow.Next = nil

    l1 := l.MergeSortList()
    l2 := fast.MergeSortList()

    return MergeTwoSortList(l1, l2)

}























//func MergeTwoSortList(l1 *IntList, l2 *IntList)*IntList {
//    var head = &IntList{}
//    var rear *IntList
//    
//    rear = head
//
//    for l1 != nil && l2 != nil {
//        if l1.Value < l2.Value {
//            rear.Next = l1
//            l1 = l1.Next
//            rear = rear.Next
//            rear.Next = nil
//        }else{
//            rear.Next = l2
//            l2 = l2.Next
//            rear = rear.Next
//            rear.Next = nil
//        }
//    }
//    if l1 != nil {
//        rear.Next = l1
//    }else if l2 != nil {
//        rear.Next = l2
//    }
//    return head.Next
//}
//
////链表排序用merge算法可以到达O(n* log n)
//func (list *IntList) MergeSortList()*IntList{
//    if list == nil || list.Next == nil {
//        return list
//    }
//
//    //从中间分为两个
//
//    fast, slow := list, list
//    for fast.Next != nil && fast.Next.Next != nil {
//        fast = fast.Next.Next
//        slow = slow.Next
//    }
//    fast = slow.Next
//    slow.Next = nil
//
//    list1 := list.MergeSortList()
//    list2 := fast.MergeSortList()
//
//    return MergeTwoSortList(list1, list2)
//}

func (head *IntList) DeleteDuplicates()*IntList {
    if head == nil || head.Next == nil {
        return head
    }
    p := head
    for p != nil {
        q := p.Next
        for q != nil && q.Value == p.Value {
            q = q.Next
        }
        p.Next = q
        p = q
    }
    return head         
}

func (head *IntList) DeleteDuplicatesAll()*IntList {
    if head == nil || head.Next == nil {
        return head
    }

    pre := &IntList{}
    pre.Next = head

    p := pre

    for p != nil && p.Next != nil {
        q := p.Next
        r := q.Next
        flag := false
        for r != nil && r.Value == q.Value {
            r = r.Next
            q = q.Next
            flag = true     //has duplicates
        }
        if flag ==  true {
            p.Next = q.Next  //delete all dups
        } else {
            p = p.Next     //nothing happend, go ahead
        }
    }
    return pre.Next         
}

func (head *IntList)IsPalindrome() bool {
    if head == nil || head.Next == nil {
        return true
    }
    li := []int{}
    p := head

    for p != nil {
        li = append(li, p.Value)
        p = p.Next
    }

    for i, j := 0, len(li)-1; i < j; i,j = i+1, j-1 {
        if li[i] != li[j] {
            return false
        }
    }
    return true
}

