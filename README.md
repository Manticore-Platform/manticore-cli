# Manticore Adversary Emulation Client Tool

## Pre-Release 

### Manticore-Cli 0.0.1

https://github.com/Manticore-Platform/manticore-cli/releases/tag/0.0.1 


- Download config.ini
- Run executable within same directory (config.ini)


## Dependencies

Install dependency with the following command 

```
go get "github.com/fatih/color"
```

## Compile

Build with the following command on the directory.

```
go build -ldflags="-s -w" .
```

## Run

Download Executable from https://github.com/Manticore-Platform/manticore-cli/tags or Compile Yourself 
Download Config.ini from https://github.com/Manticore-Platform/manticore-cli/tags or this Repository
Run with the following command in one directory

```
manticore-cli.exe
```


## Config.ini

Executable parses config.ini for taking scenarios.

public-threat-group-url : Parses Threat Groups from Manticore Public Threat Repository (https://github.com/Manticore-Platform/public-threats - Example : https://github.com/Manticore-Platform/public-threats/blob/master/ransomware/ransomware.json)

public-threat-scenarios-url : Parses Threat Scenarios from Manticore Public Threat Scenario Repository (https://github.com/Manticore-Platform/public-scenarios)

payload-directory : Parses Payload used by Manticore Threat Scenario Repository (https://github.com/Manticore-Platform/attack-arsenal/tree/master/adversary_emulation/APT29/CALDERA_DIY/evals/payloads)

```
{
"public-threat-group-url" : ["https://raw.githubusercontent.com/Manticore-Platform/public-threats/master/lolbins/lolbins.json"],
"public-threat-scenarios-url" : "https://raw.githubusercontent.com/Manticore-Platform/public-scenarios/master",
"payload-directory" : "https://raw.githubusercontent.com/mitre-attack/attack-arsenal/master/adversary_emulation/APT29/CALDERA_DIY/evals/payloads"
}

```

## Disclaimer

We are not responsible for any damage you might cause with this go tool. Use at own risk and for testing and learning only. Demo purposes only! Use this to avoid threats and make better tools against them.
