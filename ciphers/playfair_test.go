package ciphers

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type playfairEncode struct {
	key     string
	input   string
	grid    [5][5]rune
	digrams [][]rune
	output  string
}

type playfairDecode struct {
	key     string
	input   string
	grid    [5][5]rune
	digrams [][]rune
	output  string
}

type PlayfairTest struct {
	suite.Suite
	encodeCases []*playfairEncode
	decodeCases []*playfairDecode
}

func (suite *PlayfairTest) SetupTest() {
	grid := [5][5]rune{
		{'P', 'L', 'A', 'Y', 'F'},
		{'I', 'R', 'E', 'X', 'M'},
		{'B', 'C', 'D', 'G', 'H'},
		{'K', 'N', 'O', 'Q', 'S'},
		{'T', 'U', 'V', 'W', 'Z'},
	}

	digrams := [][]rune{
		{'H', 'I'}, {'D', 'E'}, {'T', 'H'},
		{'E', 'G'}, {'O', 'L'}, {'D', 'I'},
		{'N', 'T'}, {'H', 'E'}, {'T', 'R'},
		{'E', 'X'}, {'E', 'S'}, {'T', 'U'},
		{'M', 'P'},
	}

	output := "BM OD ZB XD NA BE KU DM UI XM MO UV IF"

	encodedDigrams := [][]rune{
		{'B', 'M'}, {'O', 'D'}, {'Z', 'B'},
		{'X', 'D'}, {'N', 'A'}, {'B', 'E'},
		{'K', 'U'}, {'D', 'M'}, {'U', 'I'},
		{'X', 'M'}, {'M', 'O'}, {'U', 'V'},
		{'I', 'F'},
	}

	suite.encodeCases = []*playfairEncode{
		{
			key:     "playfair example",
			input:   "hide the gold in the tree stump",
			grid:    grid,
			digrams: digrams,
			output:  output,
		},
		// ensure the cipher ignores the non-letter chars,
		// this test case should result in identical output to the first
		{
			key:     "pla  yfa - irexample",
			input:   "h*idet%7 - he gold in the tree. stump.",
			grid:    grid,
			digrams: digrams,
			output:  output,
		},
	}

	suite.decodeCases = []*playfairDecode{
		{
			key:     "playfair example",
			input:   output,
			output:  "HI DE TH EG OL DI NT HE TR EX ES TU MP",
			grid:    grid,
			digrams: encodedDigrams,
		},
	}
}

func (suite *PlayfairTest) TestGrid() {
	for _, cs := range suite.encodeCases {
		playfair := NewPlayfair(cs.key)

		suite.Equal(
			cs.grid, playfair.grid,
		)
	}
}

func (suite *PlayfairTest) TestEncode() {
	for _, cs := range suite.encodeCases {
		key, input := cs.key, cs.input
		playfair := NewPlayfair(key)
		enc, err := playfair.Encode(input)

		suite.Nil(err)
		suite.Equal(
			cs.output, enc,
		)
		suite.Equal(
			cs.digrams, playfair.digrams,
		)
	}
}

func (suite *PlayfairTest) TestDecode() {
	for _, cs := range suite.decodeCases {
		key, input := cs.key, cs.input
		playfair := NewPlayfair(key)
		dec, err := playfair.Decode(input)

		suite.Nil(err)
		suite.Equal(
			cs.output, dec,
		)

		suite.Equal(
			cs.digrams, playfair.digrams,
		)
	}
}

func TestPlayfair(t *testing.T) {
	suite.Run(t, new(PlayfairTest))
}
