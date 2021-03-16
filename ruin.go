package main

import "log"

type ruin struct {
	ruinType, aesthetic, purpose, danger, depth, threat string
}

func (r *ruin) generate() ruin {
	r.ruinType = generateRuinType()
	r.aesthetic = generateAesthetic()
	r.purpose = generatePurpose()
	r.danger = generateDanger()
	r.depth = generateDepth()
	r.threat = generateThreat()

	return *r
}

func (r *ruin) render(req string) {
	switch req {
	case "type":
		renderOutput("Type: " + r.ruinType)
	case "aesthetic":
		renderOutput("Aesthetic: " + r.aesthetic)
	case "purpose":
		renderOutput("Purpose: " + r.purpose)
	case "danger":
		renderOutput("Danger Level: " + r.danger)
	case "depth":
		renderOutput("Depth: " + r.depth)
	case "threat":
		renderOutput("Threat: " + r.threat)
	case "treasure":
		renderOutput("Treasure: TREASURE")
	default:
		renderOutput("[--- Star Ruin ---](fg:purple)")
		renderOutput("Type: " + r.ruinType)
		renderOutput("Aesthetic: " + r.aesthetic)
		renderOutput("Purpose: " + r.purpose)
		renderOutput("Danger Level: " + r.danger)
		renderOutput("Depth: " + r.depth)
		renderOutput("Threat: " + r.threat)
		renderOutput("Treasure: TREASURE")
	}
}

func generateRuinType() string {
	types, err := readNameFile("./data/ruins/type.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return types[generateNumber(0, len(types)-1)]
}

func generateAesthetic() string {
	aesthetics, err := readNameFile("./data/ruins/aesthetic.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return aesthetics[generateNumber(0, len(aesthetics)-1)]
}

func generatePurpose() string {
	purposes, err := readNameFile("./data/ruins/purpose.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return purposes[generateNumber(0, len(purposes)-1)]
}

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

func generateDepth() string {
	depths, err := readNameFile("./data/ruins/depth.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return depths[generateNumber(0, len(depths)-1)]
}

func generateThreat() string {
	threats, err := readNameFile("./data/ruins/threat.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return threats[generateNumber(0, len(threats)-1)]
}
