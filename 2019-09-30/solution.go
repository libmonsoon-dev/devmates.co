package solution

import "fmt"

func Solution(n, size int) [][]int {
	result := make([][]int, 0)
	numbers := make([]int, 0, size)
	tmp := make([]int, 0, size)

	for i := 1; i <= n; i++ {
		numbers = append(numbers, i)
	}

	var helper func(start int)

	helper = func(start int) {
		if len(tmp) == size {
			copyOfTmp := make([]int, size)
			f := copy(copyOfTmp, tmp)
			fmt.Printf("%#v\n%v\n", copyOfTmp, f)
			result = append(result, copyOfTmp)
			return
		} else {
			for i := start; i < len(numbers); i++ {
				tmp = append(tmp, numbers[i])
				helper(i + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}

	helper(0)

	return result
}
