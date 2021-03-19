package main

import (
	"log"
	"strconv"
)

type mech struct {
	weapon, system string
}

func (m *mech) generate() mech {
	m.weapon = generateMechWeapon()
	m.system = generateMechSystem()

	return *m
}

func (m *mech) render(req string) {
	switch req {
	case "weapon":
		renderOutput("Mech Weapon: " + m.weapon)
	case "system":
		renderOutput("Mech System: " + m.system)
	default:
		renderOutput("Mech Weapon: " + m.weapon)
		renderOutput("Mech System: " + m.system)
	}
}

func generateMechWeapon() string {
	rnd := generateNumber(1, 3)
	weapons, err := readNameFile("./data/mechs/weapon0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return weapons[generateNumber(0, len(weapons)-1)]
}

func generateMechSystem() string {
	rnd := generateNumber(1, 6)
	systems, err := readNameFile("./data/mechs/system0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return systems[generateNumber(0, len(systems)-1)]
}
