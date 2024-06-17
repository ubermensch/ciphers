package ciphers

import (
	"ciphers/lookup"
	"errors"
	"sync"
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
	if c.offset < 1 {
		return "", errors.New("expected positive integer offset")
	}

	runes := make([]rune, len(s))
	wg := sync.WaitGroup{}
	errCount := 0

	// encode each character in parallel
	encFunc := func(r rune, pos int, wg *sync.WaitGroup) {
		defer wg.Done()
		enc, err := c.encodeChar(r)

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

func (c *Caesar) Decode(s string) (string, error) {
	if c.offset < 1 {
		return "", errors.New("expected positive integer offset")
	}

	runes := make([]rune, len(s))
	wg := sync.WaitGroup{}
	errCount := 0

	decFunc := func(r rune, pos int, wg *sync.WaitGroup) {
		defer wg.Done()
		dec, err := c.decodeChar(r)
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

func NewCaesar(offset int) *Caesar {
	return &Caesar{
		offset:    offset,
		lowerRing: lookup.NewAlphaRing(true),
		upperRing: lookup.NewAlphaRing(false),
	}
}
