package fill

type Pixel struct {
	X int
	Y int
}

func Fill(img *[][]int, sr, sc, newColor int) {
	inputTarget := Pixel{X: sr, Y: sc}
	oldColor := (*img)[sr][sc]

	toChange := make(map[Pixel]struct{})

	// It should have been the queue, but I was too lazy to implement it
	// Also, I did not want to use libraries, because didn't want to cast interface{} to Pixel
	// Slice-based solution leads to memory leak, channel-based solution - can lead to deadlock
	toCheckQueue := make(map[Pixel]struct{})
	toCheckQueue[inputTarget] = struct{}{}

	isNeedRepaint := func(p Pixel) bool {
		_, exist := toChange[p]
		if !exist &&
			p.X >= 0 &&
			p.Y >= 0 &&
			len(*img) > p.X &&
			len((*img)[p.X]) > p.Y &&
			(*img)[p.X][p.Y] == oldColor {
			return true
		}
		return false
	}

	for {
		if len(toCheckQueue) == 0 {
			break
		}

		for p := range toCheckQueue {
			delete(toCheckQueue, p)

			if isNeedRepaint(p) {
				toChange[p] = struct{}{}

				toCheckQueue[Pixel{p.X + 1, p.Y}] = struct{}{}
				toCheckQueue[Pixel{p.X - 1, p.Y}] = struct{}{}
				toCheckQueue[Pixel{p.X, p.Y + 1}] = struct{}{}
				toCheckQueue[Pixel{p.X, p.Y - 1}] = struct{}{}

			}
		}

	}

	for p := range toChange {
		(*img)[p.X][p.Y] = newColor
	}

}
