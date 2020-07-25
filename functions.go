package main

func publicThreatGroupParser(url string,payloadDirectory string,currentDirectory string,scenarioDirectory string,scenariosOutput []ScenarioOutput){
	threatStructure, err := readPublicThreatJsonFromUrl(url)
	if err != nil {
		panic(err)
	}

	publicThreatScenarioParser(threatStructure,payloadDirectory,currentDirectory,scenarioDirectory,scenariosOutput)

}

func publicThreatScenarioParser(threat Threat,payloadDirectory string,currentDirectory string,scenarioDirectory string,scenariosOutput  []ScenarioOutput){

	for _,row := range threat.Phases {
		scenarioList, err := readThreatScenarioJsonFromUrl(scenarioDirectory+"/"+row)
		if err != nil {
			panic(err)
		}


		for _, row := range scenarioList {
			scenarioExecution(row,currentDirectory,payloadDirectory,scenariosOutput)
		}
	}
}