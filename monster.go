package main

import "log"

type monster struct {
	monsterType, aspect, bearing, size string
}

func (m *monster) generate() monster {
	m.monsterType = generateMonsterType()
	m.aspect = generateMonsterAspect()
	m.bearing = generateMonsterBearing()
	m.size = generateMonsterSize()

	return *m
}

func (m *monster) render(req string) {
	switch req {
	case "type":
		renderOutput("Type: " + m.monsterType)
	case "aspect":
		renderOutput("Aspect: " + m.aspect)
	case "bearing":
		renderOutput("Bearing: " + m.bearing)
	case "size":
		renderOutput("Size: " + m.size)
	default:
		renderOutput("--- Monster ---")
		renderOutput("Type: " + m.monsterType)
		renderOutput("Aspect: " + m.aspect)
		renderOutput("Bearing: " + m.bearing)
		renderOutput("Size: " + m.size)
		t := treasure{}
		t.generate()
		t.render("all")
	}
}

func generateMonsterType() string {
	types, err := readNameFile("./data/monsters/type.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return types[generateNumber(0, len(types)-1)]
}

func generateMonsterAspect() string {
	aspects, err := readNameFile("./data/monsters/aspect.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return aspects[generateNumber(0, len(aspects)-1)]
}

func generateMonsterBearing() string {
	bearings, err := readNameFile("./data/monsters/bearing.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return bearings[generateNumber(0, len(bearings)-1)]
}

func generateMonsterSize() string {
	sizes, err := readNameFile("./data/monsters/size.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return sizes[generateNumber(0, len(sizes)-1)]
}
