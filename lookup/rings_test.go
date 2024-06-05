package lookup

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

var (
	lowerRing = NewAlphaRing(true)
	upperRing = NewAlphaRing(false)
)

type containsTest struct {
	ring       *AlphaRing
	lookupChar rune
	contains   bool
}

type moveTest struct {
	ring   *AlphaRing
	from   rune
	offset int
	output rune
}

type AlphaRingTest struct {
	suite.Suite
	containsCases []*containsTest
	moveCases     []*moveTest
}

func (suite *AlphaRingTest) SetupTest() {
	suite.containsCases = []*containsTest{
		{
			ring:       lowerRing,
			lookupChar: 'a',
			contains:   true,
		},
		{
			ring:       lowerRing,
			lookupChar: 'z',
			contains:   true,
		},
		{
			ring:       lowerRing,
			lookupChar: 'K',
			contains:   false,
		},
		{
			ring:       lowerRing,
			lookupChar: '5',
			contains:   false,
		},
		{
			ring:       upperRing,
			lookupChar: 'A',
			contains:   true,
		},
		{
			ring:       upperRing,
			lookupChar: 'Z',
			contains:   true,
		},
		{
			ring:       upperRing,
			lookupChar: 'j',
			contains:   false,
		},
		{
			ring:       upperRing,
			lookupChar: '%',
			contains:   false,
		},
	}
	suite.moveCases = []*moveTest{
		{
			ring:   lowerRing,
			from:   'a',
			offset: 3,
			output: 'd',
		},
		{
			ring:   lowerRing,
			from:   'z',
			offset: 2,
			output: 'b',
		},
		{
			ring:   upperRing,
			from:   'G',
			offset: 4,
			output: 'K',
		},
		{
			ring:   upperRing,
			from:   'W',
			offset: 5,
			output: 'B',
		},
	}
}

func (suite *AlphaRingTest) TestContains() {
	for _, cs := range suite.containsCases {
		suite.Equal(
			cs.contains,
			cs.ring.Contains(cs.lookupChar),
		)
	}
}

func (suite *AlphaRingTest) TestMove() {
	for _, cs := range suite.moveCases {
		output, err := cs.ring.Move(cs.from, cs.offset)
		if err != nil {
			suite.Error(err)
		}
		suite.Equal(
			cs.output,
			output,
		)
	}
}

func TestRings(t *testing.T) {
	suite.Run(t, new(AlphaRingTest))
}
