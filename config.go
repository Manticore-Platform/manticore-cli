package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func configParser(config string) (string,string,string,string){
	var threatConfig ThreatConfig
	currentDirectory, err := os.Getwd()
	jsonFile, err := os.Open(currentDirectory+"/"+config)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &threatConfig)
	return threatConfig.ThreatGroupUrl,threatConfig.ThreatGroupScenarioUrl,threatConfig.PayloadDirectory,currentDirectory
}
