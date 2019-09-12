package barcode

type Barcodes []int

func (bc *Barcodes) Rearrange() {
	items := bc.CountItems()
	length := len(*bc)
	index := 0
	result := Barcodes(make([]int, length))

	for item, count := range items {
		for i := 0; i < count; i++ {
			result[index] = item
			index += 2
			if index >= length {
				index = 1
			}
		}
	}

	*bc = result
}

func (bc *Barcodes) CountItems() map[int]int {
	result := make(map[int]int)

	for _, item := range *bc {
		result[item]++
	}

	return result
}
