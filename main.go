package main

import (
	"manticore-cli/classes"
	"manticore-cli/config"
	"manticore-cli/functions"
	"manticore-cli/log"
)

func main() {
	asciiArt :=
		`
=============================================================================
 _  _   __   __ _  ____  __  ___  __  ____  ____ 
( \/ ) / _\ (  ( \(_  _)(  )/ __)/  \(  _ \(  __)
/ \/ \/    \/    /  )(   )(( (__(  O ))   / ) _) 
\_)(_/\_/\_/\_)__) (__) (__)\___)\__/(__\_)(____)	

Written By @ggsec_
=============================================================================
`
	log.PrintDescription(asciiArt)

	scenariosOutput := []classes.ScenarioOutput{}

	threatGroupUrl, threatGroupScenariosUrl, payloadDirectoryUrl, currentDirectory := config.ConfigParser("config.ini")

	for _, row := range threatGroupUrl {
		functions.PublicThreatGroupParser(row, payloadDirectoryUrl, currentDirectory, threatGroupScenariosUrl, scenariosOutput)
	}

}
