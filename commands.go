package main

import (
	"log"
	"math/rand"
	"strconv"

	ui "github.com/gizak/termui/v3"
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

	factions, err := readNameFile("./data/missions/faction.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	missions, err := readNameFile("./data/missions/mission.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	objectives, err := readNameFile("./data/missions/objective.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	oppositions, err := readNameFile("./data/missions/opposition.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	agendas, err := readNameFile("./data/missions/agenda.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	faction := factions[generateNumber(0, len(factions)-1)]
	mission := missions[generateNumber(0, len(missions)-1)]
	objective := objectives[generateNumber(0, len(objectives)-1)]
	location := generateLocation()
	aspect := generateLocationAspect()
	opposition := oppositions[generateNumber(0, len(oppositions)-1)]
	agenda := agendas[generateNumber(0, len(agendas)-1)]
	snag := generateSnag()

	renderOutput("[--- Mission Briefing ---](fg:green)")

	renderOutput("Faction: " + faction)
	renderOutput("Mission: " + mission)
	renderOutput("Objective: " + objective)
	renderOutput("Location: " + location)
	renderOutput("Location Aspect: " + aspect)
	renderOutput("Opposition: " + opposition)
	renderOutput("Agenda: " + agenda)
	renderOutput("Snag: " + snag)

	renderOutput("[--- End ---](fg:green)")

	// Update the mission in the sidebar
	missionBlock.Text = faction + "\n" + mission + "\n" + objective + "\n" + location + "\n" + aspect + "\n" + opposition + "\n" + agenda + "\n" + snag
	ui.Render(missionBlock)
}

func cmdEvent(a []string) {
	eventType := generateNumber(1, 6)

	if eventType < 5 {
		one, two := generateEvent(eventType)
		renderOutput(one + " | " + two)
	} else {
		one, two := generateEvent(generateNumber(1, 4))
		renderOutput(one + " | " + two)
		one, two = generateEvent(generateNumber(1, 4))
		renderOutput(one + " | " + two)
	}
}

func cmdRuin(a []string) {
	types, err := readNameFile("./data/ruins/type.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	aesthetics, err := readNameFile("./data/ruins/aesthetic.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	purposes, err := readNameFile("./data/ruins/purpose.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	depths, err := readNameFile("./data/ruins/depth.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	threats, err := readNameFile("./data/ruins/threat.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	danger := generateNumber(1, 20)
	var dangerResult string

	if danger < 10 {
		dangerResult = "Milk Run (5)"
	} else if danger < 15 {
		dangerResult = "Perilous (10)"
	} else {
		dangerResult = "Death Trap (15)"
	}

	renderOutput("[--- Star Ruin ---](fg:purple)")
	renderOutput("Type: " + types[generateNumber(0, len(types)-1)])
	renderOutput("Aesthetic: " + aesthetics[generateNumber(0, len(aesthetics)-1)])
	renderOutput("Purpose: " + purposes[generateNumber(0, len(purposes)-1)])
	renderOutput("Danger Level: " + dangerResult)
	renderOutput("Depth: " + depths[generateNumber(0, len(depths)-1)])
	renderOutput("Threat: " + threats[generateNumber(0, len(threats)-1)])
	generateTreasure()
}

func cmdMonster(a []string) {
	types, err := readNameFile("./data/monsters/type.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	aspects, err := readNameFile("./data/monsters/aspect.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	bearings, err := readNameFile("./data/monsters/bearing.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	sizes, err := readNameFile("./data/monsters/size.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	renderOutput("Type: " + types[generateNumber(0, len(types)-1)])
	renderOutput("Aspect: " + aspects[generateNumber(0, len(aspects)-1)])
	renderOutput("Bearing: " + bearings[generateNumber(0, len(bearings)-1)])
	renderOutput("Size: " + sizes[generateNumber(0, len(sizes)-1)])
}

func cmdTreasure(a []string) {
	renderOutput("Treasure: " + generateTreasure())
}

func cmdHazard(a []string) {
	renderOutput("Hazard: " + generateHazard())
}

func cmdGizmo(a []string) {
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

	renderOutput("[--- Gizmo ---](fg:pink)")

	renderOutput("Name: " + name)
	renderOutput("Effect: " + effect)
	renderOutput("Durability: " + durabilities[generateNumber(0, len(durabilities)-1)])

	renderOutput("[--- End ---](fg:pink)")
}

func cmdShip(a []string) {
	if a[0] != "" {
		switch a[0] {
		case "name":
			renderOutput("Ship Name: " + generateShipName())
		case "quirk":
			renderOutput("Ship Quirk: " + generateShipQuirk())
		case "perk":
			renderOutput("Ship Perk: " + generateShipPerk())
		case "origin":
			renderOutput("Ship Origin: " + generateShipOrigin())
		default:
			renderOutput("Invalid Subcommand: " + a[0])
		}
	} else {
		rndCondition := generateNumber(1, 6)
		rndType := generateNumber(1, 6)

		shipType := ""
		quirkCount := 0
		perkCount := 0

		if rndCondition < 4 {
			shipType += "Shiny "
			quirkCount = 1
		} else {
			shipType += "Scuffed "
			quirkCount = 2
		}

		if rndType < 4 {
			shipType += "Economy "
			perkCount = 1
		} else {
			shipType += "Luxury "
			perkCount = 2
		}
		shipType += "Starship"

		renderOutput("[--- Starship ---](fg:blue)")
		renderOutput("Name: " + generateShipName())
		renderOutput("Type: " + shipType)

		for i := 0; i < quirkCount; i++ {
			renderOutput("Quirk: " + generateShipQuirk())
		}

		for i := 0; i < perkCount; i++ {
			renderOutput("Perk: " + generateShipPerk())
		}

		renderOutput("Origin: " + generateShipOrigin())
	}
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
	p := generatePlanet()
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
