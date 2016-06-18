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
)

var applicationFinished sync.WaitGroup

func main() {
	startLog()
	initScreen()

	applicationConfiguration, err := readConfig()
	if err != nil {
		applicationConfiguration = getDefaultConfigFile()
	}
	initColors(applicationConfiguration)

	startCommandConsole()
	applicationFinished.Add(1)
	applicationFinished.Wait()

	defer HeaderWindow.Delete()
	defer MainWindow.Delete()
	defer CommandWindow.Delete()
	defer goncurses.End()
	defer endLogging()
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
