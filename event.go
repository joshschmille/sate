package main

import "log"

// An event contains generation for Episode Events
type event struct {
	one, two string
}

// generate generates an event
func (e *event) generate(t int) event {

	switch t {
	case 1:
		e.one, e.two = generateScuffleEvent()
	case 2:
		e.one, e.two = generateSocialEvent()
	case 3:
		e.one, e.two = generateEncounterEvent()
	case 4:
		e.one, e.two = generateDifficultyEvent()
	}

	return *e
}

// render renders the event to the game log.
func (e *event) render() {
	renderOutput(e.one + " | " + e.two)
}

// generateScuffleEvent return two strings containing scuffle values.
func generateScuffleEvent() (string, string) {
	scuffles, err := readNameFile("./data/events/scuffle.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	tactics, err := readNameFile("./data/events/tactic.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	enemy := scuffles[generateNumber(0, len(scuffles)-1)]
	tactic := tactics[generateNumber(0, len(tactics)-1)]

	return enemy, tactic
}

// generateSocialEvent return two strings containing social values.
func generateSocialEvent() (string, string) {
	socials, err := readNameFile("./data/events/social.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	bearings, err := readNameFile("./data/events/bearing.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	social := socials[generateNumber(0, len(socials)-1)]
	bearing := bearings[generateNumber(0, len(bearings)-1)]

	return social, bearing
}

// generateEncounterEvent return two strings containing encounter values.
func generateEncounterEvent() (string, string) {
	encounters, err := readNameFile("./data/events/encounter.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return encounters[generateNumber(0, len(encounters)-1)], generateFlavor()
}

// generateDifficultyEvent return two strings containing difficulty values.
func generateDifficultyEvent() (string, string) {
	difficulties, err := readNameFile("./data/events/difficulty.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return difficulties[generateNumber(0, len(difficulties)-1)], generateFlavor()
}
