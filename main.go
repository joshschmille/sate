package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var gameLog = widgets.NewList()
var missionBlock = widgets.NewParagraph()
var termWidth = 0
var termHeight = 0

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	w, h := ui.TerminalDimensions()

	termWidth = w
	termHeight = h

	ui.StyleParserColorMap = map[string]ui.Color{
		"red":    ui.Color(196),
		"blue":   ui.Color(32),
		"black":  ui.Color(16),
		"cyan":   ui.Color(87),
		"yellow": ui.Color(226),
		"orange": ui.Color(202),
		"white":  ui.ColorWhite,
		"clear":  ui.ColorClear,
		"green":  ui.Color(46),
		"purple": ui.Color(99),
		"pink":   ui.Color(219),
		"8":      ui.Color(8),
	}

	primaryColor := ui.Color(32)
	secondaryColor := ui.Color(87)

	filteredWords := [8]string{
		"<Space>",
		"<Enter>",
		"<Backspace>",
		"<Tab>",
		"<Up>",
		"<Down>",
		"<Left>",
		"<Right>",
	}

	gameLog.Title = "Game Log"
	gameLog.Rows = append(gameLog.Rows, "Welcome to [Space Aces](fg:purple): Terminal Edition!")
	gameLog.Rows = append(gameLog.Rows, "")
	gameLog.ScrollBottom()
	gameLog.SetRect(0, 0, termWidth-40, termHeight-3)
	gameLog.BorderStyle.Fg = primaryColor
	gameLog.TitleStyle.Fg = secondaryColor

	statBlock := widgets.NewParagraph()
	statBlock.Title = "Stats"
	statBlock.Text = "Name: Riya\nMoxie: +1\nSmarts: 0\nWiggles: 0\nFriends: 0\nPockets: 0\nGumption: 10/10"
	statBlock.SetRect(termWidth-40, 0, termWidth, 9)
	statBlock.BorderStyle.Fg = primaryColor
	statBlock.TitleStyle.Fg = secondaryColor

	missionBlock.Title = "Mission"
	missionBlock.Text = ""
	missionBlock.SetRect(termWidth-40, 9, termWidth, 19)
	missionBlock.BorderStyle.Fg = primaryColor
	missionBlock.TitleStyle.Fg = secondaryColor

	inputBox := widgets.NewParagraph()
	inputBox.Text = ""
	inputBox.SetRect(0, termHeight-3, termWidth, termHeight)
	inputBox.BorderStyle.Fg = primaryColor
	inputBox.TitleStyle.Fg = secondaryColor

	ui.Render(gameLog, statBlock, missionBlock, inputBox)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "<C-c>":
			return
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			termWidth = payload.Width
			termHeight = payload.Height
			gameLog.SetRect(0, 0, payload.Width-40, payload.Height-3)
			statBlock.SetRect(payload.Width-40, 0, payload.Width, 9)
			missionBlock.SetRect(payload.Width-40, 9, payload.Width, 19)
			inputBox.SetRect(0, payload.Height-3, payload.Width, payload.Height)
			ui.Clear()
			ui.Render(gameLog, statBlock, missionBlock, inputBox)
		case "<Enter>":

			parseArgs(inputBox.Text)

			inputBox.Text = ""

			ui.Render(inputBox)

		case "<Space>":
			inputBox.Text += " "
		case "<Backspace>":
			length := len(inputBox.Text)
			if length > 0 {
				inputBox.Text = inputBox.Text[:length-1]
				ui.Render(inputBox)
			}
		case "<Up>":
			// TODO: Command History.
		case "<Down>":
			// TODO: Command History.

		}

		switch e.Type {
		case ui.KeyboardEvent:
			filtered := false
			for i := 0; i < len(filteredWords); i++ {
				if e.ID == filteredWords[i] {
					filtered = true
				}
			}

			if !filtered {
				inputBox.Text += e.ID
				ui.Render(inputBox)
			}
		}
	}
}

func readNameFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func generateNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn((max+1)-min) + min
}

func Chunks(s string, chunkSize int) []string {
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string
	chunk := make([]rune, chunkSize)
	len := 0
	for _, r := range s {
		chunk[len] = r
		len++
		if len == chunkSize {
			chunks = append(chunks, string(chunk))
			len = 0
		}
	}
	if len > 0 {
		chunks = append(chunks, string(chunk[:len]))
	}
	return chunks
}

func renderOutput(s string) {
	chunked := Chunks(s, termWidth-42)
	for i := 0; i < len(chunked); i++ {
		gameLog.Rows = append(gameLog.Rows, chunked[i])
	}
	gameLog.Rows = append(gameLog.Rows, "")
	gameLog.ScrollBottom()

	ui.Render(gameLog)
}

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

func generateSnag() string {
	snags, err := readNameFile("./data/missions/snag.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return snags[generateNumber(0, len(snags)-1)]
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

func generateEvent(t int) (string, string) {
	renderOutput("[--- Event ---](fg:green)")

	switch t {
	case 1:
		return generateScuffleEvent()
	case 2:
		return generateSocialEvent()
	case 3:
		return generateEncounterEvent()
	case 4:
		return generateDifficultyEvent()
	}

	return "ERROR", "Something is very wrong with the Event Generator."
}

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

func generateEncounterEvent() (string, string) {
	encounters, err := readNameFile("./data/events/encounter.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return encounters[generateNumber(0, len(encounters)-1)], generateFlavor()
}

func generateDifficultyEvent() (string, string) {
	difficulties, err := readNameFile("./data/events/difficulty.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return difficulties[generateNumber(0, len(difficulties)-1)], generateFlavor()
}

func generateFlavor() string {
	flavors01, err := readNameFile("./data/events/flavor01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	flavors02, err := readNameFile("./data/events/flavor02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	flavors03, err := readNameFile("./data/events/flavor03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	flavorList := generateNumber(1, 3)
	flavor := ""

	switch flavorList {
	case 1:
		flavor = flavors01[generateNumber(0, len(flavors01)-1)]
	case 2:
		flavor = flavors02[generateNumber(0, len(flavors02)-1)]
	case 3:
		flavor = flavors03[generateNumber(0, len(flavors03)-1)]
	}

	return flavor
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

func generateTreasure() string {
	aspects, err := readNameFile("./data/monsters/treasure01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	features, err := readNameFile("./data/monsters/treasure02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	forms, err := readNameFile("./data/monsters/treasure03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	aspectId := generateNumber(0, len(aspects)-1)
	featureId := generateNumber(0, len(features)-1)
	formId := generateNumber(0, len(forms)-1)

	return aspects[aspectId] + " " + features[featureId] + " " + forms[formId]
}

func cmdHazard(a []string) {
	renderOutput("Hazard: " + generateHazard())
}

func generateHazard() string {
	hazardType := generateNumber(1, 3)
	hazards, err := readNameFile("./data/monsters/hazard0" + strconv.Itoa(hazardType) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return hazards[generateNumber(0, len(hazards)-1)]
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

func generateShipName() string {
	names1, err := readNameFile("./data/shipnames/01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names2, err := readNameFile("./data/shipnames/02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names3, err := readNameFile("./data/shipnames/03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names4, err := readNameFile("./data/shipnames/04.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names5, err := readNameFile("./data/shipnames/05.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names6, err := readNameFile("./data/shipnames/06.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names7, err := readNameFile("./data/shipnames/07.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	output := ""
	nameType := generateNumber(0, 2)

	switch nameType {
	case 0:

		rnd := generateNumber(0, len(names3))
		rnd2 := generateNumber(0, len(names4))

		output = names3[rnd] + names4[rnd2]
	case 1:

		rnd := generateNumber(0, len(names1))
		rnd2 := generateNumber(0, len(names2))

		output = "The " + names1[rnd] + " " + names2[rnd2]
	case 2:

		rnd := generateNumber(0, len(names5))
		rnd2 := generateNumber(0, len(names6))
		rnd3 := generateNumber(0, len(names7))

		output = "The " + names5[rnd] + " " + names6[rnd2] + " " + names7[rnd3]
	}

	return output
}

func generateShipPerk() string {
	rnd := generateNumber(1, 6)
	perks, err := readNameFile("./data/ships/perk0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return perks[generateNumber(0, len(perks)-1)]
}

func generateShipQuirk() string {
	rnd := generateNumber(1, 3)
	quirks, err := readNameFile("./data/ships/quirk0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return quirks[generateNumber(0, len(quirks)-1)]
}

func generateShipOrigin() string {
	origins, err := readNameFile("./data/ships/origin.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return origins[generateNumber(0, len(origins)-1)]
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

func generateSuddenEvent() (string, string) {
	rnd := generateNumber(1, 6)
	one, two := "", ""
	switch rnd {
	case 1:
		one, two = generateScuffleEvent()
	case 2:
		one, two = generateEncounterEvent()
	case 3:
		one = generateLocationAspect()
		two = generateFlavor()
	case 4:
		one, two = generateDifficultyEvent()
	case 5:
		one, two = generateSocialEvent()
	case 6:
		one = "Snag"
		two = generateSnag()
	}
	return one, two
}

func cmdPlanet(a []string) {
	renderOutput("Planet Type: " + generatePlanetType())
	renderOutput("Species: " + generateSpecies())

	rnd := generateNumber(1, 6)
	if rnd < 4 {
		renderOutput("Culture: " + generateCulture())
	} else if rnd < 6 {
		renderOutput("Culture: " + generateCulture() + " [-](fg:green) " + generateCulture())
	} else {
		renderOutput("Culture: " + generateCulture() + " [-><-](fg:red) " + generateCulture())
	}

	renderOutput("Feature: " + generateFeature())
	renderOutput("Aspect: " + generateFeatureAspect())

	rnd2 := generateNumber(1, 6)
	if rnd2 < 4 {
		renderOutput("Pickle: N/A")
	} else if rnd2 < 6 {
		renderOutput("Pickle: " + generatePickle())
	} else {
		renderOutput("Pickles: " + generatePickle() + " | " + generatePickle())
	}
}

func generatePlanetType() string {
	rnd := generateNumber(1, 3)
	types, err := readNameFile("./data/planets/type0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return types[generateNumber(0, len(types)-1)]
}

func generateSpecies() string {
	rnd := generateNumber(1, 2)
	prefixes, err := readNameFile("./data/planets/prefix0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	rnd2 := generateNumber(1, 2)
	suffixes, err := readNameFile("./data/planets/suffix0" + strconv.Itoa(rnd2) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return prefixes[generateNumber(0, len(prefixes)-1)] + suffixes[generateNumber(0, len(suffixes)-1)]
}

func generateCulture() string {
	rnd := generateNumber(1, 3)
	cultures, err := readNameFile("./data/planets/culture0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return cultures[generateNumber(0, len(cultures)-1)]
}

func generateFeature() string {
	rnd := generateNumber(1, 3)
	features, err := readNameFile("./data/planets/feature0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return features[generateNumber(0, len(features)-1)]
}

func generateFeatureAspect() string {
	rnd := generateNumber(1, 3)
	aspects, err := readNameFile("./data/planets/aspect0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return aspects[generateNumber(0, len(aspects)-1)]
}

func generatePickle() string {
	pickles, err := readNameFile("./data/planets/pickle.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return pickles[generateNumber(0, len(pickles)-1)]
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

func generateWeather() string {
	weathers, err := readNameFile("./data/spaceencounters/weather.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return weathers[generateNumber(0, len(weathers)-1)]
}

func generateDistressSignal() string {
	distresses, err := readNameFile("./data/spaceencounters/distress.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return distresses[generateNumber(0, len(distresses)-1)]
}

func generateAnotherShip() string {
	rnd := generateNumber(1, 2)
	ships, err := readNameFile("./data/spaceencounters/ship0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return ships[generateNumber(0, len(ships)-1)]
}

func generateShipStatus() string {
	rnd := generateNumber(1, 2)
	statuses, err := readNameFile("./data/spaceencounters/status0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return statuses[generateNumber(0, len(statuses)-1)]
}

func generateCreature() (string, string) {
	creatures, err := readNameFile("./data/spaceencounters/creature.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	bearings, err := readNameFile("./data/spaceencounters/bearing.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return creatures[generateNumber(0, len(creatures)-1)], bearings[generateNumber(0, len(bearings)-1)]
}

func generateIssue() (string, string) {
	issues, err := readNameFile("./data/spaceencounters/issue.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	severities, err := readNameFile("./data/spaceencounters/severity.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return severities[generateNumber(0, len(severities)-1)], issues[generateNumber(0, len(issues)-1)]
}

func generateStrangeEncounter() string {
	stranges, err := readNameFile("./data/spaceencounters/strange.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return stranges[generateNumber(0, len(stranges)-1)]
}

func parseArgs(s string) {
	if len(s) > 0 {
		all := strings.Fields(s)
		cmd := all[0]
		args := make([]string, len(all))

		copy(args, all[1:])

		renderOutput("[>](fg:cyan) [" + cmd + " " + strings.Join(args, " ") + "](fg:8)")

		switch cmd {
		case "roll":
			cmdRoll(args)
		case "name":
			cmdName(args)
		case "likely":
			cmdLikely(args)
		case "possibly":
			cmdPossibly(args)
		case "unlikely":
			cmdUnlikely(args)
		case "mission":
			cmdMission(args)
		case "event":
			cmdEvent(args)
		case "ruin":
			cmdRuin(args)
		case "monster":
			cmdMonster(args)
		case "treasure":
			cmdTreasure(args)
		case "hazard":
			cmdHazard(args)
		case "gizmo":
			cmdGizmo(args)
		case "ship":
			cmdShip(args)
		case "explore":
			cmdExplore(args)
		case "planet":
			cmdPlanet(args)
		case "navigate":
			cmdNavigate(args)
		default:
			renderOutput("[Invalid Command.](fg:red)")
		}
	}
}
