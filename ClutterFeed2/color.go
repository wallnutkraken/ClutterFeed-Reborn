package main

import (
	"fmt"
	"os"

	"github.com/rthornton128/goncurses"
)

/* Constants for color pairs that have meaning after initColors() finishes */
const (
	IDENTIFIER_PAIR = 101
	OWNTWEET_PAIR
	MENTION_PAIR
	COMMAND_PAIR
	ERROR_PAIR
	WARNING_PAIR
)

const (
	C_IDENTIFIER = 51
	C_OWNTWEET
	C_MENTION
	C_COMMAND
	C_ERROR
	C_WARNING
)

func initColors(config ClutterFeedConfigFile) {
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

	err = initSingleColor(config.Warning, C_WARNING.WARNING_PAIR)
	errToStderr(err)
}

func initSingleColor(colorInfo ColorSetting, colorId int16, pairId int16) error {
	err := goncurses.InitColor(colorId, colorInfo.Red, colorInfo.Green, colorInfo.Blue)
	if err != nil {
		return err
	}
	return goncurses.InitPair(pairId, colorId, goncurses.C_BLACK)
}
