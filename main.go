package main

import (
	"github.com/mkideal/cli"
	"log"
	"os"
	"time"
)

type argT struct {
	Endpoint string `cli:"*e,endpoint" usage:"URL with the endpoint used for check"`
	Tries    int    `cli:"t,tries" usage:"number of tries before fail" dft:"10"`
	Delay    int64  `cli:"d,delay" usage:"delay between tries (in seconds)" dft:"20"`
}

var ARGS *argT

func main() {
	GetArgs()

	log.Printf("Requesting at %s\n", FormatURL(ARGS.Endpoint))

	for i := 0; i < ARGS.Tries; i++ {
		log.Printf("#%d request\n", i+1)
		if IsOK(ARGS.Endpoint) {
			os.Exit(0)
		}

		time.Sleep(time.Duration(ARGS.Delay) * time.Second)
	}

	os.Exit(1)
}

func GetArgs() {
	if cli.Run(new(argT), func(ctx *cli.Context) error {
		ARGS = ctx.Argv().(*argT)
		return nil
	}) == 1 {
		os.Exit(1)
	}
}
