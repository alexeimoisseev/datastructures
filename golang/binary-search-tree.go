package main

import "fmt"

type BinarySearchTree struct {
    Val string
    Left *BinarySearchTree
    Right *BinarySearchTree
}

func (T *BinarySearchTree) Insert(val string) {
    if (val < T.Val) {
        if (T.Left != nil) {
            T.Left.Insert(val)
        } else {
            T.Left = &BinarySearchTree{val, nil, nil}
        }
    } else {
        if (T.Right != nil) {
            T.Right.Insert(val)
        } else {
            T.Right = &BinarySearchTree{val, nil, nil}
        }
    }
}

func (T *BinarySearchTree) String() (string) {
    if (T == nil) {
        return "nil"
    }
    res := "{" + T.Val +
    " Left: " + T.Left.String() +
    ", Right: " + T.Right.String() +
    "}"
    return res
}


func (T *BinarySearchTree) Find(val string) *BinarySearchTree {
    if (T == nil) {
        return nil
    }
    if (T.Val == val) {
        return T
    }
    if (val < T.Val) {
        return T.Left.Find(val)
    } else {
        return T.Right.Find(val)
    }
}

func (T *BinarySearchTree) Traverse(f func(*BinarySearchTree)) {
    if (T == nil) {
        return
    }
    T.Left.Traverse(f)
    f(T)
    T.Right.Traverse(f)
}

/*
func main() {

    root := &BinarySearchTree{"5", nil, nil}
    root.Insert("2")
    root.Insert("3")
    root.Insert("1")
    root.Insert("6")
    root.Insert("8")
    root.Insert("4")
    root.Insert("7")

    fmt.Println(root)
    six := root.Find("6")
    fmt.Println(six)

    root.Traverse(func (T *BinarySearchTree) {
        fmt.Println(T.Val)
    })
}
*/
