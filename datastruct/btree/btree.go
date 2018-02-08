package btree

import  "fmt"


type TreeNode struct {
    Val       int
    Left, Right  *TreeNode
}

func reverseSlice(sli []int) {
    length := len(sli)
    for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
        sli[i], sli[j] = sli[j], sli[i]
    }
}

//func add(t *tree, value int) *tree {
//    if t == nil {
//        t = &tree{}
//        t.value = value
//        return t
//    }
//
//    if value < t.value {
//        t.left = add(t.left, value)
//    } else {
//        t.right = add(t.right, value)
//    }
//    return t
//}
//
//func  appendValues(values[]int, t *tree)[]int {
//    if t != nil {
//        values = appendValues(values, t.left)
//        values = append(values, t.value)
//        values = appendValues(values, t.right)
//    }
//    return values
//}
//
//func Sort(values []int) {
//    var root *tree
//    for _, v := range(values) {
//        root = add(root, v)
//    }
//
//    appendValues(values[:0], root)
//}

func(self *TreeNode)InorderTravel(){
    if self != nil {
        self.Left.InorderTravel()
        //fmt.Println(self.Val)
        self.Right.InorderTravel()
    }
}

func(root*TreeNode)PreOrderTravel()[]int{
    if root == nil {
        return nil
    }
    ans := []int{}
    stack := []*TreeNode{}

    for root != nil || len(stack) != 0 {
        for root != nil {
            ans = append(ans, root.Value)
            stack = append(stack, root)
            root = root.Left
        }

        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        root = root.Right
    }
    return ans
}

func(root*TreeNode)InorderTravel2()[]int{
    if root == nil {
        return nil
    }
    ans := []int{}
    stack := []*TreeNode{}

    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        ans = append(ans, root.Value)
        root = root.Right
    }
    return ans
    
}


//func (root*TreeNode)PreOrderTravel()[]int {
//    if root == nil {
//        return nil
//    }
//
//    stack := []*TreeNode{}
//    ans := []int{}
//
//    for root != nil || len(stack) != 0 {
//        for root != nil {
//            stack = append(stack, root)
//            ans = append(ans, root.Val)
//            root = root.Left
//        }
//
//        q := stack[len(stack)-1]
//        stack = stack[:len(stack)-1]
//        root = q.Right
//    }
//    return ans
//}


//func(root *TreeNode)InorderTravel2()[]int {
//    list := []int{}
//    if root == nil {
//        return nil
//    }
//
//    stack := []*TreeNode{}
//
//    for root != nil || len(stack) != 0 {
//        for root != nil {
//            stack = append(stack, root)
//            root = root.Left
//        }
//
//        q := stack[len(stack)-1]       //出栈
//        stack = stack[:len(stack)-1]
//
//        list = append(list, q.Val)
//        root = q.Right
//    }
//
//    return list
//}


//注意看，pre的顺序是根左右
//       post的顺序是左右根
//如果把post倒着看 根右左 ,可以从pre的访问中得到启发：

func (root*TreeNode)PostOrderTravel()[]int{
    if root == nil {
        return nil
    }

    stack := []*TreeNode{}
    ans := []int{}
    
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            ans = append(ans, root.Val)
            root = root.Right    //注意和pre的不一样这里压入right
        }

        q := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        root = q.Left //转入left
    }
   //此时ans中的数据和post顺序是刚好想法的，需要reverse一下即可
    reverseSlice(ans)

    return ans
}

func (root*TreeNode)LevelOrderTravel()[]int{
    if root == nil {
        return nil
    }

    //level遍历用队列
    queue := []*TreeNode{}
    ans := []int{}

    queue = append(queue, root)

    for len(queue) != 0 {
        cur_len := len(queue)

        for i := 0; i < cur_len; i++ {     //每轮输出一次

            p := queue[0]               //每次都是队首出列
            ans = append(ans, p.Val)
            queue = queue[1:]

            if p.Left != nil {
                queue = append(queue, p.Left)
            }
            if p.Right != nil {
                queue = append(queue, p.Right)
            }
        }
    }

    return ans
}

func (root*TreeNode)IsLeaf()bool{
    if root != nil && root.Left == nil && root.Right == nil {
        return true
    }
    return false
}


func HasPathSum(root *TreeNode, sum int) bool {

    if root == nil {
        return false
    }

    if root.IsLeaf() && sum == root.Val {
        return true
    }

    return  HasPathSum(root.Left, sum-root.Val) || HasPathSum(root.Right, sum-root.Val)
}

func SumSlice(ans []int) int{
    sum := 0
    for i := 0; i < len(ans); i++ {
        sum += ans[i]
    }
    return sum
}




//回溯的写法
//func findAllPathSum(root *TreeNode, sum int, ans *[]int, res *[][]int) {
//    if root == nil {
//        return 
//    }
//
//    *ans = append(*ans, root.Val)
//
//    if root.Val == sum && root.Left == nil && root.Right == nil {
//
//        candy := make([]int, len(*ans)) //非常重要
//        copy(candy, *ans)
//
//        *res = append(*res, candy)
//        //fmt.Printf("%+v\n", *ans)
//        *ans = (*ans)[:len(*ans)-1]     //卸去叶子继续
//
//        return
//    }
//    findAllPathSum(root.Left, sum-root.Val, ans, res)
//    findAllPathSum(root.Right, sum-root.Val, ans, res)
//
//    *ans = (*ans)[:len(*ans)-1]    //卸去此根继续
//}
//
//func AllPathSum(root*TreeNode, sum int)[][]int{
//    if root == nil {
//        return nil
//    }
//
//    ans := []int{}
//    res := [][]int{}
//
//    findAllPathSum(root, sum, &ans, &res)
//
//    return res
//}
func findAllPathSum(root *TreeNode, sum int, ans *[]int, res *[][]int){
    if root == nil {
        return
    }
    *ans = append(*ans, root.Value)
    if root.Value == sum && root.Left == nil && root.Right == nil {
        candy := make([]int, len(*ans))
        copy(candy, *ans)
        *res = append(*res, candy)

        *ans = (*ans)[:len(*ans)-1]
        return
    }
    findAllPathSum(root.Left, sum-root.Value, ans, res)
    findAllPathSum(root.Right, sum-root.Value, ans, res)
    *ans = (*ans)[:len(*ans)-1]
}

func AllPathSum(root*TreeNode, sum int)[][]int{
    if root == nil {
        return nil
    }
    ans := []int{}
    res := [][]int{}
    
    findAllPathSum(root, sum, &ans, &res)
}


func createBstHelper(list []int, low, high int)*TreeNode {
    if low > high {
        return nil
    }
    mid := (low+high)/2
    root := &TreeNode{list[mid], nil, nil}
    root.Left = createBstHelper(list, low, mid-1)
    root.Right = createBstHelper(list, mid+1, high)
    return root
}

func CreateBST(list []int)*TreeNode{
    if len(list) == 0 {
        return nil
    }
    return createBstHelper(list, 0, len(list)-1)
}


//对于btree的递归，有这样的
//格式：先判断是否已经满足条件，然后看是否已经不满足条件 ，然后再左右递归
    //或先判断是否已经不满足条件，然后看是否已经满足条件，然后再左右递归




func CreateBtreFromInorderAndPreOrder(pre, in []int)*TreeNode{
    if len(pre) == 0 {
        return nil
    }
    root := &TreeNode{pre[0], nil, nil}
    var i int
    for i =0;i<len(in);i=i+1{
        if in[i] == root.Val{
            break
        }
    }
    
    //切片的这种行为真是太好用了，随便就可以创建array，让代码变得和思维一样清晰
    root.Left = CreateBtreFromInorderAndPreOrder(pre[1:i+1], in[:i])
    root.Right = CreateBtreFromInorderAndPreOrder(pre[i+1:], in[i+1:])

    return root    
}


func CreateBtreFromInorderAndPostOrder(post, in []int)*TreeNode{
    if len(post) == 0 {
        return nil
    }
    root := &TreeNode{post[len(post)-1], nil, nil}
    var i int
    for i =0;i<len(in);i=i+1{
        if in[i] == root.Val{
            break
        }
    }
    
    //切片的这种行为真是太好用了，随便就可以创建array，让代码变得和思维一样清晰
    root.Right = CreateBtreFromInorderAndPostOrder(post[i:len(post)-1], in[i+1:])
    root.Left = CreateBtreFromInorderAndPostOrder(post[:i], in[:i])

    return root    
}










