package main

import (
	"flag"
	"io"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

var debug bool

func init() {

	log.SetOutput(io.MultiWriter(&lumberjack.Logger{
		Filename:   "user.log",
		MaxSize:    1, // Mb
		MaxBackups: 2,
		MaxAge:     1, // Days
		Compress:   true,
	}))

	flag.BoolVar(&debug, "debug", true, "Run in debug mode")

	flag.Parse()

	if debug {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
		log.Println("Started in debug mode")
	}

}

func main() {

	log.Println("Starting logging")

	for i := 0; i < 12000; i++ {//to reach 1 mb
		log.Println("Continuing logging information")
    log.Println("Continuing logging information")
    log.Println("Continuing logging information")
    log.Println("Continuing logging information")
    log.Println("Continuing logging information")
	}

	log.Println("Finished logging")

}
