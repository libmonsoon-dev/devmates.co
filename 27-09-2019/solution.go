package solution

import "image"

func Solution(matrix [][]int) int {
	isValidSquare := func(upperLeft image.Point, size int) bool {
		for x := upperLeft.X; x < upperLeft.X+size; x++ {
			for y := upperLeft.Y; y < upperLeft.Y+size; y++ {
				if matrix[x][y] == 0 {
					return false
				}
			}
		}
		return true
	}
	maxSideSize := 0

	xLen := len(matrix)
	for x := 0; x < xLen; x++ {
		yLen := len(matrix[x])
		for y := 0; y < yLen; y++ {
			if matrix[x][y] == 1 {
				valid := true
				currentSize := 1

				if currentSize > maxSideSize {
					maxSideSize = currentSize
				}

				newPtr := image.Pt(x, y)

				for valid {
					currentSize++
					newPtr.X += currentSize
					newPtr.Y += currentSize

					if newPtr.X <= xLen && newPtr.Y <= yLen {
						valid = isValidSquare(image.Pt(x, y), currentSize)

						if valid && currentSize > maxSideSize {
							maxSideSize = currentSize
						}
					} else {
						valid = false
					}
				}
			}
		}
	}

	return maxSideSize * maxSideSize
}
