package ciphers

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type vigenereEncode struct {
	keys            []string
	inputs          []string
	expectedOutputs []string
}

type vigenereDecode struct {
	keys            []string
	inputs          []string
	expectedOutputs []string
}

type VigenereTest struct {
	suite.Suite
	vigenereEncode *vigenereEncode
	vigenereDecode *vigenereDecode
}

func (suite *VigenereTest) SetupTest() {
	suite.vigenereEncode = &vigenereEncode{
		keys: []string{
			"lemon",
		},
		inputs: []string{
			"attackatdawn",
		},
		expectedOutputs: []string{
			"lxfopvefrnhr",
		},
	}

	suite.vigenereDecode = &vigenereDecode{
		keys: []string{
			"lemon",
		},
		inputs: []string{
			"lxfopvefrnhr",
		},
		expectedOutputs: []string{
			"attackatdawn",
		},
	}
}

func (suite *VigenereTest) TestEncoding() {
	for i, key := range suite.vigenereEncode.keys {
		vg := NewVigenere(key)

		suite.Equal(
			suite.vigenereEncode.expectedOutputs[i],
			vg.Encode(suite.vigenereEncode.inputs[i]),
		)
	}
}

func (suite *VigenereTest) TestDecoding() {
	for i, key := range suite.vigenereDecode.keys {
		vg := NewVigenere(key)

		suite.Equal(
			suite.vigenereDecode.expectedOutputs[i],
			vg.Decode(suite.vigenereDecode.inputs[i]),
		)
	}
}

func TestVigenere(t *testing.T) {
	suite.Run(t, new(VigenereTest))
}
