package main

import (
	"log"
	"math/rand"
	"strconv"

	ui "github.com/gizak/termui/v3"
)

// cmdRoll generates and renders a D20 and D6 roll.
func cmdRoll(a []string) {
	d20 := generateNumber(1, 20)
	d6 := generateNumber(1, 6)
	extra := ""

	if d6 < 3 {
		extra = " [COST]"
	} else if d6 > 4 {
		extra = " [BENEFIT]"
	}

	output := "D20: " + strconv.Itoa(d20) + " D6: " + strconv.Itoa(d6) + extra

	renderOutput(output, "", "clear")
}

// cmdLog outputs all content after the command to the game log.
func cmdLog(a []string) {
	renderOutput(combineArgsToString(a[0:]), "logentry", "cyan")
}

// cmdName generates and renders a random character name.
func cmdName(a []string) {
	names, err := readNameFile("./data/pc/character.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	rnd := rand.Intn(len(names))

	renderOutput(names[rnd], "", "clear")
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

	renderOutput(output, "", "clear")
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

	renderOutput(output, "", "clear")
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

	renderOutput(output, "", "clear")
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
		renderOutput("All of a sudden...", "h1", "pink")
		renderOutput(one+" | "+two, "", "clear")
	} else if rnd < 5 {
		renderOutput("Feature of Interest", "h1", "pink")
		renderOutput("Feature: "+generateFeature(), "", "clear")
		renderOutput("Aspect: "+generateFeatureAspect(), "", "clear")
	} else {
		renderOutput("All of a sudden...", "h1", "pink")
		renderOutput(one+" | "+two, "", "clear")
		renderOutput("Feature of Interest", "", "clear")
		renderOutput("Feature: "+generateFeature(), "", "clear")
		renderOutput("Aspect: "+generateFeatureAspect(), "", "clear")
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
	case "skill":
		player.setAttribute("skill", combineArgsToString(a[1:]))
	case "style":
		player.setAttribute("style", combineArgsToString(a[1:]))
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
	case "help":
		renderOutput("The 'character' command is used to modify the fields related to your character.", "", "")
		renderOutput("To use it, simply choose what field you wish to change, and supply it and its new value. For example:", "", "")
		renderOutput("> character moxie +2", "", "8")
		renderOutput("Result:", "", "")
		renderOutput("The Moxie field in the Character block is set to '+2'.", "", "blue")
	default:
		renderOutput("Invalid subcommand: "+a[0], "error", "red")
		renderOutput("Try 'character help' for more info.", "info", "yellow")
	}
}

func cmdSkill(a []string) {
	renderOutput(generateSkill(), "", "clear")
}

// cmdHelp renders help data to the gamelog.
func cmdHelp(a []string) {
	lines, err := readNameFile("./data/help.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i := 0; i < len(lines); i++ {
		renderOutput(lines[i], "", "clear")
	}
}

func cmdLipsum(a []string) {
	lipsum := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum in egestas urna. Nullam sollicitudin id velit ut viverra. Curabitur facilisis massa non posuere consequat. Sed et massa porta, venenatis orci eget, lacinia ante. Phasellus laoreet mi ut purus elementum, et pharetra nisi dapibus. Suspendisse nisi velit, mollis eu tristique sed, porta quis nisl. Aliquam sed mattis quam. Morbi malesuada odio ut sagittis bibendum. Sed facilisis urna justo, non dapibus lacus accumsan a. Pellentesque est arcu, scelerisque quis enim sit amet, mattis interdum tortor.
Pellentesque elit libero, tempor sit amet fringilla non, rutrum laoreet nisi. Phasellus sed auctor lectus. Nulla facilisi. Quisque scelerisque faucibus risus, eget rhoncus mauris. Etiam in blandit dolor, nec pulvinar ex. Aenean volutpat facilisis lacus id posuere. Aenean egestas ac quam at lacinia. Duis sapien augue, faucibus sit amet venenatis at, fermentum nec odio. In maximus auctor libero, non pharetra erat rutrum eu. Fusce ornare suscipit mauris eu hendrerit.`

	renderOutput(lipsum, "", "orange")
}

func cmdNote(a []string) {
	switch a[0] {
	case "clear":
		scratchPad.Text = ""
		ui.Render(scratchPad)
	case "save":
		// TODO: Save to a provided file name, and clear the note block.
	case "help":
		renderOutput("Valid Subcommands:", "", "")
		renderOutput("clear - Clears out the note section. (No Undo)", "", "")
	default:
		renderOutput("Invalid subcommand: "+a[0], "error", "red")
		renderOutput("Try 'note help' for more info.", "info", "yellow")
	}
}

func cmdHeat(a []string) {
	value, err := strconv.Atoi(a[0])
	if err != nil {
		renderOutput("Heat Value must be an integer.", "error", "red")
	} else {
		gameLog.BorderStyle.Fg = primaryColor
		statBlock.BorderStyle.Fg = primaryColor
		heatGauge.BorderStyle.Fg = primaryColor
		missionBlock.BorderStyle.Fg = primaryColor
		scratchPad.BorderStyle.Fg = primaryColor
		heatGauge.Label = "Heat Level: " + a[0]
		heatGauge.Percent = value * 5

		if value < 6 {
			heatGauge.BarColor = ui.ColorBlue
		} else if value < 11 {
			heatGauge.BarColor = ui.ColorGreen
		} else if value < 16 {
			heatGauge.BarColor = ui.ColorYellow
		} else if value < 21 {
			heatGauge.BarColor = ui.ColorRed
		} else {
			heatGauge.BarColor = ui.ColorRed
			gameLog.BorderStyle.Fg = ui.ColorRed
			statBlock.BorderStyle.Fg = ui.ColorRed
			heatGauge.BorderStyle.Fg = ui.ColorRed
			missionBlock.BorderStyle.Fg = ui.ColorRed
			scratchPad.BorderStyle.Fg = ui.ColorRed
			ui.Render(gameLog, statBlock, heatGauge, missionBlock, scratchPad)
		}

	}
	ui.Render(gameLog, statBlock, heatGauge, missionBlock, scratchPad)
}
