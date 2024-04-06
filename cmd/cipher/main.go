package main

import (
	"ciphers/ciphers"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "caesar",
				Aliases: []string{"cs"},
				Usage:   "encode or decode with caesar cipher",
				Subcommands: []*cli.Command{
					{
						Name:    "encode",
						Aliases: []string{"e"},
						Usage:   "encode a string",
						Action: func(cCtx *cli.Context) error {
							str := cCtx.Args().Get(0)
							offset, convErr := strconv.Atoi(cCtx.Args().Get(1))
							if convErr != nil {
								panic("could not get offset")
							}
							if offset == 0 {
								offset = 3
							}
							caesar := ciphers.NewCaesar(offset)
							encoded, encodeErr := caesar.Encode(str)
							if encodeErr != nil {
								panic("could not encode string")
							}
							fmt.Println(encoded)
							return nil
						},
					},
					{
						Name:    "decode",
						Aliases: []string{"d"},
						Usage:   "decode a string",
						Action: func(cCtx *cli.Context) error {
							str := cCtx.Args().Get(0)
							offset, convErr := strconv.Atoi(cCtx.Args().Get(1))
							if convErr != nil {
								panic("could not get offset")
							}
							if offset == 0 {
								offset = 3
							}
							caesar := ciphers.NewCaesar(offset)
							encoded, encodeErr := caesar.Decode(str)
							if encodeErr != nil {
								panic("could not decode string")
							}
							fmt.Println(encoded)
							return nil
						},
					},
				},
			},
			{
				Name:    "vigenere",
				Aliases: []string{"vg"},
				Usage:   "encode or decode with Vigenère cipher",
				Subcommands: []*cli.Command{
					{
						Name:    "encode",
						Aliases: []string{"e"},
						Usage:   "encode a string",
						Action: func(cCtx *cli.Context) error {
							str := cCtx.Args().Get(0)
							key := cCtx.Args().Get(1)
							vig := ciphers.NewVigenere(key)
							encoded, encodeErr := vig.Encode(str)
							if encodeErr != nil {
								panic("could not encode string")
							}
							fmt.Println(encoded)
							return nil
						},
					},
					{
						Name:    "decode",
						Aliases: []string{"d"},
						Usage:   "decode a string",
						Action: func(cCtx *cli.Context) error {
							str := cCtx.Args().Get(0)
							key := cCtx.Args().Get(1)
							vig := ciphers.NewVigenere(key)
							encoded, encodeErr := vig.Decode(str)
							if encodeErr != nil {
								panic("could not decode string")
							}
							fmt.Println(encoded)
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
