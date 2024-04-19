package ciphers

import (
	"ciphers/lookup"
)

type Caesar struct {
	offset    int
	lowerRing *lookup.AlphaRing
	upperRing *lookup.AlphaRing
	Encoder
	Decoder
}

func (c *Caesar) encodeByte(b byte) byte {
	var encoded byte
	var err error

	switch {
	case c.lowerRing.Contains(b):
		encoded, err = c.lowerRing.Move(b, c.offset)
	case c.upperRing.Contains(b):
		encoded, err = c.upperRing.Move(b, c.offset)
	default:
		encoded, err = b, nil
	}

	if err != nil {
		panic("error encoding byte")
	}

	return encoded
}

func (c *Caesar) decodeByte(b byte) byte {
	var decoded byte
	var err error

	switch {
	case c.lowerRing.Contains(b):
		decoded, err = c.lowerRing.Move(b, -c.offset)
	case c.upperRing.Contains(b):
		decoded, err = c.upperRing.Move(b, -c.offset)
	default:
		decoded, err = b, nil
	}

	if err != nil {
		panic("error decoding byte")
	}

	return decoded
}

func (c *Caesar) Encode(s string) string {
	var runes []byte
	for _, curr := range s {
		runes = append(runes, c.encodeByte(byte(curr)))
	}
	return string(runes)
}

func (c *Caesar) Decode(s string) string {
	var runes []byte
	for _, curr := range s {
		runes = append(runes, c.decodeByte(byte(curr)))
	}
	return string(runes)
}

func NewCaesar(offset int) *Caesar {
	return &Caesar{
		offset:    offset,
		lowerRing: lookup.NewAlphaRing(true),
		upperRing: lookup.NewAlphaRing(false),
	}
}
