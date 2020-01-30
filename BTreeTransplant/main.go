package main

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}
	return root
}

func BTreeSearchItem(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == data {
		return root
	} else if data < root.Data {
		root.Left.Parent = root
		return BTreeSearchItem(root.Left, data)
	} else {
		root.Right.Parent = root
		return BTreeSearchItem(root.Right, data)
	}
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	res := BTreeSearchItem(root, node.Data)
	if res.Parent.Right.Data == res.Data {
		res.Parent.Right = rplc
	} else if res.Parent.Left.Data == res.Data {
		res.Parent.Left = rplc
	}
	return root
}

func main() {
/*
	root := &TreeNode{Data: "01"}
	BTreeInsertData(root, "07")
	BTreeInsertData(root, "05")
	BTreeInsertData(root, "12")
	BTreeInsertData(root, "10")
	BTreeInsertData(root, "02")
	BTreeInsertData(root, "03")
	node := BTreeSearchItem(root, "12")
	replacement := &TreeNode{Data: "55"}
	BTreeInsertData(replacement, "60")
	BTreeInsertData(replacement, "33")
	BTreeInsertData(replacement, "12")
	BTreeInsertData(replacement, "15")
	root = BTreeTransplant(root, node, replacement)
	BTreeApplyInorder(root, fmt.Println)
*/
/*
	root := &TreeNode{Data: "33"}
	BTreeInsertData(root, "5")
	BTreeInsertData(root, "52")
	BTreeInsertData(root, "20")
	BTreeInsertData(root, "31")
	BTreeInsertData(root, "13")
	BTreeInsertData(root, "11")
	node := BTreeSearchItem(root, "20")
	replacement := &TreeNode{Data: "55"}
	BTreeInsertData(replacement, "60")
	BTreeInsertData(replacement, "33")
	BTreeInsertData(replacement, "12")
	BTreeInsertData(replacement, "15")
	root = BTreeTransplant(root, node, replacement)
	BTreeApplyInorder(root, fmt.Println)
*/
	root := &TreeNode{Data: "03"}
	BTreeInsertData(root, "39")
	BTreeInsertData(root, "99")
	BTreeInsertData(root, "44")
	BTreeInsertData(root, "11")
	BTreeInsertData(root, "14")
	BTreeInsertData(root, "11")
	node := BTreeSearchItem(root, "11")
	replacement := &TreeNode{Data: "55"}
	BTreeInsertData(replacement, "60")
	BTreeInsertData(replacement, "33")
	BTreeInsertData(replacement, "12")
	BTreeInsertData(replacement, "15")
	root = BTreeTransplant(root, node, replacement)
	BTreeApplyInorder(root, fmt.Println)

}
