package main
import (
    "fmt"
)
var nMaxLength int
type Node struct {
    Left *Node
    Right *Node
    Value int
    leftMaxLength int
    rightMaxLength int
}
var nodes map[int]*Node

func findMaxLen(root *Node ) {
    if root == nil {
        return 
    }
    if root.Left == nil {
        root.leftMaxLength = 0
    }
    if root.Right == nil {
        root.rightMaxLength = 0
    }
    if root.Left != nil {
        findMaxLen(root.Left)
    }
    if root.Right != nil {
        findMaxLen(root.Right)
    }
    if root.Left != nil {
        if root.Left.leftMaxLength > root.Left.rightMaxLength {
            root.leftMaxLength = root.Left.leftMaxLength + 1
        }else {
            root.leftMaxLength = root.Left.rightMaxLength + 1
        }
    }
    if root.Right != nil {
        if root.Right.leftMaxLength > root.Right.rightMaxLength {
            root.rightMaxLength = root.Right.leftMaxLength + 1
        }else {
            root.rightMaxLength = root.Right.rightMaxLength + 1
        }
    }
    if root.leftMaxLength + root.rightMaxLength > nMaxLength {
        nMaxLength = root.leftMaxLength + root.rightMaxLength
    }
}
func createNode(fa, lch, rch int) {
    var root, l, r *Node
    var ok bool
    root, ok = nodes[fa]
    if !ok {
        root = &Node{}
        root.Value = fa
        root.Right = nil
        root.Left = nil
        nodes[fa] = root
    }
    if lch != 0 {
        l, ok = nodes[lch]
        if !ok {
            l = &Node{}
            l.Value = lch
            l.Right = nil
            l.Left = nil
            nodes[lch] = l
        }
        root.Left = l
    }
    if rch != 0 {
        r, ok = nodes[rch]
        if !ok {
            r = &Node{}
            r.Value = rch
            r.Right = nil
            r.Left = nil
            nodes[rch] = r
        }
        root.Right = r
    }
}

func main() {
    nodes = make(map[int]*Node)
    root := &Node{Value:1,}
    nodes[1] = root
    createNode(1,2,3)
    createNode(2,4,5)
    createNode(4,0,0)
    createNode(5,0,0)
    createNode(3,6,7)
    createNode(6,0,0)
    createNode(7,0,0)
    findMaxLen(root)
    fmt.Println("result ",nMaxLength + 1)
}
