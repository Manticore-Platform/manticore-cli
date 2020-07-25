package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func readPublicThreatJsonFromUrl(url string) (Threat, error){
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error")
	}

	defer resp.Body.Close()
	var threatStructure Threat
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &threatStructure); err != nil {
		fmt.Println("Error")
	}

	return threatStructure, nil
}

func readThreatScenarioJsonFromUrl(url string) ([]Scenario, error){
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error")
	}

	defer resp.Body.Close()
	var scenarioList []Scenario
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

