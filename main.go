package main

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
	printDescription(asciiArt)

	scenariosOutput := []ScenarioOutput{}

	threatGroupUrl,threatGroupScenariosUrl,payloadDirectoryUrl,currentDirectory := configParser("config.ini")

	publicThreatGroupParser(threatGroupUrl,payloadDirectoryUrl,currentDirectory,threatGroupScenariosUrl,scenariosOutput)


}
