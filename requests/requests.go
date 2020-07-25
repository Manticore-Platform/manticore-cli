package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"manticore-cli/classes"
	"manticore-cli/log"
	"net/http"
	"os"
)

func ReadPublicThreatJsonFromUrl(url string) (classes.Threat, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error")
	}

	defer resp.Body.Close()
	var threatStructure classes.Threat
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &threatStructure); err != nil {
		fmt.Println("Error")
	}

	return threatStructure, nil
}

func ReadThreatScenarioJsonFromUrl(url string) ([]classes.Scenario, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error")
	}

	defer resp.Body.Close()
	var scenarioList []classes.Scenario
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &scenarioList); err != nil {
		fmt.Println("Error")
	}

	return scenarioList, nil
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.PrintStdErr("Can not reach threat repository.")
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
