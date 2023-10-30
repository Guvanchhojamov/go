package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	// First application
	l := log.New(os.Stderr, "hw1- ", log.Flags())
	if time, err := ntp.Time("0.beevik-ntp.pool.ntp.org"); err != nil {
		l.Fatalln(err)
	} else {
		fmt.Print(time.Format("Monday January 02 15:05:20 -0700 MST 2006"))
	}
}
