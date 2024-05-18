// https://en.wikipedia.org/wiki/Playfair_cipher
package ciphers

import (
	"ciphers/lookup"
	"slices"
	"strings"
	"unicode"
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

func gridFromKey(startKey string) [5][5]rune {
	key := startKey
	var grid = [5][5]rune{}
	var usedChars = []rune{}
	fillerFrom := 0
	filler := lookup.NewAlphaRing(false)
	i, j := 0, 0

	nextGridPos := func() {
		if j == 4 {
			j = 0
			i++
			return
		}
		j++
	}

	nextAlpha := func() rune {
		next, err := filler.Move('A', fillerFrom)
		if err != nil {
			panic(err)
		}
		fillerFrom++
		return rune(next)
	}

	// TODO - refactor this loop:
	// * first, get nextA (from key if present, otherwise nextAlpha())
	// * keep getting nextA while:
	//    1) nextA is not a letter
	//    2) nextA is present in usedChars
	//    3) nextA is 'I' or 'J' and 'I' or 'J' present in usedChars
	for i < 5 {
		var next rune

		if len(key) > 0 {
			// If we still have chars of the key, get the next one not
			// already present in usedChars
			next = rune(key[0])
			key = key[1:]
			for slices.Contains(usedChars, next) || !unicode.IsLetter(next) {
				if len(key) > 0 {
					next = rune(key[0])
					key = key[1:]
				} else {
					next = nextAlpha()
				}
			}
		} else {
			// If we've finished the key chars, get the next alphabetical
			// character that isn't already present in the key, treating
			// 'I' and 'J' as interchangeable
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

func getDigrams(input string) [][]rune {
	var str = strings.ToUpper(input)
	digrams := [][]rune{}

	takeTwo := func() {
		digrams = append(digrams, []rune{rune(str[0]), rune(str[1])})
		str = str[2:]
	}

	takeOne := func() {
		digrams = append(digrams, []rune{rune(str[0]), 'X'})
		str = str[1:]
	}

	for len(str) > 0 {
		// If there are at least 2 chars left
		if len(str) > 1 {
			// ...and they are not the same, form a digram from them
			if str[0] != str[1] {
				takeTwo()
			} else {
				// ...if they are the same (e.g. double 'e' in 'tree'), pad with 'X'
				takeOne()
			}
		} else {
			takeOne()
		}
	}

	return digrams
}

func NewPlayfair(key string, input string) *Playfair {
	// build grid from key
	grid := gridFromKey(strings.ToUpper(key))

	// build digrams from input
	digrams := getDigrams(strings.ToUpper(input))

	return &Playfair{
		key:     key,
		grid:    grid,
		digrams: digrams,
	}
}
