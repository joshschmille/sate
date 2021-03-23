package main

import (
	"log"
	"math/rand"
	"strconv"
)

func cmdRoll(a []string) {
	output := ""

	output += "D20: " + strconv.Itoa(generateNumber(1, 20)) + " D6: " + strconv.Itoa(generateNumber(1, 6))

	renderOutput(output)
}

func cmdName(a []string) {
	names, err := readNameFile("./data/character.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	rnd := rand.Intn(len(names))

	renderOutput(names[rnd])
}

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

func cmdMission(a []string) {
	m := mission{}
	m.generate()
	m.render(a[0])
}

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

func cmdRuin(a []string) {
	r := ruin{}
	r.generate()
	r.render(a[0])
}

func cmdMonster(a []string) {
	m := monster{}
	m.generate()
	m.render(a[0])
}

func cmdTreasure(a []string) {
	t := treasure{}
	t.generate()
	t.render(a[0])
}

func cmdHazard(a []string) {
	h := hazard{}
	h.generate()
	h.render()
}

func cmdGizmo(a []string) {
	g := gizmo{}
	g.generate()
	g.render(a[0])
}

func cmdShip(a []string) {
	s := ship{}
	s.generate()
	s.render(a[0])
}

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

func cmdPlanet(a []string) {
	p := planet{}
	p.generate()
	p.render(a[0])
}

func cmdNavigate(a []string) {
	e := encounter{}
	e.generate()
	e.render()
}

func cmdSector(a []string) {
	s := sector{}
	s.generate()
	s.render(a[0])
}

func cmdNpc(a []string) {
	n := npc{}
	n.generate()
	n.render(a[0])
}

func cmdMech(a []string) {
	m := mech{}
	m.generate()
	m.render(a[0])
}

func cmdMassiveMonster(a []string) {
	mm := massivemonster{}
	mm.generate()
	mm.render()
}

func cmdBeasty(a []string) {
	b := beasty{}
	b.generate()
	b.render(a[0])
}

func cmdMacguffin(a []string) {
	m := macguffin{}
	m.generate()
	m.render(a[0])
}

func cmdBackstory(a []string) {
	bs := backstory{}
	bs.generate()
	bs.render()
}
