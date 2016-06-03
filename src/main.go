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

var (
	SIZE_X int
	SIZE_Y int
)

func main() {
	initScreen()
}

func initScreen() []*goncurses.Window {
	stdscr, err := goncurses.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	windows := make([]*goncurses.Window, 3)
	/* Order is from top to bottom for the slice */

	if goncurses.HasColors() == false {
		fmt.Fprintf(os.Stderr, "W: Console does not support color")
	}

	SIZE_Y, SIZE_X = stdscr.MaxYX()

	/* Create header window */
	headerWindow, err := goncurses.NewWindow(1, SIZE_X, 0, 0)
	fatalErrorCheck(err)
	windows[0] = headerWindow

	/* The second window being the main timeline window */
	timelineWindow, err := goncurses.NewWindow(SIZE_Y-3, SIZE_X, 1, 0)
	fatalErrorCheck(err)
	windows[1] = timelineWindow

	/* And the final command window */
	commandWindow, err := goncurses.NewWindow(2, SIZE_X, SIZE_Y-2, 0)
	fatalErrorCheck(err)
	windows[2] = commandWindow

	return windows
}

func fatalErrorCheck(err error) {
	if err != nil {
		goncurses.End()
		fmt.Println("Fatal error:", err)
		os.Exit(1)
	}
}
