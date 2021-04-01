package main

import (
	"log"
	"strconv"
)

// A massivemonster contains generation for Massive Monsters
type massivemonster struct {
	class                                 int
	mmType, weakSpot, motivation, zEffect string
	abilities, natures                    []string
}

// generate generates a massive monster.
func (mm *massivemonster) generate() massivemonster {
	mm.class = generateNumber(1, 3)
	mm.mmType = generateMMType()
	mm.weakSpot = generateWeakSpot()
	mm.motivation = generateMotivation()
	mm.zEffect = generateZEffect()

	mm.abilities = generateAbilities(mm.class)
	mm.natures = generateNatures(mm.class)

	return *mm
}

// render renders the massive monster to the game log.
func (mm *massivemonster) render() {

	renderOutput("Massive Monster", "h1", "red")
	renderOutput(mm.mmType, "", "clear")
	renderOutput("Class: "+strconv.Itoa(mm.class), "", "clear")
	renderOutput("Motivation: "+mm.motivation, "", "clear")

	renderOutput("Abilities", "h2", "clear")
	if len(mm.abilities) > 0 {
		for i := 0; i < len(mm.abilities); i++ {
			renderOutput(mm.abilities[i], "listitem", "blue")
		}
	}

	renderOutput("Natures", "h2", "clear")
	if len(mm.natures) > 0 {
		for i := 0; i < len(mm.natures); i++ {
			renderOutput(mm.natures[i], "listitem", "green")
		}
	}
	renderOutput("Weak Spot: "+mm.weakSpot, "", "yellow")
	renderOutput("Incursion Zone Effect", "h2", "purple")
	renderOutput(mm.zEffect, "listitem", "clear")

}

// generateMMType returns a string containing a massive monster type value.
func generateMMType() string {
	rnd := generateNumber(1, 2)
	types, err := readNameFile("./data/massivemonsters/type0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	rnd2 := generateNumber(1, 2)
	forms, err := readNameFile("./data/massivemonsters/form0" + strconv.Itoa(rnd2) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return types[generateNumber(0, len(types)-1)] + " " + forms[generateNumber(0, len(forms)-1)]
}

// generateAbilities returns a string slice containing 1-3 ability values.
func generateAbilities(c int) []string {
	output := []string{}

	for i := 0; i < c; i++ {

		rnd := generateNumber(1, 2)
		elements, err := readNameFile("./data/massivemonsters/abilityelement0" + strconv.Itoa(rnd) + ".names")
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}

		rnd2 := generateNumber(1, 2)
		types, err := readNameFile("./data/massivemonsters/abilitytype0" + strconv.Itoa(rnd2) + ".names")
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}

		output = append(output, elements[generateNumber(0, len(elements)-1)]+" "+types[generateNumber(0, len(types)-1)])
	}

	return output
}

// generateNatures returns a string slice containing 1-3 nature values.
func generateNatures(c int) []string {
	output := []string{}

	for i := 0; i < c; i++ {
		rnd := generateNumber(1, 2)
		natures, err := readNameFile("./data/massivemonsters/nature0" + strconv.Itoa(rnd) + ".names")
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		output = append(output, natures[generateNumber(0, len(natures)-1)])
	}

	return output
}

// generateWeakSpot returns a string containing a weak spot value.
func generateWeakSpot() string {
	spots, err := readNameFile("./data/massivemonsters/weakspot.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return spots[generateNumber(0, len(spots)-1)]
}

// generateMotivation returns a string containing a motivation value.
func generateMotivation() string {
	motivations, err := readNameFile("./data/massivemonsters/motivation.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return motivations[generateNumber(0, len(motivations)-1)]
}

// generateZEffect returns a string containing a ZEffect value.
func generateZEffect() string {
	effects, err := readNameFile("./data/massivemonsters/zeffect.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return effects[generateNumber(0, len(effects)-1)]
}
