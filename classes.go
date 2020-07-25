package main


type Threat struct {
	Description        string       `json:"description"`
	Id                string          `json:"id"`
	Name               string            `json:"name"`
	Phases               []string            `json:"phases"`

}

type ThreatConfig struct {
	ThreatGroupUrl        string       `json:"public-threat-group-url"`
	ThreatGroupScenarioUrl                string          `json:"public-threat-scenarios-url"`
	PayloadDirectory               string            `json:"payload-directory"`

}

type Scenario struct {
	Id        string       `json:"id"`
	Name                string          `json:"name"`
	Description               string            `json:"description"`
	Tactic               string            `json:"tactic"`
	Technique            map[string]string `json:"technique"`
	Platform            string `json:"platform"`
	Interface         string `json:"interface"`
	Command         string `json:"command"`
	Dependencies         string `json:"dependencies"`
	ScenarioType     string  `json:"type"`
	ScenarioInputArgument string `json:"input_argument"`
	ScenarioOutputArgument string `json:"output_argument"`

}

type ScenarioOutput struct {
	Id string
	Name string
	Description string
	ScenarioCommand string
	ScenarioOutput string
	ScenarioErr string
	ScenarioResult bool
}



