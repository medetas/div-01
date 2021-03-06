package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	iter := l.Head
	l.Head = nil
	for iter != nil {
		if data_ref != iter.Data {
			ListPushBack(l, iter.Data)
		}
		iter = iter.Next
	}
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		return
	}
	iterator := l.Head
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n

}

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
    a := &List{}

    fmt.Println("----normal state----")
    PrintList(a)
    ListRemoveIf(a, 1)
    fmt.Println("------answer-----")
    PrintList(a)
    fmt.Println()

    b := &List{}

    fmt.Println("----normal state----")
    PrintList(b)
    ListRemoveIf(b, 96)
    fmt.Println("------answer-----")
    PrintList(b)
    fmt.Println()

    c := &List{}

    fmt.Println("----normal state----")

    ListPushBack(c, 98)
    ListPushBack(c, 98)
    ListPushBack(c, 33)
    ListPushBack(c, 34)
    ListPushBack(c, 33)
    ListPushBack(c, 34)
    ListPushBack(c, 33)
    ListPushBack(c, 89)
    ListPushBack(c, 33)

    PrintList(c)
    ListRemoveIf(c, 34)
    fmt.Println("------answer-----")
    PrintList(c)
    fmt.Println()

    d := &List{}

    fmt.Println("----normal state----")

    ListPushBack(d, 79)
    ListPushBack(d, 74)
    ListPushBack(d, 99)
    ListPushBack(d, 79)
    ListPushBack(d, 7)

    PrintList(d)
    ListRemoveIf(d, 99)
    fmt.Println("------answer-----")
    PrintList(d)
    fmt.Println()

    e := &List{}

    fmt.Println("----normal state----")

    ListPushBack(e, 56)
    ListPushBack(e, 93)
    ListPushBack(e, 68)
    ListPushBack(e, 56)
    ListPushBack(e, 87)
    ListPushBack(e, 68)
    ListPushBack(e, 56)
    ListPushBack(e, 68)

    PrintList(e)
    ListRemoveIf(e, 68)
    fmt.Println("------answer-----")
    PrintList(e)
    fmt.Println()

    f := &List{}

    fmt.Println("----normal state----")

    ListPushBack(f, "mvkUxbqhQve4l")
    ListPushBack(f, "4Zc4t hnf SQ")
    ListPushBack(f, "q2If E8BPuX")

    PrintList(f)
    ListRemoveIf(f, "4Zc4t hnf SQ")
    fmt.Println("------answer-----")
    PrintList(f)
    fmt.Println()
}
