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

var (
	configPath = "clutterfeed.conf" /* TODO: Low priority: Use something like */
	/* ~/.config/ClutterFeed or %AppData%/ClutterFeed to store this file. Ideally, */
	/* there should exist a way to get the config folder for the current user. */
	/* That would be the best cross-platform solution */
)

func readConfig() (ClutterFeedConfigFile, error) {
	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			defaultConf := getDefaultConfigFile()
			return defaultConf, writeConfig(defaultConf)
		}
		return ClutterFeedConfigFile{}, err
	}
	decoder := json.NewDecoder(file)
	loadedConfig := ClutterFeedConfigFile{}

	err = decoder.Decode(&loadedConfig)

	return loadedConfig, err
}

func writeConfig(settings ClutterFeedConfigFile) error {
	var file *os.File
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		/* Create the config file if it does not already exist */
		file, err = os.Create(configPath)
		if err != nil {
			return err
		}
	}
	encoder := json.NewEncoder(file)
	err := encoder.Encode(settings)
	return err
}

func getDefaultConfigFile() ClutterFeedConfigFile {
	/* Oh boy here we go. The reason this exists is because you cannot have a constant */
	/* from a struct literal, so we have this instead. What an ugly workaround. */
	defaultConfigFile := ClutterFeedConfigFile{
		ApiTokenPair{},
		ApiTokenPair{"g66DNw8cKonlyAdqMO2XBw", "XRpxFt8KSHFvKbHKVq7tIxWpsKsOHj7Bda5XriPQ2Zg"},
		ColorSetting{30, 20, 140},
		ColorSetting{0, 150, 20},
		ColorSetting{250, 180, 30},
		ColorSetting{175, 0, 200},
		ColorSetting{255, 0, 0},
		ColorSetting{255, 225, 0},
	}

	return defaultConfigFile
}
