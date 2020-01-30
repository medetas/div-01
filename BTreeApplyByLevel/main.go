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

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error))  {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	n := root
	for n != nil {
		f(n.Data)
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
		if len(q) == 1 {
			return
		}
		q = q[1:]
		n = q[0]	
	}	
}

func main() {

  /*  root := &TreeNode{Data: "01"}
    BTreeInsertData(root, "07")
    BTreeInsertData(root, "12")
    BTreeInsertData(root, "10")
    BTreeInsertData(root, "05")
    BTreeInsertData(root, "02")
    BTreeInsertData(root, "03")
    BTreeApplyByLevel(root, fmt.Println)

    // 01
    // 07
    // 05
    // 12
    // 02
    // 10
    // 03
*/
	root := &TreeNode{Data: "01"}
	BTreeInsertData(root, "07")
	BTreeInsertData(root, "12")
	BTreeInsertData(root, "10")
	BTreeInsertData(root, "05")
	BTreeInsertData(root, "02")
	BTreeInsertData(root, "03")
	BTreeApplyByLevel(root, fmt.Print)
	fmt.Println()

    // 01070512021003

}
