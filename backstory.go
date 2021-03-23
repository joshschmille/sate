package main

import "strconv"

type backstory struct {
	origin  int
	origin1 planet
	origin2 string
	origin3 ship

	quirk, demeanor, goal, object string
	hig1, hig2, hig3              int

	faction1, faction2 string
}

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

func (b *backstory) render() {
	renderOutput("--- Backstory ---")
	renderOutput("[Place of Origin](fg:cyan)")
	if b.origin < 4 {
		b.origin1.render("all")
	} else if b.origin < 6 {
		renderOutput("Space Station: " + b.origin2)
	} else {
		b.origin3.render("all")
	}
	renderOutput("[Quirk & Demeanor](fg:cyan)")
	renderOutput("Quirk: " + b.quirk + " | Demeanor: " + b.demeanor)

	renderOutput("[Early Life](fg:cyan)")
	renderOutput("- Goal: " + b.goal + " | Object: " + b.object)

	howditgo()

	renderOutput("[First Steps](fg:cyan)")
	renderOutput("You worked for " + b.faction1)

	howditgo()

	renderOutput("[And Then...](fg:cyan)")
	renderOutput("You worked for " + b.faction2)

	howditgo()
}

func howditgo() {
	rnd := generateNumber(1, 6)
	switch rnd {
	case 1:
		generateWoe()
		generateWoe()
	case 2:
		generateWoe()
	case 3:
		renderOutput("[Macguffin](fg:yellow) - 'macguffin' to generate.")
	case 4:
		renderOutput("[Macguffin](fg:yellow) - 'macguffin' to generate.")
	case 5:
		generateWoo()
	case 6:
		generateWoo()
		generateWoo()
	}
}

func generateWoo() {
	renderOutput("[WOO](fg:green)")
	rnd := generateNumber(1, 6)
	switch rnd {
	case 1:
		renderOutput("Developed a Forte")
		renderOutput(generateNpcForte())
	case 2:
		renderOutput("Made a Friend")
		n := npc{}
		n.generate()
		n.render("all")
	case 3:
		renderOutput(generateFaction() + " owes you a favor.")
	case 4:
		renderOutput("Acquired a Gizmo")
		g := gizmo{}
		g.generate()
		g.render("all")
	case 5:
		renderOutput("Woo'd a Beasty")
		b := beasty{}
		b.generate()
		b.render("all")
	case 6:
		renderOutput("Acquired a Mech or Ship")
		renderOutput("Mech")
		m := mech{}
		m.generate()
		m.render("all")

		renderOutput("Ship")
		s := ship{}
		s.generate()
		s.render("all")
	}
}

func generateWoe() {
	renderOutput("[WOE](fg:red)")
	rnd := generateNumber(1, 6)
	switch rnd {
	case 1:
		renderOutput("Developed a Flaw")
		renderOutput(generateNpcFlaw())
	case 2:
		renderOutput("Made a Frenemy")
		n := npc{}
		n.generate()
		n.render("all")
	case 3:
		renderOutput("Indebted to " + generateFaction())
	case 4:
		renderOutput("Injured (-1 Max Gumption)")
	case 5:
		renderOutput("Robbed! (Lose a Woo!)")
	case 6:
		renderOutput("Imprisoned for " + strconv.Itoa(generateNumber(1, 6)) + " years.")
	}
}
