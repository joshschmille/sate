package main

import (
	"log"
	"strconv"
)

// A hazard contains generation for Star Ruins & Space Hulks
type hazard struct {
	hazard string
}

// generate generates a hazard.
func (h *hazard) generate() hazard {
	hazardType := generateNumber(1, 3)
	hazards, err := readNameFile("./data/monsters/hazard0" + strconv.Itoa(hazardType) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	h.hazard = hazards[generateNumber(0, len(hazards)-1)]

	return *h
}

// render renders the hazard to the game log.
func (h *hazard) render() {
	renderOutput("--- Hazard ---")
	renderOutput("Hazard: " + h.hazard)
}
