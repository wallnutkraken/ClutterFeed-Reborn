package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rthornton128/goncurses"
)

var (
	currentConsoleCommand string
	horizontalStrPosition int
	verticalStrPosition   int

	canAcceptInput bool
	exiting        bool
)

func startCommandConsole() error {
	/* Initialize the global vars */
	canAcceptInput = true
	exiting = false

	if CommandWindow == nil {
		return errors.New("Command window is not initialized")
	}

	inputChannel := make(chan goncurses.Key)
	go getInput(inputChannel)
	go handleInput(inputChannel)

	drawConsole()
	return nil
}

/* Grabs the cursor and places it whereever it should be in the current string */
func grabCommandCursor() {
	/* 8 characters are reserved for the console counter and such */
	xPlacement := 8 + horizontalStrPosition
	CommandWindow.Move(verticalStrPosition, xPlacement)
	CommandWindow.Refresh()
}

func drawConsole() {
	CommandWindow.Clear()

	CommandWindow.AttrOn(goncurses.A_BOLD)

	CommandWindow.ColorOn(COMMAND_PAIR)
	CommandWindow.Print("[")
	CommandWindow.ColorOff(COMMAND_PAIR)

	CommandWindow.Print(fmt.Sprintf("%03d", horizontalStrPosition))

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
