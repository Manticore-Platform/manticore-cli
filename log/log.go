package log

import (
	"github.com/fatih/color"
	"manticore-cli/classes"
	"strconv"
)

func PrintDescription(out string) {
	color.Yellow("%s", out)

}
func printStdOut(out string) {
	color.Green("%s", out)

}

func PrintStdErr(err string) {
	color.Red("%s", err)
}

func printFunc(err string) {
	color.White("%s", err)
}

func printInfo(out string) {
	color.Blue("%s", out)

}

func PrintScenariosOutput(scenariosOutput []classes.ScenarioOutput) {

	for _, row := range scenariosOutput {
		printInfo("=============================================================================\n")
		printFunc("ID :" + row.Id)
		printFunc("Name : " + row.Name)
		printFunc("Description : " + row.Description)
		printInfo("=============================================================================\n")
		printInfo("Scenario Command : " + row.ScenarioCommand)
		printStdOut("Scenario Output : " + row.ScenarioOutput)
		PrintStdErr("Scenario StdErr: " + row.ScenarioErr)
		PrintDescription("Scenario Result: " + strconv.FormatBool(row.ScenarioResult))
		printInfo("=============================================================================\n")

	}
}
