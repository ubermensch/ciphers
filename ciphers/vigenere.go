package ciphers

import "unicode"

// https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher
type Vigenere struct {
	key string
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
	offset := v.offset(keyByte)
	newByteInt := int(b) + offset
	if newByteInt < 97 || newByteInt > 122 || (newByteInt > 90 && newByteInt < 97) {
		newByteInt = newByteInt - 26
	}

	newByte := byte(newByteInt)
	return newByte
}

func (v *Vigenere) decodeByte(b byte, keyByte byte) byte {
	offset := v.offset(keyByte)
	bInt := int(b)
	newByteInt := bInt - offset
	if bInt > 96 && bInt < 123 && newByteInt < 97 {
		newByteInt = newByteInt + 26
	}

	if bInt > 64 && bInt < 91 && newByteInt < 65 {
		newByteInt = newByteInt + 26
	}

	newByte := byte(newByteInt)
	return newByte
}

func (v *Vigenere) Encode(s string) string {
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
	return &Vigenere{key: key}
}