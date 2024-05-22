package main

import (
	"ciphers/ciphers"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
)

func vigenere() *cli.Command {
	return &cli.Command{
		Name:    "vigenere",
		Aliases: []string{"vg"},
		Usage:   "encode or decode with Vigenère cipher",
		Subcommands: []*cli.Command{
			{
				Name:    "encode",
				Aliases: []string{"e"},
				Usage:   "with string to encode and key string",
				Action: func(cCtx *cli.Context) error {
					str := cCtx.Args().Get(0)
					key := cCtx.Args().Get(1)
					vig := ciphers.NewVigenere(key)
					encoded := vig.Encode(str)

					fmt.Println(encoded)
					return nil
				},
			},
			{
				Name:    "decode",
				Aliases: []string{"d"},
				Usage:   "with string to decode and key string",
				Action: func(cCtx *cli.Context) error {
					str := cCtx.Args().Get(0)
					key := cCtx.Args().Get(1)
					vig := ciphers.NewVigenere(key)
					decoded := vig.Decode(str)

					fmt.Println(decoded)
					return nil
				},
			},
		},
	}
}

func caesar() *cli.Command {
	return &cli.Command{
		Name:    "caesar",
		Aliases: []string{"cs"},
		Usage:   "encode or decode with Caesar cipher",
		Subcommands: []*cli.Command{
			{
				Name:    "encode",
				Aliases: []string{"e"},
				Usage:   "with string to encode and positive integer offset",
				Action: func(cCtx *cli.Context) error {
					str := cCtx.Args().Get(0)
					offset, convErr := strconv.Atoi(cCtx.Args().Get(1))
					if convErr != nil || offset < 1 {
						panic("expected positive integer offset")
					}

					caesar := ciphers.NewCaesar(offset)
					encoded := caesar.Encode(str)
					fmt.Println(encoded)
					return nil
				},
			},
			{
				Name:    "decode",
				Aliases: []string{"d"},
				Usage:   "with string to decode and positive integer offset",
				Action: func(cCtx *cli.Context) error {
					str := cCtx.Args().Get(0)
					offset, convErr := strconv.Atoi(cCtx.Args().Get(1))
					if convErr != nil || offset < 1 {
						panic("expected positive integer offset")
					}

					caesar := ciphers.NewCaesar(offset)
					decoded := caesar.Decode(str)
					fmt.Println(decoded)
					return nil
				},
			},
		},
	}
}

func playfair() *cli.Command {
	return &cli.Command{
		Name:    "playfair",
		Aliases: []string{"pf"},
		Usage:   "encode or decode with Playfair cipher",
		Subcommands: []*cli.Command{
			{
				Name:    "encode",
				Aliases: []string{"e"},
				Usage:   "with string to encode and key string",
				Action: func(cCtx *cli.Context) error {
					str := cCtx.Args().Get(0)
					key := cCtx.Args().Get(1)
					pf := ciphers.NewPlayfair(key)
					encoded := pf.Encode(str)

					fmt.Println(encoded)
					return nil
				},
			},
			{
				Name:    "decode",
				Aliases: []string{"d"},
				Usage:   "with string to decode and key string",
				Action: func(cCtx *cli.Context) error {
					str := cCtx.Args().Get(0)
					key := cCtx.Args().Get(1)
					pf := ciphers.NewPlayfair(key)
					decoded := pf.Decode(str)

					fmt.Println(decoded)
					return nil
				},
			},
		},
	}
}

func main() {
	app := &cli.App{
		Name:    "cipher",
		Suggest: true,
		Usage:   "encode and decode in various historical ciphers",
		Commands: []*cli.Command{
			caesar(),
			vigenere(),
			playfair(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
