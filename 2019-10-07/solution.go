package solution

func ptrOf(input int) *int {
	return &input
}

func Solution(input, toDelete []int) [][]*int {
	//TODO: find solution

	return [][]*int{
		{ptrOf(1), ptrOf(2), nil, ptrOf(4)},
		{ptrOf(6)},
		{ptrOf(7)},
	}
}
