package ciphers

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type caesarEncode struct {
	offsets         []int
	inputs          []string
	expectedOutputs []string
}

type caesarDecode struct {
	offsets         []int
	inputs          []string
	expectedOutputs []string
}

type CaesarTest struct {
	suite.Suite
	caesarEncode *caesarEncode
	caesarDecode *caesarDecode
}

func (suite *CaesarTest) SetupTest() {
	suite.caesarEncode = &caesarEncode{
		offsets: []int{
			7, 2, 10, 3, 5,
		},
		inputs: []string{
			"The gauls are in full retreat. Tomorrow we press the advantage!",
			"",
			"FORWARDDDDD",
			"3%^&@3#(6",
			"Today, the future is tomorrow. Tomorrow, the past is the future.",
		},
		expectedOutputs: []string{
			"Aol nhbsz hyl pu mbss ylaylha. Avtvyyvd dl wylzz aol hkchuahnl!",
			"",
			"PYBGKBNNNNN",
			"3%^&@3#(6",
			"Ytifd, ymj kzyzwj nx ytrtwwtb. Ytrtwwtb, ymj ufxy nx ymj kzyzwj.",
		},
	}

	suite.caesarDecode = &caesarDecode{
		offsets: []int{
			7, 2, 10, 3,
		},
		inputs: []string{
			"Aol nhbsz hyl pu mbss ylaylha. Avtvyyvd dl wylzz aol hkchuahnl!",
			"",
			"PYBGKBNNNNN",
			"3%^&@3#(6",
		},
		expectedOutputs: []string{
			"The gauls are in full retreat. Tomorrow we press the advantage!",
			"",
			"FORWARDDDDD",
			"3%^&@3#(6",
		},
	}
}

func (suite *CaesarTest) TestEncoding() {
	for i, offset := range suite.caesarEncode.offsets {
		cs := NewCaesar(offset)

		enc, err := cs.Encode(suite.caesarEncode.inputs[i])
		suite.Nil(err)

		suite.Equal(
			suite.caesarEncode.expectedOutputs[i],
			enc,
		)
	}
}

func (suite *CaesarTest) TestDecoding() {
	for i, offset := range suite.caesarDecode.offsets {
		cs := NewCaesar(offset)

		dec, err := cs.Decode(suite.caesarDecode.inputs[i])
		suite.Nil(err)

		suite.Equal(
			suite.caesarDecode.expectedOutputs[i],
			dec,
		)
	}
}

func TestCaesar(t *testing.T) {
	suite.Run(t, new(CaesarTest))
}
