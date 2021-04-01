package main

import "log"

// A monster contains generation for Star Ruins & Space Hulks
type monster struct {
	monsterType, aspect, bearing, size string
}

// generate generates a monster.
func (m *monster) generate() monster {
	m.monsterType = generateMonsterType()
	m.aspect = generateMonsterAspect()
	m.bearing = generateMonsterBearing()
	m.size = generateMonsterSize()

	return *m
}

// render renders the monster to the game log.
func (m *monster) render(req string) {
	switch req {
	case "type":
		renderOutput(m.monsterType, "", "clear")
	case "aspect":
		renderOutput(m.aspect, "", "clear")
	case "bearing":
		renderOutput(m.bearing, "", "clear")
	case "size":
		renderOutput(m.size, "", "clear")
	default:
		renderOutput("Monster", "h1", "orange")
		renderOutput("Type: "+m.monsterType, "", "clear")
		renderOutput("Aspect: "+m.aspect, "", "clear")
		renderOutput("Bearing: "+m.bearing, "", "clear")
		renderOutput("Size: "+m.size, "", "clear")
		t := treasure{}
		t.generate()
		renderOutput("Treasure: "+t.aspect+" "+t.feature+" "+t.form, "", "magenta")
		//t.render("all")
	}
}

// generateMonsterType returns a string containing a monster type value.
func generateMonsterType() string {
	types, err := readNameFile("./data/monsters/type.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return types[generateNumber(0, len(types)-1)]
}

// generateMonsterAspect returns a string containing a monster aspect value.
func generateMonsterAspect() string {
	aspects, err := readNameFile("./data/monsters/aspect.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return aspects[generateNumber(0, len(aspects)-1)]
}

// generateMonsterBearing returns a string containing a monster bearing value.
func generateMonsterBearing() string {
	bearings, err := readNameFile("./data/monsters/bearing.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return bearings[generateNumber(0, len(bearings)-1)]
}

// generateMonsterSize returns a string containing a monster size value.
func generateMonsterSize() string {
	sizes, err := readNameFile("./data/monsters/size.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return sizes[generateNumber(0, len(sizes)-1)]
}
