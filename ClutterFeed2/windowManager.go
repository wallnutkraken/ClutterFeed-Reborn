package main

import (
	"errors"
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

/* Message severity for writeMessage() */
const (
	DEFAULT = 0
	WARNING = 1
	ERROR   = 2
)

func drawHeader() {
	HeaderWindow.Clear()

	nameStr := "ClutterFeed " + CF_RELEASE + " (" + CF_VERSION + ")"
	loggedInStrLong := "Logged in as: (@" + TEMP_username + ")"
	loggedInStrShort := "@" + TEMP_username
	loggedInStrToWrite := loggedInStrLong /* Default string to write if there */
	/* are no problems */

	HeaderWindow.Print(nameStr)

	if SIZE_X-len(nameStr) < len(loggedInStrLong) {
		loggedInStrToWrite = loggedInStrShort
	}

	err := printAtEnd(HeaderWindow, loggedInStrToWrite)

	if err != nil {
		err = printAtEnd(HeaderWindow, loggedInStrShort)
	}

	HeaderWindow.Refresh()
}

func printAtEnd(window *goncurses.Window, content string) error {
	_, x := window.MaxYX()
	y, _ := window.CursorYX()

	if len(content)+1 > x {
		return errors.New("Content is too long.")
	}
	startingPosition := x - len(content)
	window.MovePrint(y, startingPosition, content)

	return nil
}

/* Writes a message to the Main Window. */
func writeMessage(content string, severity int) {
	var color int16 = WHITE_PAIR

	/* A prefix to the message like [W] (for warnings) or [E] */
	var prefix string = "      "

	switch severity {
	case WARNING:
		color = WARNING_PAIR
		prefix = "  [W] "
	case ERROR:
		color = ERROR_PAIR
		prefix = "  [E] "
	}

	MainWindow.ColorOn(color)
	MainWindow.Print(prefix)
	MainWindow.AttrOn(goncurses.A_BOLD)
	MainWindow.Print(content)
	MainWindow.AttrOff(goncurses.A_BOLD)
	MainWindow.ColorOff(color)
	MainWindow.Print("\n")

	MainWindow.Refresh()
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
	CommandWindow.Keypad(true)
	CommandWindow.ScrollOk(true) /* Would prevent crashes from happening when the */
	/* terminal window is at a very small size */
	fatalErrorCheck(err)

	/* And finally, create a goroutine and a channel for when the terminal is resized */
	/* and another one for when it's interrupted*/
	resizeChannel := make(chan os.Signal)
	signal.Notify(resizeChannel, syscall.SIGWINCH)
	go onResize(resizeChannel)

	interruptChannel := make(chan os.Signal)
	signal.Notify(interruptChannel, syscall.SIGINT)
	go onInterrupt(interruptChannel)
}

/* Makes sure that if an interrupt happens, curses would exit properly*/
func onInterrupt(interruptChannel chan os.Signal) {
	<-interruptChannel
	goncurses.End()
	os.Exit(1)
}

func onResize(resizeChannel chan os.Signal) {
	for {
		<-resizeChannel
		/* End curses to update terminal info. */
		goncurses.End()
		goncurses.Update()

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
		drawHeader()
		grabCommandCursor()
	}
}
