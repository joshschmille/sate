package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

// generateNumber returns an int within the min and max provided.
func generateNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn((max+1)-min) + min
}

// generateFlavor returns a string contain a flavor value.
func generateFlavor() string {
	rnd := generateNumber(1, 3)
	flavors, err := readNameFile("./data/events/flavor0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return flavors[generateNumber(0, len(flavors)-1)]
}

// generateSuddenEvent returns two strings containing event values.
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
