package log

import (
	"fmt"
	"log"
	"os"
)

var l *log.Logger

// Init creates logger with filename <fname> passed to it
func Init(fname string) error {
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_APPEND, 0x644)
	if err != nil {
		return err
	}
	l = log.New(f, "", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

// Print prints message to logger if it exests or to stdio
func Print(msg ...interface{}) {
	if l == nil {
		fmt.Println(msg...)
	} else {
		l.Print(msg...)
	}
}

// Printf prints formatted message to logger if it exests or to stdio
// It always add \n if printing to stdio
func Printf(msg string, vals ...interface{}) {
	if l == nil {
		fmt.Println(fmt.Sprintf(msg, vals...))
	} else {
		l.Printf(msg, vals...)
	}
}
