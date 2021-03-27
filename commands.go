package main

import (
	"log"
	"math/rand"
	"strconv"
)

// cmdRoll generates and renders a D20 and D6 roll.
func cmdRoll(a []string) {
	output := ""

	output += "D20: " + strconv.Itoa(generateNumber(1, 20)) + " D6: " + strconv.Itoa(generateNumber(1, 6))

	renderOutput(output)
}

// cmdLog outputs all content after the command to the game log.
func cmdLog(a []string) {
	renderOutput(combineArgsToString(a[0:]))
}

// cmdName generates and renders a random character name.
func cmdName(a []string) {
	names, err := readNameFile("./data/character.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	rnd := rand.Intn(len(names))

	renderOutput(names[rnd])
}

// cmdLikely uses "Ask The AI" to generate a response.
func cmdLikely(a []string) {
	output := ""

	if generateNumber(1, 20) > 5 {
		output += "Yes, "
	} else {
		output += "No, "
	}

	if generateNumber(1, 6) > 2 {
		output += "and..."
	} else {
		output += "but..."
	}

	renderOutput(output)
}

// cmdPossibly uses "Ask The AI" to generate a response.
func cmdPossibly(a []string) {
	output := ""

	if generateNumber(1, 20) > 10 {
		output += "Yes, "
	} else {
		output += "No, "
	}

	if generateNumber(1, 6) > 2 {
		output += "and..."
	} else {
		output += "but..."
	}

	renderOutput(output)
}

// cmdUnlikely uses "Ask The AI" to generate a response.
func cmdUnlikely(a []string) {
	output := ""

	if generateNumber(1, 20) > 15 {
		output += "Yes, "
	} else {
		output += "No, "
	}

	if generateNumber(1, 6) > 2 {
		output += "and..."
	} else {
		output += "but..."
	}

	renderOutput(output)
}

// cmdMission generates a mission, and renders it based on user args.
func cmdMission(a []string) {
	m := mission{}
	m.generate()
	m.render(a[0])
}

// cmdEvent generates an event, and renders it.
func cmdEvent(a []string) {
	eventType := generateNumber(1, 6)
	if eventType < 5 {
		e := event{}
		e.generate(eventType)
		e.render()
	} else {
		e := event{}
		e.generate(generateNumber(1, 4))
		e.render()
		e2 := event{}
		e2.generate(generateNumber(1, 4))
		e2.render()
	}
}

// cmdRuin generates a ruin, and renders it based on user args.
func cmdRuin(a []string) {
	r := ruin{}
	r.generate()
	r.render(a[0])
}

// cmdMonster generates a monster, and renders it based on user args.
func cmdMonster(a []string) {
	m := monster{}
	m.generate()
	m.render(a[0])
}

// cmdTreasure generates a treasure, and renders it based on user args.
func cmdTreasure(a []string) {
	t := treasure{}
	t.generate()
	t.render(a[0])
}

// cmdHazard generates a hazard, and renders it.
func cmdHazard(a []string) {
	h := hazard{}
	h.generate()
	h.render()
}

// cmdGizmo generates a gizmo, and renders it based on user args.
func cmdGizmo(a []string) {
	g := gizmo{}
	g.generate()
	g.render(a[0])
}

// cmdShip generates a ship, and renders it based on user args.
func cmdShip(a []string) {
	s := ship{}
	s.generate()
	s.render(a[0])
}

// cmdExplore generates a planetary exploration event, and
// renders it.
func cmdExplore(a []string) {
	rnd := generateNumber(1, 6)
	one, two := generateSuddenEvent()
	if rnd < 3 {
		renderOutput("All of a sudden...")
		renderOutput(one + " | " + two)
	} else if rnd < 5 {
		renderOutput("Feature of Interest")
		renderOutput("Feature: " + generateFeature())
		renderOutput("Aspect: " + generateFeatureAspect())
	} else {
		renderOutput("All of a sudden...")
		renderOutput(one + " | " + two)
		renderOutput("Feature of Interest")
		renderOutput("Feature: " + generateFeature())
		renderOutput("Aspect: " + generateFeatureAspect())
	}
}

// cmdPlanet generates a planet, and renders it based on user args.
func cmdPlanet(a []string) {
	p := planet{}
	p.generate()
	p.render(a[0])
}

// cmdNavigate generates a space encounter, and renders it.
func cmdNavigate(a []string) {
	e := encounter{}
	e.generate()
	e.render()
}

// cmdSector generates a sector object, and renders it based on user args.
func cmdSector(a []string) {
	s := sector{}
	s.generate()
	s.render(a[0])
}

// cmdNpc generates an NPC, and renders it based on user args.
func cmdNpc(a []string) {
	n := npc{}
	n.generate()
	n.render(a[0])
}

// cmdMech generates a mech, and renders it based on user args.
func cmdMech(a []string) {
	m := mech{}
	m.generate()
	m.render(a[0])
}

// cmdMassiveMonster generates a massive monster, and renders
// it based on user args.
func cmdMassiveMonster(a []string) {
	mm := massivemonster{}
	mm.generate()
	mm.render()
}

// cmdBeasty generates a beasty, and renders it based on user args.
func cmdBeasty(a []string) {
	b := beasty{}
	b.generate()
	b.render(a[0])
}

// cmdMacguffin generates a macguffin, and renders it based on user args.
func cmdMacguffin(a []string) {
	m := macguffin{}
	m.generate()
	m.render(a[0])
}

// cmdBackstory generates a backstory, and renders it based on user args.
func cmdBackstory(a []string) {
	bs := backstory{}
	bs.generate()
	bs.render()
}

// cmdCharacter is used to modify character data.
func cmdCharacter(a []string) {
	switch a[0] {
	case "name":
		player.setAttribute("name", combineArgsToString(a[1:]))
	case "moxie":
		player.setAttribute("moxie", combineArgsToString(a[1:]))
	case "smarts":
		player.setAttribute("smarts", combineArgsToString(a[1:]))
	case "wiggles":
		player.setAttribute("wiggles", combineArgsToString(a[1:]))
	case "friends":
		player.setAttribute("friends", combineArgsToString(a[1:]))
	case "pockets":
		player.setAttribute("pockets", combineArgsToString(a[1:]))
	case "gumption":
		player.setAttribute("gumption", combineArgsToString(a[1:]))
	default:
		renderOutput("Invalid subcommand: " + a[0])
	}
}

// cmdHelp renders help data to the gamelog.
func cmdHelp(a []string) {
	lines, err := readNameFile("./data/help.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i := 0; i < len(lines); i++ {
		renderOutput(lines[i])
	}
}
