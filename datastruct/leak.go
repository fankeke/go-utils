package main

type Node struct {
    next    *Node
    payload [64]byte
}

func main() {
    curr := new(Node)
    for {
        curr.next = new(Node)
        curr = curr.next
    }
}
