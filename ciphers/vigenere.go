package ciphers

import (
	"ciphers/lookup"
	"errors"
	"sync"
)

// https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher
type Vigenere struct {
	key       string
	lowerRing *lookup.AlphaRing
	upperRing *lookup.AlphaRing
	Encoder
	Decoder
}

func (v *Vigenere) offset(c rune) int {
	switch {
	case c <= 122 && c >= 97:
		// lower case alpha
		encInt := int(c) - 97
		return encInt
	case c <= 90 && c >= 65:
		// upper case alpha
		encInt := int(c) - 65
		return encInt
	default:
		return 0
	}
}

func (v *Vigenere) encodeChar(c rune, keyRune rune) (rune, error) {
	var encoded rune
	var err error

	offset := v.offset(keyRune)
	switch {
	case v.lowerRing.Contains(c):
		encoded, err = v.lowerRing.Move(c, offset)
	case v.upperRing.Contains(c):
		encoded, err = v.upperRing.Move(c, offset)
	default:
		encoded, err = c, nil
	}

	if err != nil {
		return 0, err
	}

	return encoded, nil
}

func (v *Vigenere) decodeChar(c rune, keyRune rune) (rune, error) {
	var decoded rune
	var err error

	offset := v.offset(keyRune)
	switch {
	case v.lowerRing.Contains(c):
		decoded, err = v.lowerRing.Move(c, -offset)
	case v.upperRing.Contains(c):
		decoded, err = v.upperRing.Move(c, -offset)
	default:
		decoded, err = c, nil
	}

	if err != nil {
		return 0, err
	}

	return decoded, nil
}

func (v *Vigenere) Encode(s string) (string, error) {
	if len(v.key) == 0 {
		return "", errors.New("empty key")
	}

	// initialize the encoded runes as the string to encode
	runes := []rune(s)
	wg := sync.WaitGroup{}
	errCount := 0

	// encode each character in parallel
	encFunc := func(r rune, pos int, wg *sync.WaitGroup) {
		defer wg.Done()
		// key repeats until it's the same length as string
		// to encrypt. e.g. input string `attackatdawn` and key
		// `LEMON` gives padded key `LEMONLEMONLE`.
		keyRune := []rune(v.key)[pos%len(v.key)]

		enc, err := v.encodeChar(r, keyRune)
		if err != nil {
			errCount++
		}

		runes[pos] = enc
	}

	for i, curr := range s {
		wg.Add(1)
		go encFunc(curr, i, &wg)
	}

	wg.Wait()
	if errCount > 0 {
		return "", errors.New("encoding failed")
	}

	return string(runes), nil
}

func (v *Vigenere) Decode(s string) (string, error) {
	if len(v.key) == 0 {
		return "", errors.New("empty key")
	}

	// initialize the encoded runes as the string to decode
	runes := []rune(s)
	wg := sync.WaitGroup{}
	errCount := 0

	decFunc := func(r rune, pos int, wg *sync.WaitGroup) {
		defer wg.Done()
		// key repeats until it's the same length as string
		// to encrypt. e.g. input string `attackatdawn` and key
		// `LEMON` gives padded key `LEMONLEMONLE`.
		keyRune := []rune(v.key)[pos%len(v.key)]
		dec, err := v.decodeChar(r, keyRune)
		if err != nil {
			errCount++
		}

		runes[pos] = dec
	}

	for i, curr := range s {
		wg.Add(1)
		go decFunc(curr, i, &wg)
	}

	wg.Wait()
	if errCount > 0 {
		return "", errors.New("decoding failed")
	}

	return string(runes), nil
}

func NewVigenere(key string) *Vigenere {
	return &Vigenere{
		key:       key,
		lowerRing: lookup.NewAlphaRing(true),
		upperRing: lookup.NewAlphaRing(false),
	}
}
