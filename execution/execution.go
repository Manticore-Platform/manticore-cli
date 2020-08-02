package execution

import (
	"bytes"
	"fmt"
	"manticore-cli/classes"
	"manticore-cli/log"
	"manticore-cli/requests"
	"os/exec"
	"strings"
)

type PowerShell struct {
	powerShell string
}

func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
}

func (p *PowerShell) Execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

func psh_execute(command string) (string, string) {
	posh := New()
	stdout, stderr, err := posh.Execute(command)

	if err != nil {
		fmt.Println(err)
		return stdout, "Command is unsuccesful"
	}
	return stdout, stderr
}

func cmd_execute(command string) (string, string) {
	cmd := exec.Command("cmd.exe", "/c", command)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		return string(output), "Command is unsuccesful"
	} else {
		return string(output), ""

	}

}

func ScenarioExecution(scenario classes.Scenario, currentDirectory string, payloadDirectory string, scenarioOutput []classes.ScenarioOutput) {
	if scenario.Dependencies != "" {
		downloadFileDirectory := currentDirectory + "/" + scenario.Dependencies
		fileUrl := payloadDirectory + "/" + scenario.Dependencies
		requests.DownloadFile(downloadFileDirectory, fileUrl)
	}
	var commandComplete = scenario.Command
	var commandOutput = ""
	var commandError = ""
	if scenario.Interface == "psh" {
		commandOutput, commandError = psh_execute(scenario.Command)
	} else {
		if scenario.ScenarioInputArgument != "" {
			command := strings.ReplaceAll(scenario.Command, "input_argument", scenario.ScenarioInputArgument)
			commandLast := strings.ReplaceAll(command, "output_argument", scenario.ScenarioOutputArgument)
			commandComplete = commandLast
			commandOutput, commandError = cmd_execute(commandLast)
		} else {
			commandOutput, commandError = cmd_execute(scenario.Command)
		}
	}
	var commandResult = true
	if commandError != "" {
		commandResult = false
	}
	n := classes.ScenarioOutput{Id: scenario.Id, Name: scenario.Name, Description: scenario.Description, ScenarioCommand: commandComplete, ScenarioOutput: commandOutput, ScenarioErr: commandError, ScenarioResult: commandResult}
	scenarioOutput = append(scenarioOutput, n)
	log.PrintScenariosOutput(scenarioOutput)
}
