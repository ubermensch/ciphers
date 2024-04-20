package ciphers

import (
	"ciphers/lookup"
	"unicode"
)

// https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher
type Vigenere struct {
	key       string
	lowerRing *lookup.AlphaRing
	upperRing *lookup.AlphaRing
	Encoder
	Decoder
}

func (v *Vigenere) offset(b byte) int {
	switch {
	case b <= 122 && b >= 97:
		// lower case alpha
		encInt := int(b) - 97
		return encInt
	case b <= 90 && b >= 65:
		// upper case alpha
		encInt := int(b) - 65
		return encInt
	default:
		return 0
	}
}

func (v *Vigenere) encodeByte(b byte, keyByte byte) byte {
	var encoded byte
	var err error

	offset := v.offset(keyByte)
	switch {
	case v.lowerRing.Contains(b):
		encoded, err = v.lowerRing.Move(b, offset)
	case v.upperRing.Contains(b):
		encoded, err = v.upperRing.Move(b, offset)
	default:
		encoded, err = b, nil
	}

	if err != nil {
		panic("error encoding byte")
	}

	return encoded
}

func (v *Vigenere) decodeByte(b byte, keyByte byte) byte {
	var decoded byte
	var err error

	offset := v.offset(keyByte)
	switch {
	case v.lowerRing.Contains(b):
		decoded, err = v.lowerRing.Move(b, -offset)
	case v.upperRing.Contains(b):
		decoded, err = v.upperRing.Move(b, -offset)
	default:
		decoded, err = b, nil
	}

	if err != nil {
		panic("error decoding byte")
	}

	return decoded
}

func (v *Vigenere) Encode(s string) string {
	var runes []byte
	for i, curr := range s {
		// key repeats until it's the same length as string
		// to encrypt. e.g. input string `attackatdawn` and key
		// `LEMON` gives padded key `LEMONLEMONLE`.
		keyByte := []byte(v.key)[i%len(v.key)]
		nextByte := v.encodeByte(byte(curr), keyByte)
		runes = append(runes, nextByte)
	}
	return string(runes)
}

func (v *Vigenere) Decode(s string) string {
	var runes []byte
	for i, curr := range s {
		if !unicode.IsLetter(curr) {
			runes = append(runes, byte(curr))
			continue
		}
		// key repeats until it's the same length as string
		// to encrypt. e.g. input string `attackatdawn` and key
		// `LEMON` gives padded key `LEMONLEMONLE`.
		keyByte := []byte(v.key)[i%len(v.key)]
		runes = append(runes, v.decodeByte(byte(curr), keyByte))
	}
	return string(runes)
}

func NewVigenere(key string) *Vigenere {
	return &Vigenere{
		key:       key,
		lowerRing: lookup.NewAlphaRing(true),
		upperRing: lookup.NewAlphaRing(false),
	}
}
