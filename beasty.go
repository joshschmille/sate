package main

import (
	"log"
	"strconv"
)

// A beasty contains generation for Besties & Beasties
type beasty struct {
	appearance, size, gumption, personality, trait, ability string
}

// generate generates a beasty
func (b *beasty) generate() beasty {
	b.appearance = generateBeastyAppearance()
	b.size = generateBeastySize()
	b.gumption = generateBeastyGumption(b.size)
	b.personality = generateBeastyPersonality()
	b.trait = generateBeastyTrait()
	b.ability = generateBeastyAbility()

	return *b
}

// render renders the backstory to the game log.
func (b *beasty) render(req string) {
	switch req {
	case "appearance":
		renderOutput(b.appearance, "", "clear")
	case "size":
		renderOutput("Size: "+b.size, "", "clear")
	case "personality":
		renderOutput("Personality: "+b.personality, "", "clear")
	case "trait":
		renderOutput("Notable Trait: "+b.trait, "", "clear")
	case "ability":
		renderOutput("Special Ability: "+b.ability, "", "clear")
	default:
		renderOutput("--- Beasty ---", "", "clear")
		renderOutput(b.appearance, "", "clear")
		renderOutput("Size: "+b.size, "", "clear")
		renderOutput("Gumption: "+b.gumption, "", "clear")
		renderOutput("Personality: "+b.personality, "", "clear")
		renderOutput("Notable Trait: "+b.trait, "", "clear")
		renderOutput("Special Ability: "+b.ability, "", "clear")
	}
}

// generateBeastyAppearance returns a string containing an appearance value.
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

// generateBeastySize returns a string containing a size value.
func generateBeastySize() string {
	sizes, err := readNameFile("./data/beasties/size.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return sizes[generateNumber(0, len(sizes)-1)]
}

// generateBeastyGumption returns a string containing a gumption value
// based on the provided string (size)
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

// generateBeastyPersonality returns a string containing a personality value.
func generateBeastyPersonality() string {
	personalities, err := readNameFile("./data/beasties/personality.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return personalities[generateNumber(0, len(personalities)-1)]
}

// generateBeastyTrait returns a string containing a trait value.
func generateBeastyTrait() string {
	rnd := generateNumber(1, 3)
	traits, err := readNameFile("./data/beasties/trait0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return traits[generateNumber(0, len(traits)-1)]
}

// generateBeastyAbility returns a string containing an ability value.
func generateBeastyAbility() string {
	rnd := generateNumber(1, 3)
	abilities, err := readNameFile("./data/beasties/ability0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return abilities[generateNumber(0, len(abilities)-1)]
}
