package main

import (
	"log"
	"strconv"
)

// A mech contains generation for Mighty Mechs
type mech struct {
	weapon, system string
}

// generate generate a mech.
func (m *mech) generate() mech {
	m.weapon = generateMechWeapon()
	m.system = generateMechSystem()

	return *m
}

// render renders the mech to the game log.
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

// generateMechWeapon returns a string containing a mech weapon value.
func generateMechWeapon() string {
	rnd := generateNumber(1, 3)
	weapons, err := readNameFile("./data/mechs/weapon0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return weapons[generateNumber(0, len(weapons)-1)]
}

// generateMechSystem returns a string containing a mech system value.
func generateMechSystem() string {
	rnd := generateNumber(1, 6)
	systems, err := readNameFile("./data/mechs/system0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return systems[generateNumber(0, len(systems)-1)]
}
