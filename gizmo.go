package main

import "log"

type gizmo struct {
	name, effect, durability string
}

func (g *gizmo) generate() gizmo {

	types01, err := readNameFile("./data/gizmos/type01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	types02, err := readNameFile("./data/gizmos/type02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	types03, err := readNameFile("./data/gizmos/type03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	prefixes01, err := readNameFile("./data/gizmos/prefix01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	prefixes02, err := readNameFile("./data/gizmos/prefix02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	forms01, err := readNameFile("./data/gizmos/form01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	forms02, err := readNameFile("./data/gizmos/form02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	forms03, err := readNameFile("./data/gizmos/form03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	effects01, err := readNameFile("./data/gizmos/effect01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	effects02, err := readNameFile("./data/gizmos/effect02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	effects03, err := readNameFile("./data/gizmos/effect03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	aspects01, err := readNameFile("./data/gizmos/aspect01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	aspects02, err := readNameFile("./data/gizmos/aspect02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	aspects03, err := readNameFile("./data/gizmos/aspect03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	durabilities, err := readNameFile("./data/gizmos/durability.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	name := ""
	effect := ""

	typeRnd := generateNumber(1, 6)
	prefixRnd := generateNumber(1, 6)
	formRnd := generateNumber(1, 6)
	effectRnd := generateNumber(1, 6)
	aspectRnd := generateNumber(1, 6)

	if typeRnd < 3 {
		name += types01[generateNumber(0, len(types01)-1)] + " "
	} else if typeRnd < 5 {
		name += types02[generateNumber(0, len(types02)-1)] + " "
	} else {
		name += types03[generateNumber(0, len(types03)-1)] + " "
	}

	if prefixRnd < 4 {
		name += prefixes01[generateNumber(0, len(prefixes01)-1)]
	} else {
		name += prefixes02[generateNumber(0, len(prefixes02)-1)]
	}

	if formRnd < 3 {
		name += forms01[generateNumber(0, len(forms01)-1)]
	} else if formRnd < 5 {
		name += forms02[generateNumber(0, len(forms02)-1)]
	} else {
		name += forms03[generateNumber(0, len(forms03)-1)]
	}

	if effectRnd < 3 {
		effect += effects01[generateNumber(0, len(effects01)-1)] + " "
	} else if effectRnd < 5 {
		effect += effects02[generateNumber(0, len(effects02)-1)] + " "
	} else {
		effect += effects03[generateNumber(0, len(effects03)-1)] + " "
	}

	if aspectRnd < 3 {
		effect += aspects01[generateNumber(0, len(aspects01)-1)]
	} else if aspectRnd < 5 {
		effect += aspects02[generateNumber(0, len(aspects02)-1)]
	} else {
		effect += aspects03[generateNumber(0, len(aspects03)-1)]
	}

	g.name = name
	g.effect = effect
	g.durability = durabilities[generateNumber(0, len(durabilities)-1)]

	return *g
}

func (g *gizmo) render(req string) {
	switch req {
	case "name":
		renderOutput("Name: " + g.name)
	case "effect":
		renderOutput("Effect: " + g.effect)
	case "durability":
		renderOutput("Durability: " + g.durability)
	default:
		renderOutput("[--- Gizmo ---](fg:pink)")
		renderOutput("Name: " + g.name)
		renderOutput("Effect: " + g.effect)
		renderOutput("Durability: " + g.durability)
	}

}
