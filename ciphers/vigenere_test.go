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
			"b",
			"tomorrow",
		},
		inputs: []string{
			"attackatdawn",
			"abcdefg",
			"To whom does one belong? To the land.",
		},
		expectedOutputs: []string{
			"lxfopvefrnhr",
			"bcdefgh",
			"Mc kyfa wcqg fba pqzfeu? Ha kys eozr.",
		},
	}

	suite.vigenereDecode = &vigenereDecode{
		keys: []string{
			"lemon",
			"b",
			"tomorrow",
		},
		inputs: []string{
			"lxfopvefrnhr",
			"bcdefgh",
			"Mc kyfa wcqg fba pqzfeu? Ha kys eozr.",
		},
		expectedOutputs: []string{
			"attackatdawn",
			"abcdefg",
			"To whom does one belong? To the land.",
		},
	}
}

func (suite *VigenereTest) TestEncoding() {
	for i, key := range suite.vigenereEncode.keys {
		vg := NewVigenere(key)

		enc, err := vg.Encode(suite.vigenereEncode.inputs[i])

		suite.Nil(err)
		suite.Equal(
			suite.vigenereEncode.expectedOutputs[i],
			enc,
		)
	}
}

func (suite *VigenereTest) TestDecoding() {
	for i, key := range suite.vigenereDecode.keys {
		vg := NewVigenere(key)

		dec, err := vg.Decode(suite.vigenereDecode.inputs[i])

		suite.Nil(err)
		suite.Equal(
			suite.vigenereDecode.expectedOutputs[i],
			dec,
		)
	}
}

func (suite *VigenereTest) TestErrors() {
	// should error with empty string key
	vg := NewVigenere("")
	_, err := vg.Encode("this won't work")
	suite.NotNil(err)
	suite.Equal("empty key", err.Error())
}

func TestVigenere(t *testing.T) {
	suite.Run(t, new(VigenereTest))
}
