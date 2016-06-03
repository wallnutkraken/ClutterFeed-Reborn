/* Main file for ClutterFeed 2; return of the console */
package main

import (
	"fmt"
	"os"

	"github.com/rthornton128/goncurses"
)

const (
	CF_VERSION = "2.0-DEV"
	CF_RELEASE = "TBD"
)

func main() {
	initScreen()

	defer HeaderWindow.Delete()
	defer MainWindow.Delete()
	defer CommandWindow.Delete()
	defer goncurses.End()
}

func fatalErrorCheck(err error) {
	if err != nil {
		goncurses.End()
		fmt.Println("Fatal error:", err)
		os.Exit(1)
	}
}
