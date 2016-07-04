package main

import (
	"github.com/rthornton128/goncurses"
)

func backspace() {
	if len(currentConsoleCommand) == 0 || horizontalStrPosition == 0 {
		return
	}
	if len(currentConsoleCommand) == horizontalStrPosition {
		/* Remove last character */
		currentConsoleCommand = currentConsoleCommand[0 : horizontalStrPosition-1]
		horizontalStrPosition--
	} else {
		preCut := currentConsoleCommand[0 : horizontalStrPosition-1]
		postCut := currentConsoleCommand[horizontalStrPosition:len(currentConsoleCommand)]
		currentConsoleCommand = preCut + postCut
		horizontalStrPosition--
	}
}

func goLeft() {
	if horizontalStrPosition == 0 || len(currentConsoleCommand) == 0 {
		return
	}
	horizontalStrPosition--
	updateCursorChannel <- true
}

func handleInput(in chan goncurses.Key) {
	CommandWindow.Keypad(true) /* Will allow us to use the keypad in the console */
	for {
		gotChar := <-in
		if gotChar == goncurses.KEY_RETURN || gotChar == goncurses.KEY_ENTER {
			/* finished command */
			parseCommandText()
		} else if int(gotChar) == 127 {
			/* 127 is backspace */
			backspace()
		} else if gotChar == goncurses.KEY_LEFT {
			/* Move cursor position left */
			goLeft()
		} else {
			currentConsoleCommand += string(rune(gotChar))
			horizontalStrPosition++
		}
		drawConsole()
	}
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
