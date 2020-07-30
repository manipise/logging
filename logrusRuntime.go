package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	var setEnv int
	Loginfo("user", false)
	Dologging()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Print("Enter your option(0:=json,1=text): ")
		fmt.Scanln(&setEnv)
		defineFormat(setEnv)
		Dologging()

	}()
	wg.Wait()

}
func Dologging() {

	log.Error("Some error, but continue.")
	// just a warning, nothing to worry about ... yet
	log.Warn("Warning....crash to ground imminent.")
	//Some general info
	log.Info("Info of the functions")
	//needed info in debug mood
	log.Debug("some info for debug")
}

func defineFormat(env int) {
	if env == 0 {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}
}

func Loginfo(svcname string, showStdout bool) {
	logdir := "./log/"
	os.Mkdir(logdir, 0775)
	t := time.Now()
	sdate := t.Format("20060102")
	logfile := logdir + sdate + "_" + svcname + ".log"
	logFile, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	var mw io.Writer
	if showStdout {
		mw = io.MultiWriter(os.Stdout, logFile)
	} else {
		mw = io.Writer(logFile)
	}
	log.SetOutput(mw)
}
