package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"manticore-cli/classes"
	"manticore-cli/log"
	"os"
)

func ConfigParser(config string) (string, string, string, string) {
	var threatConfig classes.ThreatConfig
	currentDirectory, err := os.Getwd()

	if _, err := os.Stat(currentDirectory + "/" + config); err == nil {
		jsonFile, err := os.Open(currentDirectory + "/" + config)
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &threatConfig)
		return threatConfig.ThreatGroupUrl, threatConfig.ThreatGroupScenarioUrl, threatConfig.PayloadDirectory, currentDirectory

	} else if os.IsNotExist(err) {
		log.PrintStdErr("Config File does not exist - Please Follow https://github.com/Manticore-Platform/manticore-cli for Config.ini")
		os.Exit(3)
	}
	log.PrintStdErr(err.Error())
	return "", "", "", ""
}
