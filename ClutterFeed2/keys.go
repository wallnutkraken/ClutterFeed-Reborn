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
	grabCommandCursor()
}

func goRight() {
	if horizontalStrPosition == len(currentConsoleCommand) || len(currentConsoleCommand) == 0 {
		return
	}
	horizontalStrPosition++
	grabCommandCursor()
}

func goLeft() {
	if horizontalStrPosition == 0 || len(currentConsoleCommand) == 0 {
		return
	}
	horizontalStrPosition--
	grabCommandCursor()
}

func addCharacter(characterKey rune) {
	char := string(characterKey) /* To easily concat the char to the current string */
	if len(currentConsoleCommand) == horizontalStrPosition {
		currentConsoleCommand += char
	} else {
		if horizontalStrPosition == 0 {
			currentConsoleCommand = char + currentConsoleCommand
		} else {
			currentConsoleCommand = currentConsoleCommand[0:horizontalStrPosition] +
				char + currentConsoleCommand[horizontalStrPosition:len(currentConsoleCommand)]
		}
	}
	horizontalStrPosition++
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
		} else if gotChar == goncurses.KEY_RIGHT {
			/* Move the cursor position right */
			goRight()
		} else {
			addCharacter(rune(gotChar))
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
