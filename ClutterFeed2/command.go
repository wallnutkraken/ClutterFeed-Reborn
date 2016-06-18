package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rthornton128/goncurses"
)

var (
	currentConsoleCommand string
	canAcceptInput        bool
	exiting               bool
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

/* Checks if a key is a valid key someone could input into the screen, and not */
/* something like an artifact from resizing */
func isValidKey(key goncurses.Key) bool {
	if key == 410 { /* Resizing artifact */
		return false
	}

	return true
}

func getInput(in chan goncurses.Key) {
	for exiting == false {
		pressedKey := CommandWindow.GetChar()
		if canAcceptInput && isValidKey(pressedKey) {
			in <- pressedKey
		}
	}
}

/* Grabs the cursor and places it whereever it should be in the current string */
func grabCommandCursor() {
	CommandWindow.Refresh()
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

func handleInput(in chan goncurses.Key) {
	for {
		gotChar := <-in
		if gotChar == goncurses.KEY_RETURN || gotChar == goncurses.KEY_ENTER {
			/* finished command */
			parseCommandText()
		} else if int(gotChar) == 127 {
			/* 127 is backspace */
			if len(currentConsoleCommand) > 0 {
				/* Removes last character, if it exists */
				currentConsoleCommand = currentConsoleCommand[0 : len(currentConsoleCommand)-1]
			}
		} else {
			currentConsoleCommand += string(rune(gotChar)) /* Heh, adding a char */
			/* to a string in Go isn't the most pleasant thing ever */
		}
		drawConsole()
	}
}

func parseCommandText() {
	currentConsoleCommand = strings.Trim(currentConsoleCommand, " ")
	if strings.HasPrefix(currentConsoleCommand, "/") {
		commands()
	}

	/* Reset command once we're done */
	currentConsoleCommand = ""
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
