package main

import "log"

// A treasure contains generation for Star Ruins & Space Hulks
type treasure struct {
	aspect, feature, form string
}

// generate generates a treasure.
func (t *treasure) generate() treasure {
	t.aspect = generateTreasureAspect()
	t.feature = generateTreasureFeature()
	t.form = generateTreasureForm()

	return *t
}

// render renders the treasure to the game log.
func (t *treasure) render(req string) {
	switch req {
	case "aspect":
		renderOutput(t.aspect, "", "clear")
	case "feature":
		renderOutput(t.feature, "", "clear")
	case "form":
		renderOutput(t.form, "", "clear")
	case "notitle":
		renderOutput(t.aspect+" "+t.feature+" "+t.form, "", "clear")
	default:
		renderOutput("Treasure", "h1", "yellow")
		renderOutput(t.aspect+" "+t.feature+" "+t.form, "", "clear")
	}
}

// generateTreasureAspect returns a string containing a treasure aspect value.
func generateTreasureAspect() string {
	aspects, err := readNameFile("./data/monsters/treasure01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return aspects[generateNumber(0, len(aspects)-1)]
}

// generateTreasureFeature returns a string containing a treasure feature value.
func generateTreasureFeature() string {
	features, err := readNameFile("./data/monsters/treasure02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return features[generateNumber(0, len(features)-1)]
}

// generateTreasureForm returns a string containing a treasure form value.
func generateTreasureForm() string {
	forms, err := readNameFile("./data/monsters/treasure03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return forms[generateNumber(0, len(forms)-1)]
}
