package main

import (
	"log"
	"strconv"
)

// A ship contains generation for Starship Shenanigans
type ship struct {
	shipType, name, origin string
	perks, quirks          []string
}

// generate generates a ship.
func (s *ship) generate() ship {
	s.name = generateShipName()
	s.origin = generateShipOrigin()

	rndCondition := generateNumber(1, 6)
	rndType := generateNumber(1, 6)

	shipType := ""
	quirkCount := 0
	perkCount := 0

	if rndCondition < 4 {
		shipType += "Shiny "
		quirkCount = 1
	} else {
		shipType += "Scuffed "
		quirkCount = 2
	}

	if rndType < 4 {
		shipType += "Economy "
		perkCount = 1
	} else {
		shipType += "Luxury "
		perkCount = 2
	}
	shipType += "Starship"

	s.shipType = shipType

	for i := 0; i < quirkCount; i++ {
		s.quirks = append(s.quirks, generateShipQuirk())
	}

	for i := 0; i < perkCount; i++ {
		s.perks = append(s.perks, generateShipPerk())
	}

	return *s
}

// render renders the ship to the game log.
func (s *ship) render(req string) {
	switch req {
	case "name":
		renderOutput("Name: " + s.name)
	case "perk":
		renderOutput("Perk: " + s.perks[0])
	case "quirk":
		renderOutput("Quirk: " + s.quirks[0])
	case "origin":
		renderOutput("Origin: " + s.origin)
	default:
		renderOutput("[--- Starship ---](fg:blue)")
		renderOutput(s.shipType)
		renderOutput("Name: " + s.name)
		for i := 0; i < len(s.perks); i++ {
			renderOutput("Perk: " + s.perks[i])
		}
		for i := 0; i < len(s.quirks); i++ {
			renderOutput("Quirk: " + s.quirks[i])
		}
		renderOutput("Origin: " + s.origin)
	}
}

// generateShipName returns a string containing a ship name, pieced together
// in a variety of ways.
func generateShipName() string {
	names1, err := readNameFile("./data/shipnames/01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names2, err := readNameFile("./data/shipnames/02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names3, err := readNameFile("./data/shipnames/03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names4, err := readNameFile("./data/shipnames/04.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names5, err := readNameFile("./data/shipnames/05.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names6, err := readNameFile("./data/shipnames/06.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names7, err := readNameFile("./data/shipnames/07.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	nameType := generateNumber(0, 2)

	switch nameType {
	case 0:
		return names3[generateNumber(0, len(names3)-1)] + names4[generateNumber(0, len(names4)-1)]
	case 1:
		return "The " + names1[generateNumber(0, len(names1)-1)] + " " + names2[generateNumber(0, len(names2)-1)]
	case 2:
		return "The " + names5[generateNumber(0, len(names5)-1)] + " " + names6[generateNumber(0, len(names6)-1)] + " " + names7[generateNumber(0, len(names7)-1)]
	}

	return "Something is wrong with generating a ship name."
}

// generateShipPerk returns a string containing a ship perk value.
func generateShipPerk() string {
	rnd := generateNumber(1, 6)
	perks, err := readNameFile("./data/ships/perk0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return perks[generateNumber(0, len(perks)-1)]
}

// generateShipQuirk returns a string containing a ship quirk value.
func generateShipQuirk() string {
	rnd := generateNumber(1, 3)
	quirks, err := readNameFile("./data/ships/quirk0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return quirks[generateNumber(0, len(quirks)-1)]
}

// generateShipOrigin returns a string containing a ship origin value.
func generateShipOrigin() string {
	origins, err := readNameFile("./data/ships/origin.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return origins[generateNumber(0, len(origins)-1)]
}
