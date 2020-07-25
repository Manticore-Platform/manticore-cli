package main

import (
	"github.com/fatih/color"
	"strconv"

)

func printDescription(out string){
	color.Yellow("%s", out)

}
func printStdOut(out string){
	color.Green("%s", out)

}

func printStdErr(err string) {
	color.Red("%s", err)
}

func printFunc(err string) {
	color.White("%s", err)
}


func printInfo(out string){
	color.Blue("%s", out)

}

func printScenariosOutput(scenariosOutput  []ScenarioOutput) {


	for _,row := range scenariosOutput {
		printInfo("=============================================================================\n")
		printFunc("ID :" + row.Id)
		printFunc("Name : "+ row.Name)
		printFunc("Description : "+row.Description)
		printInfo("=============================================================================\n")
		printInfo("Scenario Command : "+row.ScenarioCommand)
		printStdOut("Scenario Output : " +row.ScenarioOutput)
		printStdErr("Scenario StdErr: " + row.ScenarioErr)
		printDescription("Scenario Result: " + strconv.FormatBool(row.ScenarioResult))
		printInfo("=============================================================================\n")

	}
}
