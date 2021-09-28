package main

import (
	"crypto/rand"
	"fmt"
	"log"
	mrand "math/rand"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	chars   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
	symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

var (
	app = &cli.App{
		Name:  "Random string generator",
		Usage: "generate some random strings why dont you!",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "chars",
				Aliases: []string{"c"},
				Value:   8,
			},
			&cli.IntFlag{
				Name:    "digits",
				Aliases: []string{"d"},
				Value:   8,
			},
			&cli.IntFlag{
				Name:    "symbols",
				Aliases: []string{"s"},
				Value:   8,
			},
		},
		Action: func(c *cli.Context) error {
			var (
				numChars   = c.Int("chars")
				numDigits  = c.Int("digits")
				numSymbols = c.Int("symbols")
			)

			for _, v := range []int{numChars, numDigits, numDigits} {
				if v < 0 {
					log.Fatal("argument values cannot be negative")
				}
			}

			printRandStr(numChars, numDigits, numSymbols)
			return nil
		},
	}
)

func main() {
	_ = app.Run(os.Args)
}

func printRandStr(numChars, numDigits, numSymbols int) {
	var (
		strLen = numChars + numDigits + numSymbols
		err    error
	)

	if strLen <= 0 {
		log.Fatal("nothing to do")
	}

	var (
		randBytes = make([]byte, strLen)
	)

	_, err = rand.Reader.Read(randBytes)
	if err != nil {
		log.Fatal("unable to generate random sequence")
	}

	for k, v := range randBytes {
		dict := ""
		switch {
		case k > (numChars+numDigits)-1:
			dict = symbols
		case k > numChars-1:
			dict = digits
		default:
			dict = chars
		}
		randBytes[k] = dict[v%byte(len(dict))]
	}

	mrand.Seed(time.Now().Unix())
	mrand.Shuffle(strLen,
		func(i, j int) {
			randBytes[i], randBytes[j] = randBytes[j], randBytes[i]
		})

	fmt.Println(string(randBytes))
}
