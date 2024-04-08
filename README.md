# ciphers

A fun little CLI program to encode/decode strings in various famous historical cipher schemes.

## build

Run `go build -o bin/cipher cmd/cipher/main.go`

## run

`bin/cipher`

will provide command descriptions, but basically the scheme is:

`bin/cipher <cipher name> <encode | decode> "text to encode/decode" <offset, key, etc. (depending on cipher scheme)>`

Have fun!
