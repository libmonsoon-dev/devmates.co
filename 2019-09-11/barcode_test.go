package barcode

import (
	"reflect"
	"testing"
)

func hasTwoEqualAdjacentBarcode(code Barcodes) bool {
	for i := 1; i < len(code)-1; i++ {
		if code[i-1] == code[i] || code[i] == code[i+1] {
			return true
		}
	}

	return false
}

func sameItems(a, b Barcodes) bool {
	return reflect.DeepEqual(
		a.CountItems(),
		b.CountItems(),
	)
}

func TestBarcode_Rearrange(t *testing.T) {
	tests := []Barcodes{
		{1, 1, 1, 2, 2, 2},
		{1, 1, 1, 1, 2, 2, 3, 3},
	}
	for _, orig := range tests {
		code := Barcodes(append([]int(nil), orig...)) //copy of orig

		code.Rearrange()

		if hasTwoEqualAdjacentBarcode(code) {
			t.Errorf("In %v two adjacent barcodes are equal.", code)
		}

		if !sameItems(code, orig) {
			t.Errorf("Items in %v and %v are different", orig, code)
		}

	}
}
