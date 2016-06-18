package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rthornton128/goncurses"
)

var (
	SIZE_X int
	SIZE_Y int
)

var (
	HeaderWindow  *goncurses.Window
	MainWindow    *goncurses.Window
	CommandWindow *goncurses.Window
)

func drawHeader(username string) {
	panic("Not implemented")
}

func initScreen() {
	stdscr, err := goncurses.Init()
	goncurses.Echo(false)
	goncurses.NewLines(true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if goncurses.HasColors() == false {
		fmt.Fprintf(os.Stderr, "W: Console does not support color")
	} else {
		goncurses.StartColor()
	}

	SIZE_Y, SIZE_X = stdscr.MaxYX()

	/* Create header window */
	HeaderWindow, err = goncurses.NewWindow(1, SIZE_X, 0, 0)
	HeaderWindow.AttrOn(goncurses.A_BOLD)
	fatalErrorCheck(err)

	/* The second window being the main timeline window */
	MainWindow, err = goncurses.NewWindow(SIZE_Y-3, SIZE_X, 1, 0)
	fatalErrorCheck(err)
	MainWindow.ScrollOk(true)

	/* And the final command window */
	CommandWindow, err = goncurses.NewWindow(2, SIZE_X, SIZE_Y-2, 0)
	fatalErrorCheck(err)
	CommandWindow.Keypad(true) /* Will allow us to use the keypad in the console */
	fatalErrorCheck(err)

	/* And finally, create a goroutine and a channel for when the terminal is resized */
	resizeChannel := make(chan os.Signal)
	signal.Notify(resizeChannel, syscall.SIGWINCH)
	go onResize(resizeChannel)
}

func onResize(resizeChannel chan os.Signal) {
	for {
		<-resizeChannel
		SIZE_Y, SIZE_X = goncurses.StdScr().MaxYX()

		HeaderWindow.Resize(1, SIZE_X)
		HeaderWindow.Move(0, 0)
		HeaderWindow.Refresh()

		MainWindow.Resize(SIZE_Y-3, SIZE_X)
		MainWindow.Move(1, 0)
		MainWindow.Refresh()

		CommandWindow.Resize(2, SIZE_X)
		CommandWindow.Move(SIZE_Y-2, 0)
		CommandWindow.Refresh()

		/* Todo: handle redrawing */
	}
}
