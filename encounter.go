package main

import (
	"log"
	"strconv"
)

// An encounter contains generation for Hyperspace Hubris
type encounter struct {
	encounterType, condition, output string
}

// generate generates an encounter
func (e *encounter) generate() encounter {
	rnd := generateNumber(1, 6)
	if rnd < 3 {
		e.condition = generateWeather()
	} else {
		e.condition = "Smooth Sailing"
	}

	rnd2 := generateNumber(1, 6)
	switch rnd2 {
	case 1:
		e.encounterType = "The Opposition"
		e.output = e.encounterType
	case 2:
		e.encounterType = "Distress Signal"
		e.output = generateDistressSignal()
	case 3:
		e.encounterType = "Another Ship"
		e.output = generateAnotherShip() + " | " + generateShipStatus()
	case 4:
		creature, bearing := generateCreature()
		e.encounterType = "Space Creature"
		e.output = creature + " | " + bearing
	case 5:
		severity, issue := generateIssue()
		e.encounterType = "Onboard Issues"
		e.output = severity + " " + issue
	case 6:
		e.encounterType = "Strange Encounter"
		e.output = generateStrangeEncounter()
	}

	return *e
}

// render renders the encounter to the game log.
func (e *encounter) render() {
	renderOutput("[--- "+e.encounterType+" ---](fg:blue)", "", "clear")
	renderOutput(e.output, "", "clear")
}

// generateWeather returns a string containing a weather value.
func generateWeather() string {
	weathers, err := readNameFile("./data/spaceencounters/weather.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return weathers[generateNumber(0, len(weathers)-1)]
}

// generateDistressSignal returns a string containing a distress signal value.
func generateDistressSignal() string {
	distresses, err := readNameFile("./data/spaceencounters/distress.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	trap := ""

	rnd := generateNumber(1, 6)
	if rnd < 3 {
		trap = " | [TRAP](fg:red)"
	}
	return distresses[generateNumber(0, len(distresses)-1)] + trap
}

// generateAnotherShip returns a string containing a ship value.
func generateAnotherShip() string {
	rnd := generateNumber(1, 2)
	ships, err := readNameFile("./data/spaceencounters/ship0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return ships[generateNumber(0, len(ships)-1)]
}

// generateShipStatus returns a string containing a status value.
func generateShipStatus() string {
	rnd := generateNumber(1, 2)
	statuses, err := readNameFile("./data/spaceencounters/status0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return statuses[generateNumber(0, len(statuses)-1)]
}

// generateCreature returns two strings containing creature values.
func generateCreature() (string, string) {
	creatures, err := readNameFile("./data/spaceencounters/creature.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	bearings, err := readNameFile("./data/spaceencounters/bearing.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return creatures[generateNumber(0, len(creatures)-1)], bearings[generateNumber(0, len(bearings)-1)]
}

// generateIssue returns two strings containing issue values.
func generateIssue() (string, string) {
	issues, err := readNameFile("./data/spaceencounters/issue.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	severities, err := readNameFile("./data/spaceencounters/severity.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return severities[generateNumber(0, len(severities)-1)], issues[generateNumber(0, len(issues)-1)]
}

// generateStrangeEncounter returns a string containing a strange value.
func generateStrangeEncounter() string {
	stranges, err := readNameFile("./data/spaceencounters/strange.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return stranges[generateNumber(0, len(stranges)-1)]
}
