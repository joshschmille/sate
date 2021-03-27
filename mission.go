package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
)

// A mission contains generation for Mission Generator
type mission struct {
	faction, mission, objective, location, aspect, opposition, agenda, snag string
}

// generate generates a mission.
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

// render renders the mission to the game log and the mission block.
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

		missionBlock.Text = m.faction + "\n" + m.mission + "\n" + m.objective + "\n" + m.location + "\n" + m.aspect + "\n" + m.opposition + "\n" + m.agenda + "\n" + m.snag
		ui.Render(missionBlock)
	}
}

// generateFaction returns a string containing a faction value.
func generateFaction() string {
	factions, err := readNameFile("./data/missions/faction.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return factions[generateNumber(0, len(factions)-1)]
}

// generateMission returns a string containing a mission value.
func generateMission() string {
	missions, err := readNameFile("./data/missions/mission.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return missions[generateNumber(0, len(missions)-1)]
}

// generateObjective returns a string containing an objective value.
func generateObjective() string {
	objectives, err := readNameFile("./data/missions/objective.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return objectives[generateNumber(0, len(objectives)-1)]
}

// generateLocation returns a string containing a location value.
func generateLocation() string {
	locations, err := readNameFile("./data/missions/location.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return locations[generateNumber(0, len(locations)-1)]
}

// generateLocationAspect returns a string containing a location aspect value.
func generateLocationAspect() string {
	aspects, err := readNameFile("./data/missions/aspect.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return aspects[generateNumber(0, len(aspects)-1)]
}

// generateOpposition returns a string containing a opposition value.
func generateOpposition() string {
	oppositions, err := readNameFile("./data/missions/opposition.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return oppositions[generateNumber(0, len(oppositions)-1)]
}

// generateAgenda returns a string containing a agenda value.
func generateAgenda() string {
	agendas, err := readNameFile("./data/missions/agenda.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return agendas[generateNumber(0, len(agendas)-1)]
}

// generateSnag returns a string containing a snag value.
func generateSnag() string {
	snags, err := readNameFile("./data/missions/snag.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return snags[generateNumber(0, len(snags)-1)]
}
