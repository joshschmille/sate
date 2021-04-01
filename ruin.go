package main

import "log"

// A ruin contains generation for Star Ruins & Space Hulks
type ruin struct {
	ruinType, aesthetic, purpose, danger, depth, threat string
	treasure                                            treasure
}

// generate generates a ruin
func (r *ruin) generate() ruin {
	r.ruinType = generateRuinType()
	r.aesthetic = generateAesthetic()
	r.purpose = generatePurpose()
	r.danger = generateDanger()
	r.depth = generateDepth()
	r.threat = generateThreat()
	r.treasure.generate()

	return *r
}

// render renders the ruin to the game log.
func (r *ruin) render(req string) {
	switch req {
	case "type":
		renderOutput("Type: "+r.ruinType, "", "clear")
	case "aesthetic":
		renderOutput("Aesthetic: "+r.aesthetic, "", "clear")
	case "purpose":
		renderOutput("Purpose: "+r.purpose, "", "clear")
	case "danger":
		renderOutput("Danger Level: "+r.danger, "", "clear")
	case "depth":
		renderOutput("Depth: "+r.depth, "", "clear")
	case "threat":
		renderOutput("Threat: "+r.threat, "", "clear")
	case "treasure":
		r.treasure.render("all")
	default:
		renderOutput("[--- Star Ruin ---](fg:purple)", "", "clear")
		renderOutput("Type: "+r.ruinType, "", "clear")
		renderOutput("Aesthetic: "+r.aesthetic, "", "clear")
		renderOutput("Purpose: "+r.purpose, "", "clear")
		renderOutput("Danger Level: "+r.danger, "", "clear")
		renderOutput("Depth: "+r.depth, "", "clear")
		renderOutput("Threat: "+r.threat, "", "clear")
		r.treasure.render("all")
	}
}

// generateRuinType returns a string containing a ruin type value.
func generateRuinType() string {
	types, err := readNameFile("./data/ruins/type.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return types[generateNumber(0, len(types)-1)]
}

// generateAesthetic returns a string containing a aesthetic value.
func generateAesthetic() string {
	aesthetics, err := readNameFile("./data/ruins/aesthetic.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return aesthetics[generateNumber(0, len(aesthetics)-1)]
}

// generatePurpose returns a string containing a purpose value.
func generatePurpose() string {
	purposes, err := readNameFile("./data/ruins/purpose.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return purposes[generateNumber(0, len(purposes)-1)]
}

// generateDanger returns a string containing a danger value.
func generateDanger() string {
	danger := generateNumber(1, 20)

	if danger < 10 {
		return "Milk Run (5)"
	} else if danger < 15 {
		return "Perilous (10)"
	} else {
		return "Death Trap (15)"
	}
}

// generateDepth returns a string containing a depth value.
func generateDepth() string {
	depths, err := readNameFile("./data/ruins/depth.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return depths[generateNumber(0, len(depths)-1)]
}

// generateThreat returns a string containing a threat value.
func generateThreat() string {
	threats, err := readNameFile("./data/ruins/threat.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return threats[generateNumber(0, len(threats)-1)]
}
