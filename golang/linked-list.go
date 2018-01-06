package main

import "fmt"

type LinkedList struct {
    Data interface{}
    Next *LinkedList
}

func (L *LinkedList) AppendToTail(data interface{}) {
    if (L.Next == nil) {
        L.Next = &LinkedList{data, nil}
    } else {
        L.Next.AppendToTail(data)
    }
}

func (L *LinkedList) Find(data interface{}) *LinkedList {
    if (L == nil) {
        return nil
    }
    if (L.Data == data) {
        return L
    }
    return L.Next.Find(data)
}

func (L *LinkedList) Invert() *LinkedList {
    var prev *LinkedList = nil
    current := L
    var next *LinkedList
    for current != nil {
        next = current.Next
        current.Next = prev
        prev = current
        current = next
    }
    return prev
}

func (L *LinkedList) DeleteNode(data interface{}) *LinkedList {
    var head = L
    if (L == nil) {
        return nil
    }
    if (L.Data == data) {
        return L.Next
    }
    for head.Next != nil {
        if (head.Next.Data == data) {
            head.Next = head.Next.Next
            return L
        }
        head = head.Next
    }
    return L
}

func (L *LinkedList) String() string {
    if (L == nil) {
        return "nil"
    }
    return L.Data.(string) + " -> " + L.Next.String()
}
func main() {
    head := &LinkedList{"1", nil}
    head.Next = &LinkedList{"2", nil}
    head.AppendToTail("3")
    head.AppendToTail("4")
    head.AppendToTail("5")
    head.AppendToTail("6")
    fmt.Println(head)
    inverted := head.Invert()
    fmt.Println(inverted)
    inverted.DeleteNode("4")
    fmt.Println(inverted)
    newHead := inverted.DeleteNode("6")
    fmt.Println(newHead)
    three := newHead.Find("3")
    fmt.Println(three)
}
