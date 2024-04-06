package ciphers

type Vigenere struct {
	key string
	Encoder
	Decoder
}

func (v *Vigenere) offset(b byte) int {
	switch {
	case b <= 122 && b >= 97:
		// lower case alpha
		return int(b) - 97
	case b <= 90 && b >= 65:
		return int(b) - 65
	default:
		return 0
	}
}

func (v *Vigenere) encodeByte(b byte, keyByte byte) byte {
	return byte(int(b) + v.offset(keyByte))
}

func (v *Vigenere) decodeByte(b byte, keyByte byte) byte {
	return byte(int(b) - v.offset(keyByte))
}

func (v *Vigenere) Encode(s string) (string, error) {
	var runes []byte
	for i, curr := range s {
		// key repeats until it's the same length as string
		// to encrypt. e.g. input string `attackatdawn` and key
		// `LEMON` gives padded key `LEMONLEMONLE`.
		keyByte := []byte(v.key)[i%len(v.key)]
		runes = append(runes, v.encodeByte(byte(curr), keyByte))
	}
	return string(runes), nil
}

func (v *Vigenere) Decode(s string) (string, error) {
	var runes []byte
	for i, curr := range s {
		// key repeats until it's the same length as string
		// to encrypt. e.g. input string `attackatdawn` and key
		// `LEMON` gives padded key `LEMONLEMONLE`.
		keyByte := []byte(v.key)[i%len(v.key)]
		runes = append(runes, v.decodeByte(byte(curr), keyByte))
	}
	return string(runes), nil
}

func NewVigenere(key string) *Vigenere {
	return &Vigenere{key: key}
}
