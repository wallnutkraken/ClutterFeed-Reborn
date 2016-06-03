package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigCreation(t *testing.T) {
	genericColor := ColorSetting{1, 1, 1}
	initialConfig := ClutterFeedConfigFile{
		ApiTokenPair{"usertoken", "usersecret"},
		ApiTokenPair{"apptoken", "appsecret"},
		genericColor,
		genericColor,
		genericColor,
		genericColor,
		genericColor,
		genericColor,
	}
	err := writeConfig(initialConfig)
	if err != nil {
		assert.Fail(t, "Failed at creating config file")
	}

	newConfig, err := readConfig()
	if err != nil {
		assert.Fail(t, "Failed at reading made config file")
	}

	assert.Equal(t, initialConfig, newConfig)
}
