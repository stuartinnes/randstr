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
	lower   = "abcdefghijklmnopqrstuvwxyz"
	upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
	symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

var (
	app = &cli.App{
		Name: "Random string generator",
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name:    "upper",
				Usage:   "upper case characters",
				Aliases: []string{"u"},
				Value:   4,
			},
			&cli.UintFlag{
				Name:    "lower",
				Usage:   "dower case characters",
				Aliases: []string{"l"},
				Value:   4,
			},
			&cli.UintFlag{
				Name:    "digits",
				Usage:   "digits",
				Aliases: []string{"d"},
				Value:   4,
			},
			&cli.UintFlag{
				Name:    "symbols",
				Usage:   "symbols",
				Aliases: []string{"s"},
				Value:   4,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println(getRandStr(c.Uint("u"), c.Uint("l"), c.Uint("d"), c.Uint("s")))
			return nil
		},
	}
)

func main() {
	_ = app.Run(os.Args)
}

func getRandStr(up, low, dig, sym uint) string {
	var (
		//the overall length of our generated string - converted to an INT for later
		//use in shuffle
		strLen = int(up + low + dig + sym)
		err    error

		//ranges are used to determine which dictionary to use within
		//our main generator below - converted to an INT which avoids conversion
		//within our loop
		symbolRange = int(up + low + dig)
		digitRange  = int(up + low)
		lowRange    = int(up)

		//default dictionary value
		dict string

		//a container for our random bytes
		randBytes = make([]byte, strLen)
	)

	if strLen == 0 {
		log.Fatal("nothing to do")
	}

	_, err = rand.Reader.Read(randBytes)
	if err != nil {
		log.Fatal("unable to generate random sequence")
	}

	for k, v := range randBytes {
		switch {
		case k >= symbolRange:
			dict = symbols
		case k >= digitRange:
			dict = digits
		case k >= lowRange:
			dict = lower
		default:
			dict = upper
		}

		//switch out the byte value with something from one of our dictionaries
		randBytes[k] = dict[v%byte(len(dict))]
	}

	//now we can shuffle our ordered bytes slice
	mrand.Seed(time.Now().Unix())
	mrand.Shuffle(strLen,
		func(i, j int) {
			randBytes[i], randBytes[j] = randBytes[j], randBytes[i]
		})

	return string(randBytes)
}
