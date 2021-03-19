package main

import (
	"log"
	"strconv"
)

type massivemonster struct {
	class                                 int
	mmType, weakSpot, motivation, zEffect string
	abilities, natures                    []string
}

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

func (mm *massivemonster) render() {

	renderOutput("--- Massive Monster ---")
	renderOutput(mm.mmType)
	renderOutput("Class: " + strconv.Itoa(mm.class))
	renderOutput("Motivation: " + mm.motivation)

	renderOutput("Abilities")
	if len(mm.abilities) > 0 {
		for i := 0; i < len(mm.abilities); i++ {
			renderOutput("- " + mm.abilities[i])
		}
	}

	renderOutput("Natures")
	if len(mm.natures) > 0 {
		for i := 0; i < len(mm.natures); i++ {
			renderOutput("- " + mm.natures[i])
		}
	}
	renderOutput("Weak Spot: " + mm.weakSpot)
	renderOutput("Incursion Zone Effect")
	renderOutput("- " + mm.zEffect)

}

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

func generateWeakSpot() string {
	spots, err := readNameFile("./data/massivemonsters/weakspot.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return spots[generateNumber(0, len(spots)-1)]
}

func generateMotivation() string {
	motivations, err := readNameFile("./data/massivemonsters/motivation.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return motivations[generateNumber(0, len(motivations)-1)]
}

func generateZEffect() string {
	effects, err := readNameFile("./data/massivemonsters/zeffect.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return effects[generateNumber(0, len(effects)-1)]
}
