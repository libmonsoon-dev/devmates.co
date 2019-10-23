package solution

import "fmt"

var lookup = intToRoman{
	1:    []byte("I"),
	4:    []byte("IV"),
	5:    []byte("V"),
	9:    []byte("IX"),
	10:   []byte("X"),
	40:   []byte("XL"),
	50:   []byte("L"),
	90:   []byte("XC"),
	100:  []byte("C"),
	400:  []byte("CD"),
	500:  []byte("D"),
	900:  []byte("CM"),
	1000: []byte("M"),
}

type intToRoman map[int][]byte

func (lookup intToRoman) Get(key int) []byte {
	val, ok := lookup[key]

	if !ok {
		panic(fmt.Sprintf("intToRoman.Get(%v) item not found", key))
	}

	return val

}
