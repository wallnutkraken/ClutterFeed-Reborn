package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rthornton128/goncurses"
)

var (
	HeaderWindow  *goncurses.Window
	MainWindow    *goncurses.Window
	CommandWindow *goncurses.Window
)

func addStringAt(window *goncurses.Window, content string, ypos int, xpos int) {
	window.Move(ypos, xpos)
	addString(window, content)
}

func addString(window *goncurses.Window, content string) {
	var currentChar goncurses.Char
	allChars := []rune(content)[:]
	for i := range allChars {
		currentChar = goncurses.Char(allChars[i])
		window.AddChar(currentChar)
	}
}

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
	}

	SIZE_Y, SIZE_X = stdscr.MaxYX()

	/* Create header window */
	HeaderWindow, err := goncurses.NewWindow(1, SIZE_X, 0, 0)
	fatalErrorCheck(err)

	/* The second window being the main timeline window */
	MainWindow, err := goncurses.NewWindow(SIZE_Y-3, SIZE_X, 1, 0)
	fatalErrorCheck(err)
	MainWindow.ScrollOk(true)

	/* And the final command window */
	CommandWindow, err := goncurses.NewWindow(2, SIZE_X, SIZE_Y-2, 0)
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
		/* do resizing stuff */
	}
}
