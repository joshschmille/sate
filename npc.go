package main

import (
	"log"
	"strconv"
)

type npc struct {
	quirk, forte, demeanor, flaw, goal, object, secret string
}

func (n *npc) generate() npc {
	n.quirk = generateNpcQuirk()
	n.forte = generateNpcForte()
	n.demeanor = generateNpcDemeanor()
	n.flaw = generateNpcFlaw()
	n.goal = generateNpcGoal()
	n.object = generateNpcObject()
	n.secret = generateNpcSecret()

	return *n
}

func (n *npc) render(req string) {

	switch req {
	case "quirk":
		renderOutput("Quirk: " + n.quirk)
	case "forte":
		renderOutput("Forte: " + n.forte)
	case "demeanor":
		renderOutput("Demeanor: " + n.demeanor)
	case "flaw":
		renderOutput("Flaw: " + n.flaw)
	case "goal":
		renderOutput("Goal: " + n.goal)
	case "object":
		renderOutput("Object: " + n.object)
	case "secret":
		renderOutput("Secret: " + n.secret)
	default:
		renderOutput("--- NPC ---")
		renderOutput("Quirk: " + n.quirk)
		renderOutput("Forte: " + n.forte)
		renderOutput("Demeanor: " + n.demeanor)
		renderOutput("Flaw: " + n.flaw)
		renderOutput("Goal: " + n.goal)
		renderOutput("Object: " + n.object)
		renderOutput("Secret: " + n.secret)
	}
}

func generateIntensity() string {
	rnd := generateNumber(1, 3)
	switch rnd {
	case 1:
		return "Slightly"
	case 2:
		return "Somewhat"
	case 3:
		return "Extremely"
	}

	return "Error: generateIntensity()"
}

func generateNpcQuirk() string {
	rnd := generateNumber(1, 3)
	quirks, err := readNameFile("./data/npcs/quirk0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return generateIntensity() + " " + quirks[generateNumber(0, len(quirks)-1)]
}

func generateNpcForte() string {
	rnd := generateNumber(1, 2)
	fortes, err := readNameFile("./data/npcs/forte0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return generateIntensity() + " " + fortes[generateNumber(0, len(fortes)-1)]
}

func generateNpcDemeanor() string {
	rnd := generateNumber(1, 3)
	demeanors, err := readNameFile("./data/npcs/demeanor0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return generateIntensity() + " " + demeanors[generateNumber(0, len(demeanors)-1)]
}

func generateNpcFlaw() string {
	rnd := generateNumber(1, 2)
	flaws, err := readNameFile("./data/npcs/flaw0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return generateIntensity() + " " + flaws[generateNumber(0, len(flaws)-1)]
}

func generateNpcGoal() string {
	goals, err := readNameFile("./data/npcs/goal.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return goals[generateNumber(0, len(goals)-1)]
}

func generateNpcObject() string {
	objects, err := readNameFile("./data/npcs/object.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return objects[generateNumber(0, len(objects)-1)]
}

func generateNpcSecret() string {
	secrets, err := readNameFile("./data/npcs/secret.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return secrets[generateNumber(0, len(secrets)-1)]
}
