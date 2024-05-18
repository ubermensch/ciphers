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
	playfairEncode *playfairEncode
}

func (suite *PlayfairTest) SetupTest() {
	suite.playfairEncode = &playfairEncode{
		key:   "playfair example",
		input: "hide the gold in the tree stump",
		expectedGrid: [5][5]rune{
			{'P', 'L', 'A', 'Y', 'F'},
			{'I', 'R', 'E', 'X', 'M'},
			{'B', 'C', 'D', 'G', 'H'},
			{'K', 'N', 'O', 'Q', 'S'},
			{'T', 'U', 'V', 'W', 'Z'},
		},
		expectedDigrams: [][]rune{
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
		},
	}
}

func (suite *PlayfairTest) TestGrid() {
	key, input := suite.playfairEncode.key, suite.playfairEncode.input
	playfair := NewPlayfair(key, input)

	suite.Equal(
		suite.playfairEncode.expectedGrid, playfair.grid,
	)
}

func (suite *PlayfairTest) TestDigrams() {
	key, input := suite.playfairEncode.key, suite.playfairEncode.input
	playfair := NewPlayfair(key, input)

	suite.Equal(
		suite.playfairEncode.expectedDigrams, playfair.digrams,
	)
}

func TestPlayfair(t *testing.T) {
	suite.Run(t, new(PlayfairTest))
}
