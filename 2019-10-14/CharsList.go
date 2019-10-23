package solution

type CharsList struct {
	chars       []rune
	lenToDelete int
}

func (c *CharsList) String() string {
	return string(c.chars)
}

func (c *CharsList) RemoveDuplicates() bool {
	var repeatingChar rune
	var charCount int

	for index, char := range c.chars {
		if char == repeatingChar {
			charCount++

			if charCount == c.lenToDelete {
				c.removeItems(index+1-c.lenToDelete, index+1)
				return true
			}
		} else {
			repeatingChar = char
			charCount = 1
		}

	}

	return false
}

func (c *CharsList) removeItems(fromIndex, toIndex int) {
	//TODO: use linked list
	c.chars = append(c.chars[:fromIndex], c.chars[toIndex:]...)
}
