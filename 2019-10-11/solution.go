package solution

import (
	"math"
	"strconv"
	"strings"
)

func convertToMinutes(point string) int {
	date := strings.Split(point, ":")

	hours, err := strconv.Atoi(date[0])
	if err != nil {
		panic(err)
	}

	minutes, err := strconv.Atoi(date[1])
	if err != nil {
		panic(err)
	}

	return hours*60 + minutes
}

func Solution(input []string) int {
	minutes := [24 * 60]bool{}
	currentMinute := 0
	previousMinute := 0

	for i := 0; i < len(input); i++ {
		currentMinute = convertToMinutes(input[i])

		if minutes[currentMinute] {
			return 0
		}

		minutes[currentMinute] = true
	}

	minDiff := math.MaxInt64
	first := math.MaxInt64
	last := math.MaxInt64

	for currentMinute := 0; currentMinute < len(minutes); currentMinute++ {
		if minutes[currentMinute] {
			if first != math.MaxInt64 {
				minDiff = int(math.Min(
					float64(minDiff),
					float64(currentMinute-previousMinute),
				))
			}

			first = int(math.Min(
				float64(first),
				float64(currentMinute),
			))

			last = int(math.Min(
				float64(last),
				float64(currentMinute),
			))

			previousMinute = currentMinute
		}
	}

	minDiff = int(math.Min(
		float64(minDiff),
		float64(60*24-last+first),
	))

	return minDiff
}
