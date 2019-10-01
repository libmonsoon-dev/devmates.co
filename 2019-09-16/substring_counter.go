package substring_counter

func Count(str string) int {
	var result int
	w := struct {
		Start  int
		End    int
		Window map[string]struct{}
	}{
		Window: make(map[string]struct{}),
	}
	strLen := len(str)

	for w.Start < strLen && w.End < strLen {
		end := string(str[w.End])
		if _, ok := w.Window[end]; ok {
			delete(w.Window, string(str[w.Start]))
			w.Start++
		} else {
			w.Window[end] = struct{}{}
			w.End++
			result = max(result, w.End-w.Start)
		}
	}

	return result

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
