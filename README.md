# ciphers

A fun little CLI program to encode/decode strings in various historical cipher schemes.

So far, the following are implemented:

* [Caesar](https://en.wikipedia.org/wiki/Caesar_cipher)
* [Vigenère](https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher)
* [Playfair](https://en.wikipedia.org/wiki/Playfair_cipher)

## build 🛠️

Run `go build -o bin/cipher cmd/cipher/main.go`

## run 🏃‍♂️‍➡️

`bin/cipher`

```
NAME:
   cipher - encode and decode in various historical ciphers

USAGE:
   cipher [global options] command [command options]

COMMANDS:
   caesar, cs    encode or decode with Caesar cipher
   vigenere, vg  encode or decode with Vigenère cipher
   playfair, pf  encode or decode with Playfair cipher
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

Have fun! 🥳
