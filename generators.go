package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

func generateNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn((max+1)-min) + min
}

func generateFlavor() string {
	flavors01, err := readNameFile("./data/events/flavor01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	flavors02, err := readNameFile("./data/events/flavor02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	flavors03, err := readNameFile("./data/events/flavor03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	flavorList := generateNumber(1, 3)
	flavor := ""

	switch flavorList {
	case 1:
		flavor = flavors01[generateNumber(0, len(flavors01)-1)]
	case 2:
		flavor = flavors02[generateNumber(0, len(flavors02)-1)]
	case 3:
		flavor = flavors03[generateNumber(0, len(flavors03)-1)]
	}

	return flavor
}

func generateSuddenEvent() (string, string) {
	rnd := generateNumber(1, 6)
	one, two := "", ""
	switch rnd {
	case 1:
		one, two = generateScuffleEvent()
	case 2:
		one, two = generateEncounterEvent()
	case 3:
		one = generateLocationAspect()
		two = generateFlavor()
	case 4:
		one, two = generateDifficultyEvent()
	case 5:
		one, two = generateSocialEvent()
	case 6:
		one = "Snag"
		two = generateSnag()
	}
	return one, two
}

func generateWeather() string {
	weathers, err := readNameFile("./data/spaceencounters/weather.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return weathers[generateNumber(0, len(weathers)-1)]
}

func generateDistressSignal() string {
	distresses, err := readNameFile("./data/spaceencounters/distress.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return distresses[generateNumber(0, len(distresses)-1)]
}

func generateAnotherShip() string {
	rnd := generateNumber(1, 2)
	ships, err := readNameFile("./data/spaceencounters/ship0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return ships[generateNumber(0, len(ships)-1)]
}

func generateShipStatus() string {
	rnd := generateNumber(1, 2)
	statuses, err := readNameFile("./data/spaceencounters/status0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return statuses[generateNumber(0, len(statuses)-1)]
}

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

func generateStrangeEncounter() string {
	stranges, err := readNameFile("./data/spaceencounters/strange.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return stranges[generateNumber(0, len(stranges)-1)]
}
