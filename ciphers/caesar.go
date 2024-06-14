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

func (c *Caesar) encodeChar(b rune) (rune, error) {
	var encoded rune
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
		return 0, err
	}

	return encoded, nil
}

func (c *Caesar) decodeChar(b rune) (rune, error) {
	var decoded rune
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
		return 0, err
	}

	return decoded, nil
}

func (c *Caesar) Encode(s string) (string, error) {
	var runes []rune
	for _, curr := range s {
		enc, err := c.encodeChar(curr)
		if err != nil {
			return "", err
		}
		runes = append(runes, enc)
	}
	return string(runes), nil
}

func (c *Caesar) Decode(s string) (string, error) {
	var runes []rune
	for _, curr := range s {
		dec, err := c.decodeChar(curr)
		if err != nil {
			return "", err
		}
		runes = append(runes, dec)
	}
	return string(runes), nil
}

func NewCaesar(offset int) *Caesar {
	return &Caesar{
		offset:    offset,
		lowerRing: lookup.NewAlphaRing(true),
		upperRing: lookup.NewAlphaRing(false),
	}
}
