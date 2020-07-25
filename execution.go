package main

import (
	"fmt"
	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
	"os/exec"
	"strings"
)

func psh_execute(command string) (string,string){

	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	// ... and interact with it
	stdout, stderr, err := shell.Execute(command)
	if err != nil {
		panic(err)
	}

	return stdout,stderr
}

func cmd_execute(command string) (string,string){
	cmd := exec.Command("cmd.exe", "/c", command)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		return string(output),"Command is unsuccesful"
	} else {
		return string(output),""

	}


}

func scenarioExecution(scenario Scenario,currentDirectory string, payloadDirectory string,scenarioOutput []ScenarioOutput)  {
	if scenario.Dependencies != "" {
		downloadFileDirectory := currentDirectory + "/" + scenario.Dependencies
		fileUrl := payloadDirectory + "/" + scenario.Dependencies
		DownloadFile(downloadFileDirectory, fileUrl)
	}
	var commandComplete = scenario.Command
	var commandOutput = ""
	var commandError = ""
	if(scenario.Interface == "psh") {
		commandOutput,commandError=psh_execute(scenario.Command)
	} else {
		if(scenario.ScenarioInputArgument != ""){
			command := strings.ReplaceAll(scenario.Command, "input_argument", scenario.ScenarioInputArgument)
			commandLast := strings.ReplaceAll(command, "output_argument", scenario.ScenarioOutputArgument)
			commandComplete = commandLast
			commandOutput,commandError=cmd_execute(commandLast)
		} else {
			commandOutput,commandError=cmd_execute(scenario.Command)
		}
	}
	var commandResult = true
	if commandError != "" {
		commandResult = false
	}
	n := ScenarioOutput{Id: scenario.Id, Name: scenario.Name, Description: scenario.Description,ScenarioCommand: commandComplete,ScenarioOutput: commandOutput, ScenarioErr: commandError, ScenarioResult: commandResult}
	scenarioOutput = append(scenarioOutput,n)
	printScenariosOutput(scenarioOutput)
}
