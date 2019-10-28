package solution

import "fmt"

type operation bool
type operationList []operation

func (o operation) String() string {
	if o {
		return "+"
	}
	return "-"
}

const (
	add operation = true
	sub operation = false
)

var passableOperations = [...]operation{add, sub}

func (ops operationList) Allay(numbers []int) int {
	if len(ops) != len(numbers) {
		panic(fmt.Sprintf("len(operationList) != len(numbers) (%v != %v)", len(ops), len(numbers)))
	}

	var result int

	for i := 0; i < len(ops); i++ {
		if ops[i] == add {
			result += numbers[i]
		} else {
			result -= numbers[i]
		}
	}

	return result
}

func operationGenerator(operationsLen int) <-chan operationList {
	ch := make(chan operationList)

	go func(ch chan<- operationList) {
		defer close(ch)

		pools := make([][]operation, operationsLen)

		for i := 0; i < operationsLen; i++ {
			pools[i] = passableOperations[:]
		}
		fmt.Println(pools)

		result := make([]operationList, 0, 1)
		result = append(result, make(operationList, 0, operationsLen))

		for _, pool := range pools {
			tmp := make([]operationList, 0, len(pools)*len(result))

			for _, y := range pool {
				for _, x := range result {
					item := make(operationList, 0, operationsLen)

					item = append(item, x...)
					item = append(item, y)

					tmp = append(tmp, item)
				}
			}
			result = tmp
		}

		for _, prod := range result {
			ch <- prod
		}
	}(ch)

	return ch
}

func Solution(list []int, target int) int {
	result := 0

	for operations := range operationGenerator(len(list)) {
		fmt.Println(operations, operations.Allay(list))
		if operations.Allay(list) == target {
			result++
		}
	}

	return result
}
