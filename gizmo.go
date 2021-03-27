package main

import (
	"log"
	"strconv"
)

// A gizmo contains generation for Gizmos & Gadgets
type gizmo struct {
	name, effect, durability string
}

// generate generates a gizmo.
func (g *gizmo) generate() gizmo {

	typeRnd := generateNumber(1, 3)
	prefixRnd := generateNumber(1, 2)
	formRnd := generateNumber(1, 3)
	effectRnd := generateNumber(1, 3)
	aspectRnd := generateNumber(1, 3)

	types, err := readNameFile("./data/gizmos/type0" + strconv.Itoa(typeRnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	prefixes, err := readNameFile("./data/gizmos/prefix0" + strconv.Itoa(prefixRnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	forms, err := readNameFile("./data/gizmos/form0" + strconv.Itoa(formRnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	effects, err := readNameFile("./data/gizmos/effect0" + strconv.Itoa(effectRnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	aspects, err := readNameFile("./data/gizmos/aspect0" + strconv.Itoa(aspectRnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	durabilities, err := readNameFile("./data/gizmos/durability.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	name := ""
	effect := ""

	name += types[generateNumber(0, len(types)-1)] + " "
	name += prefixes[generateNumber(0, len(prefixes)-1)]
	name += forms[generateNumber(0, len(forms)-1)]

	effect += effects[generateNumber(0, len(effects)-1)] + " "
	effect += aspects[generateNumber(0, len(aspects)-1)]

	g.name = name
	g.effect = effect
	g.durability = durabilities[generateNumber(0, len(durabilities)-1)]

	return *g
}

// render renders the gizmo to the game log.
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
