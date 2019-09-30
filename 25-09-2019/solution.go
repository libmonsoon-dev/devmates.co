package solution

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func Solution(input []*int) int {
	max := func(first int, list ...int) int {
		result := first
		for i := 0; i < len(list); i++ {
			if list[i] > result {
				result = list[i]
			}
		}
		return result
	}

	maxSum := MinInt

	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil || root.Value == nil {
			return 0
		}
		value := *root.Value
		leftSum := dfs(root.Left)
		rightSum := dfs(root.Right)

		maxSum = max(maxSum, leftSum+value+rightSum)

		return max(0, value+leftSum, value+rightSum)
	}
	builder := &TreeBuilder{}
	builder.InsertLevelOrder(input)
	dfs(builder.Root)

	return maxSum
}

type TreeBuilder struct {
	Root *TreeNode
}

func (t *TreeBuilder) InsertLevelOrder(list []*int) {
	t.Root = t.insertLevelOrder(list, t.Root, 0, len(list))
}

func (t *TreeBuilder) insertLevelOrder(list []*int, root *TreeNode, i, n int) *TreeNode {
	if i < n {
		temp := &TreeNode{Value: list[i]}
		root = temp

		root.Left = t.insertLevelOrder(list, root.Left, 2*i+1, n)
		root.Right = t.insertLevelOrder(list, root.Right, 2*i+2, n)
	}
	return root
}

type TreeNode struct {
	Value *int
	Left  *TreeNode
	Right *TreeNode
}

func Ptr(x int) *int {
	return &x
}
