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
		renderOutput("Macguffin - 'macguffin' or 'mg' to generate.", "", "yellow")
	case 4:
		renderOutput("Macguffin - 'macguffin' or 'mg' to generate.", "", "yellow")
	case 5:
		generateWoo()
	case 6:
		generateWoo()
		generateWoo()
	}
}

// generateWoo generates a Woo, and renders it to the game log.
func generateWoo() {
	rnd := generateNumber(1, 6)
	switch rnd {
	case 1:
		renderOutput("Woo: Developed a Forte", "", "green")
		renderOutput(generateNpcForte(), "", "clear")
	case 2:
		renderOutput("Woo: Made a Friend", "", "green")
		n := npc{}
		n.generate()
		n.render("all")
	case 3:
		renderOutput("Woo: "+generateFaction()+" owes you a favor.", "", "green")
	case 4:
		renderOutput("Woo: Acquired a Gizmo", "", "green")
		g := gizmo{}
		g.generate()
		//renderOutput("Gizmo", "h2", "clear")
		g.render("notitle")
	case 5:
		renderOutput("Woo: Woo'd a Beasty", "", "green")
		b := beasty{}
		b.generate()
		b.render("all")
	case 6:
		renderOutput("Woo: Acquired a Mech or Ship", "", "green")
		renderOutput("Mech", "h3", "clear")
		m := mech{}
		m.generate()
		m.render("notitle")

		renderOutput("Ship", "h3", "clear")
		s := ship{}
		s.generate()
		s.render("notitle")
	}
}

// generateWoo generates a Woe, and renders it to the game log.
func generateWoe() {
	rnd := generateNumber(1, 6)
	switch rnd {
	case 1:
		renderOutput("Woe: Developed a Flaw", "", "red")
		renderOutput(generateNpcFlaw(), "", "clear")
	case 2:
		renderOutput("Woe: Made a Frenemy", "", "red")
		n := npc{}
		n.generate()
		n.render("all")
	case 3:
		renderOutput("Woe: Indebted to "+generateFaction(), "", "red")
	case 4:
		renderOutput("Woe: Injured (-1 Max Gumption)", "", "red")
	case 5:
		renderOutput("Woe: Robbed! (Lose a Woo!)", "", "red")
	case 6:
		renderOutput("Woe: Imprisoned for "+strconv.Itoa(generateNumber(1, 6))+" years.", "", "red")
	}
}
