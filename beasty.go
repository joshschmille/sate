package main

import (
	"log"
	"strconv"
)

type beasty struct {
	appearance, size, gumption, personality, trait, ability string
}

func (b *beasty) generate() beasty {
	b.appearance = generateBeastyAppearance()
	b.size = generateBeastySize()
	b.gumption = generateBeastyGumption(b.size)
	b.personality = generateBeastyPersonality()
	b.trait = generateBeastyTrait()
	b.ability = generateBeastyAbility()

	return *b
}

func (b *beasty) render(req string) {
	switch req {
	case "appearance":
		renderOutput(b.appearance)
	case "size":
		renderOutput("Size: " + b.size)
	case "personality":
		renderOutput("Personality: " + b.personality)
	case "trait":
		renderOutput("Notable Trait: " + b.trait)
	case "ability":
		renderOutput("Special Ability: " + b.ability)
	default:
		renderOutput("--- Beasty ---")
		renderOutput(b.appearance)
		renderOutput("Size: " + b.size)
		renderOutput("Gumption: " + b.gumption)
		renderOutput("Personality: " + b.personality)
		renderOutput("Notable Trait: " + b.trait)
		renderOutput("Special Ability: " + b.ability)
	}
}

func generateBeastyAppearance() string {
	var output string
	rnd := generateNumber(1, 6)
	if rnd < 5 {
		rnd := generateNumber(1, 3)
		appearances, err := readNameFile("./data/beasties/appearance0" + strconv.Itoa(rnd) + ".names")
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}

		output = appearances[generateNumber(0, len(appearances)-1)]
	} else {
		rnd := generateNumber(1, 3)
		appearances, err := readNameFile("./data/beasties/appearance0" + strconv.Itoa(rnd) + ".names")
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		rnd2 := generateNumber(1, 3)
		appearances2, err := readNameFile("./data/beasties/appearance0" + strconv.Itoa(rnd2) + ".names")
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}

		output = appearances[generateNumber(0, len(appearances)-1)] + " | " + appearances2[generateNumber(0, len(appearances2)-1)]
	}

	return output
}

func generateBeastySize() string {
	sizes, err := readNameFile("./data/beasties/size.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return sizes[generateNumber(0, len(sizes)-1)]
}

func generateBeastyGumption(s string) string {
	gumption := generateNumber(1, 6)
	switch s {
	case "Small":
		gumption += 1
	case "Medium":
		gumption += 3
	case "Large":
		gumption += 5
	case "Hugemongous":
		gumption += 8
	}

	return strconv.Itoa(gumption)
}

func generateBeastyPersonality() string {
	personalities, err := readNameFile("./data/beasties/personality.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return personalities[generateNumber(0, len(personalities)-1)]
}

func generateBeastyTrait() string {
	rnd := generateNumber(1, 3)
	traits, err := readNameFile("./data/beasties/trait0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return traits[generateNumber(0, len(traits)-1)]
}

func generateBeastyAbility() string {
	rnd := generateNumber(1, 3)
	abilities, err := readNameFile("./data/beasties/ability0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return abilities[generateNumber(0, len(abilities)-1)]
}