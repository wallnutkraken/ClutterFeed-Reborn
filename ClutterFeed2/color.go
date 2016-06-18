package main

import (
	"github.com/rthornton128/goncurses"
)

const RGB_TO_CURSES_MULTIPLIER = 1000.0 / 255.0

/* Constants for color pairs that have meaning after initColors() finishes */
const (
	IDENTIFIER_PAIR = 101
	OWNTWEET_PAIR
	MENTION_PAIR
	COMMAND_PAIR
	ERROR_PAIR
	WARNING_PAIR
	WHITE_PAIR
)

const (
	C_IDENTIFIER = 91
	C_OWNTWEET
	C_MENTION
	C_COMMAND
	C_ERROR
	C_WARNING
)

func initColors(config ClutterFeedConfigFile) {
	/* Init default black/white pair */
	goncurses.InitPair(WHITE_PAIR, goncurses.C_WHITE, goncurses.C_BLACK)

	err := initSingleColor(config.Identifier, C_IDENTIFIER, IDENTIFIER_PAIR)
	errToStderr(err)

	err = initSingleColor(config.OwnTweet, C_OWNTWEET, OWNTWEET_PAIR)
	errToStderr(err)

	err = initSingleColor(config.Mention, C_MENTION, MENTION_PAIR)
	errToStderr(err)

	err = initSingleColor(config.CommandBox, C_COMMAND, COMMAND_PAIR)
	errToStderr(err)

	err = initSingleColor(config.Error, C_ERROR, ERROR_PAIR)
	errToStderr(err)

	err = initSingleColor(config.Warning, C_WARNING, WARNING_PAIR)
	errToStderr(err)
}

func initSingleColor(colorInfo ColorSetting, colorId int16, pairId int16) error {
	red := int16(RGB_TO_CURSES_MULTIPLIER * float32(colorInfo.Red))
	green := int16(RGB_TO_CURSES_MULTIPLIER * float32(colorInfo.Green))
	blue := int16(RGB_TO_CURSES_MULTIPLIER * float32(colorInfo.Blue))
	err := goncurses.InitColor(colorId, red, green, blue)
	if err != nil {
		return err
	}
	return goncurses.InitPair(pairId, colorId, goncurses.C_BLACK)
}
