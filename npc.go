package main

import (
	"log"
	"strconv"
)

// A npc contains generation for Friends & Frenemies
type npc struct {
	quirk, forte, demeanor, flaw, goal, object, secret string
}

// generate generates a npc.
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

// render renders the npc to the game log.
func (n *npc) render(req string) {

	switch req {
	case "quirk":
		renderOutput(n.quirk, "", "clear")
	case "forte":
		renderOutput(n.forte, "", "clear")
	case "demeanor":
		renderOutput(n.demeanor, "", "clear")
	case "flaw":
		renderOutput(n.flaw, "", "clear")
	case "goal":
		renderOutput(n.goal, "", "clear")
	case "object":
		renderOutput(n.object, "", "clear")
	case "secret":
		renderOutput(n.secret, "", "clear")
	case "notitle":
		renderOutput("Quirk: "+n.quirk, "", "clear")
		renderOutput("Forte: "+n.forte, "", "clear")
		renderOutput("Demeanor: "+n.demeanor, "", "clear")
		renderOutput("Flaw: "+n.flaw, "", "clear")
		renderOutput("Goal: "+n.goal, "", "clear")
		renderOutput("Object: "+n.object, "", "clear")
		renderOutput("Secret: "+n.secret, "", "clear")
	default:
		renderOutput("NPC", "h1", "purple")
		renderOutput("Quirk: "+n.quirk, "", "clear")
		renderOutput("Forte: "+n.forte, "", "clear")
		renderOutput("Demeanor: "+n.demeanor, "", "clear")
		renderOutput("Flaw: "+n.flaw, "", "clear")
		renderOutput("Goal: "+n.goal, "", "clear")
		renderOutput("Object: "+n.object, "", "clear")
		renderOutput("Secret: "+n.secret, "", "clear")
	}
}

// generateIntensity returns a string containing an intensity value.
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

// generateNpcQuirk returns a string containing a npc quirk value.
func generateNpcQuirk() string {
	rnd := generateNumber(1, 3)
	quirks, err := readNameFile("./data/npcs/quirk0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return generateIntensity() + " " + quirks[generateNumber(0, len(quirks)-1)]
}

// generateNpcForte returns a string containing a npc forte value.
func generateNpcForte() string {
	rnd := generateNumber(1, 2)
	fortes, err := readNameFile("./data/npcs/forte0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return generateIntensity() + " " + fortes[generateNumber(0, len(fortes)-1)]
}

// generateNpcDemeanor returns a string containing a npc demeanor value.
func generateNpcDemeanor() string {
	rnd := generateNumber(1, 3)
	demeanors, err := readNameFile("./data/npcs/demeanor0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return generateIntensity() + " " + demeanors[generateNumber(0, len(demeanors)-1)]
}

// generateNpcFlaw returns a string containing a npc flaw value.
func generateNpcFlaw() string {
	rnd := generateNumber(1, 2)
	flaws, err := readNameFile("./data/npcs/flaw0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return generateIntensity() + " " + flaws[generateNumber(0, len(flaws)-1)]
}

// generateNpcGoal returns a string containing a npc goal value.
func generateNpcGoal() string {
	goals, err := readNameFile("./data/npcs/goal.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return goals[generateNumber(0, len(goals)-1)]
}

// generateNpcObject returns a string containing a npc object value.
func generateNpcObject() string {
	objects, err := readNameFile("./data/npcs/object.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return objects[generateNumber(0, len(objects)-1)]
}

// generateNpcSecret returns a string containing a npc secret value.
func generateNpcSecret() string {
	secrets, err := readNameFile("./data/npcs/secret.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return secrets[generateNumber(0, len(secrets)-1)]
}
