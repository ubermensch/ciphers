package ciphers

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type EncodeTestSuite struct {
	suite.Suite
	Offsets         []int
	Inputs          []string
	ExpectedOutputs []string
}

type DecodeTestSuite struct {
	suite.Suite
	Offsets         []int
	Inputs          []string
	ExpectedOutputs []string
}

func (suite *EncodeTestSuite) SetupTest() {
	suite.Offsets = []int{
		7, 2, 10,
	}
	suite.Inputs = []string{
		"The gauls are in full retreat. Tomorrow we press the advantage!",
		"",
		"FORWARDDDDD",
	}
	suite.ExpectedOutputs = []string{
		"Aol nhbsz hyl pu mbss ylaylha. Avtvyyvd dl wylzz aol hkchuahnl!",
		"",
		"PYBGKBNNNNN",
	}
}

func (suite *EncodeTestSuite) TestEncoding() {
	for i, offset := range suite.Offsets {
		cs := NewCaesar(offset)

		suite.Equal(suite.ExpectedOutputs[i], cs.Encode(suite.Inputs[i]))
	}
}

func TestCaesar_Encode(t *testing.T) {
	suite.Run(t, new(EncodeTestSuite))
}
