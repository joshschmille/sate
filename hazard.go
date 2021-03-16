package main

import (
	"log"
	"strconv"
)

type hazard struct {
	hazard string
}

func (h *hazard) generate() hazard {
	hazardType := generateNumber(1, 3)
	hazards, err := readNameFile("./data/monsters/hazard0" + strconv.Itoa(hazardType) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	h.hazard = hazards[generateNumber(0, len(hazards)-1)]

	return *h
}

func (h *hazard) render() {
	renderOutput("--- Hazard ---")
	renderOutput("Hazard: " + h.hazard)
}
