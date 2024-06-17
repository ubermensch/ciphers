// https://en.wikipedia.org/wiki/Playfair_cipher
package ciphers

import (
	"ciphers/lookup"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"sync"
	"unicode"
)

type gridFuncs struct {
	// given a row or column position, returns the new one
	shiftPos func(int) int
	// given the positions of the digram elements, returns the new column value for each
	shiftRectangle func([2]int, [2]int) (int, int)
}

var encodeFuncs = gridFuncs{
	shiftPos: func(j int) int {
		shifted := j + 1
		if shifted >= 5 {
			shifted = 0
		}
		return shifted
	},
	shiftRectangle: func(firstPos [2]int, secondPos [2]int) (int, int) {
		return secondPos[1], firstPos[1]
	},
}

var decodeFuncs = gridFuncs{
	shiftPos: func(j int) int {
		shifted := j - 1
		if shifted < 0 {
			shifted = 4
		}
		return shifted
	},
	shiftRectangle: func(firstPos [2]int, secondPos [2]int) (int, int) {
		return secondPos[1], firstPos[1]
	},
}

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

func (p *Playfair) encodeDigram(dg *[]rune) (*[]rune, error) {
	var encodedDigram *[]rune
	firstPos, secondPos := p.digramPos(*dg)
	switch {
	case p.isRow(*dg):
		encodedDigram = &[]rune{
			p.grid[firstPos[0]][encodeFuncs.shiftPos(firstPos[1])],
			p.grid[secondPos[0]][encodeFuncs.shiftPos(secondPos[1])],
		}
	case p.isColumn(*dg):
		encodedDigram = &[]rune{
			p.grid[encodeFuncs.shiftPos(firstPos[0])][firstPos[1]],
			p.grid[encodeFuncs.shiftPos(secondPos[0])][secondPos[1]],
		}
	case p.isRectangle(*dg):
		firstNewCol, secondNewCol := encodeFuncs.shiftRectangle(firstPos, secondPos)
		encodedDigram = &[]rune{
			p.grid[firstPos[0]][firstNewCol],
			p.grid[secondPos[0]][secondNewCol],
		}
	default:
		return nil, errors.New(
			fmt.Sprintf("digram does not form recognized shape: %s", string(*dg)),
		)
	}

	return encodedDigram, nil
}

func (p *Playfair) decodeDigram(dg *[]rune) (*[]rune, error) {
	var decodedDigram *[]rune
	firstPos, secondPos := p.digramPos(*dg)
	switch {
	case p.isRow(*dg):
		decodedDigram = &[]rune{
			p.grid[firstPos[0]][decodeFuncs.shiftPos(firstPos[1])],
			p.grid[secondPos[0]][decodeFuncs.shiftPos(secondPos[1])],
		}
	case p.isColumn(*dg):
		decodedDigram = &[]rune{
			p.grid[decodeFuncs.shiftPos(firstPos[0])][firstPos[1]],
			p.grid[decodeFuncs.shiftPos(secondPos[0])][secondPos[1]],
		}
	case p.isRectangle(*dg):
		firstNewCol, secondNewCol := decodeFuncs.shiftRectangle(firstPos, secondPos)
		decodedDigram = &[]rune{
			p.grid[firstPos[0]][firstNewCol],
			p.grid[secondPos[0]][secondNewCol],
		}
	default:
		return nil, errors.New(
			fmt.Sprintf("digram does not form recognized shape: %s", string(*dg)),
		)
	}

	return decodedDigram, nil
}

func (p *Playfair) Encode(input string) (string, error) {
	if len(p.key) == 0 {
		return "", errors.New("empty key")
	}

	// build digrams from input
	p.digrams = getDigrams(input)

	encodedDigrams := make([]string, len(p.digrams))
	wg := sync.WaitGroup{}
	errCount := 0

	encFunc := func(dg []rune, pos int, wg *sync.WaitGroup) {
		defer wg.Done()
		enc, err := p.encodeDigram(&dg)
		if err != nil {
			errCount++
		}
		e := *enc
		encStr := fmt.Sprintf(`%s%s`, string(e[0]), string(e[1]))
		encodedDigrams[pos] = encStr
	}

	for i, dg := range p.digrams {
		wg.Add(1)
		go encFunc(dg, i, &wg)
	}

	wg.Wait()
	if errCount > 0 {
		return "", errors.New("encoding failed")
	}

	return strings.Join(encodedDigrams, " "), nil
}

func (p *Playfair) Decode(input string) (string, error) {
	if len(p.key) == 0 {
		return "", errors.New("empty key")
	}

	// build digrams from input
	p.digrams = getDigrams(input)

	decodedDigrams := make([]string, len(p.digrams))
	wg := sync.WaitGroup{}
	errCount := 0

	decFunc := func(dg []rune, pos int, wg *sync.WaitGroup) {
		defer wg.Done()
		dec, err := p.decodeDigram(&dg)
		if err != nil {
			errCount++
		}
		d := *dec
		decStr := fmt.Sprintf(`%s%s`, string(d[0]), string(d[1]))
		decodedDigrams[pos] = decStr
	}

	for i, dg := range p.digrams {
		wg.Add(1)
		go decFunc(dg, i, &wg)
	}

	wg.Wait()
	if errCount > 0 {
		return "", errors.New("decoding failed")
	}

	return strings.Join(decodedDigrams, " "), nil
}

func NewPlayfair(key string) *Playfair {
	// build grid from key
	grid := gridFromKey(key)

	return &Playfair{
		key:  key,
		grid: grid,
	}
}
