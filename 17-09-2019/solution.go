package solution

import (
	"math"
	"sort"
)

type task struct {
	Task  string
	Count int
}

type taskCounts []task

func (t taskCounts) Len() int {
	return len(t)
}

func (t taskCounts) Less(i, j int) bool {
	return t[i].Count < t[j].Count
}

func (t taskCounts) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func Solution(tasks []string, n int) int {
	tasksLookup := make(map[string]int)

	for i := 0; i < len(tasks); i++ {
		tasksLookup[tasks[i]]++
	}

	taskCount := make(taskCounts, 0, len(tasksLookup))

	for key, value := range tasksLookup {
		taskCount = append(
			taskCount,
			task{
				Task:  key,
				Count: value,
			},
		)
	}

	sort.Sort(taskCount)

	maxTaskCount := taskCount[0].Count - 1
	idleSlots := maxTaskCount * n

	for i := 1; i < len(taskCount); i++ {
		idleSlots -= int(math.Min(float64(taskCount[i].Count), float64(maxTaskCount)))
	}

	if idleSlots > 0 {
		return idleSlots + len(tasks)
	}
	return len(tasks)
}
