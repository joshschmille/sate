package main

import (
	"log"
	"math/rand"
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
