package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
)

type mission struct {
	faction, mission, objective, location, aspect, opposition, agenda, snag string
}

func (m *mission) generate() mission {
	m.faction = generateFaction()
	m.mission = generateMission()
	m.objective = generateObjective()
	m.location = generateLocation()
	m.aspect = generateLocationAspect()
	m.opposition = generateOpposition()
	m.agenda = generateAgenda()
	m.snag = generateSnag()
	return *m
}

func (m *mission) render(req string) {
	switch req {
	case "faction":
		renderOutput("Faction: " + m.faction)
	case "mission":
		renderOutput("Mission: " + m.mission)
	case "objective":
		renderOutput("Objective: " + m.objective)
	case "location":
		renderOutput("Location: " + m.location)
	case "aspect":
		renderOutput("Location Aspect: " + m.aspect)
	case "opposition":
		renderOutput("Opposition: " + m.opposition)
	case "agenda":
		renderOutput("Agenda: " + m.agenda)
	case "snag":
		renderOutput("Snag: " + m.snag)
	default:
		renderOutput("[--- Mission Briefing ---](fg:green)")

		renderOutput("Faction: " + m.faction)
		renderOutput("Mission: " + m.mission)
		renderOutput("Objective: " + m.objective)
		renderOutput("Location: " + m.location)
		renderOutput("Location Aspect: " + m.aspect)
		renderOutput("Opposition: " + m.opposition)
		renderOutput("Agenda: " + m.agenda)
		renderOutput("Snag: " + m.snag)

		renderOutput("[--- End ---](fg:green)")

		missionBlock.Text = m.faction + "\n" + m.mission + "\n" + m.objective + "\n" + m.location + "\n" + m.aspect + "\n" + m.opposition + "\n" + m.agenda + "\n" + m.snag
		ui.Render(missionBlock)
	}
}

func generateFaction() string {
	factions, err := readNameFile("./data/missions/faction.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return factions[generateNumber(0, len(factions)-1)]
}

func generateMission() string {
	missions, err := readNameFile("./data/missions/mission.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return missions[generateNumber(0, len(missions)-1)]
}

func generateObjective() string {
	objectives, err := readNameFile("./data/missions/objective.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return objectives[generateNumber(0, len(objectives)-1)]
}

func generateLocation() string {
	locations, err := readNameFile("./data/missions/location.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return locations[generateNumber(0, len(locations)-1)]
}

func generateLocationAspect() string {
	aspects, err := readNameFile("./data/missions/aspect.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return aspects[generateNumber(0, len(aspects)-1)]
}

func generateOpposition() string {
	oppositions, err := readNameFile("./data/missions/opposition.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return oppositions[generateNumber(0, len(oppositions)-1)]
}

func generateAgenda() string {
	agendas, err := readNameFile("./data/missions/agenda.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return agendas[generateNumber(0, len(agendas)-1)]
}

func generateSnag() string {
	snags, err := readNameFile("./data/missions/snag.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return snags[generateNumber(0, len(snags)-1)]
}
