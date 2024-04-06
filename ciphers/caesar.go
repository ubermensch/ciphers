package ciphers

type Caesar struct {
	offset int
	Encoder
	Decoder
}

func (c *Caesar) encodeByte(b byte) byte {
	switch {
	case b <= 122 && b >= 97:
		// lower case alpha
		encInt := int(b) + c.offset
		if encInt > 122 {
			// loop back around z -> a
			encInt = encInt - 26
		}
		return byte(encInt)
	case b <= 90 && b >= 65:
		// upper case alpha
		encInt := int(b) + c.offset
		if encInt > 90 {
			// loop back around Z -> A
			encInt = encInt - 26
		}
		return byte(encInt)
	default:
		return b
	}
}

func (c *Caesar) decodeByte(b byte) byte {
	switch {
	case b <= 122 && b >= 97:
		// lower case alpha
		encInt := int(b) - c.offset
		if encInt < 97 {
			// loop back around a -> z
			encInt = encInt + 26
		}
		return byte(encInt)
	case b <= 90 && b >= 65:
		// upper case alpha
		encInt := int(b) - c.offset
		if encInt < 65 {
			// loop back around Z -> A
			encInt = encInt + 26
		}
		return byte(encInt)
	default:
		return b
	}
}

func (c *Caesar) Encode(s string) (string, error) {
	var runes []byte
	for _, curr := range s {
		runes = append(runes, c.encodeByte(byte(curr)))
	}
	return string(runes), nil
}

func (c *Caesar) Decode(s string) (string, error) {
	var runes []byte
	for _, curr := range s {
		runes = append(runes, c.decodeByte(byte(curr)))
	}
	return string(runes), nil
}

func NewCaesar(offset int) *Caesar {
	return &Caesar{offset: offset}
}
