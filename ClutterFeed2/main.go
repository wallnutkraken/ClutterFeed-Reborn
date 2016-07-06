/* Main file for ClutterFeed 2; return of the console */
package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/rthornton128/goncurses"
)

const (
	CF_VERSION = "2.0-DEV"
	CF_RELEASE = "TBD"
	DEBUG      = true
)

const TEMP_username = "username" /* ONLY TEMPORARY */
/* Will be removed once logging in with Twitter is added */

var applicationFinished sync.WaitGroup

func main() {
	initScreen()

	applicationConfiguration, err := readConfig()
	if err != nil {
		applicationConfiguration = getDefaultConfigFile()
	}
	initColors(applicationConfiguration)
	drawHeader()

	startCommandConsole()
	applicationFinished.Add(1)
	applicationFinished.Wait()

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

func errToStderr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
	}
}
