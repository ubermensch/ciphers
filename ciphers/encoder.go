package ciphers

type Encoder interface {
	Encode(string) string
}

type Decoder interface {
	Decode(string) string
}
