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
	rnd := generateNumber(1, 6)
	if rnd < 3 {
		renderOutput("Condition: " + generateWeather())
	} else {
		renderOutput("Condition: Smooth Sailing")
	}

	rnd2 := generateNumber(1, 6)
	switch rnd2 {
	case 1:
		renderOutput("Encounter: The Opposition")
	case 2:
		renderOutput("Distress Signal")
		renderOutput(generateDistressSignal())
	case 3:
		renderOutput("Another Ship")
		renderOutput("Ship: " + generateAnotherShip())
		renderOutput("Ship Status: " + generateShipStatus())
	case 4:
		creature, bearing := generateCreature()
		renderOutput("Space Creature")
		renderOutput(creature + " | " + bearing)
	case 5:
		severity, issue := generateIssue()
		renderOutput("Onboard Issues")
		renderOutput(severity + " " + issue)
	case 6:
		renderOutput("Strange Encounter")
		renderOutput(generateStrangeEncounter())
	}
}

func cmdSector(a []string) {
	rnd := generateNumber(1, 6)
	switch rnd {
	case 1:
		rndPlanet := generateNumber(1, 6)
		if rndPlanet < 4 {
			p := planet{}
			p.generate()
			p.render("all")
		} else if rndPlanet < 6 {
			p1 := planet{}
			p1.generate()
			p1.render("all")

			p2 := planet{}
			p2.generate()
			p2.render("all")
		} else {
			//1d6 planetoids
		}
	case 2:
		rndOutpost := generateNumber(1, 6)
		if rndOutpost < 4 {
			//single
		} else if rndOutpost < 6 {
			//twin
		} else {
			//1d6 planetoids
		}
	case 3:
		//nebula
	case 4:
		//asteroid
	case 5:
		//badlands
	case 6:
		//anomaly
	}
}
