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
	Delay    int    `cli:"d,delay" usage:"delay between tries (in seconds)" dft:"20"`
}

var ARGS *argT

func main() {
	GetArgs()

	log.Printf("Requesting at %s\n", FormatURL(ARGS.Endpoint))
	go AntiLock()

	for i := 0; i < ARGS.Tries; i++ {
		log.Printf("request #%d\n", i+1)
		go MakeRequest()

		time.Sleep(time.Duration(ARGS.Delay) * time.Second)
	}
	time.Sleep(time.Duration(ARGS.Delay) * time.Second)

	log.Printf("REQUESTS FAILED ")
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

func AntiLock() {
	time.Sleep(time.Duration(ARGS.Delay*ARGS.Tries) * time.Second)
	log.Printf("REQUESTS LOCKED")
	os.Exit(1)
}

func MakeRequest() {
	if IsOK(ARGS.Endpoint) {
		log.Printf("REQUEST SUCCEED")
		os.Exit(0)
	}
}
