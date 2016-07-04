package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rthornton128/goncurses"
)

var (
	currentConsoleCommand string
	horizontalStrPosition int

	canAcceptInput bool
	exiting        bool
)

var debugFile *os.File
var updateCursorChannel chan bool

func startCommandConsole() error {
	/* Initialize the global vars */
	canAcceptInput = true
	exiting = false

	debugFile, _ = os.Create("debug.conf")

	if CommandWindow == nil {
		return errors.New("Command window is not initialized")
	}
	inputChannel := make(chan goncurses.Key)
	updateCursorChannel = make(chan bool)

	go grabCommandCursor()
	go getInput(inputChannel)
	go handleInput(inputChannel)

	drawConsole()
	return nil
}

/* Grabs the cursor and places it whereever it should be in the current string */
func grabCommandCursor() {
	var value bool
	for {
		/* This is not the most pleasant way to use this function */
		/* However, running this in a goroutine and waiting for */
		/* input on a channel is the only method that I've tried */
		/* which worked to update the command cursor thing. */
		value = <-updateCursorChannel
		if value == true {
			/* 8 characters are reserved for the console counter and such */
			xPlacement := 8 + horizontalStrPosition
			verticalStrPosition := 0
			y, x := CommandWindow.MaxYX()
			debugFile.WriteString("Attempting to move cursor. Scr size:" + strconv.Itoa(x) + ", " + strconv.Itoa(y) + ". New pos: " + strconv.Itoa(xPlacement) + ", " + strconv.Itoa(verticalStrPosition) + "\n")
			if SIZE_X < 8+horizontalStrPosition {
				verticalStrPosition = 1
			}
			CommandWindow.Move(verticalStrPosition, xPlacement)
			CommandWindow.Refresh()
		}
	}
}

func drawConsole() {
	CommandWindow.Clear()

	CommandWindow.AttrOn(goncurses.A_BOLD)

	CommandWindow.ColorOn(COMMAND_PAIR)
	CommandWindow.Print("[")
	CommandWindow.ColorOff(COMMAND_PAIR)

	CommandWindow.Print(fmt.Sprintf("%03d", len(currentConsoleCommand)))

	CommandWindow.ColorOn(WARNING_PAIR)
	CommandWindow.Print("] > ")
	CommandWindow.AttrOff(goncurses.A_BOLD)

	CommandWindow.ColorOff(WARNING_PAIR)

	CommandWindow.Print(currentConsoleCommand)
	CommandWindow.Refresh()
}

/* Says whether the cursor was moved */
func isCursorPosMoved() bool {
	panic("Not implemented")
}

func parseCommandText() {
	currentConsoleCommand = strings.Trim(currentConsoleCommand, " ")
	if strings.HasPrefix(currentConsoleCommand, "/") {
		commands()
	}

	/* Reset command once we're done */
	currentConsoleCommand = ""
	horizontalStrPosition = 0
}

/* Function that handles what happens when a / command is inputted */
func commands() {
	command := currentConsoleCommand[1:len(currentConsoleCommand)] /* remove / */

	if strings.EqualFold(command, "exit") {
		exiting = true
		applicationFinished.Done()
	} else {
		writeMessage("No such command.", ERROR)
		goncurses.Update()
	}
}
