package ciphers

type Encoder interface {
	Encode(string) (string, error)
}

type Decoder interface {
	Decode(string) (string, error)
}
