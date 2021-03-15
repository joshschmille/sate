package main

import (
	"log"
	"strconv"
)

type planet struct {
	planetType, species, culture, feature, aspect, pickle string
}

func (p *planet) generate() planet {
	p.planetType = generatePlanetType()
	p.species = generateSpecies()
	p.culture = generateCulture()
	p.feature = generateFeature()
	p.aspect = generateFeatureAspect()
	p.pickle = generatePickle()

	return *p
}

func (p *planet) render(req string) {
	switch req {
	case "type":
		renderOutput("Planet Type: " + p.planetType) //render type
	case "species":
		renderOutput("Planet Species: " + p.species) //render species
	case "culture":
		renderOutput("Planet Culture: " + p.culture) //render culture
	case "feature":
		renderOutput("Planet Feature: " + p.feature) //render feature
	case "aspect":
		renderOutput("Planet Aspect: " + p.aspect) //render aspect
	case "pickle":
		renderOutput("Planet Pickle: " + p.pickle) //render pickle
	default:
		renderOutput("Planet Type: " + p.planetType) //render type
		renderOutput("Planet Species: " + p.species) //render species
		renderOutput("Planet Culture: " + p.culture) //render culture
		renderOutput("Planet Feature: " + p.feature) //render feature
		renderOutput("Planet Aspect: " + p.aspect)   //render aspect
		renderOutput("Planet Pickle: " + p.pickle)   //render pickle
	}
}

func generatePlanetType() string {
	rnd := generateNumber(1, 3)
	types, err := readNameFile("./data/planets/type0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return types[generateNumber(0, len(types)-1)]
}

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

func generateFeature() string {
	rnd := generateNumber(1, 3)
	features, err := readNameFile("./data/planets/feature0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return features[generateNumber(0, len(features)-1)]
}

func generateFeatureAspect() string {
	rnd := generateNumber(1, 3)
	aspects, err := readNameFile("./data/planets/aspect0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return aspects[generateNumber(0, len(aspects)-1)]
}

func generatePickle() string {
	pickles, err := readNameFile("./data/planets/pickle.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return pickles[generateNumber(0, len(pickles)-1)]
}
