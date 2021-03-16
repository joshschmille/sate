package main

import "log"

type treasure struct {
	aspect, feature, form string
}

func (t *treasure) generate() treasure {
	t.aspect = generateTreasureAspect()
	t.feature = generateTreasureFeature()
	t.form = generateTreasureForm()

	return *t
}

func (t *treasure) render(req string) {
	switch req {
	case "aspect":
		renderOutput(t.aspect)
	case "feature":
		renderOutput(t.feature)
	case "form":
		renderOutput(t.form)
	default:
		renderOutput("--- Treasure ---")
		renderOutput(t.aspect + " " + t.feature + " " + t.form)
	}
}

func generateTreasureAspect() string {
	aspects, err := readNameFile("./data/monsters/treasure01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return aspects[generateNumber(0, len(aspects)-1)]
}

func generateTreasureFeature() string {
	features, err := readNameFile("./data/monsters/treasure02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return features[generateNumber(0, len(features)-1)]
}

func generateTreasureForm() string {
	forms, err := readNameFile("./data/monsters/treasure03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return forms[generateNumber(0, len(forms)-1)]
}
