package binary_tree

import (
	tree "devmates.co/09-09-2019"
)

func createTree(rootValue int, items ...*int) *tree.Node {
	root := &tree.Node{Value: rootValue}
	sameLevelNodes := []*tree.Node{root}
	i := 0

	processItems := func(node *tree.Node, nextLevel *[]*tree.Node, left bool) {
		if i < len(items) && items[i] != nil {
			child := &tree.Node{Value: *items[i]}

			if left {
				node.AddLeft(child)
			} else {
				node.AddRight(child)
			}

			*nextLevel = append(*nextLevel, child)
		} else {
			*nextLevel = append(*nextLevel, nil)
		}
		i++
	}

	for i < len(items) {
		nextLevel := make([]*tree.Node, 0, len(sameLevelNodes)*2)

		for _, node := range sameLevelNodes {
			processItems(node, &nextLevel, true)
			processItems(node, &nextLevel, false)

			sameLevelNodes = nextLevel
		}

	}

	return root
}

func GetRightValues(items ...*int) []int {
	currentNode := createTree(*items[0], items[1:]...)

	result := make([]int, 0)

	for {
		if currentNode == nil {
			break
		}
		result = append(result, currentNode.Value)
		currentNode = currentNode.Right
	}

	return result
}
