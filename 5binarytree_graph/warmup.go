package bg

import "fmt"

func PreorderPrint(node *TreeNode, layer int) {
	if node == nil {
		return
	}
	for i := 0; i < layer; i++ {
		fmt.Printf("----")
	}
	fmt.Println(node.val)
	PreorderPrint(node.left, layer+1)
	PreorderPrint(node.right, layer+1)
}
