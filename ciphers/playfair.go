// https://en.wikipedia.org/wiki/Playfair_cipher
package ciphers

import (
	"ciphers/lookup"
	"regexp"
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

// Playfair doesn't handle non-letter chars and expects upper case.
func prepareInput(input string) string {
	rx := regexp.MustCompile(`\W+`)
	return strings.ToUpper(rx.ReplaceAllString(input, ``))
}

func gridFromKey(startKey string) [5][5]rune {
	key := prepareInput(startKey)
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

	nextFiller := func() rune {
		next, err := filler.Move('A', fillerFrom)
		if err != nil {
			panic(err)
		}
		fillerFrom++
		return rune(next)
	}

	nextChar := func() rune {
		var next rune

		if len(key) > 0 {
			next = rune(key[0])
			key = key[1:]
		} else {
			next = nextFiller()
		}
		return next
	}

	// Is the given rune present in usedChars, and is it one of the interchangeable
	// chars ('I' or 'J')
	isInterchangeable := func(c rune) bool {
		return slices.ContainsFunc(usedChars, func(u rune) bool {
			inter := []rune{'I', 'J'}
			return slices.Contains(inter, c) && slices.Contains(inter, u)
		})
	}

	for i < 5 {
		// get next (from key if present, otherwise nextFiller())
		next := nextChar()

		// keep getting next while:
		//    1) next is not a letter
		//    2) next is present in usedChars
		//    3) next is 'I' or 'J' and 'I' or 'J' present in usedChars
		for !unicode.IsLetter(next) || slices.Contains(usedChars, next) || isInterchangeable(next) {
			next = nextChar()
		}

		grid[i][j] = next
		usedChars = append(usedChars, next)

		// move the grid pointers one space to right or to newline if filled
		// current []rune{} slice
		nextGridPos()
	}

	return grid
}

func getDigrams(input string) [][]rune {
	str := prepareInput(input)
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

func (p *Playfair) Encode(input string) (string, error) {
	return "", nil
}

func (p *Playfair) Decode(input string) (string, error) {
	return "", nil
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
