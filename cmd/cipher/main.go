package main

import (
	"ciphers/ciphers"
	"errors"
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
					encoded, err := vig.Encode(str)
					if err != nil {
						return errors.New("could not encode: " + err.Error())
					}

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
					decoded, err := vig.Decode(str)
					if err != nil {
						return errors.New("could not decode: " + err.Error())
					}

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
					var str string
					offsetIdx := 1
					if len(cCtx.String("input-file")) != 0 {
						bytes, err := os.ReadFile(cCtx.String("input-file"))
						if err != nil {
							return errors.New("could not read from input file: " + err.Error())
						}
						str = string(bytes[:])
						offsetIdx = 0
					} else {
						str = cCtx.Args().Get(0)
					}
					offset, convErr := strconv.Atoi(cCtx.Args().Get(offsetIdx))
					if convErr != nil {
						return errors.New("expected positive integer offset")
					}

					cs := ciphers.NewCaesar(offset)
					encoded, err := cs.Encode(str)
					if err != nil {
						return errors.New("could not encode: " + err.Error())
					}

					if len(cCtx.String("output-file")) != 0 {
						writeErr := os.WriteFile(cCtx.String("output-file"), []byte(encoded), 0644)
						if writeErr != nil {
							return errors.New("could not write to output file: " + err.Error())
						}
					} else {
						fmt.Println(encoded)
					}

					return nil
				},
			},
			{
				Name:    "decode",
				Aliases: []string{"d"},
				Usage:   "with string to decode and positive integer offset",
				Action: func(cCtx *cli.Context) error {
					var str string
					offsetIdx := 1
					if len(cCtx.String("input-file")) != 0 {
						bytes, err := os.ReadFile(cCtx.String("input-file"))
						if err != nil {
							return errors.New("could not read from input file: " + err.Error())
						}
						str = string(bytes[:])
						offsetIdx = 0
					} else {
						str = cCtx.Args().Get(0)
					}
					offset, convErr := strconv.Atoi(cCtx.Args().Get(offsetIdx))
					if convErr != nil {
						return errors.New("expected positive integer offset")
					}

					cs := ciphers.NewCaesar(offset)
					decoded, err := cs.Decode(str)
					if err != nil {
						return errors.New("could not decode: " + err.Error())
					}

					if len(cCtx.String("output-file")) != 0 {
						writeErr := os.WriteFile(cCtx.String("output-file"), []byte(decoded), 0644)
						if writeErr != nil {
							return errors.New("could not write to output file: " + err.Error())
						}
					} else {
						fmt.Println(decoded)
					}

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
					encoded, err := pf.Encode(str)
					if err != nil {
						return errors.New("could not encode: " + err.Error())
					}

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
					decoded, err := pf.Decode(str)
					if err != nil {
						return errors.New("could not decode: " + err.Error())
					}

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
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "input-file", Aliases: []string{"if"}},
			&cli.StringFlag{Name: "output-file", Aliases: []string{"of"}},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
