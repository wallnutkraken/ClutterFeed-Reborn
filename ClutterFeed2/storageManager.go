package main

import (
	"encoding/json"
	"os"
)

type ApiTokenPair struct {
	Token  string
	Secret string
}

type ColorSetting struct {
	Red   byte
	Green byte
	Blue  byte
}

type ClutterFeedConfigFile struct {
	User       ApiTokenPair
	App        ApiTokenPair
	Identifier ColorSetting
	OwnTweet   ColorSetting
	Mention    ColorSetting
	CommandBox ColorSetting
	Error      ColorSetting
	Warning    ColorSetting
}

func readConfig() (ClutterFeedConfigFile, error) {
	configPath := "clutterfeed.conf" /* TODO: Low priority: Use something like */
	/* ~/.config/ClutterFeed or %AppData%/ClutterFeed to store this file. Ideally, */
	/* there should exist a way to get the config folder for the current user. */
	/* That would be the best cross-platform solution */
	file, err := os.Open(configPath)
	if err != nil {
		return ClutterFeedConfigFile{}, err
	}
	decoder := json.NewDecoder(file)
	loadedConfig := ClutterFeedConfigFile{}

	err = decoder.Decode(&loadedConfig)

	return loadedConfig, err
}

func writeConfig(settings ClutterFeedConfigFile) error {
	panic("Not implemented")
}
