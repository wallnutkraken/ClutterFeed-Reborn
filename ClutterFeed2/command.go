package main

import (
	"errors"
	"fmt"

	"github.com/rthornton128/goncurses"
)

var currentConsoleCommand string

func startCommandConsole() error {
	if CommandWindow == nil {
		return errors.New("Command window is not initialized")
	}

	inputChannel := make(chan goncurses.Char)
	go getInput(inputChannel)
	go handleInput(inputChannel)

	return nil
}

func getInput(in chan goncurses.Char) {
	panic("Not implemented")
}

func drawConsole() {
	CommandWindow.Clear()

	CommandWindow.AttrOn(goncurses.A_BOLD)
	//CommandWindow.Color() TODO: do these color sets
	addString(CommandWindow, "[")
	CommandWindow.Color(goncurses.C_WHITE)
	addString(CommandWindow, fmt.Sprintf("%03d", len(currentConsoleCommand)))
	//CommandWindow.Color() see TODO above
	addString(CommandWindow, "] > ")
	CommandWindow.AttrOff(goncurses.A_BOLD)
	CommandWindow.Color(goncurses.C_WHITE)

	panic("Not done. Chill out. I really need to think this through before it becomes another shit console")
}

func handleInput(in chan goncurses.Char) {
	for {
		gotChar := <-in
		if gotChar == goncurses.KEY_RETURN || gotChar == goncurses.KEY_ENTER {
			/* finished command */
		} else {
			currentConsoleCommand += string(rune(gotChar)) /* Heh, adding a char */
			/* to a string in Go isn't the most pleasant thing ever */
		}
	}
}
