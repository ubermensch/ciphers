package ciphers

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type playfairEncode struct {
	key             string
	input           string
	expectedGrid    [5][5]rune
	expectedDigrams [][]rune
}

type PlayfairTest struct {
	suite.Suite
	encodeCases []*playfairEncode
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
		{'H', 'I'},
		{'D', 'E'},
		{'T', 'H'},
		{'E', 'G'},
		{'O', 'L'},
		{'D', 'I'},
		{'N', 'T'},
		{'H', 'E'},
		{'T', 'R'},
		{'E', 'X'},
		{'E', 'S'},
		{'T', 'U'},
		{'M', 'P'},
	}

	suite.encodeCases = []*playfairEncode{
		{
			key:             "playfair example",
			input:           "hide the gold in the tree stump",
			expectedGrid:    grid,
			expectedDigrams: digrams,
		},
		// ensure the cipher ignores the non-letter chars,
		// this test case should result in identical output to the first
		{
			key:             "pla  yfa - irexample",
			input:           "h*idet%7 - he gold in the tree. stump.",
			expectedGrid:    grid,
			expectedDigrams: digrams,
		},
	}
}

func (suite *PlayfairTest) TestGrid() {
	for _, cs := range suite.encodeCases {
		key, input := cs.key, cs.input
		playfair := NewPlayfair(key, input)

		suite.Equal(
			cs.expectedGrid, playfair.grid,
		)
	}
}

func (suite *PlayfairTest) TestDigrams() {
	for _, cs := range suite.encodeCases {
		key, input := cs.key, cs.input
		playfair := NewPlayfair(key, input)

		suite.Equal(
			cs.expectedDigrams, playfair.digrams,
		)
	}
}

func TestPlayfair(t *testing.T) {
	suite.Run(t, new(PlayfairTest))
}
