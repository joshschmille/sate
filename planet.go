package main

import (
	"log"
	"strconv"
)

// A planet contains generation for Galactic Guidebook
type planet struct {
	planetType, species, culture, feature, aspect, pickle string
}

// generate generate a planet.
func (p *planet) generate() planet {
	p.planetType = generatePlanetType(false)
	p.species = generateSpecies()
	p.culture = generateCulture()
	p.feature = generateFeature()
	p.aspect = generateFeatureAspect()
	p.pickle = generatePickle()

	return *p
}

// render renders the planet to the game log.
func (p *planet) render(req string) {
	switch req {
	case "type":
		renderOutput("Planet Type: "+p.planetType, "", "clear")
	case "species":
		renderOutput("Planet Species: "+p.species, "", "clear")
	case "culture":
		renderOutput("Planet Culture: "+p.culture, "", "clear")
	case "feature":
		renderOutput("Planet Feature: "+p.feature, "", "clear")
	case "aspect":
		renderOutput("Planet Aspect: "+p.aspect, "", "clear")
	case "pickle":
		renderOutput("Planet Pickle: "+p.pickle, "", "clear")
	default:
		renderOutput("Planet Type: "+p.planetType, "", "clear")
		renderOutput("Planet Species: "+p.species, "", "clear")
		renderOutput("Planet Culture: "+p.culture, "", "clear")
		renderOutput("Planet Feature: "+p.feature, "", "clear")
		renderOutput("Planet Aspect: "+p.aspect, "", "clear")
		renderOutput("Planet Pickle: "+p.pickle, "", "clear")
	}
}

// generatePlanetType returns a string containing a planet type value.
func generatePlanetType(block bool) string {
	rnd := generateNumber(1, 3)
	types, err := readNameFile("./data/planets/type0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	output := ""
	var rnd2 int

	if block {
		rnd2 = generateNumber(0, len(types)-2)
	} else {
		rnd2 = generateNumber(0, len(types)-1)
	}

	if types[rnd2] == "Roll Twice" {
		output = generatePlanetType(true) + " | " + generatePlanetType(true)
	} else {
		output = types[rnd2]
	}

	return output
}

// generateSpecies returns a string containing a planet species value.
func generateSpecies() string {
	rnd := generateNumber(1, 2)
	prefixes, err := readNameFile("./data/planets/prefix0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	rnd2 := generateNumber(1, 2)
	suffixes, err := readNameFile("./data/planets/suffix0" + strconv.Itoa(rnd2) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return prefixes[generateNumber(0, len(prefixes)-1)] + suffixes[generateNumber(0, len(suffixes)-1)]
}

// generateCulture returns a string containing a planet culture value.
func generateCulture() string {
	rnd := generateNumber(1, 3)
	cultures, err := readNameFile("./data/planets/culture0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	rnd2 := generateNumber(1, 6)
	if rnd2 < 4 {
		return cultures[generateNumber(0, len(cultures)-1)]
	} else if rnd2 < 6 {
		return cultures[generateNumber(0, len(cultures)-1)] + " [-](fg:green) " + cultures[generateNumber(0, len(cultures)-1)]
	} else {
		return cultures[generateNumber(0, len(cultures)-1)] + " [-><-](fg:red) " + cultures[generateNumber(0, len(cultures)-1)]
	}
}

// generateFeature returns a string containing a planet feature value.
func generateFeature() string {
	rnd := generateNumber(1, 3)
	features, err := readNameFile("./data/planets/feature0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return features[generateNumber(0, len(features)-1)]
}

// generateFeatureAspect returns a string containing a planet feature aspect value.
func generateFeatureAspect() string {
	rnd := generateNumber(1, 3)
	aspects, err := readNameFile("./data/planets/aspect0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return aspects[generateNumber(0, len(aspects)-1)]
}

// generatePickle returns a string containing a planet pickle value.
func generatePickle() string {
	pickles, err := readNameFile("./data/planets/pickle.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return pickles[generateNumber(0, len(pickles)-1)]
}
