package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rthornton128/goncurses"
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

func initScreen() []*goncurses.Window {
	stdscr, err := goncurses.Init()
	goncurses.Echo(false)
	goncurses.NewLines(true)
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
	timelineWindow.ScrollOk(true)
	windows[1] = timelineWindow

	/* And the final command window */
	commandWindow, err := goncurses.NewWindow(2, SIZE_X, SIZE_Y-2, 0)
	fatalErrorCheck(err)
	commandWindow.Keypad(true) /* Will allow us to use the keypad in the console */
	fatalErrorCheck(err)
	windows[2] = commandWindow

	/* And finally, create a goroutine and a channel for when the terminal is resized */
	resizeChannel := make(chan os.Signal)
	signal.Notify(resizeChannel, syscall.SIGWINCH)
	go onResize(resizeChannel)

	return windows
}

func onResize(resizeChannel chan os.Signal) {
	for {
		<-resizeChannel
		/* do resizing stuff */
	}
}
