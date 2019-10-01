package mutations

import (
	"strings"

	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

const strLen = 8

type str [strLen]rune

var chars = [...]rune{'A', 'C', 'G', 'T'}

func (s str) Diff(another str) int {
	diff := 0
	for i := 0; i < strLen; i++ {
		if s[i] != another[i] {
			diff++
		}
	}
	return diff
}

func toStr(input string) str {
	i := 0
	var output str
	for _, char := range []rune(input)[:8] {
		output[i] = char
		i++
	}
	return output
}

func CountMutations(start, end str, bank []str) int {
	if !include(end, bank) {
		return -1
	}

	bankSet := toSet(bank)

	type QueueItem struct {
		Mutation str
		Level    int
	}

	queue := deque.New()
	queue.PushRight(QueueItem{
		Mutation: start,
	})

	for !queue.Empty() {
		current := queue.PopLeft().(QueueItem)
		if current.Mutation == end {
			return current.Level
		}

		for i := range current.Mutation {
			for _, char := range chars {
				mutationBuilder := strings.Builder{}
				mutationBuilder.WriteString(string(current.Mutation[:i]))
				mutationBuilder.WriteRune(char)
				mutationBuilder.WriteString(string(current.Mutation[i+1:]))

				mutation := toStr(mutationBuilder.String())

				if _, ok := bankSet[mutation]; ok {
					delete(bankSet, mutation)
					queue.PushRight(QueueItem{
						Mutation: mutation,
						Level:    current.Level + 1,
					})
				}

			}
		}

	}

	panic("Fix me")

}

func include(item str, slice []str) bool {
	for _, sliceItem := range slice {
		if sliceItem == item {
			return true
		}
	}

	return false
}

func toSet(slice []str) map[str]struct{} {
	result := make(map[str]struct{}, len(slice))

	for _, item := range slice {
		result[item] = struct{}{}
	}

	return result
}

//        while queue:
//            current_mutation, current_level = queue.pop(0)
//            if current_mutation == end:
//                return current_level
//
//            for index in range(len(current_mutation)):
//                for char in 'ACGT':
//                    Mutation = current_mutation[:index] + char + current_mutation[index+1:]
//                    if Mutation in bankSet:
//                        bankSet.remove(Mutation)
//                        queue.append((Mutation,current_level+1))
//
//        return -1
