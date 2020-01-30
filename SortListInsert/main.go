package main

import "fmt"

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref}
	if l == nil {
		return n
	}
	iterator := l
	if n.Data < iterator.Data {
		n.Next = iterator
		return n
	}
	for iterator.Next != nil && iterator.Next.Data < n.Data {
		iterator = iterator.Next
	}
	temp := iterator.Next
	iterator.Next = n
	n.Next = temp
	return l
}

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {

	var link *NodeI

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)

	PrintList(link)

	link = SortListInsert(link, -2)
	link = SortListInsert(link, 2)
	link = SortListInsert(link, 564)
	PrintList(link)
}
