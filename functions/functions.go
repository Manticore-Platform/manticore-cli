package functions

import (
	"manticore-cli/classes"
	"manticore-cli/execution"
	"manticore-cli/requests"
)

func PublicThreatGroupParser(url string, payloadDirectory string, currentDirectory string, scenarioDirectory string, scenariosOutput []classes.ScenarioOutput) {
	threatStructure, err := requests.ReadPublicThreatJsonFromUrl(url)
	if err != nil {
		panic(err)
	}

	publicThreatScenarioParser(threatStructure, payloadDirectory, currentDirectory, scenarioDirectory, scenariosOutput)

}

func publicThreatScenarioParser(threat classes.Threat, payloadDirectory string, currentDirectory string, scenarioDirectory string, scenariosOutput []classes.ScenarioOutput) {

	for _, row := range threat.Phases {
		scenarioList, err := requests.ReadThreatScenarioJsonFromUrl(scenarioDirectory + "/" + row)
		if err != nil {
			panic(err)
		}

		for _, row := range scenarioList {
			execution.ScenarioExecution(row, currentDirectory, payloadDirectory, scenariosOutput)
		}
	}
}
