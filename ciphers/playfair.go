// https://en.wikipedia.org/wiki/Playfair_cipher
package ciphers

import (
	"ciphers/lookup"
	"errors"
	"fmt"
	lo "github.com/samber/lo"
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
	rx := regexp.MustCompile(`[^a-zA-Z]`)
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

func (p *Playfair) digramPos(dg []rune) ([2]int, [2]int) {
	diFirst, diSecond := [2]int{}, [2]int{}
	i, j := 0, 0

	for i < 5 {
		j = 0
		for j < 5 {
			switch {
			case p.grid[i][j] == dg[0]:
				diFirst = [2]int{i, j}
			case p.grid[i][j] == dg[1]:
				diSecond = [2]int{i, j}
			}
			j++
		}
		i++
	}

	return diFirst, diSecond
}

// do the positions of these 2 runes in the digram form a rectangle?
func (p *Playfair) isRectangle(dg []rune) bool {
	firstPos, secondPos := p.digramPos(dg)
	if firstPos[0] != secondPos[0] && firstPos[1] != secondPos[1] {
		return true
	}
	return false
}

// are the positions of these 2 runes in the digram on the same row?
func (p *Playfair) isRow(dg []rune) bool {
	firstPos, secondPos := p.digramPos(dg)
	if firstPos[0] == secondPos[0] {
		return true
	}
	return false
}

// are the positions of these 2 runes in the digram on the same column?
func (p *Playfair) isColumn(dg []rune) bool {
	firstPos, secondPos := p.digramPos(dg)
	if firstPos[1] == secondPos[1] {
		return true
	}
	return false
}

func (p *Playfair) encodeDigram(dg *[]rune) *[]rune {
	// given a row or column position (i or j), returns the new one
	shiftPos := func(j int) int {
		shifted := j + 1
		if shifted >= 5 {
			shifted = 0
		}
		return shifted
	}

	// given the positions of the digram elements, returns the new j value for each
	shiftRectangle := func(firstPos [2]int, secondPos [2]int) (int, int) {
		return secondPos[1], firstPos[1]
	}

	var encodedDigram *[]rune
	firstPos, secondPos := p.digramPos(*dg)
	switch {
	case p.isRow(*dg):
		encodedDigram = &[]rune{
			p.grid[firstPos[0]][shiftPos(firstPos[1])],
			p.grid[secondPos[0]][shiftPos(secondPos[1])],
		}
	case p.isColumn(*dg):
		encodedDigram = &[]rune{
			p.grid[shiftPos(firstPos[0])][firstPos[1]],
			p.grid[shiftPos(secondPos[0])][secondPos[1]],
		}
	case p.isRectangle(*dg):
		firstNewCol, secondNewCol := shiftRectangle(firstPos, secondPos)
		encodedDigram = &[]rune{
			p.grid[firstPos[0]][firstNewCol],
			p.grid[secondPos[0]][secondNewCol],
		}
	default:
		panic(
			errors.New(
				fmt.Sprintf("digram does not form recognized shape: %s", string(*dg)),
			),
		)
	}

	return encodedDigram
}

func (p *Playfair) decodeDigram(dg *[]rune) *[]rune {
	return &[]rune{}
}

func (p *Playfair) Encode(input string) (string, error) {
	encodedDigrams := lo.Map(p.digrams, func(dg []rune, i int) string {
		e := *p.encodeDigram(&dg)
		return fmt.Sprintf(`%s%s`, string(e[0]), string(e[1]))
	})

	return strings.Join(encodedDigrams, " "), nil
}

func (p *Playfair) Decode(input string) (string, error) {
	return "", nil
}

func NewPlayfair(key string, input string) *Playfair {
	// build grid from key
	grid := gridFromKey(key)

	// build digrams from input
	digrams := getDigrams(input)

	return &Playfair{
		key:     key,
		grid:    grid,
		digrams: digrams,
	}
}
