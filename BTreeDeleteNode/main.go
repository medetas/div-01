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

func MinVal(node *TreeNode) *TreeNode {
	current := node
	for current != nil && current.Left != nil {
		current = current.Left
	}
	return current
	
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			temp := root.Right
			root = nil
			return temp
		} else if root.Right == nil {
			temp := root.Left
			root = nil
			return temp
		}
		temp := MinVal(root.Right)
		root.Data = temp.Data
		root.Right = BTreeDeleteNode(root.Right, temp)
	}
	return root	
}

func main() {
  /*  root := &TreeNode{Data: "01"}
    BTreeInsertData(root, "07")
    BTreeInsertData(root, "05")
    BTreeInsertData(root, "12")
    BTreeInsertData(root, "02")
    BTreeInsertData(root, "10")
    BTreeInsertData(root, "03")
    node := BTreeSearchItem(root, "02")
    fmt.Println("Before delete:")
    BTreeApplyInorder(root, fmt.Println)
    root = BTreeDeleteNode(root, node)
    fmt.Println("After delete:")
    BTreeApplyInorder(root, fmt.Println)

    // 01
    // 03
    // 05
    // 07
    // 10
    // 12
*/
/*
root := &TreeNode{Data: "33"}
BTreeInsertData(root, "5")
BTreeInsertData(root, "20")
BTreeInsertData(root, "31")
BTreeInsertData(root, "13")
BTreeInsertData(root, "11")
BTreeInsertData(root, "52")
node := BTreeSearchItem(root, "20")
fmt.Println("Before delete:")
BTreeApplyInorder(root, fmt.Println)
root = BTreeDeleteNode(root, node)
fmt.Println("After delete:")
BTreeApplyInorder(root, fmt.Println)
*/
    // 11
    // 13
    // 31
    // 33
    // 5
    // 52
root := &TreeNode{Data: "03"}
BTreeInsertData(root, "39")
BTreeInsertData(root, "99")
BTreeInsertData(root, "44")
BTreeInsertData(root, "11")
BTreeInsertData(root, "14")
BTreeInsertData(root, "11")
node := BTreeSearchItem(root, "03")
fmt.Println("Before delete:")
BTreeApplyInorder(root, fmt.Println)
root = BTreeDeleteNode(root, node)
fmt.Println("After delete:")
BTreeApplyInorder(root, fmt.Println)

    // 11
    // 11
    // 14
    // 39
    // 44
    // 99

    // root := &TreeNode{Data: "03"}
    // BTreeInsertData(root, "03")
    // BTreeInsertData(root, "94")
    // BTreeInsertData(root, "19")
    // BTreeInsertData(root, "24")
    // BTreeInsertData(root, "111")
    // BTreeInsertData(root, "01")
    // node := BTreeSearchItem(root, "03")
    // fmt.Println("Before delete:")
    // BTreeApplyInorder(root, fmt.Println)
    // root = BTreeDeleteNode(root, node)
    // fmt.Println("After delete:")
    // BTreeApplyInorder(root, fmt.Println)

    // 01
    // 03
    // 111
    // 19
    // 24
    // 94
}
