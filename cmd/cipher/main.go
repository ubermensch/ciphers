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

func inputString(ctx *cli.Context) (string, error) {
	var str string
	if len(ctx.String("input-file")) != 0 {
		bytes, err := os.ReadFile(ctx.String("input-file"))
		if err != nil {
			return "", errors.New("could not read from output file: " + err.Error())
		}
		str = string(bytes[:])
	} else {
		str = ctx.Args().Get(0)
	}

	return str, nil
}

func keyOrOffsetIndex(ctx *cli.Context) int {
	offsetIdx := 1
	if len(ctx.String("input-file")) != 0 {
		offsetIdx = 0
	}
	return offsetIdx
}

func handleOutput(ctx *cli.Context, output string) error {
	if len(ctx.String("output-file")) != 0 {
		writeErr := os.WriteFile(ctx.String("output-file"), []byte(output), 0644)
		if writeErr != nil {
			return errors.New("could not write to output file: " + writeErr.Error())
		}
	} else {
		fmt.Println(output)
	}
	return nil
}

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
					keyIdx := keyOrOffsetIndex(cCtx)
					key := cCtx.Args().Get(keyIdx)

					str, err := inputString(cCtx)
					if err != nil {
						return err
					}
					vig := ciphers.NewVigenere(key)
					encoded, err := vig.Encode(str)
					if err != nil {
						return errors.New("could not encode: " + err.Error())
					}

					outputErr := handleOutput(cCtx, encoded)
					if outputErr != nil {
						return outputErr
					}

					return nil
				},
			},
			{
				Name:    "decode",
				Aliases: []string{"d"},
				Usage:   "with string to decode and key string",
				Action: func(cCtx *cli.Context) error {
					keyIdx := keyOrOffsetIndex(cCtx)
					key := cCtx.Args().Get(keyIdx)

					str, err := inputString(cCtx)
					if err != nil {
						return err
					}
					vig := ciphers.NewVigenere(key)
					decoded, err := vig.Decode(str)
					if err != nil {
						return errors.New("could not decode: " + err.Error())
					}

					outputErr := handleOutput(cCtx, decoded)
					if outputErr != nil {
						return outputErr
					}

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
					offsetIdx := keyOrOffsetIndex(cCtx)
					offset, convErr := strconv.Atoi(cCtx.Args().Get(offsetIdx))
					if convErr != nil {
						return errors.New("expected positive integer offset")
					}

					str, err := inputString(cCtx)
					if err != nil {
						return err
					}

					cs := ciphers.NewCaesar(offset)
					encoded, err := cs.Encode(str)
					if err != nil {
						return errors.New("could not encode: " + err.Error())
					}

					outputErr := handleOutput(cCtx, encoded)
					if outputErr != nil {
						return outputErr
					}

					return nil
				},
			},
			{
				Name:    "decode",
				Aliases: []string{"d"},
				Usage:   "with string to decode and positive integer offset",
				Action: func(cCtx *cli.Context) error {
					offsetIdx := keyOrOffsetIndex(cCtx)
					offset, convErr := strconv.Atoi(cCtx.Args().Get(offsetIdx))
					if convErr != nil {
						return errors.New("expected positive integer offset")
					}

					str, err := inputString(cCtx)
					if err != nil {
						return err
					}

					cs := ciphers.NewCaesar(offset)
					decoded, err := cs.Decode(str)
					if err != nil {
						return errors.New("could not decode: " + err.Error())
					}

					outputErr := handleOutput(cCtx, decoded)
					if outputErr != nil {
						return outputErr
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
					keyIdx := keyOrOffsetIndex(cCtx)
					key := cCtx.Args().Get(keyIdx)

					str, err := inputString(cCtx)
					if err != nil {
						return err
					}

					pf := ciphers.NewPlayfair(key)
					encoded, err := pf.Encode(str)

					outputErr := handleOutput(cCtx, encoded)
					if outputErr != nil {
						return outputErr
					}

					return nil
				},
			},
			{
				Name:    "decode",
				Aliases: []string{"d"},
				Usage:   "with string to decode and key string",
				Action: func(cCtx *cli.Context) error {
					keyIdx := keyOrOffsetIndex(cCtx)
					key := cCtx.Args().Get(keyIdx)

					str, err := inputString(cCtx)
					if err != nil {
						return err
					}

					pf := ciphers.NewPlayfair(key)
					decoded, err := pf.Decode(str)

					outputErr := handleOutput(cCtx, decoded)
					if outputErr != nil {
						return outputErr
					}

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
