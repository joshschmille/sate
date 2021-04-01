package main

import "strconv"

// A backstory contains generation for Backstory Baloney
type backstory struct {
	origin  int
	origin1 planet
	origin2 string
	origin3 ship

	quirk, demeanor, goal, object string
	hig1, hig2, hig3              int

	faction1, faction2 string
}

// generate generates a backstory
func (b *backstory) generate() backstory {

	b.origin = generateNumber(1, 6)

	b.origin1 = planet{}
	b.origin1.generate()

	b.origin2 = generateLocationAspect()

	b.origin3 = ship{}
	b.origin3.generate()

	b.quirk = generateNpcQuirk()
	b.demeanor = generateNpcDemeanor()

	b.goal = generateNpcGoal()
	b.object = generateNpcObject()

	b.hig1 = generateNumber(1, 6)
	b.hig2 = generateNumber(1, 6)
	b.hig3 = generateNumber(1, 6)

	b.faction1 = generateFaction()
	b.faction2 = generateFaction()

	return *b
}

// render renders the backstory to the game log.
func (b *backstory) render() {
	renderOutput("Backstory", "h1", "clear")
	renderOutput("Place of Origin", "h2", "cyan")
	if b.origin < 4 {
		b.origin1.render("all")
	} else if b.origin < 6 {
		renderOutput("Space Station: "+b.origin2, "", "clear")
	} else {
		b.origin3.render("all")
	}
	renderOutput("Quirk & Demeanor", "h2", "cyan")
	renderOutput("Quirk: "+b.quirk+" | Demeanor: "+b.demeanor, "", "clear")

	renderOutput("Early Life", "h2", "cyan")
	renderOutput("Goal: "+b.goal+" | Object: "+b.object, "", "clear")

	howdItGo()

	renderOutput("First Steps", "h2", "cyan")
	renderOutput("You worked for "+b.faction1, "", "clear")

	howdItGo()

	renderOutput("And Then...", "h2", "cyan")
	renderOutput("You worked for "+b.faction2, "", "clear")

	howdItGo()
}

// howdItGo renders the result of "How'd It Go?" to the game log.
func howdItGo() {
	renderOutput("How'd It Go?", "h2", "clear")
	rnd := generateNumber(1, 6)
	switch rnd {
	case 1:
		generateWoe()
		generateWoe()
	case 2:
		generateWoe()
	case 3:
		renderOutput("[Macguffin](fg:yellow) - 'macguffin' to generate.", "", "clear")
	case 4:
		renderOutput("[Macguffin](fg:yellow) - 'macguffin' to generate.", "", "clear")
	case 5:
		generateWoo()
	case 6:
		generateWoo()
		generateWoo()
	}
}

// generateWoo generates a Woo, and renders it to the game log.
func generateWoo() {
	prefix := "[Woo](fg:green)"
	rnd := generateNumber(1, 6)
	switch rnd {
	case 1:
		renderOutput(prefix+": Developed a Forte", "", "clear")
		renderOutput(generateNpcForte(), "", "clear")
	case 2:
		renderOutput(prefix+": Made a Friend", "", "clear")
		n := npc{}
		n.generate()
		n.render("all")
	case 3:
		renderOutput(prefix+": "+generateFaction()+" owes you a favor.", "", "clear")
	case 4:
		renderOutput(prefix+": Acquired a Gizmo", "", "clear")
		g := gizmo{}
		g.generate()
		renderOutput("-- Gizmo --", "", "clear")
		g.render("notitle")
	case 5:
		renderOutput(prefix+": Woo'd a Beasty", "", "clear")
		b := beasty{}
		b.generate()
		b.render("all")
	case 6:
		renderOutput(prefix+": Acquired a Mech or Ship", "", "clear")
		renderOutput("Mech", "", "clear")
		m := mech{}
		m.generate()
		m.render("all")

		renderOutput("Ship", "", "clear")
		s := ship{}
		s.generate()
		s.render("all")
	}
}

// generateWoo generates a Woe, and renders it to the game log.
func generateWoe() {
	//renderOutput("[WOE](fg:red)")
	prefix := "[Woe](fg:red)"
	rnd := generateNumber(1, 6)
	switch rnd {
	case 1:
		renderOutput(prefix+": Developed a Flaw", "", "clear")
		renderOutput(generateNpcFlaw(), "", "clear")
	case 2:
		renderOutput(prefix+": Made a Frenemy", "", "clear")
		n := npc{}
		n.generate()
		n.render("all")
	case 3:
		renderOutput(prefix+": Indebted to "+generateFaction(), "", "clear")
	case 4:
		renderOutput(prefix+": Injured (-1 Max Gumption)", "", "clear")
	case 5:
		renderOutput(prefix+": Robbed! (Lose a Woo!)", "", "clear")
	case 6:
		renderOutput(prefix+": Imprisoned for "+strconv.Itoa(generateNumber(1, 6))+" years.", "", "clear")
	}
}
