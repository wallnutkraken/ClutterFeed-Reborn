package main

import (
	"errors"
	"fmt"

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

func getInput(in chan goncurses.Key) {
	for exiting == false {
		pressedKey := CommandWindow.GetChar()
		if canAcceptInput {
			in <- pressedKey
		}
	}
}

func drawConsole() {
	CommandWindow.Clear()

	CommandWindow.AttrOn(goncurses.A_BOLD)
	CommandWindow.ColorOn(COMMAND_PAIR)
	addString(CommandWindow, "[")
	CommandWindow.ColorOff(COMMAND_PAIR)
	addString(CommandWindow, fmt.Sprintf("%03d", len(currentConsoleCommand)))
	CommandWindow.ColorOn(COMMAND_PAIR)
	addString(CommandWindow, "] > ")
	CommandWindow.AttrOff(goncurses.A_BOLD)
	CommandWindow.ColorOff(COMMAND_PAIR)

	addString(CommandWindow, currentConsoleCommand)
}

func handleInput(in chan goncurses.Key) {
	for {
		gotChar := <-in
		if gotChar == goncurses.KEY_RETURN || gotChar == goncurses.KEY_ENTER {
			/* finished command */
			if currentConsoleCommand == "exit" {
				close(in)
				exiting = true
				applicationFinished.Done()
			}
			currentConsoleCommand = ""
		} else {
			currentConsoleCommand += string(rune(gotChar)) /* Heh, adding a char */
			/* to a string in Go isn't the most pleasant thing ever */
			drawConsole()
		}
	}
}
