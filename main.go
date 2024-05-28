package main

import "fmt"

func main() {
	input := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	output := levelOrder(input)
	fmt.Println(output)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	resultValSlice := make([][]int, 0)
	resultNodeSlice := make([][]*TreeNode, 0)

	if root != nil {
		resultNodeSlice = append(resultNodeSlice, []*TreeNode{root})
	}

	for i := 0; i < len(resultNodeSlice); i++ {
		currentValSlice := make([]int, len(resultNodeSlice[i]))
		nextNodeSlice := make([]*TreeNode, 0)
		for j, v := range resultNodeSlice[i] {
			currentValSlice[j] = v.Val
			if v.Left != nil {
				nextNodeSlice = append(nextNodeSlice, v.Left)
			}
			if v.Right != nil {
				nextNodeSlice = append(nextNodeSlice, v.Right)
			}
		}
		if len(nextNodeSlice) > 0 {
			resultNodeSlice = append(resultNodeSlice, nextNodeSlice)
		}
		resultValSlice = append(resultValSlice, currentValSlice)
	}

	return resultValSlice
}
