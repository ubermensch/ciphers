// https://en.wikipedia.org/wiki/Playfair_cipher
package ciphers

import (
	"ciphers/lookup"
	"slices"
)

type Playfair struct {
	// cipher key used to build grid
	key string
	// 5 x 5 cipher grid built from the key
	grid [5][5]rune
	// slice of 2-character digrams from the message to encrypt/decrypt
	digrams [][]rune
	Encoder
	Decoder
}

func (p *Playfair) encodeByte(b byte) byte {
	return b
}

func (p *Playfair) decodeByte(b byte) byte {
	return b
}

func gridFromKey(key string) [5][5]rune {
	var grid = [5][5]rune{}
	var usedChars = []rune{}
	fillerFrom := 0
	filler := lookup.NewAlphaRing(false)
	i, j := 0, 0

	nextGridPos := func() {
		if i == 4 {
			i = 0
			j++
			return
		}
		i++
	}

	nextAlpha := func() rune {
		next := filler.Items.Move(fillerFrom + 1).Value.(rune)
		fillerFrom++
		return next
	}

	for j < 5 {
		var next rune

		if len(key) > 0 {
			// If we still have chars of the key, get the next one not
			// already present in usedChars
			for slices.Contains(usedChars, rune(key[0])) {
				next = rune(key[0])
				key = key[1:]
			}
		} else {
			// If we've finished the key chars, get the next alphabetical
			// character that isn't already present in the key, treating
			// 'I' and 'J' and interchangeable
			var nextA = nextAlpha()
			for slices.Contains(usedChars, nextA) {
				nextA = nextAlpha()
				if slices.ContainsFunc(usedChars, func(u rune) bool {
					inter := []rune{'I', 'J'}
					return slices.Contains(inter, nextA) && slices.Contains(inter, u)
				}) {
					nextA = nextAlpha()
				}
			}

			next = nextA
		}

		grid[i][j] = next
		usedChars = append(usedChars, next)

		nextGridPos()
	}

	return grid
}

func NewPlayfair(key string) *Playfair {
	// build grid from key

	// build digrams from key

	// build
	return &Playfair{}
}
