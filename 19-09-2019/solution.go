package solution

func RecursiveSolution(n int) []string {
	var result []string

	var helper func(current string, open, closed int)

	helper = func(current string, open, closed int) {
		if len(current) == n*2 {
			result = append(result, current)
			return
		}

		if open < n {
			helper(current+"(", open+1, closed)
		}

		if open > closed {
			helper(current+")", open, closed+1)
		}

	}

	helper("", 0, 0)

	return result
}

type resultBuilder struct {
	Current string
	Open    int
	Closed  int
}

func (r *resultBuilder) String() string {
	return r.Current
}

func Solution(n int) []string {

	builders := []resultBuilder{{}}
	var nextStepBuilders []resultBuilder

	addOpen := func(builderCopy resultBuilder) {
		builderCopy.Current += "("
		builderCopy.Open++
		nextStepBuilders = append(nextStepBuilders, builderCopy)
	}

	addClosed := func(builderCopy resultBuilder) {
		builderCopy.Current += ")"
		builderCopy.Closed++
		nextStepBuilders = append(nextStepBuilders, builderCopy)
	}

	for i := 0; i < n*2; i++ {

		for _, builder := range builders {
			if builder.Open < n {
				addOpen(builder)
			}

			if builder.Open > builder.Closed {
				addClosed(builder)
			}

		}

		builders = nextStepBuilders
		nextStepBuilders = nil
	}

	result := make([]string, 0, len(nextStepBuilders))

	for _, builder := range builders {
		result = append(result, builder.String())
	}

	return result
}
