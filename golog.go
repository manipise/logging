package main

import (
	"path/filepath"
	"runtime"
	"strings"

	"log"
)

var (
	Error   = log.New(logging{}, "ERROR: ", 0)
	Warning = log.New(logging{}, "WARN: ", 0)
	Info    = log.New(logging{}, "INFO: ", 0)
	Trace   = log.New(logging{}, "Trace: ", 0)
)

type logging struct{}

func (f logging) Write(p []byte) (n int, err error) {
	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	log.Printf("[%s:%d %s]: %s", filepath.Base(file), line, fnName, p)
	return len(p), nil
}

func main() {

	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
	Trace.Println("Printing Some standards")
}
	
