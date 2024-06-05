package lookup

import (
	"container/ring"
	"errors"
	"slices"
)

type AlphaRing struct {
	items   *ring.Ring
	lower   bool
	letters []rune
}

var lowerLetters = []rune{
	'a', 'b', 'c', 'd', 'e',
	'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o',
	'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z',
}

var upperLetters = []rune{
	'A', 'B', 'C', 'D', 'E',
	'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O',
	'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z',
}

func (r *AlphaRing) Contains(c rune) bool {
	return slices.Index(r.letters, c) > -1
}

// Returns the byte `i` positions ahead or behind the `from` byte
func (r *AlphaRing) Move(from rune, i int) (rune, error) {
	if !r.Contains(from) {
		return 0, errors.New("byte not present in AlphaRing")
	}
	iFrom := slices.Index(r.letters, from)
	result := r.items.Move(iFrom + i).Value.(rune)
	return result, nil
}

func NewAlphaRing(lower bool) *AlphaRing {
	alphaItems := func() []rune {
		if lower {
			return lowerLetters
		}
		return upperLetters
	}()

	items := ring.New(26)
	for i := 0; i < 26; i++ {
		items.Value = alphaItems[i]
		items = items.Next()
	}

	return &AlphaRing{
		items:   items,
		lower:   lower,
		letters: alphaItems,
	}
}
